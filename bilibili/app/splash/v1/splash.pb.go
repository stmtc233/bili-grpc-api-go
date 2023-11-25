// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: bilibili/app/splash/v1/splash.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ShowStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Stime int64 `protobuf:"varint,2,opt,name=stime,proto3" json:"stime,omitempty"`
	Etime int64 `protobuf:"varint,3,opt,name=etime,proto3" json:"etime,omitempty"`
}

func (x *ShowStrategy) Reset() {
	*x = ShowStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_splash_v1_splash_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowStrategy) ProtoMessage() {}

func (x *ShowStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_splash_v1_splash_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowStrategy.ProtoReflect.Descriptor instead.
func (*ShowStrategy) Descriptor() ([]byte, []int) {
	return file_bilibili_app_splash_v1_splash_proto_rawDescGZIP(), []int{0}
}

func (x *ShowStrategy) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ShowStrategy) GetStime() int64 {
	if x != nil {
		return x.Stime
	}
	return 0
}

func (x *ShowStrategy) GetEtime() int64 {
	if x != nil {
		return x.Etime
	}
	return 0
}

type SplashItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type int32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	CardType int32 `protobuf:"varint,3,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	Duration int32 `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	BeginTime int64 `protobuf:"varint,5,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	EndTime int64 `protobuf:"varint,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Thumb string `protobuf:"bytes,7,opt,name=thumb,proto3" json:"thumb,omitempty"`
	Hash string `protobuf:"bytes,8,opt,name=hash,proto3" json:"hash,omitempty"`
	LogoUrl string `protobuf:"bytes,9,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	LogoHash string `protobuf:"bytes,10,opt,name=logo_hash,json=logoHash,proto3" json:"logo_hash,omitempty"`
	VideoUrl string `protobuf:"bytes,11,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
	VideoHash string `protobuf:"bytes,12,opt,name=video_hash,json=videoHash,proto3" json:"video_hash,omitempty"`
	VideoWidth int32 `protobuf:"varint,13,opt,name=video_width,json=videoWidth,proto3" json:"video_width,omitempty"`
	VideoHeight int32 `protobuf:"varint,14,opt,name=video_height,json=videoHeight,proto3" json:"video_height,omitempty"`
	Schema string `protobuf:"bytes,15,opt,name=schema,proto3" json:"schema,omitempty"`
	SchemaTitle string `protobuf:"bytes,16,opt,name=schema_title,json=schemaTitle,proto3" json:"schema_title,omitempty"`
	SchemaPackageName string `protobuf:"bytes,17,opt,name=schema_package_name,json=schemaPackageName,proto3" json:"schema_package_name,omitempty"`
	SchemaCallupWhiteList []string `protobuf:"bytes,18,rep,name=schema_callup_whiteList,json=schemaCallupWhiteList,proto3" json:"schema_callup_whiteList,omitempty"`
	Skip int32 `protobuf:"varint,19,opt,name=skip,proto3" json:"skip,omitempty"`
	Uri string `protobuf:"bytes,20,opt,name=uri,proto3" json:"uri,omitempty"`
	UriTitle string `protobuf:"bytes,21,opt,name=uri_title,json=uriTitle,proto3" json:"uri_title,omitempty"`
	Source int32 `protobuf:"varint,22,opt,name=source,proto3" json:"source,omitempty"`
	CmMark int32 `protobuf:"varint,23,opt,name=cm_mark,json=cmMark,proto3" json:"cm_mark,omitempty"`
	AdCb string `protobuf:"bytes,24,opt,name=ad_cb,json=adCb,proto3" json:"ad_cb,omitempty"`
	ResourceId int64 `protobuf:"varint,25,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	RequestId string `protobuf:"bytes,26,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	ClientIp string `protobuf:"bytes,27,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	IsAd bool `protobuf:"varint,28,opt,name=is_ad,json=isAd,proto3" json:"is_ad,omitempty"`
	IsAdLoc bool `protobuf:"varint,29,opt,name=is_ad_loc,json=isAdLoc,proto3" json:"is_ad_loc,omitempty"`
	Extra *anypb.Any `protobuf:"bytes,30,opt,name=extra,proto3" json:"extra,omitempty"`
	CardIndex int64 `protobuf:"varint,31,opt,name=card_index,json=cardIndex,proto3" json:"card_index,omitempty"`
	ServerType int64 `protobuf:"varint,32,opt,name=server_type,json=serverType,proto3" json:"server_type,omitempty"`
	Index int64 `protobuf:"varint,33,opt,name=index,proto3" json:"index,omitempty"`
	ClickUrl string `protobuf:"bytes,34,opt,name=click_url,json=clickUrl,proto3" json:"click_url,omitempty"`
	ShowUrl string `protobuf:"bytes,35,opt,name=show_url,json=showUrl,proto3" json:"show_url,omitempty"`
	TimeTarget int32 `protobuf:"varint,36,opt,name=time_target,json=timeTarget,proto3" json:"time_target,omitempty"`
	Encryption int32 `protobuf:"varint,37,opt,name=encryption,proto3" json:"encryption,omitempty"`
	EnablePreDownload bool `protobuf:"varint,38,opt,name=enable_pre_download,json=enablePreDownload,proto3" json:"enable_pre_download,omitempty"`
	EnableBackgroundDownload bool `protobuf:"varint,39,opt,name=enable_background_download,json=enableBackgroundDownload,proto3" json:"enable_background_download,omitempty"`
}

