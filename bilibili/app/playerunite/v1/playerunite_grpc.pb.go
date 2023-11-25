// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: bilibili/app/playerunite/v1/playerunite.proto

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
	Player_PlayViewUnite_FullMethodName = "/bilibili.app.playerunite.v1.Player/PlayViewUnite"
)

// PlayerClient is the client API for Player service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerClient interface {
	// 视频地址
	PlayViewUnite(ctx context.Context, in *PlayViewUniteReq, opts ...grpc.CallOption) (*PlayViewUniteReply, error)
}

type playerClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerClient(cc grpc.ClientConnInterface) PlayerClient {
	return &playerClient{cc}
}

func (c *playerClient) PlayViewUnite(ctx context.Context, in *PlayViewUniteReq, opts ...grpc.CallOption) (*PlayViewUniteReply, error) {
	out := new(PlayViewUniteReply)
	err := c.cc.Invoke(ctx, Player_PlayViewUnite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerServer is the server API for Player service.
// All implementations must embed UnimplementedPlayerServer
// for forward compatibility
type PlayerServer interface {
	// 视频地址
	PlayViewUnite(context.Context, *PlayViewUniteReq) (*PlayViewUniteReply, error)
	mustEmbedUnimplementedPlayerServer()
}

// UnimplementedPlayerServer must be embedded to have forward compatible implementations.
type UnimplementedPlayerServer struct {
}

func (UnimplementedPlayerServer) PlayViewUnite(context.Context, *PlayViewUniteReq) (*PlayViewUniteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayViewUnite not implemented")
}
func (UnimplementedPlayerServer) mustEmbedUnimplementedPlayerServer() {}

// UnsafePlayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerServer will
// result in compilation errors.
type UnsafePlayerServer interface {
	mustEmbedUnimplementedPlayerServer()
}

func RegisterPlayerServer(s grpc.ServiceRegistrar, srv PlayerServer) {
	s.RegisterService(&Player_ServiceDesc, srv)
}

func _Player_PlayViewUnite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayViewUniteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServer).PlayViewUnite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Player_PlayViewUnite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServer).PlayViewUnite(ctx, req.(*PlayViewUniteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Player_ServiceDesc is the grpc.ServiceDesc for Player service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Player_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.playerunite.v1.Player",
	HandlerType: (*PlayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlayViewUnite",
			Handler:    _Player_PlayViewUnite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/playerunite/v1/playerunite.proto",
}