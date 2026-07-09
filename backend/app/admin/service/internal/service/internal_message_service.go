package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	paginationV1 "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	"github.com/tx7do/go-utils/aggregator"
	"github.com/tx7do/go-utils/id"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"github.com/tx7do/kratos-transport/transport/sse"
	"google.golang.org/protobuf/types/known/emptypb"

	"go-wind-admin/app/admin/service/internal/data"

	adminV1 "go-wind-admin/api/gen/go/admin/service/v1"
	authenticationV1 "go-wind-admin/api/gen/go/authentication/service/v1"
	internalMessageV1 "go-wind-admin/api/gen/go/internal_message/service/v1"

	"go-wind-admin/pkg/middleware/auth"
)

type InternalMessagePublisher interface {
	Publish(ctx context.Context, streamId sse.StreamID, event *sse.Event)
}

// noopInternalMessagePublisher 是 InternalMessagePublisher 的空操作实现。
// 当 SSE 服务未配置时（NewSseServer 返回 nil，不会调用 RegisterInternalMessagePublisher），
// 用它作为默认值，避免 sendNotification 解引用 nil 接口导致 panic。
// 此时站内信仍会落库，只是不通过 SSE 实时推送。
type noopInternalMessagePublisher struct{}

func (noopInternalMessagePublisher) Publish(_ context.Context, _ sse.StreamID, _ *sse.Event) {}

type InternalMessageService struct {
	adminV1.InternalMessageServiceHTTPServer

	log *log.Helper

	internalMessageRepo          *data.InternalMessageRepo
	internalMessageCategoryRepo  *data.InternalMessageCategoryRepo
	internalMessageRecipientRepo *data.InternalMessageRecipientRepo
	userRepo                     data.UserRepo

	internalMessagePublisher InternalMessagePublisher
	authenticator            *data.Authenticator
	clientType               authenticationV1.ClientType
}

func NewInternalMessageService(
	ctx *bootstrap.Context,
	internalMessageRepo *data.InternalMessageRepo,
	internalMessageCategoryRepo *data.InternalMessageCategoryRepo,
	internalMessageRecipientRepo *data.InternalMessageRecipientRepo,
	userRepo data.UserRepo,
	authenticator *data.Authenticator,
	clientType authenticationV1.ClientType,
) *InternalMessageService {
	return &InternalMessageService{
		log:                          ctx.NewLoggerHelper("internal-message/service/admin-service"),
		internalMessageRepo:          internalMessageRepo,
		internalMessageCategoryRepo:  internalMessageCategoryRepo,
		internalMessageRecipientRepo: internalMessageRecipientRepo,
		userRepo:                     userRepo,
		authenticator:                authenticator,
		clientType:                   clientType,
		// 默认空操作发布者：SSE 未配置时不 panic；配置后由 RegisterInternalMessagePublisher 覆盖
		internalMessagePublisher: noopInternalMessagePublisher{},
	}
}

func (s *InternalMessageService) RegisterInternalMessagePublisher(internalMessagePublisher InternalMessagePublisher) {
	s.internalMessagePublisher = internalMessagePublisher
}

func (s *InternalMessageService) HandleAuthorize(_ *http.Request, token string) error {
	//s.log.Debugf("authorizing token: %s", token)
	//s.log.Debugf("authorizing token HEADER: %s", req.Header.Get("Authorization"))

	resp, err := s.authenticator.Authenticate(context.Background(), &authenticationV1.ValidateTokenRequest{
		ClientType:    s.clientType,
		Token:         token,
		TokenCategory: authenticationV1.TokenCategory_ACCESS,
	})
	if err != nil {
		s.log.Errorf("token authentication failed: %s", err)
		return err
	}

	if resp.GetIsBlocked() {
		s.log.Warnf("token is blocked: %s", token)
		return authenticationV1.ErrorForbidden("token is blocked")
	}
	if !resp.GetIsValid() {
		s.log.Warnf("token is invalid: %s", token)
		return authenticationV1.ErrorUnauthorized("invalid token")
	}

	s.log.Debugf("token authenticated successfully, userId: [%d]", resp.GetPayload().GetUserId())

	return nil
}

