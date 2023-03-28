// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: event-management/gunk/v1/comment/all.proto

package commentpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Comments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserId  int32  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	EventID int32  `protobuf:"varint,3,opt,name=EventID,proto3" json:"EventID,omitempty"`
	Comment string `protobuf:"bytes,4,opt,name=Comment,proto3" json:"Comment,omitempty"`
}

func (x *Comments) Reset() {
	*x = Comments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comments) ProtoMessage() {}

func (x *Comments) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comments.ProtoReflect.Descriptor instead.
func (*Comments) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_comment_all_proto_rawDescGZIP(), []int{0}
}

func (x *Comments) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Comments) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Comments) GetEventID() int32 {
	if x != nil {
		return x.EventID
	}
	return 0
}

func (x *Comments) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type CommentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserId   int32  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	IsAdmin  bool   `protobuf:"varint,4,opt,name=IsAdmin,proto3" json:"IsAdmin,omitempty"`
	EventID  int32  `protobuf:"varint,5,opt,name=EventID,proto3" json:"EventID,omitempty"`
	Comment  string `protobuf:"bytes,6,opt,name=Comment,proto3" json:"Comment,omitempty"`
}

func (x *CommentList) Reset() {
	*x = CommentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentList) ProtoMessage() {}

func (x *CommentList) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentList.ProtoReflect.Descriptor instead.
func (*CommentList) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_comment_all_proto_rawDescGZIP(), []int{1}
}

func (x *CommentList) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CommentList) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CommentList) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CommentList) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *CommentList) GetEventID() int32 {
	if x != nil {
		return x.EventID
	}
	return 0
}

func (x *CommentList) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type CreateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	EventID int32  `protobuf:"varint,2,opt,name=EventID,proto3" json:"EventID,omitempty"`
	Comment string `protobuf:"bytes,3,opt,name=Comment,proto3" json:"Comment,omitempty"`
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_comment_all_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCommentRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateCommentRequest) GetEventID() int32 {
	if x != nil {
		return x.EventID
	}
	return 0
}

func (x *CreateCommentRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type CreateCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments *Comments `protobuf:"bytes,1,opt,name=Comments,proto3" json:"Comments,omitempty"`
}

func (x *CreateCommentResponse) Reset() {
	*x = CreateCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentResponse) ProtoMessage() {}

func (x *CreateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentResponse.ProtoReflect.Descriptor instead.
func (*CreateCommentResponse) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_comment_all_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCommentResponse) GetComments() *Comments {
	if x != nil {
		return x.Comments
	}
	return nil
}

type EditCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *EditCommentRequest) Reset() {
	*x = EditCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditCommentRequest) ProtoMessage() {}

func (x *EditCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditCommentRequest.ProtoReflect.Descriptor instead.
func (*EditCommentRequest) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_comment_all_proto_rawDescGZIP(), []int{4}
}

func (x *EditCommentRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type EditCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments *Comments `protobuf:"bytes,1,opt,name=Comments,proto3" json:"Comments,omitempty"`
}

func (x *EditCommentResponse) Reset() {
	*x = EditCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditCommentResponse) ProtoMessage() {}

func (x *EditCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditCommentResponse.ProtoReflect.Descriptor instead.
func (*EditCommentResponse) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_comment_all_proto_rawDescGZIP(), []int{5}
}

func (x *EditCommentResponse) GetComments() *Comments {
	if x != nil {
		return x.Comments
	}
	return nil
}

type UpdateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Comment string `protobuf:"bytes,2,opt,name=Comment,proto3" json:"Comment,omitempty"`
}

func (x *UpdateCommentRequest) Reset() {
	*x = UpdateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentRequest) ProtoMessage() {}

func (x *UpdateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentRequest.ProtoReflect.Descriptor instead.
func (*UpdateCommentRequest) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_comment_all_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCommentRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateCommentRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type UpdateCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments *Comments `protobuf:"bytes,1,opt,name=Comments,proto3" json:"Comments,omitempty"`
}

func (x *UpdateCommentResponse) Reset() {
	*x = UpdateCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentResponse) ProtoMessage() {}

