// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: bilibili/dagw/component/avatar/v1/plugin.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "https://github.com/stmtc233/bili-grpc-api-go/bilibili/dagw/component/avatar/common"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommentDoubleClickConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interaction *Interaction `protobuf:"bytes,1,opt,name=interaction,proto3" json:"interaction,omitempty"`
	AnimationScale float64 `protobuf:"fixed64,2,opt,name=animation_scale,json=animationScale,proto3" json:"animation_scale,omitempty"`
}

func (x *CommentDoubleClickConfig) Reset() {
	*x = CommentDoubleClickConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentDoubleClickConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentDoubleClickConfig) ProtoMessage() {}

func (x *CommentDoubleClickConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentDoubleClickConfig.ProtoReflect.Descriptor instead.
func (*CommentDoubleClickConfig) Descriptor() ([]byte, []int) {
	return file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *CommentDoubleClickConfig) GetInteraction() *Interaction {
	if x != nil {
		return x.Interaction
	}
	return nil
}

func (x *CommentDoubleClickConfig) GetAnimationScale() float64 {
	if x != nil {
		return x.AnimationScale
	}
	return 0
}

type GyroConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gyroscope *NFTImageV2 `protobuf:"bytes,1,opt,name=gyroscope,proto3" json:"gyroscope,omitempty"`
}

func (x *GyroConfig) Reset() {
	*x = GyroConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GyroConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GyroConfig) ProtoMessage() {}

func (x *GyroConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GyroConfig.ProtoReflect.Descriptor instead.
func (*GyroConfig) Descriptor() ([]byte, []int) {
	return file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *GyroConfig) GetGyroscope() *NFTImageV2 {
	if x != nil {
		return x.Gyroscope
	}
	return nil
}

type GyroscopeContentV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileUrl string `protobuf:"bytes,1,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	Scale float32 `protobuf:"fixed32,2,opt,name=scale,proto3" json:"scale,omitempty"`
	PhysicalOrientation []*PhysicalOrientationV2 `protobuf:"bytes,3,rep,name=physical_orientation,json=physicalOrientation,proto3" json:"physical_orientation,omitempty"`
}

func (x *GyroscopeContentV2) Reset() {
	*x = GyroscopeContentV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GyroscopeContentV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GyroscopeContentV2) ProtoMessage() {}

func (x *GyroscopeContentV2) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GyroscopeContentV2.ProtoReflect.Descriptor instead.
func (*GyroscopeContentV2) Descriptor() ([]byte, []int) {
	return file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *GyroscopeContentV2) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *GyroscopeContentV2) GetScale() float32 {
	if x != nil {
		return x.Scale
	}
	return 0
}

func (x *GyroscopeContentV2) GetPhysicalOrientation() []*PhysicalOrientationV2 {
	if x != nil {
		return x.PhysicalOrientation
	}
	return nil
}

type GyroscopeEntityV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayType string `protobuf:"bytes,1,opt,name=display_type,json=displayType,proto3" json:"display_type,omitempty"`
	Contents []*GyroscopeContentV2 `protobuf:"bytes,2,rep,name=contents,proto3" json:"contents,omitempty"`
}

func (x *GyroscopeEntityV2) Reset() {
	*x = GyroscopeEntityV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GyroscopeEntityV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GyroscopeEntityV2) ProtoMessage() {}

func (x *GyroscopeEntityV2) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GyroscopeEntityV2.ProtoReflect.Descriptor instead.
func (*GyroscopeEntityV2) Descriptor() ([]byte, []int) {
	return file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *GyroscopeEntityV2) GetDisplayType() string {
	if x != nil {
		return x.DisplayType
	}
	return ""
}

func (x *GyroscopeEntityV2) GetContents() []*GyroscopeContentV2 {
	if x != nil {
		return x.Contents
	}
	return nil
}

type Interaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NftId string `protobuf:"bytes,1,opt,name=nft_id,json=nftId,proto3" json:"nft_id,omitempty"`
	Enabled bool `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Itype string `protobuf:"bytes,3,opt,name=itype,proto3" json:"itype,omitempty"`
	MetadataUrl string `protobuf:"bytes,4,opt,name=metadata_url,json=metadataUrl,proto3" json:"metadata_url,omitempty"`
}

