// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: photo.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetPhotoByIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetPhotoByIDsRequest) Reset() {
	*x = GetPhotoByIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPhotoByIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhotoByIDsRequest) ProtoMessage() {}

func (x *GetPhotoByIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_photo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhotoByIDsRequest.ProtoReflect.Descriptor instead.
func (*GetPhotoByIDsRequest) Descriptor() ([]byte, []int) {
	return file_photo_proto_rawDescGZIP(), []int{0}
}

func (x *GetPhotoByIDsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetPhotoByIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Photos []*PhotoDetail `protobuf:"bytes,1,rep,name=photos,proto3" json:"photos,omitempty"`
}

func (x *GetPhotoByIDsResponse) Reset() {
	*x = GetPhotoByIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPhotoByIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhotoByIDsResponse) ProtoMessage() {}

func (x *GetPhotoByIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_photo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhotoByIDsResponse.ProtoReflect.Descriptor instead.
func (*GetPhotoByIDsResponse) Descriptor() ([]byte, []int) {
	return file_photo_proto_rawDescGZIP(), []int{1}
}

func (x *GetPhotoByIDsResponse) GetPhotos() []*PhotoDetail {
	if x != nil {
		return x.Photos
	}
	return nil
}

type PhotoDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Small  *PhotoSize `protobuf:"bytes,3,opt,name=small,proto3" json:"small,omitempty"`
	Medium *PhotoSize `protobuf:"bytes,4,opt,name=medium,proto3" json:"medium,omitempty"`
}

func (x *PhotoDetail) Reset() {
	*x = PhotoDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoDetail) ProtoMessage() {}

func (x *PhotoDetail) ProtoReflect() protoreflect.Message {
	mi := &file_photo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoDetail.ProtoReflect.Descriptor instead.
func (*PhotoDetail) Descriptor() ([]byte, []int) {
	return file_photo_proto_rawDescGZIP(), []int{2}
}

func (x *PhotoDetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PhotoDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PhotoDetail) GetSmall() *PhotoSize {
	if x != nil {
		return x.Small
	}
	return nil
}

func (x *PhotoDetail) GetMedium() *PhotoSize {
	if x != nil {
		return x.Medium
	}
	return nil
}

type PhotoSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  int32 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height int32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *PhotoSize) Reset() {
	*x = PhotoSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoSize) ProtoMessage() {}

func (x *PhotoSize) ProtoReflect() protoreflect.Message {
	mi := &file_photo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoSize.ProtoReflect.Descriptor instead.
func (*PhotoSize) Descriptor() ([]byte, []int) {
	return file_photo_proto_rawDescGZIP(), []int{3}
}

func (x *PhotoSize) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *PhotoSize) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_photo_proto protoreflect.FileDescriptor

var file_photo_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x43,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x05, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x12,
	0x28, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x53, 0x69, 0x7a,
	0x65, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x22, 0x39, 0x0a, 0x09, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x32, 0x53, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x4a, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x73, 0x12, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x42,
	0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x42, 0x79, 0x49, 0x44,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_photo_proto_rawDescOnce sync.Once
	file_photo_proto_rawDescData = file_photo_proto_rawDesc
)

func file_photo_proto_rawDescGZIP() []byte {
	file_photo_proto_rawDescOnce.Do(func() {
		file_photo_proto_rawDescData = protoimpl.X.CompressGZIP(file_photo_proto_rawDescData)
	})
	return file_photo_proto_rawDescData
}

var file_photo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_photo_proto_goTypes = []interface{}{
	(*GetPhotoByIDsRequest)(nil),  // 0: proto.GetPhotoByIDsRequest
	(*GetPhotoByIDsResponse)(nil), // 1: proto.GetPhotoByIDsResponse
	(*PhotoDetail)(nil),           // 2: proto.PhotoDetail
	(*PhotoSize)(nil),             // 3: proto.PhotoSize
}
var file_photo_proto_depIdxs = []int32{
	2, // 0: proto.GetPhotoByIDsResponse.photos:type_name -> proto.PhotoDetail
	3, // 1: proto.PhotoDetail.small:type_name -> proto.PhotoSize
	3, // 2: proto.PhotoDetail.medium:type_name -> proto.PhotoSize
	0, // 3: proto.Photo.GetPhotoByIDs:input_type -> proto.GetPhotoByIDsRequest
	1, // 4: proto.Photo.GetPhotoByIDs:output_type -> proto.GetPhotoByIDsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_photo_proto_init() }
func file_photo_proto_init() {
	if File_photo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_photo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPhotoByIDsRequest); i {
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
		file_photo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPhotoByIDsResponse); i {
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
		file_photo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoDetail); i {
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
		file_photo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoSize); i {
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
			RawDescriptor: file_photo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_photo_proto_goTypes,
		DependencyIndexes: file_photo_proto_depIdxs,
		MessageInfos:      file_photo_proto_msgTypes,
	}.Build()
	File_photo_proto = out.File
	file_photo_proto_rawDesc = nil
	file_photo_proto_goTypes = nil
	file_photo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PhotoClient is the client API for Photo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PhotoClient interface {
	GetPhotoByIDs(ctx context.Context, in *GetPhotoByIDsRequest, opts ...grpc.CallOption) (*GetPhotoByIDsResponse, error)
}

type photoClient struct {
	cc grpc.ClientConnInterface
}

func NewPhotoClient(cc grpc.ClientConnInterface) PhotoClient {
	return &photoClient{cc}
}

func (c *photoClient) GetPhotoByIDs(ctx context.Context, in *GetPhotoByIDsRequest, opts ...grpc.CallOption) (*GetPhotoByIDsResponse, error) {
	out := new(GetPhotoByIDsResponse)
	err := c.cc.Invoke(ctx, "/proto.Photo/GetPhotoByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhotoServer is the server API for Photo service.
type PhotoServer interface {
	GetPhotoByIDs(context.Context, *GetPhotoByIDsRequest) (*GetPhotoByIDsResponse, error)
}

// UnimplementedPhotoServer can be embedded to have forward compatible implementations.
type UnimplementedPhotoServer struct {
}

func (*UnimplementedPhotoServer) GetPhotoByIDs(context.Context, *GetPhotoByIDsRequest) (*GetPhotoByIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhotoByIDs not implemented")
}

func RegisterPhotoServer(s *grpc.Server, srv PhotoServer) {
	s.RegisterService(&_Photo_serviceDesc, srv)
}

func _Photo_GetPhotoByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPhotoByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServer).GetPhotoByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Photo/GetPhotoByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServer).GetPhotoByIDs(ctx, req.(*GetPhotoByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Photo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Photo",
	HandlerType: (*PhotoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPhotoByIDs",
			Handler:    _Photo_GetPhotoByIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "photo.proto",
}
