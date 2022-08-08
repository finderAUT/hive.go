// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: autopeering/discover/proto/message.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	proto2 "github.com/finderAUT/hive.go/v3/autopeering/peer/proto"
	proto1 "github.com/finderAUT/hive.go/v3/autopeering/peer/service/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// version number and network ID to classify the protocol
	Version   uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	NetworkId uint32 `protobuf:"varint,2,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// unix time
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// endpoint of the sender; port and string form of the return IP address (e.g. "192.0.2.1", "[2001:db8::1]")
	SrcAddr string `protobuf:"bytes,4,opt,name=src_addr,json=srcAddr,proto3" json:"src_addr,omitempty"`
	SrcPort uint32 `protobuf:"varint,5,opt,name=src_port,json=srcPort,proto3" json:"src_port,omitempty"`
	// string form of receiver's IP
	// This provides a way to discover the the external address (after NAT).
	DstAddr string `protobuf:"bytes,6,opt,name=dst_addr,json=dstAddr,proto3" json:"dst_addr,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autopeering_discover_proto_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_autopeering_discover_proto_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_autopeering_discover_proto_message_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Ping) GetNetworkId() uint32 {
	if x != nil {
		return x.NetworkId
	}
	return 0
}

func (x *Ping) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Ping) GetSrcAddr() string {
	if x != nil {
		return x.SrcAddr
	}
	return ""
}

func (x *Ping) GetSrcPort() uint32 {
	if x != nil {
		return x.SrcPort
	}
	return 0
}

func (x *Ping) GetDstAddr() string {
	if x != nil {
		return x.DstAddr
	}
	return ""
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hash of the ping packet
	ReqHash []byte `protobuf:"bytes,1,opt,name=req_hash,json=reqHash,proto3" json:"req_hash,omitempty"`
	// services supported by the sender
	Services *proto1.ServiceMap `protobuf:"bytes,2,opt,name=services,proto3" json:"services,omitempty"`
	// string form of receiver's IP
	// This should mirror the source IP of the Ping's IP packet. It provides a way to discover the the external address (after NAT).
	DstAddr string `protobuf:"bytes,3,opt,name=dst_addr,json=dstAddr,proto3" json:"dst_addr,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autopeering_discover_proto_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_autopeering_discover_proto_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_autopeering_discover_proto_message_proto_rawDescGZIP(), []int{1}
}

func (x *Pong) GetReqHash() []byte {
	if x != nil {
		return x.ReqHash
	}
	return nil
}

func (x *Pong) GetServices() *proto1.ServiceMap {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Pong) GetDstAddr() string {
	if x != nil {
		return x.DstAddr
	}
	return ""
}

type DiscoveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unix time
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *DiscoveryRequest) Reset() {
	*x = DiscoveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autopeering_discover_proto_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryRequest) ProtoMessage() {}

func (x *DiscoveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autopeering_discover_proto_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryRequest.ProtoReflect.Descriptor instead.
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return file_autopeering_discover_proto_message_proto_rawDescGZIP(), []int{2}
}

func (x *DiscoveryRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type DiscoveryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hash of the corresponding request
	ReqHash []byte `protobuf:"bytes,1,opt,name=req_hash,json=reqHash,proto3" json:"req_hash,omitempty"`
	// list of peers
	Peers []*proto2.Peer `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *DiscoveryResponse) Reset() {
	*x = DiscoveryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autopeering_discover_proto_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryResponse) ProtoMessage() {}

func (x *DiscoveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autopeering_discover_proto_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryResponse.ProtoReflect.Descriptor instead.
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return file_autopeering_discover_proto_message_proto_rawDescGZIP(), []int{3}
}

func (x *DiscoveryResponse) GetReqHash() []byte {
	if x != nil {
		return x.ReqHash
	}
	return nil
}

func (x *DiscoveryResponse) GetPeers() []*proto2.Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

var File_autopeering_discover_proto_message_proto protoreflect.FileDescriptor

var file_autopeering_discover_proto_message_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x75, 0x74, 0x6f, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x61, 0x75, 0x74, 0x6f, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70,
	0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x61, 0x75, 0x74, 0x6f, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x73, 0x72, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x73, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x22, 0x6b, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x72,
	0x65, 0x71, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2d, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x22, 0x30, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x51, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x71, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x65, 0x65, 0x72, 0x73, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x74, 0x61, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x70, 0x65, 0x65, 0x72, 0x69,
	0x6e, 0x67, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_autopeering_discover_proto_message_proto_rawDescOnce sync.Once
	file_autopeering_discover_proto_message_proto_rawDescData = file_autopeering_discover_proto_message_proto_rawDesc
)

func file_autopeering_discover_proto_message_proto_rawDescGZIP() []byte {
	file_autopeering_discover_proto_message_proto_rawDescOnce.Do(func() {
		file_autopeering_discover_proto_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_autopeering_discover_proto_message_proto_rawDescData)
	})
	return file_autopeering_discover_proto_message_proto_rawDescData
}

var file_autopeering_discover_proto_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_autopeering_discover_proto_message_proto_goTypes = []interface{}{
	(*Ping)(nil),              // 0: proto.Ping
	(*Pong)(nil),              // 1: proto.Pong
	(*DiscoveryRequest)(nil),  // 2: proto.DiscoveryRequest
	(*DiscoveryResponse)(nil), // 3: proto.DiscoveryResponse
	(*proto1.ServiceMap)(nil), // 4: proto.ServiceMap
	(*proto2.Peer)(nil),       // 5: proto.Peer
}
var file_autopeering_discover_proto_message_proto_depIdxs = []int32{
	4, // 0: proto.Pong.services:type_name -> proto.ServiceMap
	5, // 1: proto.DiscoveryResponse.peers:type_name -> proto.Peer
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_autopeering_discover_proto_message_proto_init() }
func file_autopeering_discover_proto_message_proto_init() {
	if File_autopeering_discover_proto_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_autopeering_discover_proto_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_autopeering_discover_proto_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
		file_autopeering_discover_proto_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoveryRequest); i {
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
		file_autopeering_discover_proto_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoveryResponse); i {
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
			RawDescriptor: file_autopeering_discover_proto_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_autopeering_discover_proto_message_proto_goTypes,
		DependencyIndexes: file_autopeering_discover_proto_message_proto_depIdxs,
		MessageInfos:      file_autopeering_discover_proto_message_proto_msgTypes,
	}.Build()
	File_autopeering_discover_proto_message_proto = out.File
	file_autopeering_discover_proto_message_proto_rawDesc = nil
	file_autopeering_discover_proto_message_proto_goTypes = nil
	file_autopeering_discover_proto_message_proto_depIdxs = nil
}
