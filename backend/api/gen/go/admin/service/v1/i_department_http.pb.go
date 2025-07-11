// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             (unknown)
// source: admin/service/v1/i_department.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-admin/api/gen/go/user/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationDepartmentServiceCreate = "/admin.service.v1.DepartmentService/Create"
const OperationDepartmentServiceDelete = "/admin.service.v1.DepartmentService/Delete"
const OperationDepartmentServiceGet = "/admin.service.v1.DepartmentService/Get"
const OperationDepartmentServiceList = "/admin.service.v1.DepartmentService/List"
const OperationDepartmentServiceUpdate = "/admin.service.v1.DepartmentService/Update"

type DepartmentServiceHTTPServer interface {
	// Create 创建部门
	Create(context.Context, *v11.CreateDepartmentRequest) (*emptypb.Empty, error)
	// Delete 删除部门
	Delete(context.Context, *v11.DeleteDepartmentRequest) (*emptypb.Empty, error)
	// Get 查询部门详情
	Get(context.Context, *v11.GetDepartmentRequest) (*v11.Department, error)
	// List 查询部门列表
	List(context.Context, *v1.PagingRequest) (*v11.ListDepartmentResponse, error)
	// Update 更新部门
	Update(context.Context, *v11.UpdateDepartmentRequest) (*emptypb.Empty, error)
}

func RegisterDepartmentServiceHTTPServer(s *http.Server, srv DepartmentServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/departments", _DepartmentService_List4_HTTP_Handler(srv))
	r.GET("/admin/v1/departments/{id}", _DepartmentService_Get4_HTTP_Handler(srv))
	r.POST("/admin/v1/departments", _DepartmentService_Create2_HTTP_Handler(srv))
	r.PUT("/admin/v1/departments/{data.id}", _DepartmentService_Update2_HTTP_Handler(srv))
	r.DELETE("/admin/v1/departments/{id}", _DepartmentService_Delete2_HTTP_Handler(srv))
}

func _DepartmentService_List4_HTTP_Handler(srv DepartmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDepartmentServiceList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.ListDepartmentResponse)
		return ctx.Result(200, reply)
	}
}

func _DepartmentService_Get4_HTTP_Handler(srv DepartmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.GetDepartmentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDepartmentServiceGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*v11.GetDepartmentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Department)
		return ctx.Result(200, reply)
	}
}

func _DepartmentService_Create2_HTTP_Handler(srv DepartmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.CreateDepartmentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDepartmentServiceCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*v11.CreateDepartmentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _DepartmentService_Update2_HTTP_Handler(srv DepartmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.UpdateDepartmentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDepartmentServiceUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*v11.UpdateDepartmentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _DepartmentService_Delete2_HTTP_Handler(srv DepartmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.DeleteDepartmentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDepartmentServiceDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*v11.DeleteDepartmentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type DepartmentServiceHTTPClient interface {
	Create(ctx context.Context, req *v11.CreateDepartmentRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Delete(ctx context.Context, req *v11.DeleteDepartmentRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Get(ctx context.Context, req *v11.GetDepartmentRequest, opts ...http.CallOption) (rsp *v11.Department, err error)
	List(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *v11.ListDepartmentResponse, err error)
	Update(ctx context.Context, req *v11.UpdateDepartmentRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type DepartmentServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewDepartmentServiceHTTPClient(client *http.Client) DepartmentServiceHTTPClient {
	return &DepartmentServiceHTTPClientImpl{client}
}

func (c *DepartmentServiceHTTPClientImpl) Create(ctx context.Context, in *v11.CreateDepartmentRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/departments"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDepartmentServiceCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DepartmentServiceHTTPClientImpl) Delete(ctx context.Context, in *v11.DeleteDepartmentRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/departments/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDepartmentServiceDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DepartmentServiceHTTPClientImpl) Get(ctx context.Context, in *v11.GetDepartmentRequest, opts ...http.CallOption) (*v11.Department, error) {
	var out v11.Department
	pattern := "/admin/v1/departments/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDepartmentServiceGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DepartmentServiceHTTPClientImpl) List(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*v11.ListDepartmentResponse, error) {
	var out v11.ListDepartmentResponse
	pattern := "/admin/v1/departments"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDepartmentServiceList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DepartmentServiceHTTPClientImpl) Update(ctx context.Context, in *v11.UpdateDepartmentRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/departments/{data.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDepartmentServiceUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
