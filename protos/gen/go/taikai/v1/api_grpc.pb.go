// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: taikai/v1/api.proto

package taikaiv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Api_UpsertHellos_FullMethodName = "/taikai.v1.Api/UpsertHellos"
	Api_DeleteHellos_FullMethodName = "/taikai.v1.Api/DeleteHellos"
	Api_ListHellos_FullMethodName   = "/taikai.v1.Api/ListHellos"
	Api_GetHellos_FullMethodName    = "/taikai.v1.Api/GetHellos"
	Api_Healthy_FullMethodName      = "/taikai.v1.Api/Healthy"
	Api_Ready_FullMethodName        = "/taikai.v1.Api/Ready"
)

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	UpsertHellos(ctx context.Context, in *UpsertHellosRequest, opts ...grpc.CallOption) (*Hellos, error)
	DeleteHellos(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	ListHellos(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Hellos, error)
	GetHellos(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hellos, error)
	// Health check
	Healthy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Readiness check
	Ready(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) UpsertHellos(ctx context.Context, in *UpsertHellosRequest, opts ...grpc.CallOption) (*Hellos, error) {
	out := new(Hellos)
	err := c.cc.Invoke(ctx, Api_UpsertHellos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteHellos(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, Api_DeleteHellos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListHellos(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Hellos, error) {
	out := new(Hellos)
	err := c.cc.Invoke(ctx, Api_ListHellos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetHellos(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hellos, error) {
	out := new(Hellos)
	err := c.cc.Invoke(ctx, Api_GetHellos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Healthy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Api_Healthy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Ready(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Api_Ready_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations should embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	UpsertHellos(context.Context, *UpsertHellosRequest) (*Hellos, error)
	DeleteHellos(context.Context, *DeleteRequest) (*DeleteResponse, error)
	ListHellos(context.Context, *ListRequest) (*Hellos, error)
	GetHellos(context.Context, *GetRequest) (*Hellos, error)
	// Health check
	Healthy(context.Context, *Empty) (*Empty, error)
	// Readiness check
	Ready(context.Context, *Empty) (*Empty, error)
}

// UnimplementedApiServer should be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) UpsertHellos(context.Context, *UpsertHellosRequest) (*Hellos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertHellos not implemented")
}
func (UnimplementedApiServer) DeleteHellos(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHellos not implemented")
}
func (UnimplementedApiServer) ListHellos(context.Context, *ListRequest) (*Hellos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHellos not implemented")
}
func (UnimplementedApiServer) GetHellos(context.Context, *GetRequest) (*Hellos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHellos not implemented")
}
func (UnimplementedApiServer) Healthy(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthy not implemented")
}
func (UnimplementedApiServer) Ready(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ready not implemented")
}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_UpsertHellos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertHellosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpsertHellos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_UpsertHellos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpsertHellos(ctx, req.(*UpsertHellosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteHellos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteHellos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_DeleteHellos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteHellos(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ListHellos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ListHellos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_ListHellos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ListHellos(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetHellos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetHellos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetHellos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetHellos(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Healthy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Healthy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_Healthy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Healthy(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_Ready_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Ready(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "taikai.v1.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertHellos",
			Handler:    _Api_UpsertHellos_Handler,
		},
		{
			MethodName: "DeleteHellos",
			Handler:    _Api_DeleteHellos_Handler,
		},
		{
			MethodName: "ListHellos",
			Handler:    _Api_ListHellos_Handler,
		},
		{
			MethodName: "GetHellos",
			Handler:    _Api_GetHellos_Handler,
		},
		{
			MethodName: "Healthy",
			Handler:    _Api_Healthy_Handler,
		},
		{
			MethodName: "Ready",
			Handler:    _Api_Ready_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taikai/v1/api.proto",
}
