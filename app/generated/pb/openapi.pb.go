// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: openapi.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_openapi_proto protoreflect.FileDescriptor

var file_openapi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x98, 0x02, 0x92, 0x41, 0x8c, 0x01, 0x12,
	0x89, 0x01, 0x0a, 0x12, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x20, 0x48, 0x54,
	0x54, 0x50, 0x20, 0x41, 0x50, 0x49, 0x22, 0x2e, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61,
	0x6e, 0x65, 0x6c, 0x12, 0x21, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65,
	0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0x3e, 0x0a, 0x0a, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x20, 0x32, 0x2e, 0x30, 0x12, 0x30, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x73, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x2d, 0x32, 0x2e,
	0x30, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x0a, 0x0d, 0x63, 0x6f, 0x6d,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x42, 0x0c, 0x4f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70,
	0x62, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x61,
	0x6e, 0x65, 0x6c, 0xca, 0x02, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0xe2,
	0x02, 0x15, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x61,
	0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_openapi_proto_goTypes = []interface{}{}
var file_openapi_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_openapi_proto_init() }
func file_openapi_proto_init() {
	if File_openapi_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_openapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_openapi_proto_goTypes,
		DependencyIndexes: file_openapi_proto_depIdxs,
	}.Build()
	File_openapi_proto = out.File
	file_openapi_proto_rawDesc = nil
	file_openapi_proto_goTypes = nil
	file_openapi_proto_depIdxs = nil
}
