// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: bilibili/app/playurl/v1/playurl.proto

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
	PlayURL_PlayURL_FullMethodName      = "/bilibili.app.playurl.v1.PlayURL/PlayURL"
	PlayURL_Project_FullMethodName      = "/bilibili.app.playurl.v1.PlayURL/Project"
	PlayURL_PlayView_FullMethodName     = "/bilibili.app.playurl.v1.PlayURL/PlayView"
	PlayURL_PlayConfEdit_FullMethodName = "/bilibili.app.playurl.v1.PlayURL/PlayConfEdit"
	PlayURL_PlayConf_FullMethodName     = "/bilibili.app.playurl.v1.PlayURL/PlayConf"
)

// PlayURLClient is the client API for PlayURL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayURLClient interface {
	// 视频地址
	PlayURL(ctx context.Context, in *PlayURLReq, opts ...grpc.CallOption) (*PlayURLReply, error)
	// 投屏地址
	Project(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*ProjectReply, error)
	// 播放页信息
	PlayView(ctx context.Context, in *PlayViewReq, opts ...grpc.CallOption) (*PlayViewReply, error)
	// 编辑播放界面配置
	PlayConfEdit(ctx context.Context, in *PlayConfEditReq, opts ...grpc.CallOption) (*PlayConfEditReply, error)
	// 获取播放界面配置
	PlayConf(ctx context.Context, in *PlayConfReq, opts ...grpc.CallOption) (*PlayConfReply, error)
}

type playURLClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayURLClient(cc grpc.ClientConnInterface) PlayURLClient {
	return &playURLClient{cc}
}

func (c *playURLClient) PlayURL(ctx context.Context, in *PlayURLReq, opts ...grpc.CallOption) (*PlayURLReply, error) {
	out := new(PlayURLReply)
	err := c.cc.Invoke(ctx, PlayURL_PlayURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playURLClient) Project(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*ProjectReply, error) {
	out := new(ProjectReply)
	err := c.cc.Invoke(ctx, PlayURL_Project_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playURLClient) PlayView(ctx context.Context, in *PlayViewReq, opts ...grpc.CallOption) (*PlayViewReply, error) {
	out := new(PlayViewReply)
	err := c.cc.Invoke(ctx, PlayURL_PlayView_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playURLClient) PlayConfEdit(ctx context.Context, in *PlayConfEditReq, opts ...grpc.CallOption) (*PlayConfEditReply, error) {
	out := new(PlayConfEditReply)
	err := c.cc.Invoke(ctx, PlayURL_PlayConfEdit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playURLClient) PlayConf(ctx context.Context, in *PlayConfReq, opts ...grpc.CallOption) (*PlayConfReply, error) {
	out := new(PlayConfReply)
	err := c.cc.Invoke(ctx, PlayURL_PlayConf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayURLServer is the server API for PlayURL service.
// All implementations must embed UnimplementedPlayURLServer
// for forward compatibility
type PlayURLServer interface {
	// 视频地址
	PlayURL(context.Context, *PlayURLReq) (*PlayURLReply, error)
	// 投屏地址
	Project(context.Context, *ProjectReq) (*ProjectReply, error)
	// 播放页信息
	PlayView(context.Context, *PlayViewReq) (*PlayViewReply, error)
	// 编辑播放界面配置
	PlayConfEdit(context.Context, *PlayConfEditReq) (*PlayConfEditReply, error)
	// 获取播放界面配置
	PlayConf(context.Context, *PlayConfReq) (*PlayConfReply, error)
	mustEmbedUnimplementedPlayURLServer()
}

// UnimplementedPlayURLServer must be embedded to have forward compatible implementations.
type UnimplementedPlayURLServer struct {
}

func (UnimplementedPlayURLServer) PlayURL(context.Context, *PlayURLReq) (*PlayURLReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayURL not implemented")
}
func (UnimplementedPlayURLServer) Project(context.Context, *ProjectReq) (*ProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Project not implemented")
}
func (UnimplementedPlayURLServer) PlayView(context.Context, *PlayViewReq) (*PlayViewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayView not implemented")
}
func (UnimplementedPlayURLServer) PlayConfEdit(context.Context, *PlayConfEditReq) (*PlayConfEditReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayConfEdit not implemented")
}
func (UnimplementedPlayURLServer) PlayConf(context.Context, *PlayConfReq) (*PlayConfReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayConf not implemented")
}
func (UnimplementedPlayURLServer) mustEmbedUnimplementedPlayURLServer() {}

// UnsafePlayURLServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayURLServer will
// result in compilation errors.
type UnsafePlayURLServer interface {
	mustEmbedUnimplementedPlayURLServer()
}

func RegisterPlayURLServer(s grpc.ServiceRegistrar, srv PlayURLServer) {
	s.RegisterService(&PlayURL_ServiceDesc, srv)
}

func _PlayURL_PlayURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayURLServer).PlayURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlayURL_PlayURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayURLServer).PlayURL(ctx, req.(*PlayURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayURL_Project_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayURLServer).Project(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlayURL_Project_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayURLServer).Project(ctx, req.(*ProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayURL_PlayView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayViewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayURLServer).PlayView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlayURL_PlayView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayURLServer).PlayView(ctx, req.(*PlayViewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayURL_PlayConfEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayConfEditReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayURLServer).PlayConfEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlayURL_PlayConfEdit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayURLServer).PlayConfEdit(ctx, req.(*PlayConfEditReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayURL_PlayConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayConfReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayURLServer).PlayConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlayURL_PlayConf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayURLServer).PlayConf(ctx, req.(*PlayConfReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayURL_ServiceDesc is the grpc.ServiceDesc for PlayURL service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayURL_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.playurl.v1.PlayURL",
	HandlerType: (*PlayURLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlayURL",
			Handler:    _PlayURL_PlayURL_Handler,
		},
		{
			MethodName: "Project",
			Handler:    _PlayURL_Project_Handler,
		},
		{
			MethodName: "PlayView",
			Handler:    _PlayURL_PlayView_Handler,
		},
		{
			MethodName: "PlayConfEdit",
			Handler:    _PlayURL_PlayConfEdit_Handler,
		},
		{
			MethodName: "PlayConf",
			Handler:    _PlayURL_PlayConf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bilibili/app/playurl/v1/playurl.proto",
}