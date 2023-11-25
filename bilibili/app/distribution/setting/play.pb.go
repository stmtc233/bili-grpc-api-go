// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: bilibili/app/distribution/setting/play.proto

package setting

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "https://github.com/stmtc233/bili-grpc-api-go/bilibili/app/distribution/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 云端保存的播放器配置
type CloudPlayConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 启用杜比全景声
	EnablePanorama *v1.BoolValue `protobuf:"bytes,1,opt,name=enable_panorama,json=enablePanorama,proto3" json:"enable_panorama,omitempty"`
	// 启用杜比音效
	EnableDolby *v1.BoolValue `protobuf:"bytes,2,opt,name=enable_dolby,json=enableDolby,proto3" json:"enable_dolby,omitempty"`
	// 启用震动
	EnableShake *v1.BoolValue `protobuf:"bytes,3,opt,name=enable_shake,json=enableShake,proto3" json:"enable_shake,omitempty"`
	// 启用后台播放
	EnableBackground *v1.BoolValue `protobuf:"bytes,4,opt,name=enable_background,json=enableBackground,proto3" json:"enable_background,omitempty"`
	// 启用HIRES
	EnableLossLess *v1.BoolValue `protobuf:"bytes,5,opt,name=enable_loss_less,json=enableLossLess,proto3" json:"enable_loss_less,omitempty"`
}

func (x *CloudPlayConfig) Reset() {
	*x = CloudPlayConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_distribution_setting_play_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudPlayConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudPlayConfig) ProtoMessage() {}

func (x *CloudPlayConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_distribution_setting_play_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudPlayConfig.ProtoReflect.Descriptor instead.
func (*CloudPlayConfig) Descriptor() ([]byte, []int) {
	return file_bilibili_app_distribution_setting_play_proto_rawDescGZIP(), []int{0}
}

func (x *CloudPlayConfig) GetEnablePanorama() *v1.BoolValue {
	if x != nil {
		return x.EnablePanorama
	}
	return nil
}

func (x *CloudPlayConfig) GetEnableDolby() *v1.BoolValue {
	if x != nil {
		return x.EnableDolby
	}
	return nil
}

func (x *CloudPlayConfig) GetEnableShake() *v1.BoolValue {
	if x != nil {
		return x.EnableShake
	}
	return nil
}

func (x *CloudPlayConfig) GetEnableBackground() *v1.BoolValue {
	if x != nil {
		return x.EnableBackground
	}
	return nil
}

func (x *CloudPlayConfig) GetEnableLossLess() *v1.BoolValue {
	if x != nil {
		return x.EnableLossLess
	}
	return nil
}

// 播放器策略配置
type PlayConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	ShouldAutoPlay *v1.BoolValue `protobuf:"bytes,1,opt,name=should_auto_play,json=shouldAutoPlay,proto3" json:"should_auto_play,omitempty"`
	//
	ShouldAutoFullscreen *v1.BoolValue `protobuf:"bytes,2,opt,name=should_auto_fullscreen,json=shouldAutoFullscreen,proto3" json:"should_auto_fullscreen,omitempty"`
	//
	EnablePlayurlHttps *v1.BoolValue `protobuf:"bytes,3,opt,name=enable_playurl_https,json=enablePlayurlHttps,proto3" json:"enable_playurl_https,omitempty"`
	//
	EnableDanmakuInteraction *v1.BoolValue `protobuf:"bytes,4,opt,name=enable_danmaku_interaction,json=enableDanmakuInteraction,proto3" json:"enable_danmaku_interaction,omitempty"`
	//
	SmallScreenStatus *v1.Int64Value `protobuf:"bytes,5,opt,name=small_screen_status,json=smallScreenStatus,proto3" json:"small_screen_status,omitempty"`
	//
	PlayerCodecModeKey *v1.Int64Value `protobuf:"bytes,6,opt,name=player_codec_mode_key,json=playerCodecModeKey,proto3" json:"player_codec_mode_key,omitempty"`
	//
	EnableGravityRotateScreen *v1.BoolValue `protobuf:"bytes,7,opt,name=enable_gravity_rotate_screen,json=enableGravityRotateScreen,proto3" json:"enable_gravity_rotate_screen,omitempty"`
	//
	EnableDanmakuMonospaced *v1.BoolValue `protobuf:"bytes,8,opt,name=enable_danmaku_monospaced,json=enableDanmakuMonospaced,proto3" json:"enable_danmaku_monospaced,omitempty"`
	//
	EnableEditSubtitle *v1.BoolValue `protobuf:"bytes,9,opt,name=enable_edit_subtitle,json=enableEditSubtitle,proto3" json:"enable_edit_subtitle,omitempty"`
	//
	EnableSubtitle *v1.BoolValue `protobuf:"bytes,10,opt,name=enable_subtitle,json=enableSubtitle,proto3" json:"enable_subtitle,omitempty"`
	//
	ColorFilter *v1.Int64Value `protobuf:"bytes,11,opt,name=color_filter,json=colorFilter,proto3" json:"color_filter,omitempty"`
	//
	ShouldAutoStory *v1.BoolValue `protobuf:"bytes,12,opt,name=should_auto_story,json=shouldAutoStory,proto3" json:"should_auto_story,omitempty"`
	//
	LandscapeAutoStory *v1.BoolValue `protobuf:"bytes,13,opt,name=landscape_auto_story,json=landscapeAutoStory,proto3" json:"landscape_auto_story,omitempty"`
	//
	VolumeBalance *v1.BoolValue `protobuf:"bytes,14,opt,name=volume_balance,json=volumeBalance,proto3" json:"volume_balance,omitempty"`
}

