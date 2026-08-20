package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image/png"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go-wind-admin/app/admin/service/internal/data"
	"go-wind-admin/app/admin/service/internal/data/ent/privacy"
	"go-wind-admin/pkg/middleware/auth"

	"github.com/tx7do/go-crud/viewer"

	"github.com/pquerna/otp"
	otpTotp "github.com/pquerna/otp/totp"

	authenticationV1 "go-wind-admin/api/gen/go/authentication/service/v1"
)

const (
	// mfaTotpIssuer TOTP otpauth URI 中的发行方标识（认证器 App 内展示归属）。
	mfaTotpIssuer = "GoWindAdmin"
	// mfaTotpSkew 允许的时间窗口偏移（±1 个 30s 周期），防客户端/服务端时钟小幅漂移。
	mfaTotpSkew = 1
)

// MfaService 实现 adminV1.MfaServiceHTTPServer。
//
// 本轮仅落地 TOTP：
//   - 管理面（GetMFAStatus/ListEnrolledMethods/StartEnrollMethod/ConfirmEnrollMethod/DisableMFA/RevokeMFADevice）
//     需登录态，operator 从 auth.FromContext(ctx) 取得，强制只能操作本人因子。
//   - 登录挑战面（VerifyMFAChallenge）免鉴权，operation_id 由登录流程签发并存入
//     MfaChallengeCache（含 UserTokenPayload + ClientType），验证通过后用 authenticator 签发真 token。
//
// 非 TOTP 方法、StartMFAChallenge、备份码相关 RPC 本轮返回 UNIMPLEMENTED。
type MfaService struct {
	log *log.Helper

	mfaFactorRepo     *data.UserMfaFactorRepo
	mfaChallengeCache *data.MfaChallengeCache
	authenticator     *data.Authenticator
}

func NewMfaService(
	ctx *bootstrap.Context,
	mfaFactorRepo *data.UserMfaFactorRepo,
	mfaChallengeCache *data.MfaChallengeCache,
	authenticator *data.Authenticator,
) *MfaService {
	return &MfaService{
		log:               ctx.NewLoggerHelper("mfa/service/admin-service"),
		mfaFactorRepo:     mfaFactorRepo,
		mfaChallengeCache: mfaChallengeCache,
		authenticator:     authenticator,
	}
}

// GetMFAStatus 查询当前登录用户 MFA 总览。
func (s *MfaService) GetMFAStatus(ctx context.Context, _ *authenticationV1.GetMFAStatusRequest) (*authenticationV1.GetMFAStatusResponse, error) {
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	uid := operator.GetUserId()
	tid := operator.GetTenantId()
	infos, err := s.mfaFactorRepo.ListByUser(ctx, tid, uid)
	if err != nil {
		return nil, authenticationV1.ErrorInternalServerError("query mfa status failed")
	}
	hasTotp := false
	for _, i := range infos {
		if i.Method == authenticationV1.MFAMethod_TOTP && i.Enabled {
			hasTotp = true
			break
		}
	}
	resp := &authenticationV1.GetMFAStatusResponse{
		Enabled:     hasTotp,
		Enforcement: authenticationV1.MFAEnforcement_MFA_NOT_REQUIRED,
	}
	if hasTotp {
		resp.Enforcement = authenticationV1.MFAEnforcement_MFA_REQUIRED
		resp.Enrolled = buildEnrolledProto(infos)
	}
	return resp, nil
}

// ListEnrolledMethods 列出当前登录用户的 MFA 因子（不含 secret）。
func (s *MfaService) ListEnrolledMethods(ctx context.Context, _ *authenticationV1.ListEnrolledMethodsRequest) (*authenticationV1.ListEnrolledMethodsResponse, error) {
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	infos, err := s.mfaFactorRepo.ListByUser(ctx, operator.GetTenantId(), operator.GetUserId())
	if err != nil {
		return nil, authenticationV1.ErrorInternalServerError("list mfa methods failed")
	}
	return &authenticationV1.ListEnrolledMethodsResponse{
		Items: buildEnrolledProto(infos),
	}, nil
}

