// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: protocol/collaborate/collaborate.proto

package collaborate

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

// 返回的错误信息
type RetMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RetMsg) Reset() {
	*x = RetMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_collaborate_collaborate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetMsg) ProtoMessage() {}

func (x *RetMsg) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_collaborate_collaborate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetMsg.ProtoReflect.Descriptor instead.
func (*RetMsg) Descriptor() ([]byte, []int) {
	return file_protocol_collaborate_collaborate_proto_rawDescGZIP(), []int{0}
}

func (x *RetMsg) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RetMsg) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// 通用协议
type GeneralReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GeneralReq) Reset() {
	*x = GeneralReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_collaborate_collaborate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralReq) ProtoMessage() {}

func (x *GeneralReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_collaborate_collaborate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralReq.ProtoReflect.Descriptor instead.
func (*GeneralReq) Descriptor() ([]byte, []int) {
	return file_protocol_collaborate_collaborate_proto_rawDescGZIP(), []int{1}
}

func (x *GeneralReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GeneralResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *RetMsg `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GeneralResp) Reset() {
	*x = GeneralResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_collaborate_collaborate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralResp) ProtoMessage() {}

func (x *GeneralResp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_collaborate_collaborate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralResp.ProtoReflect.Descriptor instead.
func (*GeneralResp) Descriptor() ([]byte, []int) {
	return file_protocol_collaborate_collaborate_proto_rawDescGZIP(), []int{2}
}

func (x *GeneralResp) GetMsg() *RetMsg {
	if x != nil {
		return x.Msg
	}
	return nil
}

// 好友信息
type FriendRecordData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FriendId int64  `protobuf:"varint,2,opt,name=friendId,proto3" json:"friendId,omitempty"`
	IsTop    bool   `protobuf:"varint,3,opt,name=isTop,proto3" json:"isTop,omitempty"`
	IsBlack  bool   `protobuf:"varint,4,opt,name=isBlack,proto3" json:"isBlack,omitempty"`
	Label    string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *FriendRecordData) Reset() {
	*x = FriendRecordData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_collaborate_collaborate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRecordData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRecordData) ProtoMessage() {}

func (x *FriendRecordData) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_collaborate_collaborate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRecordData.ProtoReflect.Descriptor instead.
func (*FriendRecordData) Descriptor() ([]byte, []int) {
	return file_protocol_collaborate_collaborate_proto_rawDescGZIP(), []int{3}
}

func (x *FriendRecordData) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FriendRecordData) GetFriendId() int64 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

func (x *FriendRecordData) GetIsTop() bool {
	if x != nil {
		return x.IsTop
	}
	return false
}

func (x *FriendRecordData) GetIsBlack() bool {
	if x != nil {
		return x.IsBlack
	}
	return false
}

func (x *FriendRecordData) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

// 添加好友
type AddFriendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 发起人id
	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	// 被请求添加好友id
	FriendId int64 `protobuf:"varint,2,opt,name=friendId,proto3" json:"friendId,omitempty"`
	// 发起人对好友的备注
	UserLabel string `protobuf:"bytes,3,opt,name=userLabel,proto3" json:"userLabel,omitempty"`
	// 好友对发起人的备注
	FriendLabel string `protobuf:"bytes,4,opt,name=friendLabel,proto3" json:"friendLabel,omitempty"`
}

func (x *AddFriendReq) Reset() {
	*x = AddFriendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_collaborate_collaborate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendReq) ProtoMessage() {}

func (x *AddFriendReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_collaborate_collaborate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendReq.ProtoReflect.Descriptor instead.
func (*AddFriendReq) Descriptor() ([]byte, []int) {
	return file_protocol_collaborate_collaborate_proto_rawDescGZIP(), []int{4}
}

func (x *AddFriendReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddFriendReq) GetFriendId() int64 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

func (x *AddFriendReq) GetUserLabel() string {
	if x != nil {
		return x.UserLabel
	}
	return ""
}

func (x *AddFriendReq) GetFriendLabel() string {
	if x != nil {
		return x.FriendLabel
	}
	return ""
}

type AddFriendResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *RetMsg `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AddFriendResp) Reset() {
	*x = AddFriendResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_collaborate_collaborate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendResp) ProtoMessage() {}

func (x *AddFriendResp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_collaborate_collaborate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendResp.ProtoReflect.Descriptor instead.
func (*AddFriendResp) Descriptor() ([]byte, []int) {
	return file_protocol_collaborate_collaborate_proto_rawDescGZIP(), []int{5}
}

func (x *AddFriendResp) GetMsg() *RetMsg {
	if x != nil {
		return x.Msg
	}
	return nil
}

// 删除好友
type DeleteFriendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FriendId int64 `protobuf:"varint,2,opt,name=friendId,proto3" json:"friendId,omitempty"`
}

func (x *DeleteFriendReq) Reset() {
	*x = DeleteFriendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_collaborate_collaborate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFriendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFriendReq) ProtoMessage() {}

func (x *DeleteFriendReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_collaborate_collaborate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFriendReq.ProtoReflect.Descriptor instead.
func (*DeleteFriendReq) Descriptor() ([]byte, []int) {
	return file_protocol_collaborate_collaborate_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteFriendReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteFriendReq) GetFriendId() int64 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

type DeleteFriendResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *RetMsg `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DeleteFriendResp) Reset() {
	*x = DeleteFriendResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_collaborate_collaborate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFriendResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFriendResp) ProtoMessage() {}

func (x *DeleteFriendResp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_collaborate_collaborate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFriendResp.ProtoReflect.Descriptor instead.
func (*DeleteFriendResp) Descriptor() ([]byte, []int) {
	return file_protocol_collaborate_collaborate_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteFriendResp) GetMsg() *RetMsg {
	if x != nil {
		return x.Msg
	}
	return nil
}

// 修改好友信息
type UpdateFriendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *FriendRecordData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateFriendReq) Reset() {
	*x = UpdateFriendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_collaborate_collaborate_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFriendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFriendReq) ProtoMessage() {}

func (x *UpdateFriendReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_collaborate_collaborate_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFriendReq.ProtoReflect.Descriptor instead.
func (*UpdateFriendReq) Descriptor() ([]byte, []int) {
	return file_protocol_collaborate_collaborate_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateFriendReq) GetData() *FriendRecordData {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateFriendResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  *RetMsg           `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *FriendRecordData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateFriendResp) Reset() {
	*x = UpdateFriendResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_collaborate_collaborate_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFriendResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFriendResp) ProtoMessage() {}

func (x *UpdateFriendResp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_collaborate_collaborate_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFriendResp.ProtoReflect.Descriptor instead.
func (*UpdateFriendResp) Descriptor() ([]byte, []int) {
	return file_protocol_collaborate_collaborate_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateFriendResp) GetMsg() *RetMsg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *UpdateFriendResp) GetData() *FriendRecordData {
	if x != nil {
		return x.Data
	}
	return nil
}

// 查询全部好友
type FindAllFriendsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Opt int32 `protobuf:"varint,2,opt,name=opt,proto3" json:"opt,omitempty"`
}