func (x *PlayConfig) Reset() {
	*x = PlayConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_distribution_setting_play_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayConfig) ProtoMessage() {}

func (x *PlayConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_distribution_setting_play_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayConfig.ProtoReflect.Descriptor instead.
func (*PlayConfig) Descriptor() ([]byte, []int) {
	return file_bilibili_app_distribution_setting_play_proto_rawDescGZIP(), []int{1}
}

func (x *PlayConfig) GetShouldAutoPlay() *v1.BoolValue {
	if x != nil {
		return x.ShouldAutoPlay
	}
	return nil
}

func (x *PlayConfig) GetShouldAutoFullscreen() *v1.BoolValue {
	if x != nil {
		return x.ShouldAutoFullscreen
	}
	return nil
}

func (x *PlayConfig) GetEnablePlayurlHttps() *v1.BoolValue {
	if x != nil {
		return x.EnablePlayurlHttps
	}
	return nil
}

func (x *PlayConfig) GetEnableDanmakuInteraction() *v1.BoolValue {
	if x != nil {
		return x.EnableDanmakuInteraction
	}
	return nil
}

func (x *PlayConfig) GetSmallScreenStatus() *v1.Int64Value {
	if x != nil {
		return x.SmallScreenStatus
	}
	return nil
}

func (x *PlayConfig) GetPlayerCodecModeKey() *v1.Int64Value {
	if x != nil {
		return x.PlayerCodecModeKey
	}
	return nil
}

func (x *PlayConfig) GetEnableGravityRotateScreen() *v1.BoolValue {
	if x != nil {
		return x.EnableGravityRotateScreen
	}
	return nil
}

func (x *PlayConfig) GetEnableDanmakuMonospaced() *v1.BoolValue {
	if x != nil {
		return x.EnableDanmakuMonospaced
	}
	return nil
}

func (x *PlayConfig) GetEnableEditSubtitle() *v1.BoolValue {
	if x != nil {
		return x.EnableEditSubtitle
	}
	return nil
}

func (x *PlayConfig) GetEnableSubtitle() *v1.BoolValue {
	if x != nil {
		return x.EnableSubtitle
	}
	return nil
}

func (x *PlayConfig) GetColorFilter() *v1.Int64Value {
	if x != nil {
		return x.ColorFilter
	}
	return nil
}

func (x *PlayConfig) GetShouldAutoStory() *v1.BoolValue {
	if x != nil {
		return x.ShouldAutoStory
	}
	return nil
}

func (x *PlayConfig) GetLandscapeAutoStory() *v1.BoolValue {
	if x != nil {
		return x.LandscapeAutoStory
	}
	return nil
}

func (x *PlayConfig) GetVolumeBalance() *v1.BoolValue {
	if x != nil {
		return x.VolumeBalance
	}
	return nil
}

// 灰度测试特殊功能？
type SpecificPlayConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	EnableSegmentedSection *v1.BoolValue `protobuf:"bytes,1,opt,name=enable_segmented_section,json=enableSegmentedSection,proto3" json:"enable_segmented_section,omitempty"`
}

