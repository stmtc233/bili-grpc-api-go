// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: bilibili/relation/interfaces/api.proto

package interfaces

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

type AtSearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 可以为 1 , 但是不能为 0 或空 不知道有啥用
	Mid int64 `protobuf:"varint,1,opt,name=mid,proto3" json:"mid,omitempty"`
	// 用户名搜索关键词
	Keyword string `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *AtSearchReq) Reset() {
	*x = AtSearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_relation_interfaces_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtSearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtSearchReq) ProtoMessage() {}

func (x *AtSearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_relation_interfaces_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtSearchReq.ProtoReflect.Descriptor instead.
func (*AtSearchReq) Descriptor() ([]byte, []int) {
	return file_bilibili_relation_interfaces_api_proto_rawDescGZIP(), []int{0}
}

func (x *AtSearchReq) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *AtSearchReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type AtSearchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 搜索结果分组
	Items []*AtGroup `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *AtSearchReply) Reset() {
	*x = AtSearchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_relation_interfaces_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtSearchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtSearchReply) ProtoMessage() {}

func (x *AtSearchReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_relation_interfaces_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtSearchReply.ProtoReflect.Descriptor instead.
func (*AtSearchReply) Descriptor() ([]byte, []int) {
	return file_bilibili_relation_interfaces_api_proto_rawDescGZIP(), []int{1}
}

func (x *AtSearchReply) GetItems() []*AtGroup {
	if x != nil {
		return x.Items
	}
	return nil
}

type AtGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分组类型  2: 我的关注 4:其他 ,其他自测
	GroupType int32 `protobuf:"varint,1,opt,name=group_type,json=groupType,proto3" json:"group_type,omitempty"`
	// 分组名称
	GroupName string `protobuf:"bytes,2,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	// 用户列表
	Items []*AtItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *AtGroup) Reset() {
	*x = AtGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_relation_interfaces_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtGroup) ProtoMessage() {}

func (x *AtGroup) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_relation_interfaces_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtGroup.ProtoReflect.Descriptor instead.
func (*AtGroup) Descriptor() ([]byte, []int) {
	return file_bilibili_relation_interfaces_api_proto_rawDescGZIP(), []int{2}
}

func (x *AtGroup) GetGroupType() int32 {
	if x != nil {
		return x.GroupType
	}
	return 0
}

func (x *AtGroup) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *AtGroup) GetItems() []*AtItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type AtItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mid                int64  `protobuf:"varint,1,opt,name=mid,proto3" json:"mid,omitempty"`
	Name               string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Face               string `protobuf:"bytes,3,opt,name=face,proto3" json:"face,omitempty"`
	Fans               int32  `protobuf:"varint,4,opt,name=fans,proto3" json:"fans,omitempty"`
	OfficialVerifyType int32  `protobuf:"varint,5,opt,name=official_verify_type,json=officialVerifyType,proto3" json:"official_verify_type,omitempty"`
}

func (x *AtItem) Reset() {
	*x = AtItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_relation_interfaces_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtItem) ProtoMessage() {}

func (x *AtItem) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_relation_interfaces_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtItem.ProtoReflect.Descriptor instead.
func (*AtItem) Descriptor() ([]byte, []int) {
	return file_bilibili_relation_interfaces_api_proto_rawDescGZIP(), []int{3}
}

func (x *AtItem) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *AtItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AtItem) GetFace() string {
	if x != nil {
		return x.Face
	}
	return ""
}

func (x *AtItem) GetFans() int32 {
	if x != nil {
		return x.Fans
	}
	return 0
}

func (x *AtItem) GetOfficialVerifyType() int32 {
	if x != nil {
		return x.OfficialVerifyType
	}
	return 0
}

var File_bilibili_relation_interfaces_api_proto protoreflect.FileDescriptor

var file_bilibili_relation_interfaces_api_proto_rawDesc = []byte{
	0x0a, 0x26, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x39, 0x0a, 0x0b, 0x41, 0x74, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x4e, 0x0a, 0x0d, 0x41, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x07, 0x41, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x06,
	0x41, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x61, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x66, 0x61, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x12, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x32, 0x7b, 0x0a, 0x11, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x08, 0x41,
	0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x2b, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x1a, 0x2d, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x6d, 0x74, 0x63, 0x32, 0x33, 0x33, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_relation_interfaces_api_proto_rawDescOnce sync.Once
	file_bilibili_relation_interfaces_api_proto_rawDescData = file_bilibili_relation_interfaces_api_proto_rawDesc
)

func file_bilibili_relation_interfaces_api_proto_rawDescGZIP() []byte {
	file_bilibili_relation_interfaces_api_proto_rawDescOnce.Do(func() {
		file_bilibili_relation_interfaces_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_relation_interfaces_api_proto_rawDescData)
	})
	return file_bilibili_relation_interfaces_api_proto_rawDescData
}

var file_bilibili_relation_interfaces_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bilibili_relation_interfaces_api_proto_goTypes = []interface{}{
	(*AtSearchReq)(nil),   // 0: bilibili.relation.interface.v1.AtSearchReq
	(*AtSearchReply)(nil), // 1: bilibili.relation.interface.v1.AtSearchReply
	(*AtGroup)(nil),       // 2: bilibili.relation.interface.v1.AtGroup
	(*AtItem)(nil),        // 3: bilibili.relation.interface.v1.AtItem
}
var file_bilibili_relation_interfaces_api_proto_depIdxs = []int32{
	2, // 0: bilibili.relation.interface.v1.AtSearchReply.items:type_name -> bilibili.relation.interface.v1.AtGroup
	3, // 1: bilibili.relation.interface.v1.AtGroup.items:type_name -> bilibili.relation.interface.v1.AtItem
	0, // 2: bilibili.relation.interface.v1.RelationInterface.AtSearch:input_type -> bilibili.relation.interface.v1.AtSearchReq
	1, // 3: bilibili.relation.interface.v1.RelationInterface.AtSearch:output_type -> bilibili.relation.interface.v1.AtSearchReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bilibili_relation_interfaces_api_proto_init() }
func file_bilibili_relation_interfaces_api_proto_init() {
	if File_bilibili_relation_interfaces_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_relation_interfaces_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtSearchReq); i {
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
		file_bilibili_relation_interfaces_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtSearchReply); i {
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
		file_bilibili_relation_interfaces_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtGroup); i {
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
		file_bilibili_relation_interfaces_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtItem); i {
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
			RawDescriptor: file_bilibili_relation_interfaces_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_relation_interfaces_api_proto_goTypes,
		DependencyIndexes: file_bilibili_relation_interfaces_api_proto_depIdxs,
		MessageInfos:      file_bilibili_relation_interfaces_api_proto_msgTypes,
	}.Build()
	File_bilibili_relation_interfaces_api_proto = out.File
	file_bilibili_relation_interfaces_api_proto_rawDesc = nil
	file_bilibili_relation_interfaces_api_proto_goTypes = nil
	file_bilibili_relation_interfaces_api_proto_depIdxs = nil
}