func (s *InternalMessageService) HandleSubscribe(streamID sse.StreamID, _ *sse.Subscriber) {
	s.log.Infof("subscriber [%s] connected", streamID)
}

func (s *InternalMessageService) extractRelationIDs(
	messages []*internalMessageV1.InternalMessage,
	categorySet aggregator.ResourceMap[uint32, *internalMessageV1.InternalMessageCategory],
) {
	for _, p := range messages {
		if p.GetCategoryId() > 0 {
			categorySet[p.GetCategoryId()] = nil
		}
	}
}

func (s *InternalMessageService) fetchRelationInfo(
	ctx context.Context,
	categorySet aggregator.ResourceMap[uint32, *internalMessageV1.InternalMessageCategory],
) error {
	if len(categorySet) > 0 {
		categoryIds := make([]uint32, 0, len(categorySet))
		for i := range categorySet {
			categoryIds = append(categoryIds, i)
		}

		categories, err := s.internalMessageCategoryRepo.ListCategoriesByIds(ctx, categoryIds)
		if err != nil {
			s.log.Errorf("query internal message category err: %v", err)
			return err
		}

		for _, g := range categories {
			categorySet[g.GetId()] = g
		}
	}

	return nil
}

func (s *InternalMessageService) bindRelations(
	messages []*internalMessageV1.InternalMessage,
	categorySet aggregator.ResourceMap[uint32, *internalMessageV1.InternalMessageCategory],
) {
	aggregator.Populate(
		messages,
		categorySet,
		func(ou *internalMessageV1.InternalMessage) uint32 { return ou.GetCategoryId() },
		func(ou *internalMessageV1.InternalMessage, c *internalMessageV1.InternalMessageCategory) {
			ou.CategoryName = c.Name
		},
	)
}

func (s *InternalMessageService) enrichRelations(ctx context.Context, messages []*internalMessageV1.InternalMessage) error {
	var categorySet = make(aggregator.ResourceMap[uint32, *internalMessageV1.InternalMessageCategory])
	s.extractRelationIDs(messages, categorySet)
	if err := s.fetchRelationInfo(ctx, categorySet); err != nil {
		return err
	}
	s.bindRelations(messages, categorySet)
	return nil
}

func (s *InternalMessageService) ListMessage(ctx context.Context, req *paginationV1.PagingRequest) (*internalMessageV1.ListInternalMessageResponse, error) {
	resp, err := s.internalMessageRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}

	_ = s.enrichRelations(ctx, resp.Items)

	return resp, nil
}

func (s *InternalMessageService) GetMessage(ctx context.Context, req *internalMessageV1.GetInternalMessageRequest) (*internalMessageV1.InternalMessage, error) {
	resp, err := s.internalMessageRepo.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	fakeItems := []*internalMessageV1.InternalMessage{resp}
	_ = s.enrichRelations(ctx, fakeItems)

	return resp, nil
}