func (x *SpecificPlayConfig) Reset() {
	*x = SpecificPlayConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_distribution_setting_play_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecificPlayConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecificPlayConfig) ProtoMessage() {}

func (x *SpecificPlayConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_distribution_setting_play_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecificPlayConfig.ProtoReflect.Descriptor instead.
func (*SpecificPlayConfig) Descriptor() ([]byte, []int) {
	return file_bilibili_app_distribution_setting_play_proto_rawDescGZIP(), []int{2}
}

func (x *SpecificPlayConfig) GetEnableSegmentedSection() *v1.BoolValue {
	if x != nil {
		return x.EnableSegmentedSection
	}
	return nil
}

var File_bilibili_app_distribution_setting_play_proto protoreflect.FileDescriptor

var file_bilibili_app_distribution_setting_play_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x1a, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x03, 0x0a, 0x0f, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x50, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x0f, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x6e, 0x6f, 0x72, 0x61, 0x6d, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x6e, 0x6f, 0x72, 0x61, 0x6d, 0x61, 0x12, 0x4a, 0x0a,
	0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x6f, 0x6c, 0x62, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x44, 0x6f, 0x6c, 0x62, 0x79, 0x12, 0x4a, 0x0a, 0x0c, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x54, 0x0a, 0x11, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x51, 0x0a, 0x10, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x5f, 0x6c, 0x65, 0x73, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x6f, 0x73, 0x73, 0x4c, 0x65, 0x73, 0x73, 0x22, 0x80,
	0x0a, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x51, 0x0a,
	0x10, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x70, 0x6c, 0x61,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0e, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x41, 0x75, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79,
	0x12, 0x5d, 0x0a, 0x16, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f,
	0x66, 0x75, 0x6c, 0x6c, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x73, 0x68, 0x6f, 0x75, 0x6c,
	0x64, 0x41, 0x75, 0x74, 0x6f, 0x46, 0x75, 0x6c, 0x6c, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12,
	0x59, 0x0a, 0x14, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x75, 0x72,
	0x6c, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x75, 0x72, 0x6c, 0x48, 0x74, 0x74, 0x70, 0x73, 0x12, 0x65, 0x0a, 0x1a, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x6e, 0x6d, 0x61, 0x6b, 0x75, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x44,
	0x61, 0x6e, 0x6d, 0x61, 0x6b, 0x75, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x58, 0x0a, 0x13, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5b, 0x0a, 0x15, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x63, 0x4d, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x68, 0x0a, 0x1c, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74,
	0x65, 0x5f, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x19, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x47,
	0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x53, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x12, 0x63, 0x0a, 0x19, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x6e,
	0x6d, 0x61, 0x6b, 0x75, 0x5f, 0x6d, 0x6f, 0x6e, 0x6f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x17,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x6e, 0x6d, 0x61, 0x6b, 0x75, 0x4d, 0x6f, 0x6e,
	0x6f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x12, 0x59, 0x0a, 0x14, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x64, 0x69, 0x74, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x75, 0x62, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x53, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x41, 0x75, 0x74,
	0x6f, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x59, 0x0a, 0x14, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x63,
	0x61, 0x70, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x6c,
	0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x4e, 0x0a, 0x0e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0d, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x22, 0x77, 0x0a, 0x12, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x50, 0x6c, 0x61,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x61, 0x0a, 0x18, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x16, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x65, 0x64, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x50, 0x5a, 0x4e, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x6d, 0x74, 0x63, 0x32, 0x33, 0x33, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_distribution_setting_play_proto_rawDescOnce sync.Once
	file_bilibili_app_distribution_setting_play_proto_rawDescData = file_bilibili_app_distribution_setting_play_proto_rawDesc
)

func file_bilibili_app_distribution_setting_play_proto_rawDescGZIP() []byte {
	file_bilibili_app_distribution_setting_play_proto_rawDescOnce.Do(func() {
		file_bilibili_app_distribution_setting_play_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_distribution_setting_play_proto_rawDescData)
	})
	return file_bilibili_app_distribution_setting_play_proto_rawDescData
}