func (x *SplashItem) Reset() {
	*x = SplashItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_splash_v1_splash_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplashItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplashItem) ProtoMessage() {}

func (x *SplashItem) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_splash_v1_splash_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplashItem.ProtoReflect.Descriptor instead.
func (*SplashItem) Descriptor() ([]byte, []int) {
	return file_bilibili_app_splash_v1_splash_proto_rawDescGZIP(), []int{1}
}

func (x *SplashItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SplashItem) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *SplashItem) GetCardType() int32 {
	if x != nil {
		return x.CardType
	}
	return 0
}

func (x *SplashItem) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *SplashItem) GetBeginTime() int64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *SplashItem) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *SplashItem) GetThumb() string {
	if x != nil {
		return x.Thumb
	}
	return ""
}

func (x *SplashItem) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *SplashItem) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *SplashItem) GetLogoHash() string {
	if x != nil {
		return x.LogoHash
	}
	return ""
}

func (x *SplashItem) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *SplashItem) GetVideoHash() string {
	if x != nil {
		return x.VideoHash
	}
	return ""
}

func (x *SplashItem) GetVideoWidth() int32 {
	if x != nil {
		return x.VideoWidth
	}
	return 0
}

func (x *SplashItem) GetVideoHeight() int32 {
	if x != nil {
		return x.VideoHeight
	}
	return 0
}

func (x *SplashItem) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *SplashItem) GetSchemaTitle() string {
	if x != nil {
		return x.SchemaTitle
	}
	return ""
}

func (x *SplashItem) GetSchemaPackageName() string {
	if x != nil {
		return x.SchemaPackageName
	}
	return ""
}

func (x *SplashItem) GetSchemaCallupWhiteList() []string {
	if x != nil {
		return x.SchemaCallupWhiteList
	}
	return nil
}

func (x *SplashItem) GetSkip() int32 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *SplashItem) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *SplashItem) GetUriTitle() string {
	if x != nil {
		return x.UriTitle
	}
	return ""
}

func (x *SplashItem) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *SplashItem) GetCmMark() int32 {
	if x != nil {
		return x.CmMark
	}
	return 0
}

func (x *SplashItem) GetAdCb() string {
	if x != nil {
		return x.AdCb
	}
	return ""
}

func (x *SplashItem) GetResourceId() int64 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

func (x *SplashItem) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *SplashItem) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *SplashItem) GetIsAd() bool {
	if x != nil {
		return x.IsAd
	}
	return false
}

func (x *SplashItem) GetIsAdLoc() bool {
	if x != nil {
		return x.IsAdLoc
	}
	return false
}

func (x *SplashItem) GetExtra() *anypb.Any {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *SplashItem) GetCardIndex() int64 {
	if x != nil {
		return x.CardIndex
	}
	return 0
}

func (x *SplashItem) GetServerType() int64 {
	if x != nil {
		return x.ServerType
	}
	return 0
}

