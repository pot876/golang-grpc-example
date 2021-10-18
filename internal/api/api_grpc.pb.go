// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api_grpc.proto

package api

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

type FiboRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From uint64 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To   uint64 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *FiboRequest) Reset() {
	*x = FiboRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FiboRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FiboRequest) ProtoMessage() {}

func (x *FiboRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FiboRequest.ProtoReflect.Descriptor instead.
func (*FiboRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *FiboRequest) GetFrom() uint64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *FiboRequest) GetTo() uint64 {
	if x != nil {
		return x.To
	}
	return 0
}

type FiboReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numbers []string `protobuf:"bytes,1,rep,name=numbers,proto3" json:"numbers,omitempty"`
}

func (x *FiboReply) Reset() {
	*x = FiboReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FiboReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FiboReply) ProtoMessage() {}

func (x *FiboReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FiboReply.ProtoReflect.Descriptor instead.
func (*FiboReply) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *FiboReply) GetNumbers() []string {
	if x != nil {
		return x.Numbers
	}
	return nil
}

var File_api_grpc_proto protoreflect.FileDescriptor

var file_api_grpc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x31, 0x0a, 0x0b, 0x46, 0x69, 0x62, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x09, 0x46, 0x69, 0x62, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x32, 0x6a, 0x0a, 0x04, 0x46, 0x69,
	0x62, 0x6f, 0x12, 0x2c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x62, 0x6f, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x12, 0x0c, 0x2e, 0x46, 0x69, 0x62, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x46, 0x69, 0x62, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x69, 0x62, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0c, 0x2e, 0x46, 0x69, 0x62, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x46, 0x69, 0x62, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x42, 0x17, 0x5a, 0x15, 0x66, 0x69, 0x62, 0x6f, 0x2d, 0x70,
	0x72, 0x6a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_proto_rawDescOnce sync.Once
	file_api_grpc_proto_rawDescData = file_api_grpc_proto_rawDesc
)

func file_api_grpc_proto_rawDescGZIP() []byte {
	file_api_grpc_proto_rawDescOnce.Do(func() {
		file_api_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_proto_rawDescData)
	})
	return file_api_grpc_proto_rawDescData
}

var file_api_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_grpc_proto_goTypes = []interface{}{
	(*FiboRequest)(nil), // 0: FiboRequest
	(*FiboReply)(nil),   // 1: FiboReply
}
var file_api_grpc_proto_depIdxs = []int32{
	0, // 0: Fibo.GetFiboNumbers:input_type -> FiboRequest
	0, // 1: Fibo.GetFiboNumbersStream:input_type -> FiboRequest
	1, // 2: Fibo.GetFiboNumbers:output_type -> FiboReply
	1, // 3: Fibo.GetFiboNumbersStream:output_type -> FiboReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_grpc_proto_init() }
func file_api_grpc_proto_init() {
	if File_api_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FiboRequest); i {
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
		file_api_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FiboReply); i {
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
			RawDescriptor: file_api_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_proto_goTypes,
		DependencyIndexes: file_api_grpc_proto_depIdxs,
		MessageInfos:      file_api_grpc_proto_msgTypes,
	}.Build()
	File_api_grpc_proto = out.File
	file_api_grpc_proto_rawDesc = nil
	file_api_grpc_proto_goTypes = nil
	file_api_grpc_proto_depIdxs = nil
}
