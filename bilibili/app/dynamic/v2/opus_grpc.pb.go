// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: bilibili/app/dynamic/v2/opus.proto

package v2

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
	Opus_ListFav_FullMethodName    = "/bilibili.app.dynamic.v2.Opus/ListFav"
	Opus_OpusDetail_FullMethodName = "/bilibili.app.dynamic.v2.Opus/OpusDetail"
)

// OpusClient is the client API for Opus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpusClient interface {
	ListFav(ctx context.Context, in *ListFavReq, opts ...grpc.CallOption) (*ListFavResp, error)
	OpusDetail(ctx context.Context, in *OpusDetailReq, opts ...grpc.CallOption) (*OpusDetailResp, error)
}

type opusClient struct {
	cc grpc.ClientConnInterface
}

func NewOpusClient(cc grpc.ClientConnInterface) OpusClient {
	return &opusClient{cc}
}

func (c *opusClient) ListFav(ctx context.Context, in *ListFavReq, opts ...grpc.CallOption) (*ListFavResp, error) {
	out := new(ListFavResp)
	err := c.cc.Invoke(ctx, Opus_ListFav_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opusClient) OpusDetail(ctx context.Context, in *OpusDetailReq, opts ...grpc.CallOption) (*OpusDetailResp, error) {
	out := new(OpusDetailResp)
	err := c.cc.Invoke(ctx, Opus_OpusDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpusServer is the server API for Opus service.
// All implementations must embed UnimplementedOpusServer
// for forward compatibility
type OpusServer interface {
	ListFav(context.Context, *ListFavReq) (*ListFavResp, error)
	OpusDetail(context.Context, *OpusDetailReq) (*OpusDetailResp, error)
	mustEmbedUnimplementedOpusServer()
}

// UnimplementedOpusServer must be embedded to have forward compatible implementations.
type UnimplementedOpusServer struct {
}

func (UnimplementedOpusServer) ListFav(context.Context, *ListFavReq) (*ListFavResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFav not implemented")
}
func (UnimplementedOpusServer) OpusDetail(context.Context, *OpusDetailReq) (*OpusDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpusDetail not implemented")
}
func (UnimplementedOpusServer) mustEmbedUnimplementedOpusServer() {}

// UnsafeOpusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpusServer will
// result in compilation errors.
type UnsafeOpusServer interface {
	mustEmbedUnimplementedOpusServer()
}

func RegisterOpusServer(s grpc.ServiceRegistrar, srv OpusServer) {
	s.RegisterService(&Opus_ServiceDesc, srv)
}

func _Opus_ListFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFavReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpusServer).ListFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Opus_ListFav_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpusServer).ListFav(ctx, req.(*ListFavReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Opus_OpusDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpusDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpusServer).OpusDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Opus_OpusDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpusServer).OpusDetail(ctx, req.(*OpusDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Opus_ServiceDesc is the grpc.ServiceDesc for Opus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Opus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.dynamic.v2.Opus",
	HandlerType: (*OpusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFav",
			Handler:    _Opus_ListFav_Handler,
		},
		{
			MethodName: "OpusDetail",
			Handler:    _Opus_OpusDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/dynamic/v2/opus.proto",
}
