// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.3
// source: mafia.proto

package mafia

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type FinishDayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinishDayRequest) Reset() {
	*x = FinishDayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishDayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishDayRequest) ProtoMessage() {}

func (x *FinishDayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishDayRequest.ProtoReflect.Descriptor instead.
func (*FinishDayRequest) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{2}
}

type FinishDayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinishDayResponse) Reset() {
	*x = FinishDayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishDayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishDayResponse) ProtoMessage() {}

func (x *FinishDayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishDayResponse.ProtoReflect.Descriptor instead.
func (*FinishDayResponse) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{3}
}

type ExecutePlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExecutePlayerRequest) Reset() {
	*x = ExecutePlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutePlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePlayerRequest) ProtoMessage() {}

func (x *ExecutePlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePlayerRequest.ProtoReflect.Descriptor instead.
func (*ExecutePlayerRequest) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{4}
}

type ExecutePlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExecutePlayerResponse) Reset() {
	*x = ExecutePlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutePlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutePlayerResponse) ProtoMessage() {}

func (x *ExecutePlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutePlayerResponse.ProtoReflect.Descriptor instead.
func (*ExecutePlayerResponse) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{5}
}

type KillPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *KillPlayerRequest) Reset() {
	*x = KillPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KillPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillPlayerRequest) ProtoMessage() {}

func (x *KillPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillPlayerRequest.ProtoReflect.Descriptor instead.
func (*KillPlayerRequest) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{6}
}

func (x *KillPlayerRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type KillPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KillPlayerResponse) Reset() {
	*x = KillPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KillPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillPlayerResponse) ProtoMessage() {}

func (x *KillPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillPlayerResponse.ProtoReflect.Descriptor instead.
func (*KillPlayerResponse) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{7}
}

type CheckPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *CheckPlayerRequest) Reset() {
	*x = CheckPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPlayerRequest) ProtoMessage() {}

func (x *CheckPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPlayerRequest.ProtoReflect.Descriptor instead.
func (*CheckPlayerRequest) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{8}
}

func (x *CheckPlayerRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type CheckPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsMafia bool `protobuf:"varint,1,opt,name=is_mafia,json=isMafia,proto3" json:"is_mafia,omitempty"`
}

func (x *CheckPlayerResponse) Reset() {
	*x = CheckPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPlayerResponse) ProtoMessage() {}

func (x *CheckPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPlayerResponse.ProtoReflect.Descriptor instead.
func (*CheckPlayerResponse) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{9}
}

func (x *CheckPlayerResponse) GetIsMafia() bool {
	if x != nil {
		return x.IsMafia
	}
	return false
}

type PublishCheckResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishCheckResultRequest) Reset() {
	*x = PublishCheckResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishCheckResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishCheckResultRequest) ProtoMessage() {}

func (x *PublishCheckResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishCheckResultRequest.ProtoReflect.Descriptor instead.
func (*PublishCheckResultRequest) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{10}
}

type PublishCheckResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishCheckResultResponse) Reset() {
	*x = PublishCheckResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishCheckResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishCheckResultResponse) ProtoMessage() {}

func (x *PublishCheckResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishCheckResultResponse.ProtoReflect.Descriptor instead.
func (*PublishCheckResultResponse) Descriptor() ([]byte, []int) {
	return file_mafia_proto_rawDescGZIP(), []int{11}
}

var File_mafia_proto protoreflect.FileDescriptor

var file_mafia_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x61, 0x66, 0x69, 0x61, 0x22, 0x21, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x44, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x44, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x0a, 0x14,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a,
	0x11, 0x4b, 0x69, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x4b, 0x69,
	0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2c, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x30,
	0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x66, 0x69,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4d, 0x61, 0x66, 0x69, 0x61,
	0x22, 0x1b, 0x0a, 0x19, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1c, 0x0a,
	0x1a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa5, 0x03, 0x0a, 0x05,
	0x4d, 0x61, 0x66, 0x69, 0x61, 0x12, 0x2c, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13,
	0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x44, 0x61, 0x79,
	0x12, 0x17, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x44,
	0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x66, 0x69,
	0x61, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x44, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x41, 0x0a, 0x0a, 0x4b, 0x69, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x18, 0x2e,
	0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e,
	0x4b, 0x69, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d,
	0x61, 0x66, 0x69, 0x61, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x12, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x20,
	0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x68, 0x6d, 0x65, 0x74, 0x6f, 0x76, 0x2f, 0x73,
	0x6f, 0x61, 0x2d, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2d, 0x32, 0x2f, 0x6d, 0x61,
	0x66, 0x69, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mafia_proto_rawDescOnce sync.Once
	file_mafia_proto_rawDescData = file_mafia_proto_rawDesc
)

