package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	authnEngine "github.com/tx7do/kratos-authn/engine"

	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/jwt"
)

type UserTokenCacheRepo struct {
	log *log.Helper

	rdb *redis.Client // redis客户端

	authenticator authnEngine.Authenticator // 身份验证器

	accessTokenKeyPrefix  string // 访问令牌键前缀
	refreshTokenKeyPrefix string // 刷新令牌键前缀

	accessTokenExpires  time.Duration // 访问令牌过期时间
	refreshTokenExpires time.Duration // 刷新令牌过期时间
}

func NewUserTokenCacheRepo(
	logger log.Logger,
	rdb *redis.Client,
	authenticator authnEngine.Authenticator,
	accessTokenKeyPrefix string,
	refreshTokenKeyPrefix string,
	accessTokenExpires time.Duration,
	refreshTokenExpires time.Duration,
) *UserTokenCacheRepo {
	l := log.NewHelper(log.With(logger, "module", "user-token/cache"))
	return &UserTokenCacheRepo{
		rdb:                   rdb,
		log:                   l,
		authenticator:         authenticator,
		accessTokenKeyPrefix:  accessTokenKeyPrefix,
		refreshTokenKeyPrefix: refreshTokenKeyPrefix,
		accessTokenExpires:    accessTokenExpires,
		refreshTokenExpires:   refreshTokenExpires,
	}
}

// GenerateToken 创建令牌
func (r *UserTokenCacheRepo) GenerateToken(ctx context.Context, user *userV1.User, clientId string) (accessToken string, refreshToken string, err error) {
	// 创建访问令牌
	if accessToken, err = r.GenerateAccessToken(ctx, user, clientId); accessToken == "" {
		err = errors.New("create access token failed")
		return
	}

	// 创建刷新令牌
	if refreshToken, err = r.GenerateRefreshToken(ctx, user); refreshToken == "" {
		err = errors.New("create refresh token failed")
		return
	}

	return
}

// GenerateAccessToken 创建访问令牌
func (r *UserTokenCacheRepo) GenerateAccessToken(ctx context.Context, user *userV1.User, clientId string) (accessToken string, err error) {
	if accessToken = r.createAccessJwtToken(user, clientId); accessToken == "" {
		err = errors.New("create access token failed")
		return
	}

	if err = r.setAccessTokenToRedis(ctx, user.GetId(), accessToken, r.accessTokenExpires); err != nil {
		return
	}

	return
}

// GenerateRefreshToken 创建刷新令牌
func (r *UserTokenCacheRepo) GenerateRefreshToken(ctx context.Context, user *userV1.User) (refreshToken string, err error) {
	if refreshToken = r.createRefreshToken(); refreshToken == "" {
		err = errors.New("create refresh token failed")
		return
	}

	if err = r.setRefreshTokenToRedis(ctx, user.GetId(), refreshToken, r.refreshTokenExpires); err != nil {
		return
	}

	return
}

// RemoveToken 移除所有令牌
func (r *UserTokenCacheRepo) RemoveToken(ctx context.Context, userId uint32) error {
	var err error
	if err = r.deleteAccessTokenFromRedis(ctx, userId); err != nil {
		r.log.Errorf("remove user access token failed: [%v]", err)
	}

	if err = r.deleteRefreshTokenFromRedis(ctx, userId); err != nil {
		r.log.Errorf("remove user refresh token failed: [%v]", err)
	}

	return err
}

// RemoveAccessToken 访问令牌
func (r *UserTokenCacheRepo) RemoveAccessToken(ctx context.Context, userId uint32, accessToken string) error {
	key := r.makeAccessTokenKey(userId)
	return r.delField(ctx, key, accessToken)
}

// RemoveRefreshToken 刷新令牌
func (r *UserTokenCacheRepo) RemoveRefreshToken(ctx context.Context, userId uint32, refreshToken string) error {
	key := r.makeRefreshTokenKey(userId)
	return r.delField(ctx, key, refreshToken)
}

// IsExistAccessToken 访问令牌是否存在
func (r *UserTokenCacheRepo) IsExistAccessToken(ctx context.Context, userId uint32, accessToken string) bool {
	key := r.makeAccessTokenKey(userId)
	return r.exists(ctx, key, accessToken)
}

// IsExistRefreshToken 刷新令牌是否存在
func (r *UserTokenCacheRepo) IsExistRefreshToken(ctx context.Context, userId uint32, refreshToken string) bool {
	key := r.makeRefreshTokenKey(userId)
	return r.exists(ctx, key, refreshToken)
}

// setAccessTokenToRedis 设置访问令牌
func (r *UserTokenCacheRepo) setAccessTokenToRedis(ctx context.Context, userId uint32, token string, expires time.Duration) error {
	key := r.makeAccessTokenKey(userId)
	return r.set(ctx, key, token, expires)
}

func (r *UserTokenCacheRepo) set(ctx context.Context, key string, token string, expires time.Duration) error {
	var err error
	if err = r.rdb.HSet(ctx, key, token, "").Err(); err != nil {
		return err
	}

	if expires > 0 {
		if err = r.rdb.HExpire(ctx, key, expires, token).Err(); err != nil {
			return err
		}
	}

	return nil
}

func (r *UserTokenCacheRepo) exists(ctx context.Context, key string, token string) bool {
	n, err := r.rdb.HExists(ctx, key, token).Result()
	if err != nil {
		return false
	}
	return n
}

func (r *UserTokenCacheRepo) del(ctx context.Context, key string) error {
	return r.rdb.Del(ctx, key).Err()
}

func (r *UserTokenCacheRepo) delField(ctx context.Context, key string, token string) error {
	return r.rdb.HDel(ctx, key, token).Err()
}

// deleteAccessTokenFromRedis 删除访问令牌
func (r *UserTokenCacheRepo) deleteAccessTokenFromRedis(ctx context.Context, userId uint32) error {
	key := r.makeAccessTokenKey(userId)
	return r.del(ctx, key)
}

// setRefreshTokenToRedis 设置刷新令牌
func (r *UserTokenCacheRepo) setRefreshTokenToRedis(ctx context.Context, userId uint32, token string, expires time.Duration) error {
	key := r.makeRefreshTokenKey(userId)
	return r.set(ctx, key, token, expires)
}

// deleteRefreshTokenFromRedis 删除刷新令牌
func (r *UserTokenCacheRepo) deleteRefreshTokenFromRedis(ctx context.Context, userId uint32) error {
	key := r.makeRefreshTokenKey(userId)
	return r.del(ctx, key)
}

// createAccessJwtToken 生成JWT访问令牌
func (r *UserTokenCacheRepo) createAccessJwtToken(user *userV1.User, clientId string) string {
	authClaims := jwt.NewUserTokenAuthClaims(user, clientId)

	signedToken, err := r.authenticator.CreateIdentity(*authClaims)
	if err != nil {
		r.log.Error("create access token failed: [%v]", err)
	}

	return signedToken
}

// createRefreshToken 生成刷新令牌
func (r *UserTokenCacheRepo) createRefreshToken() string {
	strUUID := uuid.New()
	return strUUID.String()
}

// makeAccessTokenKey 生成访问令牌键
func (r *UserTokenCacheRepo) makeAccessTokenKey(userId uint32) string {
	return fmt.Sprintf("%s%d", r.accessTokenKeyPrefix, userId)
}

// makeRefreshTokenKey 生成刷新令牌键
func (r *UserTokenCacheRepo) makeRefreshTokenKey(userId uint32) string {
	return fmt.Sprintf("%s%d", r.refreshTokenKeyPrefix, userId)
}
