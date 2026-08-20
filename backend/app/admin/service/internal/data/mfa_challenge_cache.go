package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/tx7do/go-utils/jwtutil"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"google.golang.org/protobuf/proto"

	authenticationV1 "go-wind-admin/api/gen/go/authentication/service/v1"
)

const (
	// MfaChallengeTTL 挑战/注册操作上下文的有效期。
	MfaChallengeTTL = 5 * time.Minute

	// 登录挑战上下文 key 前缀：mfa:login:<operation_id>
	mfaLoginChallengeKeyFmt = "mfa:login:%s"
	// 注册上下文 key 前缀：mfa:enroll：<operation_id>
	mfaEnrollChallengeKeyFmt = "mfa:enroll:%s"
)

var ErrMfaChallengeNotFound = errors.New("mfa challenge not found or expired")

// MfaLoginChallengeContext 登录挑战上下文。
// 密码校验通过且用户绑定 TOTP 后，由 doGrantTypePassword 写入；
// VerifyMFAChallenge 取出并用于签发真 token。
type MfaLoginChallengeContext struct {
	Payload    *authenticationV1.UserTokenPayload
	ClientType authenticationV1.ClientType
}

// MfaEnrollChallengeContext 注册上下文。
// StartEnrollMethod 生成 TOTP secret 后写入；ConfirmEnrollMethod 取出 secret 校验首码。
type MfaEnrollChallengeContext struct {
	Secret   string
	TenantID uint32
	UserID   uint32
}

// MfaChallengeCache MFA 操作上下文缓存。
// 所有操作均 verify-and-delete 单次有效（取即删），与 captchaClient 一致。
type MfaChallengeCache struct {
	log *log.Helper
	rdb *redis.Client
}

func NewMfaChallengeCache(ctx *bootstrap.Context, rdb *redis.Client) *MfaChallengeCache {
	return &MfaChallengeCache{
		rdb: rdb,
		log: ctx.NewLoggerHelper("mfa-challenge/cache"),
	}
}

// SetLoginChallenge 写入登录挑战上下文，返回 operation_id。
func (c *MfaChallengeCache) SetLoginChallenge(ctx context.Context, payload *authenticationV1.UserTokenPayload, clientType authenticationV1.ClientType) (string, error) {
	opId := newOperationID()
	data, err := proto.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal login challenge payload failed: %w", err)
	}
	envelope := mfaLoginChallengeEnvelope{
		Payload:    data,
		ClientType: int32(clientType),
	}
	raw, err := json.Marshal(envelope)
	if err != nil {
		return "", fmt.Errorf("marshal login challenge envelope failed: %w", err)
	}
	key := fmt.Sprintf(mfaLoginChallengeKeyFmt, opId)
	if err := c.rdb.Set(ctx, key, raw, MfaChallengeTTL).Err(); err != nil {
		c.log.Errorf("set login challenge failed: %s", err.Error())
		return "", fmt.Errorf("set login challenge failed")
	}
	return opId, nil
}

// TakeLoginChallenge 取出并删除登录挑战上下文。
func (c *MfaChallengeCache) TakeLoginChallenge(ctx context.Context, opId string) (*MfaLoginChallengeContext, error) {
	key := fmt.Sprintf(mfaLoginChallengeKeyFmt, opId)
	raw, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrMfaChallengeNotFound
		}
		c.log.Errorf("get login challenge failed: %s", err.Error())
		return nil, fmt.Errorf("get login challenge failed")
	}
	// verify-and-delete：取出后立即删除，保证单次有效
	c.rdb.Del(ctx, key)

	var envelope mfaLoginChallengeEnvelope
	if err := json.Unmarshal([]byte(raw), &envelope); err != nil {
		return nil, fmt.Errorf("unmarshal login challenge envelope failed: %w", err)
	}
	payload := &authenticationV1.UserTokenPayload{}
	if err := proto.Unmarshal(envelope.Payload, payload); err != nil {
		return nil, fmt.Errorf("unmarshal login challenge payload failed: %w", err)
	}
	return &MfaLoginChallengeContext{
		Payload:    payload,
		ClientType: authenticationV1.ClientType(envelope.ClientType),
	}, nil
}

// SetEnrollChallenge 写入注册上下文，返回 operation_id。
func (c *MfaChallengeCache) SetEnrollChallenge(ctx context.Context, secret string, tenantID, userID uint32) (string, error) {
	opId := newOperationID()
	envelope := MfaEnrollChallengeContext{
		Secret:   secret,
		TenantID: tenantID,
		UserID:   userID,
	}
	raw, err := json.Marshal(envelope)
	if err != nil {
		return "", fmt.Errorf("marshal enroll challenge envelope failed: %w", err)
	}
	key := fmt.Sprintf(mfaEnrollChallengeKeyFmt, opId)
	if err := c.rdb.Set(ctx, key, raw, MfaChallengeTTL).Err(); err != nil {
		c.log.Errorf("set enroll challenge failed: %s", err.Error())
		return "", fmt.Errorf("set enroll challenge failed")
	}
	return opId, nil
}

// TakeEnrollChallenge 取出并删除注册上下文。
func (c *MfaChallengeCache) TakeEnrollChallenge(ctx context.Context, opId string) (*MfaEnrollChallengeContext, error) {
	key := fmt.Sprintf(mfaEnrollChallengeKeyFmt, opId)
	raw, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrMfaChallengeNotFound
		}
		c.log.Errorf("get enroll challenge failed: %s", err.Error())
		return nil, fmt.Errorf("get enroll challenge failed")
	}
	// verify-and-delete
	c.rdb.Del(ctx, key)

	var envelope MfaEnrollChallengeContext
	if err := json.Unmarshal([]byte(raw), &envelope); err != nil {
		return nil, fmt.Errorf("unmarshal enroll challenge envelope failed: %w", err)
	}
	return &envelope, nil
}

// mfaLoginChallengeEnvelope 是登录挑战上下文的传输封装。
// payload 为 proto 序列化的 UserTokenPayload（含 roles/data_scope 等），clientType 为枚举值。
type mfaLoginChallengeEnvelope struct {
	Payload    []byte
	ClientType int32
}

func newOperationID() string {
	// 复用 jwtutil 的加密随机串生成（与 refresh token 同一来源），保证 operation_id
	// 不可预测且全局唯一。
	id, _ := jwtutil.NewRefreshToken()
	return id
}
