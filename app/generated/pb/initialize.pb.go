// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: initialize.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip       string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port     int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	ServerID string `protobuf:"bytes,3,opt,name=serverID,proto3" json:"serverID,omitempty"`
	Token    string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Csr      []byte `protobuf:"bytes,5,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_initialize_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_initialize_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *RegisterRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RegisterRequest) GetServerID() string {
	if x != nil {
		return x.ServerID
	}
	return ""
}

func (x *RegisterRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RegisterRequest) GetCsr() []byte {
	if x != nil {
		return x.Csr
	}
	return nil
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientCert    []byte       `protobuf:"bytes,1,opt,name=clientCert,proto3" json:"clientCert,omitempty"`
	ClusterCACert []byte       `protobuf:"bytes,2,opt,name=clusterCACert,proto3" json:"clusterCACert,omitempty"`
	Nodes         []*Node      `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
	LinkStates    []*LinkState `protobuf:"bytes,4,rep,name=linkStates,proto3" json:"linkStates,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_initialize_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_initialize_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetClientCert() []byte {
	if x != nil {
		return x.ClientCert
	}
	return nil
}

func (x *RegisterResponse) GetClusterCACert() []byte {
	if x != nil {
		return x.ClusterCACert
	}
	return nil
}

func (x *RegisterResponse) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *RegisterResponse) GetLinkStates() []*LinkState {
	if x != nil {
		return x.LinkStates
	}
	return nil
}

type GetNodesInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes      []*Node      `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	LinkStates []*LinkState `protobuf:"bytes,2,rep,name=linkStates,proto3" json:"linkStates,omitempty"`
}

func (x *GetNodesInfoResponse) Reset() {
	*x = GetNodesInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_initialize_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodesInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesInfoResponse) ProtoMessage() {}

func (x *GetNodesInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_initialize_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesInfoResponse.ProtoReflect.Descriptor instead.
func (*GetNodesInfoResponse) Descriptor() ([]byte, []int) {
	return file_initialize_proto_rawDescGZIP(), []int{2}
}

func (x *GetNodesInfoResponse) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *GetNodesInfoResponse) GetLinkStates() []*LinkState {
	if x != nil {
		return x.LinkStates
	}
	return nil
}

var File_initialize_proto protoreflect.FileDescriptor

var file_initialize_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x1a, 0x0b, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x73, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x63, 0x73, 0x72, 0x22, 0xb1,
	0x01, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x65, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x41,
	0x43, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x41, 0x43, 0x65, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x0a,
	0x6c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x6e,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x73, 0x22, 0x6f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x34, 0x0a,
	0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x4c, 0x69,
	0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x32, 0x82, 0x02, 0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa3, 0x01, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e,
	0x65, 0x6c, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x5e, 0x92, 0x41, 0x45, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x12, 0x0e, 0x4a, 0x6f, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x1a, 0x27, 0x41, 0x20, 0x6e, 0x6f, 0x64, 0x65, 0x20, 0x74, 0x72, 0x79, 0x20, 0x74, 0x6f, 0x20,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a,
	0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x12,
	0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61,
	0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x8b, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x42, 0x0f, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x61,
	0x6e, 0x65, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x4f, 0x70, 0x65,
	0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0xca, 0x02, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e,
	0x65, 0x6c, 0xe2, 0x02, 0x15, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x4f, 0x70, 0x65,
	0x6e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_initialize_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_initialize_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),      // 0: openPanel.RegisterRequest
	(*RegisterResponse)(nil),     // 1: openPanel.RegisterResponse
	(*GetNodesInfoResponse)(nil), // 2: openPanel.GetNodesInfoResponse
	(*Node)(nil),                 // 3: entpb.Node
	(*LinkState)(nil),            // 4: openPanel.LinkState
	(*emptypb.Empty)(nil),        // 5: google.protobuf.Empty
}
var file_initialize_proto_depIdxs = []int32{
	3, // 0: openPanel.RegisterResponse.nodes:type_name -> entpb.Node
	4, // 1: openPanel.RegisterResponse.linkStates:type_name -> openPanel.LinkState
	3, // 2: openPanel.GetNodesInfoResponse.nodes:type_name -> entpb.Node
	4, // 3: openPanel.GetNodesInfoResponse.linkStates:type_name -> openPanel.LinkState
	0, // 4: openPanel.InitializeService.Register:input_type -> openPanel.RegisterRequest
	5, // 5: openPanel.InitializeService.GetNodesInfo:input_type -> google.protobuf.Empty
	1, // 6: openPanel.InitializeService.Register:output_type -> openPanel.RegisterResponse
	2, // 7: openPanel.InitializeService.GetNodesInfo:output_type -> openPanel.GetNodesInfoResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_initialize_proto_init() }
func file_initialize_proto_init() {
	if File_initialize_proto != nil {
		return
	}
	file_entpb_proto_init()
	file_router_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_initialize_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
			switch v := v.(*RegisterResponse); i {
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
		file_initialize_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodesInfoResponse); i {
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
			NumMessages:   3,
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
