// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: protos/global_grpc.proto

package grpc

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

type GlobalGrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RpcKey  string `protobuf:"bytes,1,opt,name=rpc_key,json=rpcKey,proto3" json:"rpc_key,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GlobalGrpcRequest) Reset() {
	*x = GlobalGrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_global_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalGrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalGrpcRequest) ProtoMessage() {}

func (x *GlobalGrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_global_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalGrpcRequest.ProtoReflect.Descriptor instead.
func (*GlobalGrpcRequest) Descriptor() ([]byte, []int) {
	return file_protos_global_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *GlobalGrpcRequest) GetRpcKey() string {
	if x != nil {
		return x.RpcKey
	}
	return ""
}

func (x *GlobalGrpcRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GlobalGrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GlobalGrpcResponse) Reset() {
	*x = GlobalGrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_global_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalGrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalGrpcResponse) ProtoMessage() {}

func (x *GlobalGrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_global_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalGrpcResponse.ProtoReflect.Descriptor instead.
func (*GlobalGrpcResponse) Descriptor() ([]byte, []int) {
	return file_protos_global_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *GlobalGrpcResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_global_grpc_proto protoreflect.FileDescriptor

var file_protos_global_grpc_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x46, 0x0a, 0x11, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x47,
	0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x70,
	0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x70, 0x63,
	0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a,
	0x12, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb1, 0x01,
	0x0a, 0x11, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x47, 0x52, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x47, 0x52, 0x70,
	0x63, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51,
	0x0a, 0x10, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x47, 0x72, 0x70, 0x63, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_global_grpc_proto_rawDescOnce sync.Once
	file_protos_global_grpc_proto_rawDescData = file_protos_global_grpc_proto_rawDesc
)

func file_protos_global_grpc_proto_rawDescGZIP() []byte {
	file_protos_global_grpc_proto_rawDescOnce.Do(func() {
		file_protos_global_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_global_grpc_proto_rawDescData)
	})
	return file_protos_global_grpc_proto_rawDescData
}

var file_protos_global_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_global_grpc_proto_goTypes = []interface{}{
	(*GlobalGrpcRequest)(nil),  // 0: main_grpc.GlobalGrpcRequest
	(*GlobalGrpcResponse)(nil), // 1: main_grpc.GlobalGrpcResponse
}
var file_protos_global_grpc_proto_depIdxs = []int32{
	0, // 0: main_grpc.GlobalGRpcService.GlobalGRpc:input_type -> main_grpc.GlobalGrpcRequest
	0, // 1: main_grpc.GlobalGRpcService.GlobalGrpcStream:input_type -> main_grpc.GlobalGrpcRequest
	1, // 2: main_grpc.GlobalGRpcService.GlobalGRpc:output_type -> main_grpc.GlobalGrpcResponse
	1, // 3: main_grpc.GlobalGRpcService.GlobalGrpcStream:output_type -> main_grpc.GlobalGrpcResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_global_grpc_proto_init() }
func file_protos_global_grpc_proto_init() {
	if File_protos_global_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_global_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalGrpcRequest); i {
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
		file_protos_global_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalGrpcResponse); i {
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
			RawDescriptor: file_protos_global_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_global_grpc_proto_goTypes,
		DependencyIndexes: file_protos_global_grpc_proto_depIdxs,
		MessageInfos:      file_protos_global_grpc_proto_msgTypes,
	}.Build()
	File_protos_global_grpc_proto = out.File
	file_protos_global_grpc_proto_rawDesc = nil
	file_protos_global_grpc_proto_goTypes = nil
	file_protos_global_grpc_proto_depIdxs = nil
}