func (x *FindAllFriendsReq) Reset() {
	*x = FindAllFriendsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_collaborate_collaborate_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllFriendsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllFriendsReq) ProtoMessage() {}

func (x *FindAllFriendsReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_collaborate_collaborate_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllFriendsReq.ProtoReflect.Descriptor instead.
func (*FindAllFriendsReq) Descriptor() ([]byte, []int) {
	return file_protocol_collaborate_collaborate_proto_rawDescGZIP(), []int{10}
}

func (x *FindAllFriendsReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *FindAllFriendsReq) GetOpt() int32 {
	if x != nil {
		return x.Opt
	}
	return 0
}

type FindAllFriendsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  *RetMsg             `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []*FriendRecordData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *FindAllFriendsResp) Reset() {
	*x = FindAllFriendsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_collaborate_collaborate_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllFriendsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllFriendsResp) ProtoMessage() {}

func (x *FindAllFriendsResp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_collaborate_collaborate_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllFriendsResp.ProtoReflect.Descriptor instead.
func (*FindAllFriendsResp) Descriptor() ([]byte, []int) {
	return file_protocol_collaborate_collaborate_proto_rawDescGZIP(), []int{11}
}

func (x *FindAllFriendsResp) GetMsg() *RetMsg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *FindAllFriendsResp) GetData() []*FriendRecordData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_protocol_collaborate_collaborate_proto protoreflect.FileDescriptor

var file_protocol_collaborate_collaborate_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61,
	0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62,
	0x6f, 0x72, 0x61, 0x74, 0x65, 0x22, 0x36, 0x0a, 0x06, 0x52, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x1e, 0x0a,
	0x0a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x34, 0x0a,
	0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6c, 0x6c,
	0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x8c, 0x01, 0x0a, 0x10, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x73, 0x54, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x54,
	0x6f, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x36, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72,
	0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x45, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62,
	0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x44, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6c, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61,
	0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x37, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6f, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6f, 0x70, 0x74, 0x22, 0x6e,
	0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e,
	0x52, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x31, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6c, 0x6c,
	0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xcd,
	0x02, 0x0a, 0x12, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e,
	0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x1c, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6c, 0x6c,
	0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61,
	0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x1e, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x18,
	0x5a, 0x16, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6c,
	0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_collaborate_collaborate_proto_rawDescOnce sync.Once
	file_protocol_collaborate_collaborate_proto_rawDescData = file_protocol_collaborate_collaborate_proto_rawDesc
)

func file_protocol_collaborate_collaborate_proto_rawDescGZIP() []byte {
	file_protocol_collaborate_collaborate_proto_rawDescOnce.Do(func() {
		file_protocol_collaborate_collaborate_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_collaborate_collaborate_proto_rawDescData)
	})
	return file_protocol_collaborate_collaborate_proto_rawDescData
}

var file_protocol_collaborate_collaborate_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_protocol_collaborate_collaborate_proto_goTypes = []interface{}{
	(*RetMsg)(nil),             // 0: collaborate.RetMsg
	(*GeneralReq)(nil),         // 1: collaborate.GeneralReq
	(*GeneralResp)(nil),        // 2: collaborate.GeneralResp
	(*FriendRecordData)(nil),   // 3: collaborate.FriendRecordData
	(*AddFriendReq)(nil),       // 4: collaborate.AddFriendReq
	(*AddFriendResp)(nil),      // 5: collaborate.AddFriendResp
	(*DeleteFriendReq)(nil),    // 6: collaborate.DeleteFriendReq
	(*DeleteFriendResp)(nil),   // 7: collaborate.DeleteFriendResp
	(*UpdateFriendReq)(nil),    // 8: collaborate.UpdateFriendReq
	(*UpdateFriendResp)(nil),   // 9: collaborate.UpdateFriendResp
	(*FindAllFriendsReq)(nil),  // 10: collaborate.FindAllFriendsReq
	(*FindAllFriendsResp)(nil), // 11: collaborate.FindAllFriendsResp
}
var file_protocol_collaborate_collaborate_proto_depIdxs = []int32{
	0,  // 0: collaborate.GeneralResp.msg:type_name -> collaborate.RetMsg
	0,  // 1: collaborate.AddFriendResp.msg:type_name -> collaborate.RetMsg
	0,  // 2: collaborate.DeleteFriendResp.msg:type_name -> collaborate.RetMsg
	3,  // 3: collaborate.UpdateFriendReq.data:type_name -> collaborate.FriendRecordData
	0,  // 4: collaborate.UpdateFriendResp.msg:type_name -> collaborate.RetMsg
	3,  // 5: collaborate.UpdateFriendResp.data:type_name -> collaborate.FriendRecordData
	0,  // 6: collaborate.FindAllFriendsResp.msg:type_name -> collaborate.RetMsg
	3,  // 7: collaborate.FindAllFriendsResp.data:type_name -> collaborate.FriendRecordData
	4,  // 8: collaborate.CollaborateService.AddFriend:input_type -> collaborate.AddFriendReq
	6,  // 9: collaborate.CollaborateService.DeleteFriend:input_type -> collaborate.DeleteFriendReq
	8,  // 10: collaborate.CollaborateService.UpdateFriend:input_type -> collaborate.UpdateFriendReq
	10, // 11: collaborate.CollaborateService.FindAllFriends:input_type -> collaborate.FindAllFriendsReq
	5,  // 12: collaborate.CollaborateService.AddFriend:output_type -> collaborate.AddFriendResp
	7,  // 13: collaborate.CollaborateService.DeleteFriend:output_type -> collaborate.DeleteFriendResp
	9,  // 14: collaborate.CollaborateService.UpdateFriend:output_type -> collaborate.UpdateFriendResp
	11, // 15: collaborate.CollaborateService.FindAllFriends:output_type -> collaborate.FindAllFriendsResp
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_protocol_collaborate_collaborate_proto_init() }
func file_protocol_collaborate_collaborate_proto_init() {
	if File_protocol_collaborate_collaborate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_collaborate_collaborate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetMsg); i {
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
		file_protocol_collaborate_collaborate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralReq); i {
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
		file_protocol_collaborate_collaborate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralResp); i {
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
		file_protocol_collaborate_collaborate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRecordData); i {
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
		file_protocol_collaborate_collaborate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendReq); i {
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
		file_protocol_collaborate_collaborate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendResp); i {
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
		file_protocol_collaborate_collaborate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFriendReq); i {
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
		file_protocol_collaborate_collaborate_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFriendResp); i {
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
		file_protocol_collaborate_collaborate_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFriendReq); i {
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
		file_protocol_collaborate_collaborate_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFriendResp); i {
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
		file_protocol_collaborate_collaborate_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllFriendsReq); i {
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
		file_protocol_collaborate_collaborate_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllFriendsResp); i {
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
			RawDescriptor: file_protocol_collaborate_collaborate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_collaborate_collaborate_proto_goTypes,
		DependencyIndexes: file_protocol_collaborate_collaborate_proto_depIdxs,
		MessageInfos:      file_protocol_collaborate_collaborate_proto_msgTypes,
	}.Build()
	File_protocol_collaborate_collaborate_proto = out.File
	file_protocol_collaborate_collaborate_proto_rawDesc = nil
	file_protocol_collaborate_collaborate_proto_goTypes = nil
	file_protocol_collaborate_collaborate_proto_depIdxs = nil
}
