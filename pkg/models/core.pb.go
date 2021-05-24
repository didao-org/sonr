// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: core.proto

package models

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

// Traffic Direction of Session
type Direction int32

const (
	Direction_Default  Direction = 0 // From Memory/Viewing
	Direction_Incoming Direction = 1 // Incoming Transfer
	Direction_Outgoing Direction = 2 // Outgoing Transfer
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "Default",
		1: "Incoming",
		2: "Outgoing",
	}
	Direction_value = map[string]int32{
		"Default":  0,
		"Incoming": 1,
		"Outgoing": 2,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_core_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_core_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0}
}

// Sent on Data Transfer to Add piece of File - Binary
type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buffer     []byte `protobuf:"bytes,1,opt,name=buffer,proto3" json:"buffer,omitempty"`          // Binary Value of Chunk
	Size       int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`             // Size of this Chunk
	IsComplete bool   `protobuf:"varint,3,opt,name=isComplete,proto3" json:"isComplete,omitempty"` // Whether this is final chunk
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

func (x *Chunk) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Chunk) GetIsComplete() bool {
	if x != nil {
		return x.IsComplete
	}
	return false
}

type Global struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers      map[string]string `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UserPeerID string            `protobuf:"bytes,2,opt,name=userPeerID,proto3" json:"userPeerID,omitempty"`
	Sname      string            `protobuf:"bytes,3,opt,name=sname,proto3" json:"sname,omitempty"`
}

func (x *Global) Reset() {
	*x = Global{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Global) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Global) ProtoMessage() {}

func (x *Global) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Global.ProtoReflect.Descriptor instead.
func (*Global) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{1}
}

func (x *Global) GetPeers() map[string]string {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *Global) GetUserPeerID() string {
	if x != nil {
		return x.UserPeerID
	}
	return ""
}

func (x *Global) GetSname() string {
	if x != nil {
		return x.Sname
	}
	return ""
}

var File_core_proto protoreflect.FileDescriptor

var file_core_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x53, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x06, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12,
	0x28, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x1a,
	0x38, 0x0a, 0x0a, 0x50, 0x65, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x34, 0x0a, 0x09, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x75, 0x74, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x42,
	0x09, 0x5a, 0x07, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_core_proto_rawDescOnce sync.Once
	file_core_proto_rawDescData = file_core_proto_rawDesc
)

func file_core_proto_rawDescGZIP() []byte {
	file_core_proto_rawDescOnce.Do(func() {
		file_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_proto_rawDescData)
	})
	return file_core_proto_rawDescData
}

var file_core_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_core_proto_goTypes = []interface{}{
	(Direction)(0), // 0: Direction
	(*Chunk)(nil),  // 1: Chunk
	(*Global)(nil), // 2: Global
	nil,            // 3: Global.PeersEntry
}
var file_core_proto_depIdxs = []int32{
	3, // 0: Global.peers:type_name -> Global.PeersEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_core_proto_init() }
func file_core_proto_init() {
	if File_core_proto != nil {
		return
	}
	file_api_proto_init()
	file_data_proto_init()
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Global); i {
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
			RawDescriptor: file_core_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_proto_goTypes,
		DependencyIndexes: file_core_proto_depIdxs,
		EnumInfos:         file_core_proto_enumTypes,
		MessageInfos:      file_core_proto_msgTypes,
	}.Build()
	File_core_proto = out.File
	file_core_proto_rawDesc = nil
	file_core_proto_goTypes = nil
	file_core_proto_depIdxs = nil
}