func (x *SplashItem) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *SplashItem) GetClickUrl() string {
	if x != nil {
		return x.ClickUrl
	}
	return ""
}

func (x *SplashItem) GetShowUrl() string {
	if x != nil {
		return x.ShowUrl
	}
	return ""
}

func (x *SplashItem) GetTimeTarget() int32 {
	if x != nil {
		return x.TimeTarget
	}
	return 0
}

func (x *SplashItem) GetEncryption() int32 {
	if x != nil {
		return x.Encryption
	}
	return 0
}

func (x *SplashItem) GetEnablePreDownload() bool {
	if x != nil {
		return x.EnablePreDownload
	}
	return false
}

func (x *SplashItem) GetEnableBackgroundDownload() bool {
	if x != nil {
		return x.EnableBackgroundDownload
	}
	return false
}

// -响应
type SplashReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxTime int32 `protobuf:"varint,1,opt,name=max_time,json=maxTime,proto3" json:"max_time,omitempty"`
	MinInterval int32 `protobuf:"varint,2,opt,name=min_interval,json=minInterval,proto3" json:"min_interval,omitempty"`
	PullInterval int32 `protobuf:"varint,3,opt,name=pull_interval,json=pullInterval,proto3" json:"pull_interval,omitempty"`
	List []*SplashItem `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	Show []*ShowStrategy `protobuf:"bytes,5,rep,name=show,proto3" json:"show,omitempty"`
}

func (x *SplashReply) Reset() {
	*x = SplashReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_splash_v1_splash_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplashReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplashReply) ProtoMessage() {}

func (x *SplashReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_splash_v1_splash_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplashReply.ProtoReflect.Descriptor instead.
func (*SplashReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_splash_v1_splash_proto_rawDescGZIP(), []int{2}
}

func (x *SplashReply) GetMaxTime() int32 {
	if x != nil {
		return x.MaxTime
	}
	return 0
}

func (x *SplashReply) GetMinInterval() int32 {
	if x != nil {
		return x.MinInterval
	}
	return 0
}

func (x *SplashReply) GetPullInterval() int32 {
	if x != nil {
		return x.PullInterval
	}
	return 0
}

func (x *SplashReply) GetList() []*SplashItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *SplashReply) GetShow() []*ShowStrategy {
	if x != nil {
		return x.Show
	}
	return nil
}

// -请求
type SplashReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width int32 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height int32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Birth string `protobuf:"bytes,3,opt,name=birth,proto3" json:"birth,omitempty"`
	AdExtra string `protobuf:"bytes,4,opt,name=ad_extra,json=adExtra,proto3" json:"ad_extra,omitempty"`
	Network string `protobuf:"bytes,5,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *SplashReq) Reset() {
	*x = SplashReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_splash_v1_splash_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplashReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplashReq) ProtoMessage() {}

func (x *SplashReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_splash_v1_splash_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplashReq.ProtoReflect.Descriptor instead.
func (*SplashReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_splash_v1_splash_proto_rawDescGZIP(), []int{3}
}

func (x *SplashReq) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *SplashReq) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *SplashReq) GetBirth() string {
	if x != nil {
		return x.Birth
	}
	return ""
}

func (x *SplashReq) GetAdExtra() string {
	if x != nil {
		return x.AdExtra
	}
	return ""
}

func (x *SplashReq) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

var File_bilibili_app_splash_v1_splash_proto protoreflect.FileDescriptor

