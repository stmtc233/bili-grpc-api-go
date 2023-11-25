// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: bilibili/polymer/list/v1/list.proto

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
	List_FavoriteTab_FullMethodName  = "/bilibili.polymer.list.v1.List/FavoriteTab"
	List_CheckAccount_FullMethodName = "/bilibili.polymer.list.v1.List/CheckAccount"
)

// ListClient is the client API for List service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListClient interface {
	FavoriteTab(ctx context.Context, in *FavoriteTabReq, opts ...grpc.CallOption) (*FavoriteTabReply, error)
	CheckAccount(ctx context.Context, in *CheckAccountReq, opts ...grpc.CallOption) (*CheckAccountReply, error)
}

type listClient struct {
	cc grpc.ClientConnInterface
}

func NewListClient(cc grpc.ClientConnInterface) ListClient {
	return &listClient{cc}
}

func (c *listClient) FavoriteTab(ctx context.Context, in *FavoriteTabReq, opts ...grpc.CallOption) (*FavoriteTabReply, error) {
	out := new(FavoriteTabReply)
	err := c.cc.Invoke(ctx, List_FavoriteTab_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listClient) CheckAccount(ctx context.Context, in *CheckAccountReq, opts ...grpc.CallOption) (*CheckAccountReply, error) {
	out := new(CheckAccountReply)
	err := c.cc.Invoke(ctx, List_CheckAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListServer is the server API for List service.
// All implementations must embed UnimplementedListServer
// for forward compatibility
type ListServer interface {
	FavoriteTab(context.Context, *FavoriteTabReq) (*FavoriteTabReply, error)
	CheckAccount(context.Context, *CheckAccountReq) (*CheckAccountReply, error)
	mustEmbedUnimplementedListServer()
}

// UnimplementedListServer must be embedded to have forward compatible implementations.
type UnimplementedListServer struct {
}

func (UnimplementedListServer) FavoriteTab(context.Context, *FavoriteTabReq) (*FavoriteTabReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteTab not implemented")
}
func (UnimplementedListServer) CheckAccount(context.Context, *CheckAccountReq) (*CheckAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAccount not implemented")
}
func (UnimplementedListServer) mustEmbedUnimplementedListServer() {}

// UnsafeListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListServer will
// result in compilation errors.
type UnsafeListServer interface {
	mustEmbedUnimplementedListServer()
}

func RegisterListServer(s grpc.ServiceRegistrar, srv ListServer) {
	s.RegisterService(&List_ServiceDesc, srv)
}

func _List_FavoriteTab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteTabReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).FavoriteTab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: List_FavoriteTab_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).FavoriteTab(ctx, req.(*FavoriteTabReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _List_CheckAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).CheckAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: List_CheckAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).CheckAccount(ctx, req.(*CheckAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// List_ServiceDesc is the grpc.ServiceDesc for List service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var List_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.polymer.list.v1.List",
	HandlerType: (*ListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FavoriteTab",
			Handler:    _List_FavoriteTab_Handler,
		},
		{
			MethodName: "CheckAccount",
			Handler:    _List_CheckAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/polymer/list/v1/list.proto",
}
