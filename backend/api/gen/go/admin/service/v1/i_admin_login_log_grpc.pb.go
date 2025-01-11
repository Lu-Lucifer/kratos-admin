// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: admin/service/v1/i_admin_login_log.proto

package servicev1

import (
	context "context"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-admin/api/gen/go/system/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AdminLoginLogService_ListAdminLoginLog_FullMethodName   = "/admin.service.v1.AdminLoginLogService/ListAdminLoginLog"
	AdminLoginLogService_GetAdminLoginLog_FullMethodName    = "/admin.service.v1.AdminLoginLogService/GetAdminLoginLog"
	AdminLoginLogService_CreateAdminLoginLog_FullMethodName = "/admin.service.v1.AdminLoginLogService/CreateAdminLoginLog"
	AdminLoginLogService_UpdateAdminLoginLog_FullMethodName = "/admin.service.v1.AdminLoginLogService/UpdateAdminLoginLog"
	AdminLoginLogService_DeleteAdminLoginLog_FullMethodName = "/admin.service.v1.AdminLoginLogService/DeleteAdminLoginLog"
)

// AdminLoginLogServiceClient is the client API for AdminLoginLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 后台登录日志管理服务
type AdminLoginLogServiceClient interface {
	// 查询后台登录日志列表
	ListAdminLoginLog(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListAdminLoginLogResponse, error)
	// 查询后台登录日志详情
	GetAdminLoginLog(ctx context.Context, in *v11.GetAdminLoginLogRequest, opts ...grpc.CallOption) (*v11.AdminLoginLog, error)
	// 创建后台登录日志
	CreateAdminLoginLog(ctx context.Context, in *v11.CreateAdminLoginLogRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新后台登录日志
	UpdateAdminLoginLog(ctx context.Context, in *v11.UpdateAdminLoginLogRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除后台登录日志
	DeleteAdminLoginLog(ctx context.Context, in *v11.DeleteAdminLoginLogRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type adminLoginLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminLoginLogServiceClient(cc grpc.ClientConnInterface) AdminLoginLogServiceClient {
	return &adminLoginLogServiceClient{cc}
}

func (c *adminLoginLogServiceClient) ListAdminLoginLog(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListAdminLoginLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.ListAdminLoginLogResponse)
	err := c.cc.Invoke(ctx, AdminLoginLogService_ListAdminLoginLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminLoginLogServiceClient) GetAdminLoginLog(ctx context.Context, in *v11.GetAdminLoginLogRequest, opts ...grpc.CallOption) (*v11.AdminLoginLog, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.AdminLoginLog)
	err := c.cc.Invoke(ctx, AdminLoginLogService_GetAdminLoginLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminLoginLogServiceClient) CreateAdminLoginLog(ctx context.Context, in *v11.CreateAdminLoginLogRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AdminLoginLogService_CreateAdminLoginLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminLoginLogServiceClient) UpdateAdminLoginLog(ctx context.Context, in *v11.UpdateAdminLoginLogRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AdminLoginLogService_UpdateAdminLoginLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminLoginLogServiceClient) DeleteAdminLoginLog(ctx context.Context, in *v11.DeleteAdminLoginLogRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AdminLoginLogService_DeleteAdminLoginLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminLoginLogServiceServer is the server API for AdminLoginLogService service.
// All implementations must embed UnimplementedAdminLoginLogServiceServer
// for forward compatibility.
//
// 后台登录日志管理服务
type AdminLoginLogServiceServer interface {
	// 查询后台登录日志列表
	ListAdminLoginLog(context.Context, *v1.PagingRequest) (*v11.ListAdminLoginLogResponse, error)
	// 查询后台登录日志详情
	GetAdminLoginLog(context.Context, *v11.GetAdminLoginLogRequest) (*v11.AdminLoginLog, error)
	// 创建后台登录日志
	CreateAdminLoginLog(context.Context, *v11.CreateAdminLoginLogRequest) (*emptypb.Empty, error)
	// 更新后台登录日志
	UpdateAdminLoginLog(context.Context, *v11.UpdateAdminLoginLogRequest) (*emptypb.Empty, error)
	// 删除后台登录日志
	DeleteAdminLoginLog(context.Context, *v11.DeleteAdminLoginLogRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAdminLoginLogServiceServer()
}

// UnimplementedAdminLoginLogServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdminLoginLogServiceServer struct{}

func (UnimplementedAdminLoginLogServiceServer) ListAdminLoginLog(context.Context, *v1.PagingRequest) (*v11.ListAdminLoginLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAdminLoginLog not implemented")
}
func (UnimplementedAdminLoginLogServiceServer) GetAdminLoginLog(context.Context, *v11.GetAdminLoginLogRequest) (*v11.AdminLoginLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminLoginLog not implemented")
}
func (UnimplementedAdminLoginLogServiceServer) CreateAdminLoginLog(context.Context, *v11.CreateAdminLoginLogRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdminLoginLog not implemented")
}
func (UnimplementedAdminLoginLogServiceServer) UpdateAdminLoginLog(context.Context, *v11.UpdateAdminLoginLogRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdminLoginLog not implemented")
}
func (UnimplementedAdminLoginLogServiceServer) DeleteAdminLoginLog(context.Context, *v11.DeleteAdminLoginLogRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAdminLoginLog not implemented")
}
func (UnimplementedAdminLoginLogServiceServer) mustEmbedUnimplementedAdminLoginLogServiceServer() {}
func (UnimplementedAdminLoginLogServiceServer) testEmbeddedByValue()                              {}

// UnsafeAdminLoginLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminLoginLogServiceServer will
// result in compilation errors.
type UnsafeAdminLoginLogServiceServer interface {
	mustEmbedUnimplementedAdminLoginLogServiceServer()
}

func RegisterAdminLoginLogServiceServer(s grpc.ServiceRegistrar, srv AdminLoginLogServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdminLoginLogServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdminLoginLogService_ServiceDesc, srv)
}

func _AdminLoginLogService_ListAdminLoginLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminLoginLogServiceServer).ListAdminLoginLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminLoginLogService_ListAdminLoginLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminLoginLogServiceServer).ListAdminLoginLog(ctx, req.(*v1.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminLoginLogService_GetAdminLoginLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.GetAdminLoginLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminLoginLogServiceServer).GetAdminLoginLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminLoginLogService_GetAdminLoginLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminLoginLogServiceServer).GetAdminLoginLog(ctx, req.(*v11.GetAdminLoginLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminLoginLogService_CreateAdminLoginLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.CreateAdminLoginLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminLoginLogServiceServer).CreateAdminLoginLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminLoginLogService_CreateAdminLoginLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminLoginLogServiceServer).CreateAdminLoginLog(ctx, req.(*v11.CreateAdminLoginLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminLoginLogService_UpdateAdminLoginLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.UpdateAdminLoginLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminLoginLogServiceServer).UpdateAdminLoginLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminLoginLogService_UpdateAdminLoginLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminLoginLogServiceServer).UpdateAdminLoginLog(ctx, req.(*v11.UpdateAdminLoginLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminLoginLogService_DeleteAdminLoginLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.DeleteAdminLoginLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminLoginLogServiceServer).DeleteAdminLoginLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminLoginLogService_DeleteAdminLoginLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminLoginLogServiceServer).DeleteAdminLoginLog(ctx, req.(*v11.DeleteAdminLoginLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminLoginLogService_ServiceDesc is the grpc.ServiceDesc for AdminLoginLogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminLoginLogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.service.v1.AdminLoginLogService",
	HandlerType: (*AdminLoginLogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAdminLoginLog",
			Handler:    _AdminLoginLogService_ListAdminLoginLog_Handler,
		},
		{
			MethodName: "GetAdminLoginLog",
			Handler:    _AdminLoginLogService_GetAdminLoginLog_Handler,
		},
		{
			MethodName: "CreateAdminLoginLog",
			Handler:    _AdminLoginLogService_CreateAdminLoginLog_Handler,
		},
		{
			MethodName: "UpdateAdminLoginLog",
			Handler:    _AdminLoginLogService_UpdateAdminLoginLog_Handler,
		},
		{
			MethodName: "DeleteAdminLoginLog",
			Handler:    _AdminLoginLogService_DeleteAdminLoginLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/service/v1/i_admin_login_log.proto",
}