var file_bilibili_app_splash_v1_splash_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x70, 0x6c, 0x61, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x6c, 0x61, 0x73, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x73, 0x70, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x77,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0xa8, 0x09, 0x0a, 0x0a, 0x53, 0x70, 0x6c, 0x61, 0x73, 0x68, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x72, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x75, 0x70, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x43, 0x61, 0x6c, 0x6c, 0x75, 0x70, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x6b, 0x69, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x72, 0x69, 0x5f, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x72, 0x69, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6d,
	0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6d, 0x4d,
	0x61, 0x72, 0x6b, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x64, 0x5f, 0x63, 0x62, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x43, 0x62, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x18, 0x1c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x41, 0x64, 0x12, 0x1a, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x61, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x41, 0x64, 0x4c, 0x6f, 0x63, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x20, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x21, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x63,
	0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x63, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x77, 0x55, 0x72, 0x6c,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x24, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x25, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x5f,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x26, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x3c, 0x0a, 0x1a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x27, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0xe2, 0x01, 0x0a, 0x0b, 0x53, 0x70, 0x6c, 0x61, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69,
	0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x12, 0x36, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x73, 0x70, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x6c, 0x61, 0x73, 0x68,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x73, 0x68,
	0x6f, 0x77, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x70, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x04,
	0x73, 0x68, 0x6f, 0x77, 0x22, 0x84, 0x01, 0x0a, 0x09, 0x53, 0x70, 0x6c, 0x61, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x5f, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x32, 0x58, 0x0a, 0x06, 0x53,
	0x70, 0x6c, 0x61, 0x73, 0x68, 0x12, 0x4e, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x70, 0x6c,
	0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x6c, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x1a, 0x23, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x73, 0x70, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x6c, 0x61, 0x73, 0x68,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x6d, 0x74, 0x63, 0x32, 0x33, 0x33, 0x2f, 0x62, 0x69, 0x6c,
	0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x70, 0x6c, 0x61, 0x73,
	0x68, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_splash_v1_splash_proto_rawDescOnce sync.Once
	file_bilibili_app_splash_v1_splash_proto_rawDescData = file_bilibili_app_splash_v1_splash_proto_rawDesc
)

func file_bilibili_app_splash_v1_splash_proto_rawDescGZIP() []byte {
	file_bilibili_app_splash_v1_splash_proto_rawDescOnce.Do(func() {
		file_bilibili_app_splash_v1_splash_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_splash_v1_splash_proto_rawDescData)
	})
	return file_bilibili_app_splash_v1_splash_proto_rawDescData
}

var file_bilibili_app_splash_v1_splash_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bilibili_app_splash_v1_splash_proto_goTypes = []interface{}{
	(*ShowStrategy)(nil), // 0: bilibili.app.splash.v1.ShowStrategy
	(*SplashItem)(nil),   // 1: bilibili.app.splash.v1.SplashItem
	(*SplashReply)(nil),  // 2: bilibili.app.splash.v1.SplashReply
	(*SplashReq)(nil),    // 3: bilibili.app.splash.v1.SplashReq
	(*anypb.Any)(nil),    // 4: google.protobuf.Any
}
var file_bilibili_app_splash_v1_splash_proto_depIdxs = []int32{
	4, // 0: bilibili.app.splash.v1.SplashItem.extra:type_name -> google.protobuf.Any
	1, // 1: bilibili.app.splash.v1.SplashReply.list:type_name -> bilibili.app.splash.v1.SplashItem
	0, // 2: bilibili.app.splash.v1.SplashReply.show:type_name -> bilibili.app.splash.v1.ShowStrategy
	3, // 3: bilibili.app.splash.v1.Splash.List:input_type -> bilibili.app.splash.v1.SplashReq
	2, // 4: bilibili.app.splash.v1.Splash.List:output_type -> bilibili.app.splash.v1.SplashReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bilibili_app_splash_v1_splash_proto_init() }
func file_bilibili_app_splash_v1_splash_proto_init() {
	if File_bilibili_app_splash_v1_splash_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_splash_v1_splash_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowStrategy); i {
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
		file_bilibili_app_splash_v1_splash_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplashItem); i {
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
		file_bilibili_app_splash_v1_splash_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplashReply); i {
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
		file_bilibili_app_splash_v1_splash_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplashReq); i {
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
			RawDescriptor: file_bilibili_app_splash_v1_splash_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_app_splash_v1_splash_proto_goTypes,
		DependencyIndexes: file_bilibili_app_splash_v1_splash_proto_depIdxs,
		MessageInfos:      file_bilibili_app_splash_v1_splash_proto_msgTypes,
	}.Build()
	File_bilibili_app_splash_v1_splash_proto = out.File
	file_bilibili_app_splash_v1_splash_proto_rawDesc = nil
	file_bilibili_app_splash_v1_splash_proto_goTypes = nil
	file_bilibili_app_splash_v1_splash_proto_depIdxs = nil
}