func file_mafia_proto_rawDescGZIP() []byte {
	file_mafia_proto_rawDescOnce.Do(func() {
		file_mafia_proto_rawDescData = protoimpl.X.CompressGZIP(file_mafia_proto_rawDescData)
	})
	return file_mafia_proto_rawDescData
}

var file_mafia_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_mafia_proto_goTypes = []interface{}{
	(*Event)(nil),                      // 0: mafia.Event
	(*LoginRequest)(nil),               // 1: mafia.LoginRequest
	(*FinishDayRequest)(nil),           // 2: mafia.FinishDayRequest
	(*FinishDayResponse)(nil),          // 3: mafia.FinishDayResponse
	(*ExecutePlayerRequest)(nil),       // 4: mafia.ExecutePlayerRequest
	(*ExecutePlayerResponse)(nil),      // 5: mafia.ExecutePlayerResponse
	(*KillPlayerRequest)(nil),          // 6: mafia.KillPlayerRequest
	(*KillPlayerResponse)(nil),         // 7: mafia.KillPlayerResponse
	(*CheckPlayerRequest)(nil),         // 8: mafia.CheckPlayerRequest
	(*CheckPlayerResponse)(nil),        // 9: mafia.CheckPlayerResponse
	(*PublishCheckResultRequest)(nil),  // 10: mafia.PublishCheckResultRequest
	(*PublishCheckResultResponse)(nil), // 11: mafia.PublishCheckResultResponse
}
var file_mafia_proto_depIdxs = []int32{
	1,  // 0: mafia.Mafia.Login:input_type -> mafia.LoginRequest
	2,  // 1: mafia.Mafia.FinishDay:input_type -> mafia.FinishDayRequest
	4,  // 2: mafia.Mafia.ExecutePlayer:input_type -> mafia.ExecutePlayerRequest
	6,  // 3: mafia.Mafia.KillPlayer:input_type -> mafia.KillPlayerRequest
	8,  // 4: mafia.Mafia.CheckPlayer:input_type -> mafia.CheckPlayerRequest
	10, // 5: mafia.Mafia.PublishCheckResult:input_type -> mafia.PublishCheckResultRequest
	0,  // 6: mafia.Mafia.Login:output_type -> mafia.Event
	3,  // 7: mafia.Mafia.FinishDay:output_type -> mafia.FinishDayResponse
	5,  // 8: mafia.Mafia.ExecutePlayer:output_type -> mafia.ExecutePlayerResponse
	7,  // 9: mafia.Mafia.KillPlayer:output_type -> mafia.KillPlayerResponse
	9,  // 10: mafia.Mafia.CheckPlayer:output_type -> mafia.CheckPlayerResponse
	11, // 11: mafia.Mafia.PublishCheckResult:output_type -> mafia.PublishCheckResultResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_mafia_proto_init() }
func file_mafia_proto_init() {
	if File_mafia_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mafia_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_mafia_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_mafia_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishDayRequest); i {
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
		file_mafia_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishDayResponse); i {
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
		file_mafia_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutePlayerRequest); i {
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
		file_mafia_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutePlayerResponse); i {
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
		file_mafia_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KillPlayerRequest); i {
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
		file_mafia_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KillPlayerResponse); i {
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
		file_mafia_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPlayerRequest); i {
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
		file_mafia_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPlayerResponse); i {
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
		file_mafia_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishCheckResultRequest); i {
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
		file_mafia_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishCheckResultResponse); i {
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
			RawDescriptor: file_mafia_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mafia_proto_goTypes,
		DependencyIndexes: file_mafia_proto_depIdxs,
		MessageInfos:      file_mafia_proto_msgTypes,
	}.Build()
	File_mafia_proto = out.File
	file_mafia_proto_rawDesc = nil
	file_mafia_proto_goTypes = nil
	file_mafia_proto_depIdxs = nil
}
