// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: bilibili/app/search/v2/search.proto

package v2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	main "https://github.com/stmtc233/bili-grpc-api-go/bilibili/broadcast/message/main"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CancelChatTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	FromSource string `protobuf:"bytes,2,opt,name=from_source,json=fromSource,proto3" json:"from_source,omitempty"`
}

func (x *CancelChatTaskReq) Reset() {
	*x = CancelChatTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_search_v2_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelChatTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelChatTaskReq) ProtoMessage() {}

func (x *CancelChatTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_search_v2_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelChatTaskReq.ProtoReflect.Descriptor instead.
func (*CancelChatTaskReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_search_v2_search_proto_rawDescGZIP(), []int{0}
}

func (x *CancelChatTaskReq) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *CancelChatTaskReq) GetFromSource() string {
	if x != nil {
		return x.FromSource
	}
	return ""
}

type CancelChatTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CancelChatTaskReply) Reset() {
	*x = CancelChatTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_search_v2_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelChatTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelChatTaskReply) ProtoMessage() {}

func (x *CancelChatTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_search_v2_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelChatTaskReply.ProtoReflect.Descriptor instead.
func (*CancelChatTaskReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_search_v2_search_proto_rawDescGZIP(), []int{1}
}

func (x *CancelChatTaskReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetChatAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetChatAuthReq) Reset() {
	*x = GetChatAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_search_v2_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatAuthReq) ProtoMessage() {}

func (x *GetChatAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_search_v2_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatAuthReq.ProtoReflect.Descriptor instead.
func (*GetChatAuthReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_search_v2_search_proto_rawDescGZIP(), []int{2}
}

type GetChatAuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Display bool `protobuf:"varint,1,opt,name=display,proto3" json:"display,omitempty"`
	Icon string `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
	IconNight string `protobuf:"bytes,3,opt,name=icon_night,json=iconNight,proto3" json:"icon_night,omitempty"`
	JumpLink string `protobuf:"bytes,4,opt,name=jump_link,json=jumpLink,proto3" json:"jump_link,omitempty"`
	TextGuide string `protobuf:"bytes,5,opt,name=text_guide,json=textGuide,proto3" json:"text_guide,omitempty"`
	JumpLinkType int32 `protobuf:"varint,6,opt,name=jump_link_type,json=jumpLinkType,proto3" json:"jump_link_type,omitempty"`
}

func (x *GetChatAuthReply) Reset() {
	*x = GetChatAuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_search_v2_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatAuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatAuthReply) ProtoMessage() {}

func (x *GetChatAuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_search_v2_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatAuthReply.ProtoReflect.Descriptor instead.
func (*GetChatAuthReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_search_v2_search_proto_rawDescGZIP(), []int{3}
}

func (x *GetChatAuthReply) GetDisplay() bool {
	if x != nil {
		return x.Display
	}
	return false
}

func (x *GetChatAuthReply) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *GetChatAuthReply) GetIconNight() string {
	if x != nil {
		return x.IconNight
	}
	return ""
}

func (x *GetChatAuthReply) GetJumpLink() string {
	if x != nil {
		return x.JumpLink
	}
	return ""
}

func (x *GetChatAuthReply) GetTextGuide() string {
	if x != nil {
		return x.TextGuide
	}
	return ""
}

func (x *GetChatAuthReply) GetJumpLinkType() int32 {
	if x != nil {
		return x.JumpLinkType
	}
	return 0
}

type GetChatResultReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	SessionId string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	FromSource string `protobuf:"bytes,3,opt,name=from_source,json=fromSource,proto3" json:"from_source,omitempty"`
	TrackId string `protobuf:"bytes,4,opt,name=track_id,json=trackId,proto3" json:"track_id,omitempty"`
}

func (x *GetChatResultReq) Reset() {
	*x = GetChatResultReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_search_v2_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatResultReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatResultReq) ProtoMessage() {}

func (x *GetChatResultReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_search_v2_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatResultReq.ProtoReflect.Descriptor instead.
func (*GetChatResultReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_search_v2_search_proto_rawDescGZIP(), []int{4}
}

func (x *GetChatResultReq) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *GetChatResultReq) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *GetChatResultReq) GetFromSource() string {
	if x != nil {
		return x.FromSource
	}
	return ""
}

func (x *GetChatResultReq) GetTrackId() string {
	if x != nil {
		return x.TrackId
	}
	return ""
}

type SearchEggInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EggType int32 `protobuf:"varint,1,opt,name=egg_type,json=eggType,proto3" json:"egg_type,omitempty"`
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	IsCommercial int32 `protobuf:"varint,3,opt,name=is_commercial,json=isCommercial,proto3" json:"is_commercial,omitempty"`
	MaskColor string `protobuf:"bytes,4,opt,name=mask_color,json=maskColor,proto3" json:"mask_color,omitempty"`
	MaskTransparency int64 `protobuf:"varint,5,opt,name=mask_transparency,json=maskTransparency,proto3" json:"mask_transparency,omitempty"`
	Md5 string `protobuf:"bytes,6,opt,name=md5,proto3" json:"md5,omitempty"`
	ReType int32 `protobuf:"varint,7,opt,name=re_type,json=reType,proto3" json:"re_type,omitempty"`
	ReUrl string `protobuf:"bytes,8,opt,name=re_url,json=reUrl,proto3" json:"re_url,omitempty"`
	ReValue string `protobuf:"bytes,9,opt,name=re_value,json=reValue,proto3" json:"re_value,omitempty"`
	ShowCount int32 `protobuf:"varint,10,opt,name=show_count,json=showCount,proto3" json:"show_count,omitempty"`
	Size int64 `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Source int64 `protobuf:"varint,12,opt,name=source,proto3" json:"source,omitempty"`
	Url string `protobuf:"bytes,13,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *SearchEggInfo) Reset() {
	*x = SearchEggInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_search_v2_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchEggInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchEggInfo) ProtoMessage() {}

func (x *SearchEggInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_search_v2_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchEggInfo.ProtoReflect.Descriptor instead.
func (*SearchEggInfo) Descriptor() ([]byte, []int) {
	return file_bilibili_app_search_v2_search_proto_rawDescGZIP(), []int{5}
}

func (x *SearchEggInfo) GetEggType() int32 {
	if x != nil {
		return x.EggType
	}
	return 0
}

func (x *SearchEggInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SearchEggInfo) GetIsCommercial() int32 {
	if x != nil {
		return x.IsCommercial
	}
	return 0
}

func (x *SearchEggInfo) GetMaskColor() string {
	if x != nil {
		return x.MaskColor
	}
	return ""
}

func (x *SearchEggInfo) GetMaskTransparency() int64 {
	if x != nil {
		return x.MaskTransparency
	}
	return 0
}

func (x *SearchEggInfo) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *SearchEggInfo) GetReType() int32 {
	if x != nil {
		return x.ReType
	}
	return 0
}

func (x *SearchEggInfo) GetReUrl() string {
	if x != nil {
		return x.ReUrl
	}
	return ""
}

func (x *SearchEggInfo) GetReValue() string {
	if x != nil {
		return x.ReValue
	}
	return ""
}

func (x *SearchEggInfo) GetShowCount() int32 {
	if x != nil {
		return x.ShowCount
	}
	return 0
}

func (x *SearchEggInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchEggInfo) GetSource() int64 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *SearchEggInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type SearchEggInfos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EggInfo []*SearchEggInfo `protobuf:"bytes,1,rep,name=egg_info,json=eggInfo,proto3" json:"egg_info,omitempty"`
}

func (x *SearchEggInfos) Reset() {
	*x = SearchEggInfos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_search_v2_search_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchEggInfos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchEggInfos) ProtoMessage() {}

func (x *SearchEggInfos) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_search_v2_search_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchEggInfos.ProtoReflect.Descriptor instead.
func (*SearchEggInfos) Descriptor() ([]byte, []int) {
	return file_bilibili_app_search_v2_search_proto_rawDescGZIP(), []int{6}
}

func (x *SearchEggInfos) GetEggInfo() []*SearchEggInfo {
	if x != nil {
		return x.EggInfo
	}
	return nil
}

type SearchEggReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Seid string `protobuf:"bytes,2,opt,name=seid,proto3" json:"seid,omitempty"`
	Result *SearchEggInfos `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SearchEggReply) Reset() {
	*x = SearchEggReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_search_v2_search_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchEggReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchEggReply) ProtoMessage() {}

func (x *SearchEggReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_search_v2_search_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchEggReply.ProtoReflect.Descriptor instead.
func (*SearchEggReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_search_v2_search_proto_rawDescGZIP(), []int{7}
}

func (x *SearchEggReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SearchEggReply) GetSeid() string {
	if x != nil {
		return x.Seid
	}
	return ""
}

func (x *SearchEggReply) GetResult() *SearchEggInfos {
	if x != nil {
		return x.Result
	}
	return nil
}

type SearchEggReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchEggReq) Reset() {
	*x = SearchEggReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_search_v2_search_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchEggReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchEggReq) ProtoMessage() {}

func (x *SearchEggReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_search_v2_search_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchEggReq.ProtoReflect.Descriptor instead.
func (*SearchEggReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_search_v2_search_proto_rawDescGZIP(), []int{8}
}

type SubmitChatTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	SessionId string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *SubmitChatTaskReply) Reset() {
	*x = SubmitChatTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_search_v2_search_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitChatTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitChatTaskReply) ProtoMessage() {}

func (x *SubmitChatTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_search_v2_search_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitChatTaskReply.ProtoReflect.Descriptor instead.
func (*SubmitChatTaskReply) Descriptor() ([]byte, []int) {
	return file_bilibili_app_search_v2_search_proto_rawDescGZIP(), []int{9}
}

func (x *SubmitChatTaskReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SubmitChatTaskReply) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type SubmitChatTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	TrackId string `protobuf:"bytes,2,opt,name=track_id,json=trackId,proto3" json:"track_id,omitempty"`
	FromSource string `protobuf:"bytes,3,opt,name=from_source,json=fromSource,proto3" json:"from_source,omitempty"`
}

func (x *SubmitChatTaskReq) Reset() {
	*x = SubmitChatTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_app_search_v2_search_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitChatTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitChatTaskReq) ProtoMessage() {}

func (x *SubmitChatTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_app_search_v2_search_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitChatTaskReq.ProtoReflect.Descriptor instead.
func (*SubmitChatTaskReq) Descriptor() ([]byte, []int) {
	return file_bilibili_app_search_v2_search_proto_rawDescGZIP(), []int{10}
}

func (x *SubmitChatTaskReq) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SubmitChatTaskReq) GetTrackId() string {
	if x != nil {
		return x.TrackId
	}
	return ""
}

func (x *SubmitChatTaskReq) GetFromSource() string {
	if x != nil {
		return x.FromSource
	}
	return ""
}

var File_bilibili_app_search_v2_search_proto protoreflect.FileDescriptor

var file_bilibili_app_search_v2_search_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x32, 0x1a, 0x2c, 0x62,
	0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x11, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0x29, 0x0a, 0x13, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x22, 0xc1, 0x01,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x63, 0x6f, 0x6e, 0x4e, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x6a, 0x75, 0x6d, 0x70, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6a, 0x75, 0x6d, 0x70, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x65, 0x78, 0x74, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6a,
	0x75, 0x6d, 0x70, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x6a, 0x75, 0x6d, 0x70, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x83, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x22, 0xe5, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x45, 0x67, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x67, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x67, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x69, 0x73, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x73,
	0x6b, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x6d, 0x61, 0x73, 0x6b,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x6d, 0x61, 0x73, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22,
	0x52, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x67, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x12, 0x40, 0x0a, 0x08, 0x65, 0x67, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x45, 0x67, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x65, 0x67, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x78, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x67, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x65, 0x69, 0x64, 0x12, 0x3e, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x67, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x0e, 0x0a,
	0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x67, 0x67, 0x52, 0x65, 0x71, 0x22, 0x48, 0x0a,
	0x13, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x32, 0x80,
	0x04, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x68, 0x0a, 0x0e, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x29, 0x2e, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x5f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x26, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x66, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x28, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x2b, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x59, 0x0a, 0x09,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x67, 0x67, 0x12, 0x24, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x76, 0x32, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x67, 0x67, 0x52, 0x65, 0x71, 0x1a,
	0x26, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45,
	0x67, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x68, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x29, 0x2e, 0x62, 0x69, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x76, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x61, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x45, 0x5a, 0x43, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x6d, 0x74, 0x63, 0x32, 0x33, 0x33,
	0x2f, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_app_search_v2_search_proto_rawDescOnce sync.Once
	file_bilibili_app_search_v2_search_proto_rawDescData = file_bilibili_app_search_v2_search_proto_rawDesc
)

