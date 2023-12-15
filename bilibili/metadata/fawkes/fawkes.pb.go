// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: bilibili/metadata/fawkes/fawkes.proto

package fawkes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FawkesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 客户端在fawkes系统中对应的已发布最新的config版本号
	Config string `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// 客户端在fawkes系统中对应的已发布最新的ff版本号
	Ff string `protobuf:"bytes,2,opt,name=ff,proto3" json:"ff,omitempty"`
	Dd string `protobuf:"bytes,3,opt,name=dd,proto3" json:"dd,omitempty"`
}

func (x *FawkesReply) Reset() {
	*x = FawkesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_metadata_fawkes_fawkes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FawkesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FawkesReply) ProtoMessage() {}

func (x *FawkesReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_metadata_fawkes_fawkes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FawkesReply.ProtoReflect.Descriptor instead.
func (*FawkesReply) Descriptor() ([]byte, []int) {
	return file_bilibili_metadata_fawkes_fawkes_proto_rawDescGZIP(), []int{0}
}

func (x *FawkesReply) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *FawkesReply) GetFf() string {
	if x != nil {
		return x.Ff
	}
	return ""
}

func (x *FawkesReply) GetDd() string {
	if x != nil {
		return x.Dd
	}
	return ""
}

type FawkesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 客户端在fawkes系统的唯一名, 如 `android64`
	Appkey string `protobuf:"bytes,1,opt,name=appkey,proto3" json:"appkey,omitempty"`
	// 客户端在fawkes系统中的环境参数, 如 `prod`
	Env string `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
	// 启动id, 8 位 0~9, a~z 组成的字符串
	SessionId string `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *FawkesReq) Reset() {
	*x = FawkesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_metadata_fawkes_fawkes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FawkesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FawkesReq) ProtoMessage() {}

func (x *FawkesReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_metadata_fawkes_fawkes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FawkesReq.ProtoReflect.Descriptor instead.
func (*FawkesReq) Descriptor() ([]byte, []int) {
	return file_bilibili_metadata_fawkes_fawkes_proto_rawDescGZIP(), []int{1}
}

func (x *FawkesReq) GetAppkey() string {
	if x != nil {
		return x.Appkey
	}
	return ""
}

func (x *FawkesReq) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *FawkesReq) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

var File_bilibili_metadata_fawkes_fawkes_proto protoreflect.FileDescriptor

var file_bilibili_metadata_fawkes_fawkes_proto_rawDesc = []byte{
	0x0a, 0x25, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x66, 0x61, 0x77, 0x6b, 0x65, 0x73, 0x2f, 0x66, 0x61, 0x77, 0x6b, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x66, 0x61, 0x77, 0x6b, 0x65,
	0x73, 0x22, 0x45, 0x0a, 0x0b, 0x46, 0x61, 0x77, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x66, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x66, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x64, 0x22, 0x54, 0x0a, 0x09, 0x46, 0x61, 0x77, 0x6b,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6e, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x3f,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x6d,
	0x74, 0x63, 0x32, 0x33, 0x33, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x66, 0x61, 0x77, 0x6b, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_metadata_fawkes_fawkes_proto_rawDescOnce sync.Once
	file_bilibili_metadata_fawkes_fawkes_proto_rawDescData = file_bilibili_metadata_fawkes_fawkes_proto_rawDesc
)

func file_bilibili_metadata_fawkes_fawkes_proto_rawDescGZIP() []byte {
	file_bilibili_metadata_fawkes_fawkes_proto_rawDescOnce.Do(func() {
		file_bilibili_metadata_fawkes_fawkes_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_metadata_fawkes_fawkes_proto_rawDescData)
	})
	return file_bilibili_metadata_fawkes_fawkes_proto_rawDescData
}

var file_bilibili_metadata_fawkes_fawkes_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bilibili_metadata_fawkes_fawkes_proto_goTypes = []interface{}{
	(*FawkesReply)(nil), // 0: bilibili.metadata.fawkes.FawkesReply
	(*FawkesReq)(nil),   // 1: bilibili.metadata.fawkes.FawkesReq
}
var file_bilibili_metadata_fawkes_fawkes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bilibili_metadata_fawkes_fawkes_proto_init() }
func file_bilibili_metadata_fawkes_fawkes_proto_init() {
	if File_bilibili_metadata_fawkes_fawkes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_metadata_fawkes_fawkes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FawkesReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bilibili_metadata_fawkes_fawkes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FawkesReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bilibili_metadata_fawkes_fawkes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_metadata_fawkes_fawkes_proto_goTypes,
		DependencyIndexes: file_bilibili_metadata_fawkes_fawkes_proto_depIdxs,
		MessageInfos:      file_bilibili_metadata_fawkes_fawkes_proto_msgTypes,
	}.Build()
	File_bilibili_metadata_fawkes_fawkes_proto = out.File
	file_bilibili_metadata_fawkes_fawkes_proto_rawDesc = nil
	file_bilibili_metadata_fawkes_fawkes_proto_goTypes = nil
	file_bilibili_metadata_fawkes_fawkes_proto_depIdxs = nil
}
