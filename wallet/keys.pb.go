// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/wallet/keys.proto

package wallet

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

// SNID is the unique identifier of a user with Sonr Protocol
type SNID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"` // Domain Name of the node, e.g. prad.snr
	PubKey []byte `protobuf:"bytes,2,opt,name=pubKey,proto3" json:"pubKey,omitempty"` // Public Key of the node, value in the Domains second TXT record
	PeerID string `protobuf:"bytes,3,opt,name=peerID,proto3" json:"peerID,omitempty"` // Peer ID of the node, calculated from the public key
	Did    string `protobuf:"bytes,4,opt,name=did,proto3" json:"did,omitempty"`       // DID of the node, calculated from the public key
}

func (x *SNID) Reset() {
	*x = SNID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wallet_keys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SNID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SNID) ProtoMessage() {}

func (x *SNID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wallet_keys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SNID.ProtoReflect.Descriptor instead.
func (*SNID) Descriptor() ([]byte, []int) {
	return file_proto_wallet_keys_proto_rawDescGZIP(), []int{0}
}

func (x *SNID) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *SNID) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *SNID) GetPeerID() string {
	if x != nil {
		return x.PeerID
	}
	return ""
}

func (x *SNID) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

// UUID is General Message ID with Signature, ID, and Timestamp
type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`  // Signature of the message
	Value     string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`          // ID of the message
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // Unix timestamp
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wallet_keys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wallet_keys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_proto_wallet_keys_proto_rawDescGZIP(), []int{1}
}

func (x *UUID) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *UUID) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_proto_wallet_keys_proto protoreflect.FileDescriptor

var file_proto_wallet_keys_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x6b,
	0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22, 0x60, 0x0a, 0x04, 0x53, 0x4e, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_wallet_keys_proto_rawDescOnce sync.Once
	file_proto_wallet_keys_proto_rawDescData = file_proto_wallet_keys_proto_rawDesc
)

func file_proto_wallet_keys_proto_rawDescGZIP() []byte {
	file_proto_wallet_keys_proto_rawDescOnce.Do(func() {
		file_proto_wallet_keys_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_wallet_keys_proto_rawDescData)
	})
	return file_proto_wallet_keys_proto_rawDescData
}

var file_proto_wallet_keys_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_wallet_keys_proto_goTypes = []interface{}{
	(*SNID)(nil), // 0: sonr.wallet.SNID
	(*UUID)(nil), // 1: sonr.wallet.UUID
}
var file_proto_wallet_keys_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_wallet_keys_proto_init() }
func file_proto_wallet_keys_proto_init() {
	if File_proto_wallet_keys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_wallet_keys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SNID); i {
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
		file_proto_wallet_keys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
			RawDescriptor: file_proto_wallet_keys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_wallet_keys_proto_goTypes,
		DependencyIndexes: file_proto_wallet_keys_proto_depIdxs,
		MessageInfos:      file_proto_wallet_keys_proto_msgTypes,
	}.Build()
	File_proto_wallet_keys_proto = out.File
	file_proto_wallet_keys_proto_rawDesc = nil
	file_proto_wallet_keys_proto_goTypes = nil
	file_proto_wallet_keys_proto_depIdxs = nil
}