func file_bilibili_app_search_v2_search_proto_rawDescGZIP() []byte {
	file_bilibili_app_search_v2_search_proto_rawDescOnce.Do(func() {
		file_bilibili_app_search_v2_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_app_search_v2_search_proto_rawDescData)
	})
	return file_bilibili_app_search_v2_search_proto_rawDescData
}

var file_bilibili_app_search_v2_search_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_bilibili_app_search_v2_search_proto_goTypes = []interface{}{
	(*CancelChatTaskReq)(nil),   // 0: bilibili.app.search.v2.CancelChatTaskReq
	(*CancelChatTaskReply)(nil), // 1: bilibili.app.search.v2.CancelChatTaskReply
	(*GetChatAuthReq)(nil),      // 2: bilibili.app.search.v2.GetChatAuthReq
	(*GetChatAuthReply)(nil),    // 3: bilibili.app.search.v2.GetChatAuthReply
	(*GetChatResultReq)(nil),    // 4: bilibili.app.search.v2.GetChatResultReq
	(*SearchEggInfo)(nil),       // 5: bilibili.app.search.v2.SearchEggInfo
	(*SearchEggInfos)(nil),      // 6: bilibili.app.search.v2.SearchEggInfos
	(*SearchEggReply)(nil),      // 7: bilibili.app.search.v2.SearchEggReply
	(*SearchEggReq)(nil),        // 8: bilibili.app.search.v2.SearchEggReq
	(*SubmitChatTaskReply)(nil), // 9: bilibili.app.search.v2.SubmitChatTaskReply
	(*SubmitChatTaskReq)(nil),   // 10: bilibili.app.search.v2.SubmitChatTaskReq
	(*main.ChatResult)(nil),     // 11: bilibili.broadcast.message.main.ChatResult
}
var file_bilibili_app_search_v2_search_proto_depIdxs = []int32{
	5,  // 0: bilibili.app.search.v2.SearchEggInfos.egg_info:type_name -> bilibili.app.search.v2.SearchEggInfo
	6,  // 1: bilibili.app.search.v2.SearchEggReply.result:type_name -> bilibili.app.search.v2.SearchEggInfos
	0,  // 2: bilibili.app.search.v2.Search.CancelChatTask:input_type -> bilibili.app.search.v2.CancelChatTaskReq
	2,  // 3: bilibili.app.search.v2.Search.GetChatAuth:input_type -> bilibili.app.search.v2.GetChatAuthReq
	4,  // 4: bilibili.app.search.v2.Search.GetChatResult:input_type -> bilibili.app.search.v2.GetChatResultReq
	8,  // 5: bilibili.app.search.v2.Search.SearchEgg:input_type -> bilibili.app.search.v2.SearchEggReq
	10, // 6: bilibili.app.search.v2.Search.SubmitChatTask:input_type -> bilibili.app.search.v2.SubmitChatTaskReq
	1,  // 7: bilibili.app.search.v2.Search.CancelChatTask:output_type -> bilibili.app.search.v2.CancelChatTaskReply
	3,  // 8: bilibili.app.search.v2.Search.GetChatAuth:output_type -> bilibili.app.search.v2.GetChatAuthReply
	11, // 9: bilibili.app.search.v2.Search.GetChatResult:output_type -> bilibili.broadcast.message.main.ChatResult
	7,  // 10: bilibili.app.search.v2.Search.SearchEgg:output_type -> bilibili.app.search.v2.SearchEggReply
	9,  // 11: bilibili.app.search.v2.Search.SubmitChatTask:output_type -> bilibili.app.search.v2.SubmitChatTaskReply
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_bilibili_app_search_v2_search_proto_init() }
func file_bilibili_app_search_v2_search_proto_init() {
	if File_bilibili_app_search_v2_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_app_search_v2_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelChatTaskReq); i {
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
		file_bilibili_app_search_v2_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelChatTaskReply); i {
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
		file_bilibili_app_search_v2_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatAuthReq); i {
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
		file_bilibili_app_search_v2_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatAuthReply); i {
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
		file_bilibili_app_search_v2_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatResultReq); i {
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
		file_bilibili_app_search_v2_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchEggInfo); i {
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
		file_bilibili_app_search_v2_search_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchEggInfos); i {
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
		file_bilibili_app_search_v2_search_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchEggReply); i {
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
		file_bilibili_app_search_v2_search_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchEggReq); i {
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
		file_bilibili_app_search_v2_search_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitChatTaskReply); i {
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
		file_bilibili_app_search_v2_search_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitChatTaskReq); i {
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
			RawDescriptor: file_bilibili_app_search_v2_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_app_search_v2_search_proto_goTypes,
		DependencyIndexes: file_bilibili_app_search_v2_search_proto_depIdxs,
		MessageInfos:      file_bilibili_app_search_v2_search_proto_msgTypes,
	}.Build()
	File_bilibili_app_search_v2_search_proto = out.File
	file_bilibili_app_search_v2_search_proto_rawDesc = nil
	file_bilibili_app_search_v2_search_proto_goTypes = nil
	file_bilibili_app_search_v2_search_proto_depIdxs = nil
}
