// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: debug.proto

package looprpc

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

type ForceAutoLoopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ForceAutoLoopRequest) Reset() {
	*x = ForceAutoLoopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceAutoLoopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceAutoLoopRequest) ProtoMessage() {}

func (x *ForceAutoLoopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_debug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceAutoLoopRequest.ProtoReflect.Descriptor instead.
func (*ForceAutoLoopRequest) Descriptor() ([]byte, []int) {
	return file_debug_proto_rawDescGZIP(), []int{0}
}

type ForceAutoLoopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ForceAutoLoopResponse) Reset() {
	*x = ForceAutoLoopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceAutoLoopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceAutoLoopResponse) ProtoMessage() {}

func (x *ForceAutoLoopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_debug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceAutoLoopResponse.ProtoReflect.Descriptor instead.
func (*ForceAutoLoopResponse) Descriptor() ([]byte, []int) {
	return file_debug_proto_rawDescGZIP(), []int{1}
}

var File_debug_proto protoreflect.FileDescriptor

var file_debug_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c,
	0x6f, 0x6f, 0x70, 0x72, 0x70, 0x63, 0x22, 0x16, 0x0a, 0x14, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x41,
	0x75, 0x74, 0x6f, 0x4c, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17,
	0x0a, 0x15, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x4c, 0x6f, 0x6f, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x57, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x12, 0x4e, 0x0a, 0x0d, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x4c, 0x6f, 0x6f,
	0x70, 0x12, 0x1d, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x6f, 0x72, 0x63,
	0x65, 0x41, 0x75, 0x74, 0x6f, 0x4c, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x6f, 0x72, 0x63, 0x65,
	0x41, 0x75, 0x74, 0x6f, 0x4c, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x6c, 0x6f, 0x6f,
	0x70, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_debug_proto_rawDescOnce sync.Once
	file_debug_proto_rawDescData = file_debug_proto_rawDesc
)

func file_debug_proto_rawDescGZIP() []byte {
	file_debug_proto_rawDescOnce.Do(func() {
		file_debug_proto_rawDescData = protoimpl.X.CompressGZIP(file_debug_proto_rawDescData)
	})
	return file_debug_proto_rawDescData
}

var file_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_debug_proto_goTypes = []interface{}{
	(*ForceAutoLoopRequest)(nil),  // 0: looprpc.ForceAutoLoopRequest
	(*ForceAutoLoopResponse)(nil), // 1: looprpc.ForceAutoLoopResponse
}
var file_debug_proto_depIdxs = []int32{
	0, // 0: looprpc.Debug.ForceAutoLoop:input_type -> looprpc.ForceAutoLoopRequest
	1, // 1: looprpc.Debug.ForceAutoLoop:output_type -> looprpc.ForceAutoLoopResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_debug_proto_init() }
func file_debug_proto_init() {
	if File_debug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_debug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceAutoLoopRequest); i {
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
		file_debug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceAutoLoopResponse); i {
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
			RawDescriptor: file_debug_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_debug_proto_goTypes,
		DependencyIndexes: file_debug_proto_depIdxs,
		MessageInfos:      file_debug_proto_msgTypes,
	}.Build()
	File_debug_proto = out.File
	file_debug_proto_rawDesc = nil
	file_debug_proto_goTypes = nil
	file_debug_proto_depIdxs = nil
}