func (s *InternalMessageService) CreateMessage(ctx context.Context, req *internalMessageV1.CreateInternalMessageRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)

	if _, err = s.internalMessageRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *InternalMessageService) UpdateMessage(ctx context.Context, req *internalMessageV1.UpdateInternalMessageRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.Id = trans.Ptr(req.GetId())

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)
	if req.UpdateMask != nil {
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "updated_by")
	}

	if err = s.internalMessageRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *InternalMessageService) DeleteMessage(ctx context.Context, req *internalMessageV1.DeleteInternalMessageRequest) (*emptypb.Empty, error) {
	if err := s.internalMessageRepo.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// RevokeMessage 撤销某条消息
func (s *InternalMessageService) RevokeMessage(ctx context.Context, req *internalMessageV1.RevokeMessageRequest) (*emptypb.Empty, error) {
	var err error
	if err = s.internalMessageRepo.Delete(ctx, req.GetMessageId()); err != nil {
		s.log.Errorf("delete internal message failed: [%d]", req.GetMessageId())
	}

	if err = s.internalMessageRecipientRepo.RevokeMessage(ctx, req); err != nil {
		s.log.Errorf("delete internal message inbox failed: [%d][%d]", req.GetMessageId(), req.GetUserId())
	}

	return &emptypb.Empty{}, err
}

// SendMessage 发送消息
func (s *InternalMessageService) SendMessage(ctx context.Context, req *internalMessageV1.SendMessageRequest) (*internalMessageV1.SendMessageResponse, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	var msg *internalMessageV1.InternalMessage
	if msg, err = s.internalMessageRepo.Create(ctx, &internalMessageV1.CreateInternalMessageRequest{
		Data: &internalMessageV1.InternalMessage{
			Title:      req.Title,
			Content:    trans.Ptr(req.GetContent()),
			Status:     trans.Ptr(internalMessageV1.InternalMessage_PUBLISHED),
			Type:       trans.Ptr(req.GetType()),
			CategoryId: req.CategoryId,
			CreatedBy:  trans.Ptr(operator.GetUserId()),
			CreatedAt:  timeutil.TimeToTimestamppb(&now),
		},
	}); err != nil {
		s.log.Errorf("create internal message failed: %s", err)
		return nil, err
	}

	if req.GetTargetAll() {
		users, err := s.userRepo.List(ctx, &paginationV1.PagingRequest{NoPaging: trans.Ptr(true)})
		if err != nil {
			s.log.Errorf("send message failed, list users failed, %s", err)
		} else {
			for _, user := range users.Items {
				_ = s.sendNotification(ctx, msg.GetId(), user.GetId(), operator.GetUserId(), &now, msg.GetTitle(), msg.GetContent())
			}
		}
	} else {
		if req.RecipientUserId != nil {
			_ = s.sendNotification(ctx, msg.GetId(), req.GetRecipientUserId(), operator.GetUserId(), &now, msg.GetTitle(), msg.GetContent())
		} else {
			if len(req.TargetUserIds) != 0 {
				for _, uid := range req.TargetUserIds {
					_ = s.sendNotification(ctx, msg.GetId(), uid, operator.GetUserId(), &now, msg.GetTitle(), msg.GetContent())
				}
			}
		}
	}

	return &internalMessageV1.SendMessageResponse{
		MessageId: msg.GetId(),
	}, nil
}

// sendNotification 向客户端发送通知消息
func (s *InternalMessageService) sendNotification(ctx context.Context, messageId uint32, recipientUserId uint32, senderUserId uint32, now *time.Time, title, content string) error {
	recipient := &internalMessageV1.InternalMessageRecipient{
		MessageId:       trans.Ptr(messageId),
		RecipientUserId: trans.Ptr(recipientUserId),
		Status:          trans.Ptr(internalMessageV1.InternalMessageRecipient_SENT),
		CreatedBy:       trans.Ptr(senderUserId),
		CreatedAt:       timeutil.TimeToTimestamppb(now),
		Title:           trans.Ptr(title),
		Content:         trans.Ptr(content),
	}

	var err error
	var entity *internalMessageV1.InternalMessageRecipient
	if entity, err = s.internalMessageRecipientRepo.Create(ctx, recipient); err != nil {
		s.log.Errorf("send message failed, send to user failed, %s", err)
		return err
	}
	recipient.Id = entity.Id

	recipientJson, _ := json.Marshal(recipient)

	recipientStreamIds := s.authenticator.GetAccessTokens(ctx, s.clientType, recipientUserId)
	for _, streamId := range recipientStreamIds {
		s.internalMessagePublisher.Publish(ctx, sse.StreamID(streamId), &sse.Event{
			ID:    []byte(id.NewGUIDv4(false)),
			Data:  recipientJson,
			Event: []byte("notification"),
		})
	}

	return nil
}
