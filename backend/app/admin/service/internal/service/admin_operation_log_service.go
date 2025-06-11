package service

import (
	"context"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/types/known/emptypb"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

type AdminOperationLogService struct {
	adminV1.AdminOperationLogServiceHTTPServer

	log *log.Helper

	operationLogRepo *data.AdminOperationLogRepo
	apiResourceRepo  *data.ApiResourceRepo

	apis     []*adminV1.ApiResource
	apiMutex sync.RWMutex
}

func NewAdminOperationLogService(
	logger log.Logger,
	operationLogRepo *data.AdminOperationLogRepo,
	apiResourceRepo *data.ApiResourceRepo,
) *AdminOperationLogService {
	l := log.NewHelper(log.With(logger, "module", "admin-operation-log/service/admin-service"))
	return &AdminOperationLogService{
		log:              l,
		operationLogRepo: operationLogRepo,
		apiResourceRepo:  apiResourceRepo,
	}
}

func (s *AdminOperationLogService) queryApiResources(ctx context.Context, path, method string) (*adminV1.ApiResource, error) {
	if len(s.apis) == 0 {
		s.apiMutex.Lock()
		apis, err := s.apiResourceRepo.List(ctx, &pagination.PagingRequest{
			NoPaging: trans.Ptr(true),
		})
		if err != nil {
			s.apiMutex.Unlock()
			return nil, err
		}
		s.apis = apis.Items
		s.apiMutex.Unlock()
	}

	if len(s.apis) == 0 {
		return nil, adminV1.ErrorNotFound("no API resources found")
	}

	for _, api := range s.apis {
		if api.GetPath() == path && api.GetMethod() == method {
			return api, nil
		}
	}

	return nil, nil
}

func (s *AdminOperationLogService) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListAdminOperationLogResponse, error) {
	resp, err := s.operationLogRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(resp.Items); i++ {
		l := resp.Items[i]
		if l == nil {
			continue
		}
		a, _ := s.queryApiResources(ctx, l.GetPath(), l.GetMethod())
		if a != nil {
			l.ApiDescription = a.Description
			l.ApiModule = a.ModuleDescription
		}
	}

	return resp, nil
}

func (s *AdminOperationLogService) Get(ctx context.Context, req *adminV1.GetAdminOperationLogRequest) (*adminV1.AdminOperationLog, error) {
	resp, err := s.operationLogRepo.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	a, _ := s.queryApiResources(ctx, resp.GetPath(), resp.GetMethod())
	if a != nil {
		resp.ApiDescription = a.Description
		resp.ApiModule = a.ModuleDescription
	}

	return resp, nil
}

func (s *AdminOperationLogService) Create(ctx context.Context, req *adminV1.CreateAdminOperationLogRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	if err := s.operationLogRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