func (x *UpdateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentResponse.ProtoReflect.Descriptor instead.
func (*UpdateCommentResponse) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_comment_all_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCommentResponse) GetComments() *Comments {
	if x != nil {
		return x.Comments
	}
	return nil
}

type CommentListOfEVentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventID int32 `protobuf:"varint,1,opt,name=EventID,proto3" json:"EventID,omitempty"`
}

func (x *CommentListOfEVentsRequest) Reset() {
	*x = CommentListOfEVentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListOfEVentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListOfEVentsRequest) ProtoMessage() {}

func (x *CommentListOfEVentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListOfEVentsRequest.ProtoReflect.Descriptor instead.
func (*CommentListOfEVentsRequest) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_comment_all_proto_rawDescGZIP(), []int{8}
}

func (x *CommentListOfEVentsRequest) GetEventID() int32 {
	if x != nil {
		return x.EventID
	}
	return 0
}

type CommentListOfEVentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CmtOfEv []*CommentList `protobuf:"bytes,1,rep,name=CmtOfEv,proto3" json:"CmtOfEv,omitempty"`
}

func (x *CommentListOfEVentsResponse) Reset() {
	*x = CommentListOfEVentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListOfEVentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListOfEVentsResponse) ProtoMessage() {}

func (x *CommentListOfEVentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_management_gunk_v1_comment_all_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListOfEVentsResponse.ProtoReflect.Descriptor instead.
func (*CommentListOfEVentsResponse) Descriptor() ([]byte, []int) {
	return file_event_management_gunk_v1_comment_all_proto_rawDescGZIP(), []int{9}
}

func (x *CommentListOfEVentsResponse) GetCmtOfEv() []*CommentList {
	if x != nil {
		return x.CmtOfEv
	}
	return nil
}

var File_event_management_gunk_v1_comment_all_proto protoreflect.FileDescriptor

var file_event_management_gunk_v1_comment_all_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1a, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1b, 0x0a, 0x07, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28,
	0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00,
	0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0xbe, 0x01, 0x0a, 0x0b, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00,
	0x50, 0x00, 0x12, 0x1a, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1c,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1b, 0x0a, 0x07,
	0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0a, 0x08,
	0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1b, 0x0a, 0x07, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x74, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12,
	0x1b, 0x0a, 0x07, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1b, 0x0a, 0x07,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08,
	0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18,
	0x00, 0x22, 0x52, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08,
	0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x34, 0x0a, 0x12, 0x45, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x50, 0x0a, 0x13, 0x45,
	0x64, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28,
	0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x53, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1b, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00,
	0x18, 0x00, 0x22, 0x52, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06,
	0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x41, 0x0a, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x45, 0x56, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x07, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x5a, 0x0a, 0x1b, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x45, 0x56, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x43, 0x6d, 0x74, 0x4f,
	0x66, 0x45, 0x76, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08,
	0x00, 0x10, 0x00, 0x18, 0x00, 0x32, 0xa1, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02,
	0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x58, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x12, 0x5e, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x12, 0x70, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x66, 0x45, 0x56, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x66, 0x45, 0x56, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x45, 0x56, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x1a, 0x03, 0x88, 0x02, 0x00, 0x42, 0x45, 0x48, 0x01, 0x50, 0x00, 0x5a,
	0x2a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x80, 0x01, 0x00, 0x88, 0x01,
	0x00, 0x90, 0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01, 0xd0, 0x02, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_management_gunk_v1_comment_all_proto_rawDescOnce sync.Once
	file_event_management_gunk_v1_comment_all_proto_rawDescData = file_event_management_gunk_v1_comment_all_proto_rawDesc
)

func file_event_management_gunk_v1_comment_all_proto_rawDescGZIP() []byte {
	file_event_management_gunk_v1_comment_all_proto_rawDescOnce.Do(func() {
		file_event_management_gunk_v1_comment_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_management_gunk_v1_comment_all_proto_rawDescData)
	})
	return file_event_management_gunk_v1_comment_all_proto_rawDescData
}

