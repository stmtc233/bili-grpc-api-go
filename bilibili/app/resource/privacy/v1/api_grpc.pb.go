// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: bilibili/app/resource/privacy/v1/api.proto

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
	Privacy_PrivacyConfig_FullMethodName    = "/bilibili.app.resource.privacy.v1.Privacy/PrivacyConfig"
	Privacy_SetPrivacyConfig_FullMethodName = "/bilibili.app.resource.privacy.v1.Privacy/SetPrivacyConfig"
)

// PrivacyClient is the client API for Privacy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivacyClient interface {
	// 获取隐私设置
	PrivacyConfig(ctx context.Context, in *NoArgRequest, opts ...grpc.CallOption) (*PrivacyConfigReply, error)
	// 修改隐私设置
	SetPrivacyConfig(ctx context.Context, in *SetPrivacyConfigRequest, opts ...grpc.CallOption) (*NoReply, error)
}

type privacyClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivacyClient(cc grpc.ClientConnInterface) PrivacyClient {
	return &privacyClient{cc}
}

func (c *privacyClient) PrivacyConfig(ctx context.Context, in *NoArgRequest, opts ...grpc.CallOption) (*PrivacyConfigReply, error) {
	out := new(PrivacyConfigReply)
	err := c.cc.Invoke(ctx, Privacy_PrivacyConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyClient) SetPrivacyConfig(ctx context.Context, in *SetPrivacyConfigRequest, opts ...grpc.CallOption) (*NoReply, error) {
	out := new(NoReply)
	err := c.cc.Invoke(ctx, Privacy_SetPrivacyConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivacyServer is the server API for Privacy service.
// All implementations must embed UnimplementedPrivacyServer
// for forward compatibility
type PrivacyServer interface {
	// 获取隐私设置
	PrivacyConfig(context.Context, *NoArgRequest) (*PrivacyConfigReply, error)
	// 修改隐私设置
	SetPrivacyConfig(context.Context, *SetPrivacyConfigRequest) (*NoReply, error)
	mustEmbedUnimplementedPrivacyServer()
}

// UnimplementedPrivacyServer must be embedded to have forward compatible implementations.
type UnimplementedPrivacyServer struct {
}

func (UnimplementedPrivacyServer) PrivacyConfig(context.Context, *NoArgRequest) (*PrivacyConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrivacyConfig not implemented")
}
func (UnimplementedPrivacyServer) SetPrivacyConfig(context.Context, *SetPrivacyConfigRequest) (*NoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPrivacyConfig not implemented")
}
func (UnimplementedPrivacyServer) mustEmbedUnimplementedPrivacyServer() {}

// UnsafePrivacyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivacyServer will
// result in compilation errors.
type UnsafePrivacyServer interface {
	mustEmbedUnimplementedPrivacyServer()
}

func RegisterPrivacyServer(s grpc.ServiceRegistrar, srv PrivacyServer) {
	s.RegisterService(&Privacy_ServiceDesc, srv)
}

func _Privacy_PrivacyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoArgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServer).PrivacyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Privacy_PrivacyConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServer).PrivacyConfig(ctx, req.(*NoArgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Privacy_SetPrivacyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPrivacyConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServer).SetPrivacyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Privacy_SetPrivacyConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServer).SetPrivacyConfig(ctx, req.(*SetPrivacyConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Privacy_ServiceDesc is the grpc.ServiceDesc for Privacy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Privacy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.resource.privacy.v1.Privacy",
	HandlerType: (*PrivacyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrivacyConfig",
			Handler:    _Privacy_PrivacyConfig_Handler,
		},
		{
			MethodName: "SetPrivacyConfig",
			Handler:    _Privacy_SetPrivacyConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/resource/privacy/v1/api.proto",
}