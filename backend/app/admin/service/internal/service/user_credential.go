package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	authenticationV1 "kratos-admin/api/gen/go/authentication/service/v1"
)

type UserCredentialService struct {
	authenticationV1.UnimplementedUserCredentialServiceServer

	log *log.Helper

	userCredentialsRepo *data.UserCredentialRepo
}

func NewUserCredentialService(
	logger log.Logger,
	userCredentialRepo *data.UserCredentialRepo,
) *UserCredentialService {
	l := log.NewHelper(log.With(logger, "module", "user-credential/service/admin-service"))
	return &UserCredentialService{
		log:                 l,
		userCredentialsRepo: userCredentialRepo,
	}
}

func (s *UserCredentialService) List(ctx context.Context, req *pagination.PagingRequest) (*authenticationV1.ListUserCredentialResponse, error) {
	return s.userCredentialsRepo.List(ctx, req)
}

func (s *UserCredentialService) Get(ctx context.Context, req *authenticationV1.GetUserCredentialRequest) (*authenticationV1.UserCredential, error) {
	return s.userCredentialsRepo.Get(ctx, req)
}

func (s *UserCredentialService) GetByIdentifier(ctx context.Context, req *authenticationV1.GetUserCredentialByIdentifierRequest) (*authenticationV1.UserCredential, error) {
	return s.userCredentialsRepo.GetByIdentifier(ctx, req)
}

func (s *UserCredentialService) Create(ctx context.Context, req *authenticationV1.CreateUserCredentialRequest) (*emptypb.Empty, error) {
	if err := s.userCredentialsRepo.Create(ctx, req); err != nil {
		// s.log.Info(err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *UserCredentialService) Update(ctx context.Context, req *authenticationV1.UpdateUserCredentialRequest) (*emptypb.Empty, error) {
	if err := s.userCredentialsRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *UserCredentialService) Delete(ctx context.Context, req *authenticationV1.DeleteUserCredentialRequest) (*emptypb.Empty, error) {
	if err := s.userCredentialsRepo.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *UserCredentialService) VerifyCredential(ctx context.Context, req *authenticationV1.VerifyCredentialRequest) (*authenticationV1.VerifyCredentialResponse, error) {
	return s.userCredentialsRepo.VerifyCredential(ctx, req)
}

func (s *UserCredentialService) ChangeCredential(ctx context.Context, req *authenticationV1.ChangeCredentialRequest) (*emptypb.Empty, error) {
	err := s.userCredentialsRepo.ChangeCredential(ctx, req)
	return &emptypb.Empty{}, err
}

func (s *UserCredentialService) ResetCredential(_ context.Context, _ *authenticationV1.ResetCredentialRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
