// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: event-management/gunk/v1/userEvent/all.proto

package usereventpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_Going      Status = 0
	Status_MaybeGoing Status = 1
	Status_Interested Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Going",
		1: "MaybeGoing",
		2: "Interested",
	}
	Status_value = map[string]int32{
		"Going":      0,
		"MaybeGoing": 1,
		"Interested": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_event_management_gunk_v1_userEvent_all_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_event_management_gunk_v1_userEvent_all_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_userEvent_all_proto_rawDescGZIP(), []int{0}
}

type UserEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32                  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	EventID   int32                  `protobuf:"varint,2,opt,name=EventID,proto3" json:"EventID,omitempty"`
	Status    Status                 `protobuf:"varint,3,opt,name=Status,proto3,enum=usereventpb.Status" json:"Status,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *UserEvent) Reset() {
	*x = UserEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEvent) ProtoMessage() {}

func (x *UserEvent) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEvent.ProtoReflect.Descriptor instead.
func (*UserEvent) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_userEvent_all_proto_rawDescGZIP(), []int{0}
}

func (x *UserEvent) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserEvent) GetEventID() int32 {
	if x != nil {
		return x.EventID
	}
	return 0
}

func (x *UserEvent) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Going
}

func (x *UserEvent) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserEvent) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateUserEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	EventID int32  `protobuf:"varint,2,opt,name=EventID,proto3" json:"EventID,omitempty"`
	Status  Status `protobuf:"varint,3,opt,name=Status,proto3,enum=usereventpb.Status" json:"Status,omitempty"`
}

func (x *CreateUserEventRequest) Reset() {
	*x = CreateUserEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserEventRequest) ProtoMessage() {}

func (x *CreateUserEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserEventRequest.ProtoReflect.Descriptor instead.
func (*CreateUserEventRequest) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_userEvent_all_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserEventRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateUserEventRequest) GetEventID() int32 {
	if x != nil {
		return x.EventID
	}
	return 0
}

func (x *CreateUserEventRequest) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Going
}

type CreateUserEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEvent *UserEvent `protobuf:"bytes,1,opt,name=UserEvent,proto3" json:"UserEvent,omitempty"`
}

func (x *CreateUserEventResponse) Reset() {
	*x = CreateUserEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserEventResponse) ProtoMessage() {}

func (x *CreateUserEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserEventResponse.ProtoReflect.Descriptor instead.
func (*CreateUserEventResponse) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_userEvent_all_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserEventResponse) GetUserEvent() *UserEvent {
	if x != nil {
		return x.UserEvent
	}
	return nil
}

type IsEventAvailableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	EventID int32 `protobuf:"varint,2,opt,name=EventID,proto3" json:"EventID,omitempty"`
}

func (x *IsEventAvailableRequest) Reset() {
	*x = IsEventAvailableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsEventAvailableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsEventAvailableRequest) ProtoMessage() {}

func (x *IsEventAvailableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsEventAvailableRequest.ProtoReflect.Descriptor instead.
func (*IsEventAvailableRequest) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_userEvent_all_proto_rawDescGZIP(), []int{3}
}

func (x *IsEventAvailableRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *IsEventAvailableRequest) GetEventID() int32 {
	if x != nil {
		return x.EventID
	}
	return 0
}

type IsEventAvailableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEvent *UserEvent `protobuf:"bytes,1,opt,name=UserEvent,proto3" json:"UserEvent,omitempty"`
}

func (x *IsEventAvailableResponse) Reset() {
	*x = IsEventAvailableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsEventAvailableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsEventAvailableResponse) ProtoMessage() {}

func (x *IsEventAvailableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsEventAvailableResponse.ProtoReflect.Descriptor instead.
func (*IsEventAvailableResponse) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_userEvent_all_proto_rawDescGZIP(), []int{4}
}

func (x *IsEventAvailableResponse) GetUserEvent() *UserEvent {
	if x != nil {
		return x.UserEvent
	}
	return nil
}

type UpdateUserEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEvent *UserEvent `protobuf:"bytes,1,opt,name=UserEvent,proto3" json:"UserEvent,omitempty"`
}

func (x *UpdateUserEventRequest) Reset() {
	*x = UpdateUserEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserEventRequest) ProtoMessage() {}

func (x *UpdateUserEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserEventRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserEventRequest) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_userEvent_all_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserEventRequest) GetUserEvent() *UserEvent {
	if x != nil {
		return x.UserEvent
	}
	return nil
}

type UpdateUserEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEvent *UserEvent `protobuf:"bytes,1,opt,name=UserEvent,proto3" json:"UserEvent,omitempty"`
}

func (x *UpdateUserEventResponse) Reset() {
	*x = UpdateUserEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserEventResponse) ProtoMessage() {}

func (x *UpdateUserEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_userEvent_all_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserEventResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserEventResponse) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_userEvent_all_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserEventResponse) GetUserEvent() *UserEvent {
	if x != nil {
		return x.UserEvent
	}
	return nil
}

var File_event_management_gunk_v1_userEvent_all_proto protoreflect.FileDescriptor

var file_event_management_gunk_v1_userEvent_all_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1b, 0x0a, 0x07, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x50, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x50, 0x00, 0x12, 0x39, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12,
	0x39, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00,
	0x18, 0x00, 0x22, 0x8a, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08,
	0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1b, 0x0a, 0x07, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22,
	0x58, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x5a, 0x0a, 0x17, 0x49, 0x73, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x12, 0x1b, 0x0a, 0x07, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08,
	0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x59, 0x0a, 0x18, 0x49, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00,
	0x22, 0x57, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x58, 0x0a, 0x17, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42,
	0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10,
	0x00, 0x18, 0x00, 0x2a, 0x43, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a,
	0x05, 0x47, 0x6f, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x1a, 0x02, 0x08, 0x00, 0x12, 0x12, 0x0a, 0x0a,
	0x4d, 0x61, 0x79, 0x62, 0x65, 0x47, 0x6f, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x1a, 0x02, 0x08, 0x00,
	0x12, 0x12, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x65, 0x64, 0x10, 0x02,
	0x1a, 0x02, 0x08, 0x00, 0x1a, 0x02, 0x18, 0x00, 0x32, 0xd8, 0x02, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a,
	0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00,
	0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x6b, 0x0a, 0x10, 0x49, 0x73, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x49, 0x73, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e,
	0x49, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x12, 0x68, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x03,
	0x88, 0x02, 0x00, 0x42, 0x49, 0x48, 0x01, 0x50, 0x00, 0x5a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x75, 0x6e, 0x6b,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x3b, 0x75, 0x73,
	0x65, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x80, 0x01, 0x00, 0x88, 0x01, 0x00, 0x90,
	0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01, 0xd0, 0x02, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_management_gunk_v1_userEvent_all_proto_rawDescOnce sync.Once
	file_event_management_gunk_v1_userEvent_all_proto_rawDescData = file_event_management_gunk_v1_userEvent_all_proto_rawDesc
)

func file_event_management_gunk_v1_userEvent_all_proto_rawDescGZIP() []byte {
	file_event_management_gunk_v1_userEvent_all_proto_rawDescOnce.Do(func() {
		file_event_management_gunk_v1_userEvent_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_management_gunk_v1_userEvent_all_proto_rawDescData)
	})
	return file_event_management_gunk_v1_userEvent_all_proto_rawDescData
}

var (
	file_event_management_gunk_v1_userEvent_all_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_event_management_gunk_v1_userEvent_all_proto_msgTypes  = make([]protoimpl.MessageInfo, 7)
	file_event_management_gunk_v1_userEvent_all_proto_goTypes   = []interface{}{
		(Status)(0),                      // 0: usereventpb.Status
		(*UserEvent)(nil),                // 1: usereventpb.UserEvent
		(*CreateUserEventRequest)(nil),   // 2: usereventpb.CreateUserEventRequest
		(*CreateUserEventResponse)(nil),  // 3: usereventpb.CreateUserEventResponse
		(*IsEventAvailableRequest)(nil),  // 4: usereventpb.IsEventAvailableRequest
		(*IsEventAvailableResponse)(nil), // 5: usereventpb.IsEventAvailableResponse
		(*UpdateUserEventRequest)(nil),   // 6: usereventpb.UpdateUserEventRequest
		(*UpdateUserEventResponse)(nil),  // 7: usereventpb.UpdateUserEventResponse
		(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
	}
)

var file_event_management_gunk_v1_userEvent_all_proto_depIdxs = []int32{
	0,  // 0: usereventpb.UserEvent.Status:type_name -> usereventpb.Status
	8,  // 1: usereventpb.UserEvent.CreatedAt:type_name -> google.protobuf.Timestamp
	8,  // 2: usereventpb.UserEvent.UpdatedAt:type_name -> google.protobuf.Timestamp
	0,  // 3: usereventpb.CreateUserEventRequest.Status:type_name -> usereventpb.Status
	1,  // 4: usereventpb.CreateUserEventResponse.UserEvent:type_name -> usereventpb.UserEvent
	1,  // 5: usereventpb.IsEventAvailableResponse.UserEvent:type_name -> usereventpb.UserEvent
	1,  // 6: usereventpb.UpdateUserEventRequest.UserEvent:type_name -> usereventpb.UserEvent
	1,  // 7: usereventpb.UpdateUserEventResponse.UserEvent:type_name -> usereventpb.UserEvent
	2,  // 8: usereventpb.UserEventService.CreateUserEvent:input_type -> usereventpb.CreateUserEventRequest
	4,  // 9: usereventpb.UserEventService.IsEventAvailable:input_type -> usereventpb.IsEventAvailableRequest
	6,  // 10: usereventpb.UserEventService.UpdateUserEvent:input_type -> usereventpb.UpdateUserEventRequest
	3,  // 11: usereventpb.UserEventService.CreateUserEvent:output_type -> usereventpb.CreateUserEventResponse
	5,  // 12: usereventpb.UserEventService.IsEventAvailable:output_type -> usereventpb.IsEventAvailableResponse
	7,  // 13: usereventpb.UserEventService.UpdateUserEvent:output_type -> usereventpb.UpdateUserEventResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_event_management_gunk_v1_userEvent_all_proto_init() }
func file_event_management_gunk_v1_userEvent_all_proto_init() {
	if File_event_management_gunk_v1_userEvent_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_management_gunk_v1_userEvent_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEvent); i {
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
		file_event_management_gunk_v1_userEvent_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserEventRequest); i {
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
		file_event_management_gunk_v1_userEvent_all_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserEventResponse); i {
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
		file_event_management_gunk_v1_userEvent_all_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsEventAvailableRequest); i {
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
		file_event_management_gunk_v1_userEvent_all_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsEventAvailableResponse); i {
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
		file_event_management_gunk_v1_userEvent_all_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserEventRequest); i {
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
		file_event_management_gunk_v1_userEvent_all_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserEventResponse); i {
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
			RawDescriptor: file_event_management_gunk_v1_userEvent_all_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_management_gunk_v1_userEvent_all_proto_goTypes,
		DependencyIndexes: file_event_management_gunk_v1_userEvent_all_proto_depIdxs,
		EnumInfos:         file_event_management_gunk_v1_userEvent_all_proto_enumTypes,
		MessageInfos:      file_event_management_gunk_v1_userEvent_all_proto_msgTypes,
	}.Build()
	File_event_management_gunk_v1_userEvent_all_proto = out.File
	file_event_management_gunk_v1_userEvent_all_proto_rawDesc = nil
	file_event_management_gunk_v1_userEvent_all_proto_goTypes = nil
	file_event_management_gunk_v1_userEvent_all_proto_depIdxs = nil
}
