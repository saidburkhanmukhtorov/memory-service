// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: submodule-for-timecapsule/memory_service/comment.proto

package memory

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

// Comment represents a comment on a memory.
type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MemoryId  string `protobuf:"bytes,2,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	UserId    string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content   string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_memory_service_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Comment) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

func (x *Comment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Comment) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// GetCommentByIdRequest represents a request to retrieve a comment by its ID.
type GetCommentByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCommentByIdRequest) Reset() {
	*x = GetCommentByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentByIdRequest) ProtoMessage() {}

func (x *GetCommentByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCommentByIdRequest) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_memory_service_comment_proto_rawDescGZIP(), []int{1}
}

func (x *GetCommentByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteCommentRequest represents a request to delete a comment by its ID.
type DeleteCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCommentRequest) Reset() {
	*x = DeleteCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentRequest) ProtoMessage() {}

func (x *DeleteCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentRequest.ProtoReflect.Descriptor instead.
func (*DeleteCommentRequest) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_memory_service_comment_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCommentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteCommentResponse represents a response to a comment deletion request.
type DeleteCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteCommentResponse) Reset() {
	*x = DeleteCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentResponse) ProtoMessage() {}

func (x *DeleteCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentResponse.ProtoReflect.Descriptor instead.
func (*DeleteCommentResponse) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_memory_service_comment_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCommentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// GetAllCommentsRequest represents a request to retrieve all comments.
type GetAllCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit    int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	MemoryId string `protobuf:"bytes,3,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"` // Filter by memory ID
	UserId   string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`       // Filter by user ID
	Content  string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`                   // Filter by content
}

func (x *GetAllCommentsRequest) Reset() {
	*x = GetAllCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCommentsRequest) ProtoMessage() {}

func (x *GetAllCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetAllCommentsRequest) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_memory_service_comment_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllCommentsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllCommentsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllCommentsRequest) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

func (x *GetAllCommentsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetAllCommentsRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// GetAllCommentsResponse represents a response containing a list of comments.
type GetAllCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Count    int32      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllCommentsResponse) Reset() {
	*x = GetAllCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCommentsResponse) ProtoMessage() {}

func (x *GetAllCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCommentsResponse.ProtoReflect.Descriptor instead.
func (*GetAllCommentsResponse) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_memory_service_comment_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllCommentsResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *GetAllCommentsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_submodule_for_timecapsule_memory_service_comment_proto protoreflect.FileDescriptor

var file_submodule_for_timecapsule_memory_service_comment_proto_rawDesc = []byte{
	0x0a, 0x36, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2d, 0x66, 0x6f, 0x72, 0x2d,
	0x74, 0x69, 0x6d, 0x65, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2f, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x22, 0xa7, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x15, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x91,
	0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32,
	0xf1, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_submodule_for_timecapsule_memory_service_comment_proto_rawDescOnce sync.Once
	file_submodule_for_timecapsule_memory_service_comment_proto_rawDescData = file_submodule_for_timecapsule_memory_service_comment_proto_rawDesc
)

func file_submodule_for_timecapsule_memory_service_comment_proto_rawDescGZIP() []byte {
	file_submodule_for_timecapsule_memory_service_comment_proto_rawDescOnce.Do(func() {
		file_submodule_for_timecapsule_memory_service_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_submodule_for_timecapsule_memory_service_comment_proto_rawDescData)
	})
	return file_submodule_for_timecapsule_memory_service_comment_proto_rawDescData
}

var file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_submodule_for_timecapsule_memory_service_comment_proto_goTypes = []any{
	(*Comment)(nil),                // 0: memory.Comment
	(*GetCommentByIdRequest)(nil),  // 1: memory.GetCommentByIdRequest
	(*DeleteCommentRequest)(nil),   // 2: memory.DeleteCommentRequest
	(*DeleteCommentResponse)(nil),  // 3: memory.DeleteCommentResponse
	(*GetAllCommentsRequest)(nil),  // 4: memory.GetAllCommentsRequest
	(*GetAllCommentsResponse)(nil), // 5: memory.GetAllCommentsResponse
}
var file_submodule_for_timecapsule_memory_service_comment_proto_depIdxs = []int32{
	0, // 0: memory.GetAllCommentsResponse.comments:type_name -> memory.Comment
	1, // 1: memory.CommentService.GetCommentById:input_type -> memory.GetCommentByIdRequest
	2, // 2: memory.CommentService.DeleteComment:input_type -> memory.DeleteCommentRequest
	4, // 3: memory.CommentService.GetAllComments:input_type -> memory.GetAllCommentsRequest
	0, // 4: memory.CommentService.GetCommentById:output_type -> memory.Comment
	3, // 5: memory.CommentService.DeleteComment:output_type -> memory.DeleteCommentResponse
	5, // 6: memory.CommentService.GetAllComments:output_type -> memory.GetAllCommentsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_submodule_for_timecapsule_memory_service_comment_proto_init() }
func file_submodule_for_timecapsule_memory_service_comment_proto_init() {
	if File_submodule_for_timecapsule_memory_service_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Comment); i {
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
		file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetCommentByIdRequest); i {
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
		file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteCommentRequest); i {
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
		file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteCommentResponse); i {
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
		file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllCommentsRequest); i {
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
		file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllCommentsResponse); i {
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
			RawDescriptor: file_submodule_for_timecapsule_memory_service_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_submodule_for_timecapsule_memory_service_comment_proto_goTypes,
		DependencyIndexes: file_submodule_for_timecapsule_memory_service_comment_proto_depIdxs,
		MessageInfos:      file_submodule_for_timecapsule_memory_service_comment_proto_msgTypes,
	}.Build()
	File_submodule_for_timecapsule_memory_service_comment_proto = out.File
	file_submodule_for_timecapsule_memory_service_comment_proto_rawDesc = nil
	file_submodule_for_timecapsule_memory_service_comment_proto_goTypes = nil
	file_submodule_for_timecapsule_memory_service_comment_proto_depIdxs = nil
}
