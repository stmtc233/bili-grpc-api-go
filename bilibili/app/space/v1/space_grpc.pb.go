// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: bilibili/app/space/v1/space.proto

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
	Space_Archive_FullMethodName = "/bilibili.app.space.v1.Space/Archive"
)

// SpaceClient is the client API for Space service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpaceClient interface {
	Archive(ctx context.Context, in *ArchiveReq, opts ...grpc.CallOption) (*ArchiveReply, error)
}

type spaceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpaceClient(cc grpc.ClientConnInterface) SpaceClient {
	return &spaceClient{cc}
}

func (c *spaceClient) Archive(ctx context.Context, in *ArchiveReq, opts ...grpc.CallOption) (*ArchiveReply, error) {
	out := new(ArchiveReply)
	err := c.cc.Invoke(ctx, Space_Archive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpaceServer is the server API for Space service.
// All implementations must embed UnimplementedSpaceServer
// for forward compatibility
type SpaceServer interface {
	Archive(context.Context, *ArchiveReq) (*ArchiveReply, error)
	mustEmbedUnimplementedSpaceServer()
}

// UnimplementedSpaceServer must be embedded to have forward compatible implementations.
type UnimplementedSpaceServer struct {
}

func (UnimplementedSpaceServer) Archive(context.Context, *ArchiveReq) (*ArchiveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Archive not implemented")
}
func (UnimplementedSpaceServer) mustEmbedUnimplementedSpaceServer() {}

// UnsafeSpaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpaceServer will
// result in compilation errors.
type UnsafeSpaceServer interface {
	mustEmbedUnimplementedSpaceServer()
}

func RegisterSpaceServer(s grpc.ServiceRegistrar, srv SpaceServer) {
	s.RegisterService(&Space_ServiceDesc, srv)
}

func _Space_Archive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServer).Archive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Space_Archive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServer).Archive(ctx, req.(*ArchiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Space_ServiceDesc is the grpc.ServiceDesc for Space service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Space_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.space.v1.Space",
	HandlerType: (*SpaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Archive",
			Handler:    _Space_Archive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/space/v1/space.proto",
}
