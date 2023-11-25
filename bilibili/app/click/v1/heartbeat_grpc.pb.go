// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: bilibili/app/click/v1/heartbeat.proto

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
	Click_HeartBeat_FullMethodName = "/bilibili.app.click.v1.Click/HeartBeat"
)

// ClickClient is the client API for Click service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClickClient interface {
	// 客户端心跳上传
	HeartBeat(ctx context.Context, in *HeartBeatReq, opts ...grpc.CallOption) (*HeartBeatReply, error)
}

type clickClient struct {
	cc grpc.ClientConnInterface
}

func NewClickClient(cc grpc.ClientConnInterface) ClickClient {
	return &clickClient{cc}
}

func (c *clickClient) HeartBeat(ctx context.Context, in *HeartBeatReq, opts ...grpc.CallOption) (*HeartBeatReply, error) {
	out := new(HeartBeatReply)
	err := c.cc.Invoke(ctx, Click_HeartBeat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClickServer is the server API for Click service.
// All implementations must embed UnimplementedClickServer
// for forward compatibility
type ClickServer interface {
	// 客户端心跳上传
	HeartBeat(context.Context, *HeartBeatReq) (*HeartBeatReply, error)
	mustEmbedUnimplementedClickServer()
}

// UnimplementedClickServer must be embedded to have forward compatible implementations.
type UnimplementedClickServer struct {
}

func (UnimplementedClickServer) HeartBeat(context.Context, *HeartBeatReq) (*HeartBeatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (UnimplementedClickServer) mustEmbedUnimplementedClickServer() {}

// UnsafeClickServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClickServer will
// result in compilation errors.
type UnsafeClickServer interface {
	mustEmbedUnimplementedClickServer()
}

func RegisterClickServer(s grpc.ServiceRegistrar, srv ClickServer) {
	s.RegisterService(&Click_ServiceDesc, srv)
}

func _Click_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Click_HeartBeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickServer).HeartBeat(ctx, req.(*HeartBeatReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Click_ServiceDesc is the grpc.ServiceDesc for Click service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Click_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.click.v1.Click",
	HandlerType: (*ClickServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _Click_HeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/click/v1/heartbeat.proto",
}