var (
	file_event_management_gunk_v1_comment_all_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
	file_event_management_gunk_v1_comment_all_proto_goTypes  = []interface{}{
		(*Comments)(nil),                    // 0: commentpb.Comments
		(*CommentList)(nil),                 // 1: commentpb.CommentList
		(*CreateCommentRequest)(nil),        // 2: commentpb.CreateCommentRequest
		(*CreateCommentResponse)(nil),       // 3: commentpb.CreateCommentResponse
		(*EditCommentRequest)(nil),          // 4: commentpb.EditCommentRequest
		(*EditCommentResponse)(nil),         // 5: commentpb.EditCommentResponse
		(*UpdateCommentRequest)(nil),        // 6: commentpb.UpdateCommentRequest
		(*UpdateCommentResponse)(nil),       // 7: commentpb.UpdateCommentResponse
		(*CommentListOfEVentsRequest)(nil),  // 8: commentpb.CommentListOfEVentsRequest
		(*CommentListOfEVentsResponse)(nil), // 9: commentpb.CommentListOfEVentsResponse
	}
)

var file_event_management_gunk_v1_comment_all_proto_depIdxs = []int32{
	0, // 0: commentpb.CreateCommentResponse.Comments:type_name -> commentpb.Comments
	0, // 1: commentpb.EditCommentResponse.Comments:type_name -> commentpb.Comments
	0, // 2: commentpb.UpdateCommentResponse.Comments:type_name -> commentpb.Comments
	1, // 3: commentpb.CommentListOfEVentsResponse.CmtOfEv:type_name -> commentpb.CommentList
	2, // 4: commentpb.CommentService.CreateComment:input_type -> commentpb.CreateCommentRequest
	4, // 5: commentpb.CommentService.EditComment:input_type -> commentpb.EditCommentRequest
	6, // 6: commentpb.CommentService.UpdateComment:input_type -> commentpb.UpdateCommentRequest
	8, // 7: commentpb.CommentService.CommentListOfEVents:input_type -> commentpb.CommentListOfEVentsRequest
	3, // 8: commentpb.CommentService.CreateComment:output_type -> commentpb.CreateCommentResponse
	5, // 9: commentpb.CommentService.EditComment:output_type -> commentpb.EditCommentResponse
	7, // 10: commentpb.CommentService.UpdateComment:output_type -> commentpb.UpdateCommentResponse
	9, // 11: commentpb.CommentService.CommentListOfEVents:output_type -> commentpb.CommentListOfEVentsResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_event_management_gunk_v1_comment_all_proto_init() }
func file_event_management_gunk_v1_comment_all_proto_init() {
	if File_event_management_gunk_v1_comment_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_management_gunk_v1_comment_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comments); i {
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
		file_event_management_gunk_v1_comment_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentList); i {
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
		file_event_management_gunk_v1_comment_all_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentRequest); i {
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
		file_event_management_gunk_v1_comment_all_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentResponse); i {
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
		file_event_management_gunk_v1_comment_all_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditCommentRequest); i {
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
		file_event_management_gunk_v1_comment_all_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditCommentResponse); i {
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
		file_event_management_gunk_v1_comment_all_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentRequest); i {
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
		file_event_management_gunk_v1_comment_all_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentResponse); i {
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
		file_event_management_gunk_v1_comment_all_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListOfEVentsRequest); i {
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
		file_event_management_gunk_v1_comment_all_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListOfEVentsResponse); i {
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
			RawDescriptor: file_event_management_gunk_v1_comment_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_management_gunk_v1_comment_all_proto_goTypes,
		DependencyIndexes: file_event_management_gunk_v1_comment_all_proto_depIdxs,
		MessageInfos:      file_event_management_gunk_v1_comment_all_proto_msgTypes,
	}.Build()
	File_event_management_gunk_v1_comment_all_proto = out.File
	file_event_management_gunk_v1_comment_all_proto_rawDesc = nil
	file_event_management_gunk_v1_comment_all_proto_goTypes = nil
	file_event_management_gunk_v1_comment_all_proto_depIdxs = nil
}
