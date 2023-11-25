// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: bilibili/broadcast/v2/laser.proto

package v2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 服务端下发Laser事件
type LaserEventResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务id
	Taskid int64 `protobuf:"varint,1,opt,name=taskid,proto3" json:"taskid,omitempty"`
	// 指令名
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	// 指令参数json字符串
	Params string `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *LaserEventResp) Reset() {
	*x = LaserEventResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_broadcast_v2_laser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaserEventResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaserEventResp) ProtoMessage() {}

func (x *LaserEventResp) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_broadcast_v2_laser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaserEventResp.ProtoReflect.Descriptor instead.
func (*LaserEventResp) Descriptor() ([]byte, []int) {
	return file_bilibili_broadcast_v2_laser_proto_rawDescGZIP(), []int{0}
}

func (x *LaserEventResp) GetTaskid() int64 {
	if x != nil {
		return x.Taskid
	}
	return 0
}

func (x *LaserEventResp) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *LaserEventResp) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

var File_bilibili_broadcast_v2_laser_proto protoreflect.FileDescriptor

var file_bilibili_broadcast_v2_laser_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0e, 0x4c, 0x61, 0x73, 0x65, 0x72,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x32, 0x56, 0x0a, 0x05, 0x4c, 0x61, 0x73, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x0a, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x25, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x61, 0x73, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x30, 0x01, 0x42, 0x44, 0x5a, 0x42, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x6d, 0x74, 0x63, 0x32, 0x33, 0x33, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_broadcast_v2_laser_proto_rawDescOnce sync.Once
	file_bilibili_broadcast_v2_laser_proto_rawDescData = file_bilibili_broadcast_v2_laser_proto_rawDesc
)

func file_bilibili_broadcast_v2_laser_proto_rawDescGZIP() []byte {
	file_bilibili_broadcast_v2_laser_proto_rawDescOnce.Do(func() {
		file_bilibili_broadcast_v2_laser_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_broadcast_v2_laser_proto_rawDescData)
	})
	return file_bilibili_broadcast_v2_laser_proto_rawDescData
}

var file_bilibili_broadcast_v2_laser_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bilibili_broadcast_v2_laser_proto_goTypes = []interface{}{
	(*LaserEventResp)(nil), // 0: bilibili.broadcast.v2.LaserEventResp
	(*emptypb.Empty)(nil),  // 1: google.protobuf.Empty
}
var file_bilibili_broadcast_v2_laser_proto_depIdxs = []int32{
	1, // 0: bilibili.broadcast.v2.Laser.WatchEvent:input_type -> google.protobuf.Empty
	0, // 1: bilibili.broadcast.v2.Laser.WatchEvent:output_type -> bilibili.broadcast.v2.LaserEventResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bilibili_broadcast_v2_laser_proto_init() }
func file_bilibili_broadcast_v2_laser_proto_init() {
	if File_bilibili_broadcast_v2_laser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_broadcast_v2_laser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaserEventResp); i {
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
			RawDescriptor: file_bilibili_broadcast_v2_laser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_broadcast_v2_laser_proto_goTypes,
		DependencyIndexes: file_bilibili_broadcast_v2_laser_proto_depIdxs,
		MessageInfos:      file_bilibili_broadcast_v2_laser_proto_msgTypes,
	}.Build()
	File_bilibili_broadcast_v2_laser_proto = out.File
	file_bilibili_broadcast_v2_laser_proto_rawDesc = nil
	file_bilibili_broadcast_v2_laser_proto_goTypes = nil
	file_bilibili_broadcast_v2_laser_proto_depIdxs = nil
}
