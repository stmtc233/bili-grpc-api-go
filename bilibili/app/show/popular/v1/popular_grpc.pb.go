// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: bilibili/app/show/popular/v1/popular.proto

package v1

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
	Popular_Index_FullMethodName = "/bilibili.app.show.v1.Popular/Index"
)

// PopularClient is the client API for Popular service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PopularClient interface {
	// 热门列表
	Index(ctx context.Context, in *PopularResultReq, opts ...grpc.CallOption) (*PopularReply, error)
}

type popularClient struct {
	cc grpc.ClientConnInterface
}

func NewPopularClient(cc grpc.ClientConnInterface) PopularClient {
	return &popularClient{cc}
}

func (c *popularClient) Index(ctx context.Context, in *PopularResultReq, opts ...grpc.CallOption) (*PopularReply, error) {
	out := new(PopularReply)
	err := c.cc.Invoke(ctx, Popular_Index_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PopularServer is the server API for Popular service.
// All implementations must embed UnimplementedPopularServer
// for forward compatibility
type PopularServer interface {
	// 热门列表
	Index(context.Context, *PopularResultReq) (*PopularReply, error)
	mustEmbedUnimplementedPopularServer()
}

// UnimplementedPopularServer must be embedded to have forward compatible implementations.
type UnimplementedPopularServer struct {
}

func (UnimplementedPopularServer) Index(context.Context, *PopularResultReq) (*PopularReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedPopularServer) mustEmbedUnimplementedPopularServer() {}

// UnsafePopularServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PopularServer will
// result in compilation errors.
type UnsafePopularServer interface {
	mustEmbedUnimplementedPopularServer()
}

func RegisterPopularServer(s grpc.ServiceRegistrar, srv PopularServer) {
	s.RegisterService(&Popular_ServiceDesc, srv)
}

func _Popular_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PopularResultReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PopularServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Popular_Index_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PopularServer).Index(ctx, req.(*PopularResultReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Popular_ServiceDesc is the grpc.ServiceDesc for Popular service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Popular_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.show.v1.Popular",
	HandlerType: (*PopularServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _Popular_Index_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/show/popular/v1/popular.proto",
}