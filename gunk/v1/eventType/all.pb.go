// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: event-management/gunk/v1/eventType/all.proto

package eventTypepb

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

type EventTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int32                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	EventTypeName string                 `protobuf:"bytes,2,opt,name=EventTypeName,proto3" json:"EventTypeName,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
}

func (x *EventTypes) Reset() {
	*x = EventTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTypes) ProtoMessage() {}

func (x *EventTypes) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTypes.ProtoReflect.Descriptor instead.
func (*EventTypes) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_eventType_all_proto_rawDescGZIP(), []int{0}
}

func (x *EventTypes) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *EventTypes) GetEventTypeName() string {
	if x != nil {
		return x.EventTypeName
	}
	return ""
}

func (x *EventTypes) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *EventTypes) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *EventTypes) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type EventTypeFilterList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllEventType []*EventTypes `protobuf:"bytes,1,rep,name=AllEventType,proto3" json:"AllEventType,omitempty"`
	SearchTerm   string        `protobuf:"bytes,2,opt,name=SearchTerm,proto3" json:"SearchTerm,omitempty"`
}

func (x *EventTypeFilterList) Reset() {
	*x = EventTypeFilterList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTypeFilterList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTypeFilterList) ProtoMessage() {}

func (x *EventTypeFilterList) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTypeFilterList.ProtoReflect.Descriptor instead.
func (*EventTypeFilterList) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_eventType_all_proto_rawDescGZIP(), []int{1}
}

func (x *EventTypeFilterList) GetAllEventType() []*EventTypes {
	if x != nil {
		return x.AllEventType
	}
	return nil
}

func (x *EventTypeFilterList) GetSearchTerm() string {
	if x != nil {
		return x.SearchTerm
	}
	return ""
}

type CreateEventTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTypeName string `protobuf:"bytes,1,opt,name=EventTypeName,proto3" json:"EventTypeName,omitempty"`
}

func (x *CreateEventTypeRequest) Reset() {
	*x = CreateEventTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventTypeRequest) ProtoMessage() {}

func (x *CreateEventTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventTypeRequest.ProtoReflect.Descriptor instead.
func (*CreateEventTypeRequest) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_eventType_all_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEventTypeRequest) GetEventTypeName() string {
	if x != nil {
		return x.EventTypeName
	}
	return ""
}

type CreateEventTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTypes *EventTypes `protobuf:"bytes,1,opt,name=EventTypes,proto3" json:"EventTypes,omitempty"`
}

func (x *CreateEventTypeResponse) Reset() {
	*x = CreateEventTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventTypeResponse) ProtoMessage() {}

func (x *CreateEventTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventTypeResponse.ProtoReflect.Descriptor instead.
func (*CreateEventTypeResponse) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_eventType_all_proto_rawDescGZIP(), []int{3}
}

func (x *CreateEventTypeResponse) GetEventTypes() *EventTypes {
	if x != nil {
		return x.EventTypes
	}
	return nil
}

type EditEventTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *EditEventTypeRequest) Reset() {
	*x = EditEventTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditEventTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditEventTypeRequest) ProtoMessage() {}

func (x *EditEventTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditEventTypeRequest.ProtoReflect.Descriptor instead.
func (*EditEventTypeRequest) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_eventType_all_proto_rawDescGZIP(), []int{4}
}

func (x *EditEventTypeRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type EditEventTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTypes *EventTypes `protobuf:"bytes,1,opt,name=EventTypes,proto3" json:"EventTypes,omitempty"`
}

func (x *EditEventTypeResponse) Reset() {
	*x = EditEventTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditEventTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditEventTypeResponse) ProtoMessage() {}

func (x *EditEventTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditEventTypeResponse.ProtoReflect.Descriptor instead.
func (*EditEventTypeResponse) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_eventType_all_proto_rawDescGZIP(), []int{5}
}

func (x *EditEventTypeResponse) GetEventTypes() *EventTypes {
	if x != nil {
		return x.EventTypes
	}
	return nil
}

type UpdateEventTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	EventTypeName string `protobuf:"bytes,2,opt,name=EventTypeName,proto3" json:"EventTypeName,omitempty"`
}

func (x *UpdateEventTypeRequest) Reset() {
	*x = UpdateEventTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventTypeRequest) ProtoMessage() {}

func (x *UpdateEventTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventTypeRequest.ProtoReflect.Descriptor instead.
func (*UpdateEventTypeRequest) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_eventType_all_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateEventTypeRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateEventTypeRequest) GetEventTypeName() string {
	if x != nil {
		return x.EventTypeName
	}
	return ""
}

type UpdateEventTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTypes *EventTypes `protobuf:"bytes,1,opt,name=EventTypes,proto3" json:"EventTypes,omitempty"`
}

func (x *UpdateEventTypeResponse) Reset() {
	*x = UpdateEventTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventTypeResponse) ProtoMessage() {}

func (x *UpdateEventTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventTypeResponse.ProtoReflect.Descriptor instead.
func (*UpdateEventTypeResponse) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_eventType_all_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateEventTypeResponse) GetEventTypes() *EventTypes {
	if x != nil {
		return x.EventTypes
	}
	return nil
}

type DeleteEventTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteEventTypeRequest) Reset() {
	*x = DeleteEventTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventTypeRequest) ProtoMessage() {}

func (x *DeleteEventTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventTypeRequest.ProtoReflect.Descriptor instead.
func (*DeleteEventTypeRequest) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_eventType_all_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteEventTypeRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteEventTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteEventTypeResponse) Reset() {
	*x = DeleteEventTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventTypeResponse) ProtoMessage() {}

func (x *DeleteEventTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventTypeResponse.ProtoReflect.Descriptor instead.
func (*DeleteEventTypeResponse) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_eventType_all_proto_rawDescGZIP(), []int{9}
}

type EventTypeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchTerm string `protobuf:"bytes,1,opt,name=SearchTerm,proto3" json:"SearchTerm,omitempty"`
}

func (x *EventTypeListRequest) Reset() {
	*x = EventTypeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTypeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTypeListRequest) ProtoMessage() {}

func (x *EventTypeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTypeListRequest.ProtoReflect.Descriptor instead.
func (*EventTypeListRequest) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_eventType_all_proto_rawDescGZIP(), []int{10}
}

func (x *EventTypeListRequest) GetSearchTerm() string {
	if x != nil {
		return x.SearchTerm
	}
	return ""
}

type EventTypeListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTypeFilterList *EventTypeFilterList `protobuf:"bytes,1,opt,name=EventTypeFilterList,proto3" json:"EventTypeFilterList,omitempty"`
}

func (x *EventTypeListResponse) Reset() {
	*x = EventTypeListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTypeListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTypeListResponse) ProtoMessage() {}

func (x *EventTypeListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_eventType_all_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTypeListResponse.ProtoReflect.Descriptor instead.
func (*EventTypeListResponse) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_eventType_all_proto_rawDescGZIP(), []int{11}
}

func (x *EventTypeListResponse) GetEventTypeFilterList() *EventTypeFilterList {
	if x != nil {
		return x.EventTypeFilterList
	}
	return nil
}

var File_event_management_gunk_v1_eventType_all_proto protoreflect.FileDescriptor

var file_event_management_gunk_v1_eventType_all_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x2f, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x02, 0x0a,
	0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x50, 0x00, 0x12, 0x21, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x39, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x12, 0x39, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x39, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22,
	0x78, 0x0a, 0x13, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x43, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x5a,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00,
	0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x36, 0x0a, 0x14, 0x45, 0x64,
	0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00,
	0x18, 0x00, 0x22, 0x58, 0x0a, 0x15, 0x45, 0x64, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x5b, 0x0a, 0x16,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x21,
	0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x5a, 0x0a, 0x17, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08,
	0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x38, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22,
	0x21, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00,
	0x18, 0x00, 0x22, 0x3e, 0x0a, 0x14, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00,
	0x18, 0x00, 0x22, 0x6a, 0x0a, 0x15, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x13, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x32, 0x9d,
	0x04, 0x0a, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x62, 0x0a,
	0x0d, 0x45, 0x64, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x64, 0x69,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e,
	0x45, 0x64, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x12, 0x68, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x68, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x70,
	0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x62, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88,
	0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x03, 0x88, 0x02, 0x00, 0x42, 0x49,
	0x48, 0x01, 0x50, 0x00, 0x5a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x70, 0x62, 0x80, 0x01, 0x00, 0x88, 0x01, 0x00, 0x90, 0x01, 0x00, 0xb8, 0x01, 0x00,
	0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01, 0xd0, 0x02, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_event_management_gunk_v1_eventType_all_proto_rawDescOnce sync.Once
	file_event_management_gunk_v1_eventType_all_proto_rawDescData = file_event_management_gunk_v1_eventType_all_proto_rawDesc
)

func file_event_management_gunk_v1_eventType_all_proto_rawDescGZIP() []byte {
	file_event_management_gunk_v1_eventType_all_proto_rawDescOnce.Do(func() {
		file_event_management_gunk_v1_eventType_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_management_gunk_v1_eventType_all_proto_rawDescData)
	})
	return file_event_management_gunk_v1_eventType_all_proto_rawDescData
}

var (
	file_event_management_gunk_v1_eventType_all_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
	file_event_management_gunk_v1_eventType_all_proto_goTypes  = []interface{}{
		(*EventTypes)(nil),              // 0: eventTypepb.EventTypes
		(*EventTypeFilterList)(nil),     // 1: eventTypepb.EventTypeFilterList
		(*CreateEventTypeRequest)(nil),  // 2: eventTypepb.CreateEventTypeRequest
		(*CreateEventTypeResponse)(nil), // 3: eventTypepb.CreateEventTypeResponse
		(*EditEventTypeRequest)(nil),    // 4: eventTypepb.EditEventTypeRequest
		(*EditEventTypeResponse)(nil),   // 5: eventTypepb.EditEventTypeResponse
		(*UpdateEventTypeRequest)(nil),  // 6: eventTypepb.UpdateEventTypeRequest
		(*UpdateEventTypeResponse)(nil), // 7: eventTypepb.UpdateEventTypeResponse
		(*DeleteEventTypeRequest)(nil),  // 8: eventTypepb.DeleteEventTypeRequest
		(*DeleteEventTypeResponse)(nil), // 9: eventTypepb.DeleteEventTypeResponse
		(*EventTypeListRequest)(nil),    // 10: eventTypepb.EventTypeListRequest
		(*EventTypeListResponse)(nil),   // 11: eventTypepb.EventTypeListResponse
		(*timestamppb.Timestamp)(nil),   // 12: google.protobuf.Timestamp
	}
)

var file_event_management_gunk_v1_eventType_all_proto_depIdxs = []int32{
	12, // 0: eventTypepb.EventTypes.CreatedAt:type_name -> google.protobuf.Timestamp
	12, // 1: eventTypepb.EventTypes.UpdatedAt:type_name -> google.protobuf.Timestamp
	12, // 2: eventTypepb.EventTypes.DeletedAt:type_name -> google.protobuf.Timestamp
	0,  // 3: eventTypepb.EventTypeFilterList.AllEventType:type_name -> eventTypepb.EventTypes
	0,  // 4: eventTypepb.CreateEventTypeResponse.EventTypes:type_name -> eventTypepb.EventTypes
	0,  // 5: eventTypepb.EditEventTypeResponse.EventTypes:type_name -> eventTypepb.EventTypes
	0,  // 6: eventTypepb.UpdateEventTypeResponse.EventTypes:type_name -> eventTypepb.EventTypes
	1,  // 7: eventTypepb.EventTypeListResponse.EventTypeFilterList:type_name -> eventTypepb.EventTypeFilterList
	2,  // 8: eventTypepb.EventTypeService.CreateEventType:input_type -> eventTypepb.CreateEventTypeRequest
	4,  // 9: eventTypepb.EventTypeService.EditEventType:input_type -> eventTypepb.EditEventTypeRequest
	6,  // 10: eventTypepb.EventTypeService.UpdateEventType:input_type -> eventTypepb.UpdateEventTypeRequest
	8,  // 11: eventTypepb.EventTypeService.DeleteEventType:input_type -> eventTypepb.DeleteEventTypeRequest
	10, // 12: eventTypepb.EventTypeService.EventTypeList:input_type -> eventTypepb.EventTypeListRequest
	3,  // 13: eventTypepb.EventTypeService.CreateEventType:output_type -> eventTypepb.CreateEventTypeResponse
	5,  // 14: eventTypepb.EventTypeService.EditEventType:output_type -> eventTypepb.EditEventTypeResponse
	7,  // 15: eventTypepb.EventTypeService.UpdateEventType:output_type -> eventTypepb.UpdateEventTypeResponse
	9,  // 16: eventTypepb.EventTypeService.DeleteEventType:output_type -> eventTypepb.DeleteEventTypeResponse
	11, // 17: eventTypepb.EventTypeService.EventTypeList:output_type -> eventTypepb.EventTypeListResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_event_management_gunk_v1_eventType_all_proto_init() }
func file_event_management_gunk_v1_eventType_all_proto_init() {
	if File_event_management_gunk_v1_eventType_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_management_gunk_v1_eventType_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTypes); i {
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
		file_event_management_gunk_v1_eventType_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTypeFilterList); i {
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
		file_event_management_gunk_v1_eventType_all_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventTypeRequest); i {
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
		file_event_management_gunk_v1_eventType_all_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventTypeResponse); i {
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
		file_event_management_gunk_v1_eventType_all_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditEventTypeRequest); i {
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
		file_event_management_gunk_v1_eventType_all_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditEventTypeResponse); i {
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
		file_event_management_gunk_v1_eventType_all_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventTypeRequest); i {
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
		file_event_management_gunk_v1_eventType_all_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventTypeResponse); i {
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
		file_event_management_gunk_v1_eventType_all_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventTypeRequest); i {
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
		file_event_management_gunk_v1_eventType_all_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventTypeResponse); i {
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
		file_event_management_gunk_v1_eventType_all_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTypeListRequest); i {
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
		file_event_management_gunk_v1_eventType_all_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTypeListResponse); i {
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
			RawDescriptor: file_event_management_gunk_v1_eventType_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_management_gunk_v1_eventType_all_proto_goTypes,
		DependencyIndexes: file_event_management_gunk_v1_eventType_all_proto_depIdxs,
		MessageInfos:      file_event_management_gunk_v1_eventType_all_proto_msgTypes,
	}.Build()
	File_event_management_gunk_v1_eventType_all_proto = out.File
	file_event_management_gunk_v1_eventType_all_proto_rawDesc = nil
	file_event_management_gunk_v1_eventType_all_proto_goTypes = nil
	file_event_management_gunk_v1_eventType_all_proto_depIdxs = nil
}