func (x *Interaction) Reset() {
	*x = Interaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interaction) ProtoMessage() {}

func (x *Interaction) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interaction.ProtoReflect.Descriptor instead.
func (*Interaction) Descriptor() ([]byte, []int) {
	return file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *Interaction) GetNftId() string {
	if x != nil {
		return x.NftId
	}
	return ""
}

func (x *Interaction) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Interaction) GetItype() string {
	if x != nil {
		return x.Itype
	}
	return ""
}

func (x *Interaction) GetMetadataUrl() string {
	if x != nil {
		return x.MetadataUrl
	}
	return ""
}

type LiveAnimeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsLive bool `protobuf:"varint,1,opt,name=is_live,json=isLive,proto3" json:"is_live,omitempty"`
}

func (x *LiveAnimeConfig) Reset() {
	*x = LiveAnimeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveAnimeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveAnimeConfig) ProtoMessage() {}

func (x *LiveAnimeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveAnimeConfig.ProtoReflect.Descriptor instead.
func (*LiveAnimeConfig) Descriptor() ([]byte, []int) {
	return file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *LiveAnimeConfig) GetIsLive() bool {
	if x != nil {
		return x.IsLive
	}
	return false
}

type LiveAnimeItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color *common.ColorConfig `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	StartRatio float64 `protobuf:"fixed64,2,opt,name=start_ratio,json=startRatio,proto3" json:"start_ratio,omitempty"`
	EndRatio float64 `protobuf:"fixed64,3,opt,name=end_ratio,json=endRatio,proto3" json:"end_ratio,omitempty"`
	StartStroke float64 `protobuf:"fixed64,4,opt,name=start_stroke,json=startStroke,proto3" json:"start_stroke,omitempty"`
	StartOpacity float64 `protobuf:"fixed64,5,opt,name=start_opacity,json=startOpacity,proto3" json:"start_opacity,omitempty"`
	Phase int64 `protobuf:"varint,6,opt,name=phase,proto3" json:"phase,omitempty"`
}

func (x *LiveAnimeItem) Reset() {
	*x = LiveAnimeItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveAnimeItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveAnimeItem) ProtoMessage() {}

func (x *LiveAnimeItem) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveAnimeItem.ProtoReflect.Descriptor instead.
func (*LiveAnimeItem) Descriptor() ([]byte, []int) {
	return file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescGZIP(), []int{6}
}

func (x *LiveAnimeItem) GetColor() *common.ColorConfig {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *LiveAnimeItem) GetStartRatio() float64 {
	if x != nil {
		return x.StartRatio
	}
	return 0
}

func (x *LiveAnimeItem) GetEndRatio() float64 {
	if x != nil {
		return x.EndRatio
	}
	return 0
}

func (x *LiveAnimeItem) GetStartStroke() float64 {
	if x != nil {
		return x.StartStroke
	}
	return 0
}

func (x *LiveAnimeItem) GetStartOpacity() float64 {
	if x != nil {
		return x.StartOpacity
	}
	return 0
}

func (x *LiveAnimeItem) GetPhase() int64 {
	if x != nil {
		return x.Phase
	}
	return 0
}

type NFTImageV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gyroscope []*GyroscopeEntityV2 `protobuf:"bytes,1,rep,name=gyroscope,proto3" json:"gyroscope,omitempty"`
}

func (x *NFTImageV2) Reset() {
	*x = NFTImageV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NFTImageV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NFTImageV2) ProtoMessage() {}

func (x *NFTImageV2) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NFTImageV2.ProtoReflect.Descriptor instead.
func (*NFTImageV2) Descriptor() ([]byte, []int) {
	return file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescGZIP(), []int{7}
}

func (x *NFTImageV2) GetGyroscope() []*GyroscopeEntityV2 {
	if x != nil {
		return x.Gyroscope
	}
	return nil
}

type PhysicalOrientationAnimation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Bezier string `protobuf:"bytes,3,opt,name=bezier,proto3" json:"bezier,omitempty"`
}

func (x *PhysicalOrientationAnimation) Reset() {
	*x = PhysicalOrientationAnimation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhysicalOrientationAnimation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhysicalOrientationAnimation) ProtoMessage() {}

func (x *PhysicalOrientationAnimation) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhysicalOrientationAnimation.ProtoReflect.Descriptor instead.
func (*PhysicalOrientationAnimation) Descriptor() ([]byte, []int) {
	return file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescGZIP(), []int{8}
}

func (x *PhysicalOrientationAnimation) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PhysicalOrientationAnimation) GetBezier() string {
	if x != nil {
		return x.Bezier
	}
	return ""
}

type PhysicalOrientationV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Animations []*PhysicalOrientationAnimation `protobuf:"bytes,3,rep,name=animations,proto3" json:"animations,omitempty"`
}

func (x *PhysicalOrientationV2) Reset() {
	*x = PhysicalOrientationV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhysicalOrientationV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhysicalOrientationV2) ProtoMessage() {}

func (x *PhysicalOrientationV2) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhysicalOrientationV2.ProtoReflect.Descriptor instead.
func (*PhysicalOrientationV2) Descriptor() ([]byte, []int) {
	return file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescGZIP(), []int{9}
}

func (x *PhysicalOrientationV2) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PhysicalOrientationV2) GetAnimations() []*PhysicalOrientationAnimation {
	if x != nil {
		return x.Animations
	}
	return nil
}

var File_bilibili_dagw_component_avatar_v1_plugin_proto protoreflect.FileDescriptor

var file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x64, 0x61, 0x67, 0x77, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x28, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x64, 0x61, 0x67, 0x77, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x32, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2f, 0x64, 0x61, 0x67, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x2f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c,
	0x01, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x43, 0x6c, 0x69, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x57, 0x0a, 0x0b, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x64, 0x61, 0x67, 0x77,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x61,
	0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x22, 0x60, 0x0a,
	0x0a, 0x47, 0x79, 0x72, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x52, 0x0a, 0x09, 0x67,
	0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x64, 0x61, 0x67, 0x77, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4e, 0x46, 0x54, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x56, 0x32, 0x52, 0x09, 0x67, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22,
	0xb9, 0x01, 0x0a, 0x12, 0x47, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x56, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x72, 0x0a, 0x14, 0x70, 0x68, 0x79, 0x73, 0x69,
	0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x64, 0x61, 0x67, 0x77, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x52, 0x13, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c,
	0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x90, 0x01, 0x0a, 0x11,
	0x47, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56,
	0x32, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x58, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x64, 0x61, 0x67, 0x77, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x47, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x56, 0x32, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x77,
	0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a,
	0x06, 0x6e, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e,
	0x66, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x55, 0x72, 0x6c, 0x22, 0x2a, 0x0a, 0x0f, 0x4c, 0x69, 0x76, 0x65, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73,
	0x5f, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4c,
	0x69, 0x76, 0x65, 0x22, 0xf5, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x76, 0x65, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x48, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x64, 0x61, 0x67, 0x77, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6f,
	0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x6f, 0x6b, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x22, 0x67, 0x0a, 0x0a, 0x4e,
	0x46, 0x54, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x56, 0x32, 0x12, 0x59, 0x0a, 0x09, 0x67, 0x79, 0x72,
	0x6f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x62,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x64, 0x61, 0x67, 0x77, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x32, 0x52, 0x09, 0x67, 0x79, 0x72, 0x6f, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x22, 0x4a, 0x0a, 0x1c, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c,
	0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x69, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x7a, 0x69,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x65, 0x7a, 0x69, 0x65, 0x72,
	0x22, 0x93, 0x01, 0x0a, 0x15, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x66,
	0x0a, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x46, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x64, 0x61,
	0x67, 0x77, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x68,
	0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x50, 0x5a, 0x4e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x6d,
	0x74, 0x63, 0x32, 0x33, 0x33, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f,
	0x64, 0x61, 0x67, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescOnce sync.Once
	file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescData = file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDesc
)

func file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescGZIP() []byte {
	file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescOnce.Do(func() {
		file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescData)
	})
	return file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDescData
}

var file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_bilibili_dagw_component_avatar_v1_plugin_proto_goTypes = []interface{}{
	(*CommentDoubleClickConfig)(nil),     // 0: bilibili.dagw.component.avatar.v1.plugin.CommentDoubleClickConfig
	(*GyroConfig)(nil),                   // 1: bilibili.dagw.component.avatar.v1.plugin.GyroConfig
	(*GyroscopeContentV2)(nil),           // 2: bilibili.dagw.component.avatar.v1.plugin.GyroscopeContentV2
	(*GyroscopeEntityV2)(nil),            // 3: bilibili.dagw.component.avatar.v1.plugin.GyroscopeEntityV2
	(*Interaction)(nil),                  // 4: bilibili.dagw.component.avatar.v1.plugin.Interaction
	(*LiveAnimeConfig)(nil),              // 5: bilibili.dagw.component.avatar.v1.plugin.LiveAnimeConfig
	(*LiveAnimeItem)(nil),                // 6: bilibili.dagw.component.avatar.v1.plugin.LiveAnimeItem
	(*NFTImageV2)(nil),                   // 7: bilibili.dagw.component.avatar.v1.plugin.NFTImageV2
	(*PhysicalOrientationAnimation)(nil), // 8: bilibili.dagw.component.avatar.v1.plugin.PhysicalOrientationAnimation
	(*PhysicalOrientationV2)(nil),        // 9: bilibili.dagw.component.avatar.v1.plugin.PhysicalOrientationV2
	(*common.ColorConfig)(nil),           // 10: bilibili.dagw.component.avatar.common.ColorConfig
}
var file_bilibili_dagw_component_avatar_v1_plugin_proto_depIdxs = []int32{
	4,  // 0: bilibili.dagw.component.avatar.v1.plugin.CommentDoubleClickConfig.interaction:type_name -> bilibili.dagw.component.avatar.v1.plugin.Interaction
	7,  // 1: bilibili.dagw.component.avatar.v1.plugin.GyroConfig.gyroscope:type_name -> bilibili.dagw.component.avatar.v1.plugin.NFTImageV2
	9,  // 2: bilibili.dagw.component.avatar.v1.plugin.GyroscopeContentV2.physical_orientation:type_name -> bilibili.dagw.component.avatar.v1.plugin.PhysicalOrientationV2
	2,  // 3: bilibili.dagw.component.avatar.v1.plugin.GyroscopeEntityV2.contents:type_name -> bilibili.dagw.component.avatar.v1.plugin.GyroscopeContentV2
	10, // 4: bilibili.dagw.component.avatar.v1.plugin.LiveAnimeItem.color:type_name -> bilibili.dagw.component.avatar.common.ColorConfig
	3,  // 5: bilibili.dagw.component.avatar.v1.plugin.NFTImageV2.gyroscope:type_name -> bilibili.dagw.component.avatar.v1.plugin.GyroscopeEntityV2
	8,  // 6: bilibili.dagw.component.avatar.v1.plugin.PhysicalOrientationV2.animations:type_name -> bilibili.dagw.component.avatar.v1.plugin.PhysicalOrientationAnimation
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_bilibili_dagw_component_avatar_v1_plugin_proto_init() }
func file_bilibili_dagw_component_avatar_v1_plugin_proto_init() {
	if File_bilibili_dagw_component_avatar_v1_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentDoubleClickConfig); i {
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
		file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GyroConfig); i {
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
		file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GyroscopeContentV2); i {
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
		file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GyroscopeEntityV2); i {
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
		file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interaction); i {
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
		file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveAnimeConfig); i {
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
		file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveAnimeItem); i {
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
		file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NFTImageV2); i {
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
		file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhysicalOrientationAnimation); i {
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
		file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhysicalOrientationV2); i {
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
			RawDescriptor: file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_dagw_component_avatar_v1_plugin_proto_goTypes,
		DependencyIndexes: file_bilibili_dagw_component_avatar_v1_plugin_proto_depIdxs,
		MessageInfos:      file_bilibili_dagw_component_avatar_v1_plugin_proto_msgTypes,
	}.Build()
	File_bilibili_dagw_component_avatar_v1_plugin_proto = out.File
	file_bilibili_dagw_component_avatar_v1_plugin_proto_rawDesc = nil
	file_bilibili_dagw_component_avatar_v1_plugin_proto_goTypes = nil
	file_bilibili_dagw_component_avatar_v1_plugin_proto_depIdxs = nil
}