// StartEnrollMethod 开始注册 MFA 方法。仅 TOTP 本轮实现。
func (s *MfaService) StartEnrollMethod(ctx context.Context, req *authenticationV1.StartEnrollMethodRequest) (*authenticationV1.StartEnrollMethodResponse, error) {
	if req.GetMethod() != authenticationV1.MFAMethod_TOTP {
		return nil, authenticationV1.ErrorBadRequest("only TOTP is supported")
	}
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	// 预检：已绑定 TOTP 时拒绝重复注册（(tenant,user,method) 唯一约束兜底，
	// 但提前返回友好错误，避免 Confirm 阶段撞唯一索引报笼统 500）。
	// 重新绑定需先 DisableMFA 解绑。
	if has, herr := s.mfaFactorRepo.HasEnabledTotp(ctx, operator.GetTenantId(), operator.GetUserId()); herr != nil {
		return nil, authenticationV1.ErrorInternalServerError("check mfa status failed")
	} else if has {
		return nil, authenticationV1.ErrorBadRequest("totp already enrolled, disable it first")
	}

	// 生成 TOTP key：issuer 标识发行方，account 用 userId 防 App 内同名冲突。
	key, err := otpTotp.Generate(otpTotp.GenerateOpts{
		Issuer:      mfaTotpIssuer,
		AccountName: fmt.Sprintf("uid:%d", operator.GetUserId()),
	})
	if err != nil {
		s.log.Errorf("generate totp key failed: %s", err.Error())
		return nil, authenticationV1.ErrorInternalServerError("generate totp key failed")
	}

	opId, err := s.mfaChallengeCache.SetEnrollChallenge(ctx, key.Secret(), operator.GetTenantId(), operator.GetUserId())
	if err != nil {
		s.log.Errorf("set enroll challenge failed: %s", err.Error())
		return nil, authenticationV1.ErrorInternalServerError("start enroll failed")
	}

	qrUri, err := totpQrDataUri(key)
	if err != nil {
		s.log.Errorf("encode totp qr failed: %s", err.Error())
		return nil, authenticationV1.ErrorInternalServerError("encode totp qr failed")
	}

	return &authenticationV1.StartEnrollMethodResponse{
		Result: &authenticationV1.StartEnrollMethodResponse_Totp{
			Totp: &authenticationV1.TOTPResult{
				Secret:        key.Secret(),
				OtpAuthUrl:    key.URL(),
				QrCodeDataUri: qrUri,
			},
		},
		OperationId: opId,
		ExpiresAt:   timestamppb.New(time.Now().Add(data.MfaChallengeTTL)),
	}, nil
}

// ConfirmEnrollMethod 确认注册：校验首码通过则落库因子。
func (s *MfaService) ConfirmEnrollMethod(ctx context.Context, req *authenticationV1.ConfirmEnrollMethodRequest) (*authenticationV1.ConfirmEnrollMethodResponse, error) {
	if req.GetMethod() != authenticationV1.MFAMethod_TOTP {
		return nil, authenticationV1.ErrorBadRequest("only TOTP is supported")
	}
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	enrollCtx, err := s.mfaChallengeCache.TakeEnrollChallenge(ctx, req.GetOperationId())
	if err != nil {
		return nil, authenticationV1.ErrorBadRequest("invalid or expired enroll operation")
	}
	// 防 operation_id 跨用户劫持：注册上下文绑定的人必须等于当前 operator
	if enrollCtx.TenantID != operator.GetTenantId() || enrollCtx.UserID != operator.GetUserId() {
		return nil, authenticationV1.ErrorForbidden("enroll operation user mismatch")
	}

	// 校验首码：用默认 Validate（±1 窗口，等同 ValidateCustom 的默认参数语义）
	if !otpTotp.Validate(req.GetTotpCode(), enrollCtx.Secret) {
		return &authenticationV1.ConfirmEnrollMethodResponse{Success: false}, nil
	}

	factorId, err := s.mfaFactorRepo.CreateTotpFactor(ctx, enrollCtx.TenantID, enrollCtx.UserID, enrollCtx.Secret, req.GetDisplay())
	if err != nil {
		return nil, authenticationV1.ErrorInternalServerError("create mfa factor failed")
	}
	return &authenticationV1.ConfirmEnrollMethodResponse{
		Success:      true,
		CredentialId: fmt.Sprintf("%d", factorId),
	}, nil
}

