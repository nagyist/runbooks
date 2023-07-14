// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: sci.proto

package sci

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

type CreateSignedURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName        string `protobuf:"bytes,1,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	ObjectName        string `protobuf:"bytes,2,opt,name=objectName,proto3" json:"objectName,omitempty"`
	ExpirationSeconds int64  `protobuf:"varint,3,opt,name=expirationSeconds,proto3" json:"expirationSeconds,omitempty"`
}

func (x *CreateSignedURLRequest) Reset() {
	*x = CreateSignedURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sci_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSignedURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSignedURLRequest) ProtoMessage() {}

func (x *CreateSignedURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sci_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSignedURLRequest.ProtoReflect.Descriptor instead.
func (*CreateSignedURLRequest) Descriptor() ([]byte, []int) {
	return file_sci_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSignedURLRequest) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *CreateSignedURLRequest) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *CreateSignedURLRequest) GetExpirationSeconds() int64 {
	if x != nil {
		return x.ExpirationSeconds
	}
	return 0
}

type CreateSignedURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateSignedURLResponse) Reset() {
	*x = CreateSignedURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sci_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSignedURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSignedURLResponse) ProtoMessage() {}

func (x *CreateSignedURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sci_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSignedURLResponse.ProtoReflect.Descriptor instead.
func (*CreateSignedURLResponse) Descriptor() ([]byte, []int) {
	return file_sci_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSignedURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_sci_proto protoreflect.FileDescriptor

var file_sci_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x63, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x63, 0x69,
	0x2e, 0x76, 0x31, 0x22, 0x86, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c,
	0x0a, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x2b, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0x62, 0x0a, 0x0a, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x12, 0x1e, 0x2e, 0x73, 0x63, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x63, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x62, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x75, 0x73, 0x61, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x75, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x63, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sci_proto_rawDescOnce sync.Once
	file_sci_proto_rawDescData = file_sci_proto_rawDesc
)

func file_sci_proto_rawDescGZIP() []byte {
	file_sci_proto_rawDescOnce.Do(func() {
		file_sci_proto_rawDescData = protoimpl.X.CompressGZIP(file_sci_proto_rawDescData)
	})
	return file_sci_proto_rawDescData
}

var file_sci_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sci_proto_goTypes = []interface{}{
	(*CreateSignedURLRequest)(nil),  // 0: sci.v1.CreateSignedURLRequest
	(*CreateSignedURLResponse)(nil), // 1: sci.v1.CreateSignedURLResponse
}
var file_sci_proto_depIdxs = []int32{
	0, // 0: sci.v1.Controller.CreateSignedURL:input_type -> sci.v1.CreateSignedURLRequest
	1, // 1: sci.v1.Controller.CreateSignedURL:output_type -> sci.v1.CreateSignedURLResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sci_proto_init() }
func file_sci_proto_init() {
	if File_sci_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sci_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSignedURLRequest); i {
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
		file_sci_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSignedURLResponse); i {
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
			RawDescriptor: file_sci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sci_proto_goTypes,
		DependencyIndexes: file_sci_proto_depIdxs,
		MessageInfos:      file_sci_proto_msgTypes,
	}.Build()
	File_sci_proto = out.File
	file_sci_proto_rawDesc = nil
	file_sci_proto_goTypes = nil
	file_sci_proto_depIdxs = nil
}
