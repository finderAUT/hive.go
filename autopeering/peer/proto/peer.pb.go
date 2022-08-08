// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: autopeering/peer/proto/peer.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/iotaledger/hive.go/v2/autopeering/peer/service/proto"
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

// Minimal encoding of a peer
type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public key used for signing
	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// string form of the peers ID
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	// services supported by the peer
	Services *proto1.ServiceMap `protobuf:"bytes,3,opt,name=services,proto3" json:"services,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autopeering_peer_proto_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_autopeering_peer_proto_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_autopeering_peer_proto_peer_proto_rawDescGZIP(), []int{0}
}

func (x *Peer) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Peer) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Peer) GetServices() *proto1.ServiceMap {
	if x != nil {
		return x.Services
	}
	return nil
}

var File_autopeering_peer_proto_peer_proto protoreflect.FileDescriptor

var file_autopeering_peer_proto_peer_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x75, 0x74, 0x6f, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x65,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x61, 0x75, 0x74, 0x6f,
	0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12,
	0x2d, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4d, 0x61, 0x70, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x36,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x74,
	0x61, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x67, 0x6f, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x65, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_autopeering_peer_proto_peer_proto_rawDescOnce sync.Once
	file_autopeering_peer_proto_peer_proto_rawDescData = file_autopeering_peer_proto_peer_proto_rawDesc
)

func file_autopeering_peer_proto_peer_proto_rawDescGZIP() []byte {
	file_autopeering_peer_proto_peer_proto_rawDescOnce.Do(func() {
		file_autopeering_peer_proto_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_autopeering_peer_proto_peer_proto_rawDescData)
	})
	return file_autopeering_peer_proto_peer_proto_rawDescData
}

var file_autopeering_peer_proto_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_autopeering_peer_proto_peer_proto_goTypes = []interface{}{
	(*Peer)(nil),              // 0: proto.Peer
	(*proto1.ServiceMap)(nil), // 1: proto.ServiceMap
}
var file_autopeering_peer_proto_peer_proto_depIdxs = []int32{
	1, // 0: proto.Peer.services:type_name -> proto.ServiceMap
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_autopeering_peer_proto_peer_proto_init() }
func file_autopeering_peer_proto_peer_proto_init() {
	if File_autopeering_peer_proto_peer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_autopeering_peer_proto_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer); i {
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
			RawDescriptor: file_autopeering_peer_proto_peer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_autopeering_peer_proto_peer_proto_goTypes,
		DependencyIndexes: file_autopeering_peer_proto_peer_proto_depIdxs,
		MessageInfos:      file_autopeering_peer_proto_peer_proto_msgTypes,
	}.Build()
	File_autopeering_peer_proto_peer_proto = out.File
	file_autopeering_peer_proto_peer_proto_rawDesc = nil
	file_autopeering_peer_proto_peer_proto_goTypes = nil
	file_autopeering_peer_proto_peer_proto_depIdxs = nil
}