// VerifyMFAChallenge 验证登录 MFA 挑战。通过则签发真 token 并返回 LoginResponse。
// 免鉴权：operation_id 由登录流程在密码校验通过、待二次验证阶段签发。
func (s *MfaService) VerifyMFAChallenge(ctx context.Context, req *authenticationV1.VerifyMFAChallengeRequest) (*authenticationV1.LoginResponse, error) {
	// 本接口在鉴权白名单内、无 auth 中间件注入 viewer，而 ent 隐私层
	// （TenantPrivacy mixin）要求 ViewerContext 存在——与登录流程的
	// resetContextForLogin 同款处理：Noop viewer + privacy.Allow。
	// 挑战上下文绑定的 userId/tenantId 来自密码校验通过后的 tokenPayload，
	// 越权面由 operation_id 单次有效 + 归属校验兜住。
	ctx = viewer.WithContext(ctx, viewer.NewNoopContext())
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	challengeCtx, err := s.mfaChallengeCache.TakeLoginChallenge(ctx, req.GetOperationId())
	if err != nil {
		return nil, authenticationV1.ErrorBadRequest("invalid or expired mfa operation")
	}

	payload := challengeCtx.Payload
	if payload == nil {
		return nil, authenticationV1.ErrorInternalServerError("mfa challenge payload missing")
	}
	uid := payload.GetUserId()
	tid := payload.GetTenantId()

	// 取该用户 ENABLED 的 TOTP 因子并解密 secret
	factorId, plainSecret, err := s.mfaFactorRepo.FindEnabledTotpForUser(ctx, tid, uid)
	if err != nil {
		s.log.Errorf("find totp factor for login mfa failed uid=%d: %s", uid, err.Error())
		return nil, authenticationV1.ErrorForbidden("mfa verification failed")
	}

	// 校验 TOTP 码：±1 窗口（防时钟漂移），默认周期 30s、6 位、SHA1。
	// digits/algorithm 用常量：totp.Generate 产出的 key 恒为 6 位 SHA1，这里与之对齐。
	ok, verr := otpTotp.ValidateCustom(req.GetTotpCode(), plainSecret, time.Now(), otpTotp.ValidateOpts{
		Period:    30,
		Skew:      mfaTotpSkew,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	if verr != nil || !ok {
		return nil, authenticationV1.ErrorForbidden("invalid mfa code")
	}

	// 通过：更新 last_used_at（best-effort），签发真 token
	_ = s.mfaFactorRepo.UpdateLastUsed(ctx, tid, uid, factorId, time.Now())

	accessToken, refreshToken, err := s.authenticator.CreateUserToken(ctx, challengeCtx.ClientType, payload)
	if err != nil {
		return nil, err
	}

	return &authenticationV1.LoginResponse{
		TokenType:        authenticationV1.TokenType_bearer,
		AccessToken:      accessToken,
		RefreshToken:     trans.Ptr(refreshToken),
		ExpiresIn:        int64(s.authenticator.GetAccessTokenExpires(challengeCtx.ClientType).Seconds()),
		RefreshExpiresIn: trans.Ptr(int64(s.authenticator.GetRefreshTokenExpires(challengeCtx.ClientType).Seconds())),
	}, nil
}

// DisableMFA 禁用/移除当前登录用户的 MFA 凭证。
func (s *MfaService) DisableMFA(ctx context.Context, req *authenticationV1.DisableMFARequest) (*emptypb.Empty, error) {
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	factorId, err := parseFactorId(req.GetCredentialId())
	if err != nil {
		return nil, authenticationV1.ErrorBadRequest("invalid credential id")
	}
	ok, err := s.mfaFactorRepo.DeleteForUser(ctx, operator.GetTenantId(), operator.GetUserId(), factorId)
	if err != nil {
		return nil, authenticationV1.ErrorInternalServerError("disable mfa failed")
	}
	if !ok {
		return nil, authenticationV1.ErrorNotFound("mfa credential not found")
	}
	return &emptypb.Empty{}, nil
}

// RevokeMFADevice 撤销指定 MFA 凭证（按 id），强制归属校验。
func (s *MfaService) RevokeMFADevice(ctx context.Context, req *authenticationV1.RevokeMFADeviceRequest) (*emptypb.Empty, error) {
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	factorId, err := parseFactorId(req.GetCredentialId())
	if err != nil {
		return nil, authenticationV1.ErrorBadRequest("invalid credential id")
	}
	ok, err := s.mfaFactorRepo.DeleteForUser(ctx, operator.GetTenantId(), operator.GetUserId(), factorId)
	if err != nil {
		return nil, authenticationV1.ErrorInternalServerError("revoke mfa device failed")
	}
	if !ok {
		return nil, authenticationV1.ErrorNotFound("mfa credential not found")
	}
	return &emptypb.Empty{}, nil
}

// buildEnrolledProto 将仓储返回的因子元信息列表转为 proto EnrolledMethod 列表。
// 不含 secret。
func buildEnrolledProto(infos []data.EnrolledFactorInfo) []*authenticationV1.EnrolledMethod {
	out := make([]*authenticationV1.EnrolledMethod, 0, len(infos))
	for _, i := range infos {
		em := &authenticationV1.EnrolledMethod{
			Id:         fmt.Sprintf("%d", i.ID),
			Method:     i.Method,
			Display:    i.DisplayName,
			Enabled:    i.Enabled,
			CreatedAt:  nil,
			LastUsedAt: nil,
		}
		if i.CreatedAt != nil {
			em.CreatedAt = timestamppb.New(*i.CreatedAt)
		}
		if i.LastUsedAt != nil {
			em.LastUsedAt = timestamppb.New(*i.LastUsedAt)
		}
		out = append(out, em)
	}
	return out
}

// parseFactorId 将 proto 返回的字符串形式 credential_id 解析为 uint32 主键。
func parseFactorId(s string) (uint32, error) {
	v, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}

// totpQrDataUri 将 otp.Key 的二维码图渲染为 PNG data URI，供前端 <img> 展示。
// 生成器（totp.Generate）产生的 key 是标准 otpauth URI，编码为二维码后可被
// Google Authenticator 等认证器 App 扫码导入。
func totpQrDataUri(key *otp.Key) (string, error) {
	if key == nil {
		return "", fmt.Errorf("nil totp key")
	}
	img, err := key.Image(240, 240)
	if err != nil {
		return "", fmt.Errorf("render qr failed: %w", err)
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return "", fmt.Errorf("encode qr png failed: %w", err)
	}
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