var file_bilibili_app_distribution_setting_play_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bilibili_app_distribution_setting_play_proto_goTypes = []interface{}{
	(*CloudPlayConfig)(nil),    // 0: bilibili.app.distribution.setting.play.CloudPlayConfig
	(*PlayConfig)(nil),         // 1: bilibili.app.distribution.setting.play.PlayConfig
	(*SpecificPlayConfig)(nil), // 2: bilibili.app.distribution.setting.play.SpecificPlayConfig
	(*v1.BoolValue)(nil),       // 3: bilibili.app.distribution.v1.BoolValue
	(*v1.Int64Value)(nil),      // 4: bilibili.app.distribution.v1.Int64Value
}
var file_bilibili_app_distribution_setting_play_proto_depIdxs = []int32{
	3,  // 0: bilibili.app.distribution.setting.play.CloudPlayConfig.enable_panorama:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 1: bilibili.app.distribution.setting.play.CloudPlayConfig.enable_dolby:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 2: bilibili.app.distribution.setting.play.CloudPlayConfig.enable_shake:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 3: bilibili.app.distribution.setting.play.CloudPlayConfig.enable_background:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 4: bilibili.app.distribution.setting.play.CloudPlayConfig.enable_loss_less:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 5: bilibili.app.distribution.setting.play.PlayConfig.should_auto_play:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 6: bilibili.app.distribution.setting.play.PlayConfig.should_auto_fullscreen:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 7: bilibili.app.distribution.setting.play.PlayConfig.enable_playurl_https:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 8: bilibili.app.distribution.setting.play.PlayConfig.enable_danmaku_interaction:type_name -> bilibili.app.distribution.v1.BoolValue
	4,  // 9: bilibili.app.distribution.setting.play.PlayConfig.small_screen_status:type_name -> bilibili.app.distribution.v1.Int64Value
	4,  // 10: bilibili.app.distribution.setting.play.PlayConfig.player_codec_mode_key:type_name -> bilibili.app.distribution.v1.Int64Value
	3,  // 11: bilibili.app.distribution.setting.play.PlayConfig.enable_gravity_rotate_screen:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 12: bilibili.app.distribution.setting.play.PlayConfig.enable_danmaku_monospaced:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 13: bilibili.app.distribution.setting.play.PlayConfig.enable_edit_subtitle:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 14: bilibili.app.distribution.setting.play.PlayConfig.enable_subtitle:type_name -> bilibili.app.distribution.v1.BoolValue
	4,  // 15: bilibili.app.distribution.setting.play.PlayConfig.color_filter:type_name -> bilibili.app.distribution.v1.Int64Value
	3,  // 16: bilibili.app.distribution.setting.play.PlayConfig.should_auto_story:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 17: bilibili.app.distribution.setting.play.PlayConfig.landscape_auto_story:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 18: bilibili.app.distribution.setting.play.PlayConfig.volume_balance:type_name -> bilibili.app.distribution.v1.BoolValue
	3,  // 19: bilibili.app.distribution.setting.play.SpecificPlayConfig.enable_segmented_section:type_name -> bilibili.app.distribution.v1.BoolValue
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_bilibili_app_distribution_setting_play_proto_init() }
func file_bilibili_app_distribution_setting_play_proto_init() {
	if File_bilibili_app_distribution_setting_play_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_distribution_setting_play_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudPlayConfig); i {
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
		file_bilibili_app_distribution_setting_play_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayConfig); i {
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
		file_bilibili_app_distribution_setting_play_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecificPlayConfig); i {
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
			RawDescriptor: file_bilibili_app_distribution_setting_play_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_app_distribution_setting_play_proto_goTypes,
		DependencyIndexes: file_bilibili_app_distribution_setting_play_proto_depIdxs,
		MessageInfos:      file_bilibili_app_distribution_setting_play_proto_msgTypes,
	}.Build()
	File_bilibili_app_distribution_setting_play_proto = out.File
	file_bilibili_app_distribution_setting_play_proto_rawDesc = nil
	file_bilibili_app_distribution_setting_play_proto_goTypes = nil
	file_bilibili_app_distribution_setting_play_proto_depIdxs = nil
}
