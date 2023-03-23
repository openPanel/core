// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: initialize.proto

package pb

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

type InitializeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip       string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port     int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	ServerID string `protobuf:"bytes,3,opt,name=serverID,proto3" json:"serverID,omitempty"`
	Token    string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *InitializeRequest) Reset() {
	*x = InitializeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_initialize_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeRequest) ProtoMessage() {}

func (x *InitializeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_initialize_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeRequest.ProtoReflect.Descriptor instead.
func (*InitializeRequest) Descriptor() ([]byte, []int) {
	return file_initialize_proto_rawDescGZIP(), []int{0}
}

func (x *InitializeRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *InitializeRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *InitializeRequest) GetServerID() string {
	if x != nil {
		return x.ServerID
	}
	return ""
}

func (x *InitializeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type InitializeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientCert       string  `protobuf:"bytes,1,opt,name=clientCert,proto3" json:"clientCert,omitempty"`
	ClientPrivateKey string  `protobuf:"bytes,2,opt,name=clientPrivateKey,proto3" json:"clientPrivateKey,omitempty"`
	CACert           string  `protobuf:"bytes,3,opt,name=CACert,proto3" json:"CACert,omitempty"`
	Nodes            []*Node `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Error            *Error  `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *InitializeResponse) Reset() {
	*x = InitializeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_initialize_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeResponse) ProtoMessage() {}

func (x *InitializeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_initialize_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeResponse.ProtoReflect.Descriptor instead.
func (*InitializeResponse) Descriptor() ([]byte, []int) {
	return file_initialize_proto_rawDescGZIP(), []int{1}
}

func (x *InitializeResponse) GetClientCert() string {
	if x != nil {
		return x.ClientCert
	}
	return ""
}

func (x *InitializeResponse) GetClientPrivateKey() string {
	if x != nil {
		return x.ClientPrivateKey
	}
	return ""
}

func (x *InitializeResponse) GetCACert() string {
	if x != nil {
		return x.CACert
	}
	return ""
}

func (x *InitializeResponse) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *InitializeResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_initialize_proto protoreflect.FileDescriptor

var file_initialize_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x11, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc7, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x12, 0x2a, 0x0a,
	0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x41, 0x43,
	0x65, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x41, 0x43, 0x65, 0x72,
	0x74, 0x12, 0x25, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61,
	0x6e, 0x65, 0x6c, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0x5e, 0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x12, 0x1c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x2e,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x8b, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e,
	0x65, 0x6c, 0x42, 0x0f, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x4f,
	0x58, 0x58, 0xaa, 0x02, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0xca, 0x02,
	0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0xe2, 0x02, 0x15, 0x4f, 0x70, 0x65,
	0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_initialize_proto_rawDescOnce sync.Once
	file_initialize_proto_rawDescData = file_initialize_proto_rawDesc
)

func file_initialize_proto_rawDescGZIP() []byte {
	file_initialize_proto_rawDescOnce.Do(func() {
		file_initialize_proto_rawDescData = protoimpl.X.CompressGZIP(file_initialize_proto_rawDescData)
	})
	return file_initialize_proto_rawDescData
}

var file_initialize_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_initialize_proto_goTypes = []interface{}{
	(*InitializeRequest)(nil),  // 0: openPanel.InitializeRequest
	(*InitializeResponse)(nil), // 1: openPanel.InitializeResponse
	(*Node)(nil),               // 2: openPanel.Node
	(*Error)(nil),              // 3: openPanel.Error
}
var file_initialize_proto_depIdxs = []int32{
	2, // 0: openPanel.InitializeResponse.nodes:type_name -> openPanel.Node
	3, // 1: openPanel.InitializeResponse.error:type_name -> openPanel.Error
	0, // 2: openPanel.InitializeService.Initialize:input_type -> openPanel.InitializeRequest
	1, // 3: openPanel.InitializeService.Initialize:output_type -> openPanel.InitializeResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_initialize_proto_init() }
func file_initialize_proto_init() {
	if File_initialize_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_initialize_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeRequest); i {
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
		file_initialize_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeResponse); i {
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
			RawDescriptor: file_initialize_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_initialize_proto_goTypes,
		DependencyIndexes: file_initialize_proto_depIdxs,
		MessageInfos:      file_initialize_proto_msgTypes,
	}.Build()
	File_initialize_proto = out.File
	file_initialize_proto_rawDesc = nil
	file_initialize_proto_goTypes = nil
	file_initialize_proto_depIdxs = nil
}
