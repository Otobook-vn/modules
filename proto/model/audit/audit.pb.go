// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: audit.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string       `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Target string       `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Id     string       `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Author *AuditAuthor `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Action string       `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
	Data   string       `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_audit_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CreateRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *CreateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateRequest) GetAuthor() *AuditAuthor {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *CreateRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *CreateRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_audit_proto_rawDescGZIP(), []int{1}
}

type GetByTargetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source    string                 `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Target    string                 `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Id        string                 `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetByTargetRequest) Reset() {
	*x = GetByTargetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByTargetRequest) ProtoMessage() {}

func (x *GetByTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByTargetRequest.ProtoReflect.Descriptor instead.
func (*GetByTargetRequest) Descriptor() ([]byte, []int) {
	return file_audit_proto_rawDescGZIP(), []int{2}
}

func (x *GetByTargetRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *GetByTargetRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *GetByTargetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetByTargetRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type GetByTargetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*AuditDetail `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetByTargetResponse) Reset() {
	*x = GetByTargetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByTargetResponse) ProtoMessage() {}

func (x *GetByTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByTargetResponse.ProtoReflect.Descriptor instead.
func (*GetByTargetResponse) Descriptor() ([]byte, []int) {
	return file_audit_proto_rawDescGZIP(), []int{3}
}

func (x *GetByTargetResponse) GetList() []*AuditDetail {
	if x != nil {
		return x.List
	}
	return nil
}

type AuditDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Author    *AuditAuthor           `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Target    string                 `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Action    string                 `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Data      string                 `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AuditDetail) Reset() {
	*x = AuditDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditDetail) ProtoMessage() {}

func (x *AuditDetail) ProtoReflect() protoreflect.Message {
	mi := &file_audit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditDetail.ProtoReflect.Descriptor instead.
func (*AuditDetail) Descriptor() ([]byte, []int) {
	return file_audit_proto_rawDescGZIP(), []int{4}
}

func (x *AuditDetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuditDetail) GetAuthor() *AuditAuthor {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *AuditDetail) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *AuditDetail) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *AuditDetail) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AuditDetail) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type AuditAuthor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AuditAuthor) Reset() {
	*x = AuditAuthor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditAuthor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditAuthor) ProtoMessage() {}

func (x *AuditAuthor) ProtoReflect() protoreflect.Message {
	mi := &file_audit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditAuthor.ProtoReflect.Descriptor instead.
func (*AuditAuthor) Descriptor() ([]byte, []int) {
	return file_audit_proto_rawDescGZIP(), []int{5}
}

func (x *AuditAuthor) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuditAuthor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_audit_proto protoreflect.FileDescriptor

var file_audit_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x8e, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x3d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0xc7, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x64, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x0b, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x84,
	0x01, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_audit_proto_rawDescOnce sync.Once
	file_audit_proto_rawDescData = file_audit_proto_rawDesc
)

func file_audit_proto_rawDescGZIP() []byte {
	file_audit_proto_rawDescOnce.Do(func() {
		file_audit_proto_rawDescData = protoimpl.X.CompressGZIP(file_audit_proto_rawDescData)
	})
	return file_audit_proto_rawDescData
}

var file_audit_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_audit_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),         // 0: proto.CreateRequest
	(*CreateResponse)(nil),        // 1: proto.CreateResponse
	(*GetByTargetRequest)(nil),    // 2: proto.GetByTargetRequest
	(*GetByTargetResponse)(nil),   // 3: proto.GetByTargetResponse
	(*AuditDetail)(nil),           // 4: proto.AuditDetail
	(*AuditAuthor)(nil),           // 5: proto.AuditAuthor
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_audit_proto_depIdxs = []int32{
	5, // 0: proto.CreateRequest.author:type_name -> proto.AuditAuthor
	6, // 1: proto.GetByTargetRequest.timestamp:type_name -> google.protobuf.Timestamp
	4, // 2: proto.GetByTargetResponse.list:type_name -> proto.AuditDetail
	5, // 3: proto.AuditDetail.author:type_name -> proto.AuditAuthor
	6, // 4: proto.AuditDetail.createdAt:type_name -> google.protobuf.Timestamp
	0, // 5: proto.Audit.Create:input_type -> proto.CreateRequest
	2, // 6: proto.Audit.GetByTarget:input_type -> proto.GetByTargetRequest
	1, // 7: proto.Audit.Create:output_type -> proto.CreateResponse
	3, // 8: proto.Audit.GetByTarget:output_type -> proto.GetByTargetResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_audit_proto_init() }
func file_audit_proto_init() {
	if File_audit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_audit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_audit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_audit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByTargetRequest); i {
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
		file_audit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByTargetResponse); i {
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
		file_audit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditDetail); i {
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
		file_audit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditAuthor); i {
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
			RawDescriptor: file_audit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_audit_proto_goTypes,
		DependencyIndexes: file_audit_proto_depIdxs,
		MessageInfos:      file_audit_proto_msgTypes,
	}.Build()
	File_audit_proto = out.File
	file_audit_proto_rawDesc = nil
	file_audit_proto_goTypes = nil
	file_audit_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuditClient is the client API for Audit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuditClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	GetByTarget(ctx context.Context, in *GetByTargetRequest, opts ...grpc.CallOption) (*GetByTargetResponse, error)
}

type auditClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditClient(cc grpc.ClientConnInterface) AuditClient {
	return &auditClient{cc}
}

func (c *auditClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/proto.Audit/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditClient) GetByTarget(ctx context.Context, in *GetByTargetRequest, opts ...grpc.CallOption) (*GetByTargetResponse, error) {
	out := new(GetByTargetResponse)
	err := c.cc.Invoke(ctx, "/proto.Audit/GetByTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditServer is the server API for Audit service.
type AuditServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	GetByTarget(context.Context, *GetByTargetRequest) (*GetByTargetResponse, error)
}

// UnimplementedAuditServer can be embedded to have forward compatible implementations.
type UnimplementedAuditServer struct {
}

func (*UnimplementedAuditServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAuditServer) GetByTarget(context.Context, *GetByTargetRequest) (*GetByTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByTarget not implemented")
}

func RegisterAuditServer(s *grpc.Server, srv AuditServer) {
	s.RegisterService(&_Audit_serviceDesc, srv)
}

func _Audit_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Audit/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Audit_GetByTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServer).GetByTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Audit/GetByTarget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServer).GetByTarget(ctx, req.(*GetByTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Audit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Audit",
	HandlerType: (*AuditServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Audit_Create_Handler,
		},
		{
			MethodName: "GetByTarget",
			Handler:    _Audit_GetByTarget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audit.proto",
}
