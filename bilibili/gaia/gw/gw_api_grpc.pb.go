// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: bilibili/gaia/gw/gw_api.proto

package gw

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Gaia_ExClimbAppleTrees_FullMethodName = "/bilibili.gaia.gw.Gaia/ExClimbAppleTrees"
	Gaia_ExFetchPublicKey_FullMethodName  = "/bilibili.gaia.gw.Gaia/ExFetchPublicKey"
	Gaia_ExGetAxe_FullMethodName          = "/bilibili.gaia.gw.Gaia/ExGetAxe"
	Gaia_ExUploadAppList_FullMethodName   = "/bilibili.gaia.gw.Gaia/ExUploadAppList"
)

// GaiaClient is the client API for Gaia service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GaiaClient interface {
	ExClimbAppleTrees(ctx context.Context, in *GaiaEncryptMsgReq, opts ...grpc.CallOption) (*UploadAppListReply, error)
	// 拉取rsa公钥
	ExFetchPublicKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FetchPublicKeyReply, error)
	ExGetAxe(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FetchPublicKeyReply, error)
	// 应用列表上报
	ExUploadAppList(ctx context.Context, in *GaiaEncryptMsgReq, opts ...grpc.CallOption) (*UploadAppListReply, error)
}

type gaiaClient struct {
	cc grpc.ClientConnInterface
}

func NewGaiaClient(cc grpc.ClientConnInterface) GaiaClient {
	return &gaiaClient{cc}
}

func (c *gaiaClient) ExClimbAppleTrees(ctx context.Context, in *GaiaEncryptMsgReq, opts ...grpc.CallOption) (*UploadAppListReply, error) {
	out := new(UploadAppListReply)
	err := c.cc.Invoke(ctx, Gaia_ExClimbAppleTrees_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gaiaClient) ExFetchPublicKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FetchPublicKeyReply, error) {
	out := new(FetchPublicKeyReply)
	err := c.cc.Invoke(ctx, Gaia_ExFetchPublicKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gaiaClient) ExGetAxe(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FetchPublicKeyReply, error) {
	out := new(FetchPublicKeyReply)
	err := c.cc.Invoke(ctx, Gaia_ExGetAxe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gaiaClient) ExUploadAppList(ctx context.Context, in *GaiaEncryptMsgReq, opts ...grpc.CallOption) (*UploadAppListReply, error) {
	out := new(UploadAppListReply)
	err := c.cc.Invoke(ctx, Gaia_ExUploadAppList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GaiaServer is the server API for Gaia service.
// All implementations must embed UnimplementedGaiaServer
// for forward compatibility
type GaiaServer interface {
	ExClimbAppleTrees(context.Context, *GaiaEncryptMsgReq) (*UploadAppListReply, error)
	// 拉取rsa公钥
	ExFetchPublicKey(context.Context, *emptypb.Empty) (*FetchPublicKeyReply, error)
	ExGetAxe(context.Context, *emptypb.Empty) (*FetchPublicKeyReply, error)
	// 应用列表上报
	ExUploadAppList(context.Context, *GaiaEncryptMsgReq) (*UploadAppListReply, error)
	mustEmbedUnimplementedGaiaServer()
}

// UnimplementedGaiaServer must be embedded to have forward compatible implementations.
type UnimplementedGaiaServer struct {
}

func (UnimplementedGaiaServer) ExClimbAppleTrees(context.Context, *GaiaEncryptMsgReq) (*UploadAppListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExClimbAppleTrees not implemented")
}
func (UnimplementedGaiaServer) ExFetchPublicKey(context.Context, *emptypb.Empty) (*FetchPublicKeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExFetchPublicKey not implemented")
}
func (UnimplementedGaiaServer) ExGetAxe(context.Context, *emptypb.Empty) (*FetchPublicKeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExGetAxe not implemented")
}
func (UnimplementedGaiaServer) ExUploadAppList(context.Context, *GaiaEncryptMsgReq) (*UploadAppListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExUploadAppList not implemented")
}
func (UnimplementedGaiaServer) mustEmbedUnimplementedGaiaServer() {}

// UnsafeGaiaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GaiaServer will
// result in compilation errors.
type UnsafeGaiaServer interface {
	mustEmbedUnimplementedGaiaServer()
}

func RegisterGaiaServer(s grpc.ServiceRegistrar, srv GaiaServer) {
	s.RegisterService(&Gaia_ServiceDesc, srv)
}

func _Gaia_ExClimbAppleTrees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GaiaEncryptMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaiaServer).ExClimbAppleTrees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gaia_ExClimbAppleTrees_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaiaServer).ExClimbAppleTrees(ctx, req.(*GaiaEncryptMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gaia_ExFetchPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaiaServer).ExFetchPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gaia_ExFetchPublicKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaiaServer).ExFetchPublicKey(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gaia_ExGetAxe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaiaServer).ExGetAxe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gaia_ExGetAxe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaiaServer).ExGetAxe(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gaia_ExUploadAppList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GaiaEncryptMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaiaServer).ExUploadAppList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gaia_ExUploadAppList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaiaServer).ExUploadAppList(ctx, req.(*GaiaEncryptMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Gaia_ServiceDesc is the grpc.ServiceDesc for Gaia service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gaia_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.gaia.gw.Gaia",
	HandlerType: (*GaiaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExClimbAppleTrees",
			Handler:    _Gaia_ExClimbAppleTrees_Handler,
		},
		{
			MethodName: "ExFetchPublicKey",
			Handler:    _Gaia_ExFetchPublicKey_Handler,
		},
		{
			MethodName: "ExGetAxe",
			Handler:    _Gaia_ExGetAxe_Handler,
		},
		{
			MethodName: "ExUploadAppList",
			Handler:    _Gaia_ExUploadAppList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/gaia/gw/gw_api.proto",
}
