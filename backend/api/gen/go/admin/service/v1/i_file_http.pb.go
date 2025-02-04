// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.3
// - protoc             (unknown)
// source: admin/service/v1/i_file.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "kratos-admin/api/gen/go/file/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationFileServiceOssUploadUrl = "/admin.service.v1.FileService/OssUploadUrl"

type FileServiceHTTPServer interface {
	// OssUploadUrl 获取对象存储（OSS）上传用的预签名链接
	OssUploadUrl(context.Context, *v1.OssUploadUrlRequest) (*v1.OssUploadUrlResponse, error)
}

func RegisterFileServiceHTTPServer(s *http.Server, srv FileServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/admin/v1/file:upload-url", _FileService_OssUploadUrl0_HTTP_Handler(srv))
}

func _FileService_OssUploadUrl0_HTTP_Handler(srv FileServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.OssUploadUrlRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFileServiceOssUploadUrl)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OssUploadUrl(ctx, req.(*v1.OssUploadUrlRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.OssUploadUrlResponse)
		return ctx.Result(200, reply)
	}
}

type FileServiceHTTPClient interface {
	OssUploadUrl(ctx context.Context, req *v1.OssUploadUrlRequest, opts ...http.CallOption) (rsp *v1.OssUploadUrlResponse, err error)
}

type FileServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewFileServiceHTTPClient(client *http.Client) FileServiceHTTPClient {
	return &FileServiceHTTPClientImpl{client}
}

func (c *FileServiceHTTPClientImpl) OssUploadUrl(ctx context.Context, in *v1.OssUploadUrlRequest, opts ...http.CallOption) (*v1.OssUploadUrlResponse, error) {
	var out v1.OssUploadUrlResponse
	pattern := "/admin/v1/file:upload-url"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFileServiceOssUploadUrl))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
