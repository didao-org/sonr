// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/v1/schema.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LinkKind int32

const (
	LinkKind_UNKNOWN LinkKind = 0
	LinkKind_OBJECT  LinkKind = 1
	LinkKind_SCHEMA  LinkKind = 2
	LinkKind_BUCKET  LinkKind = 3
)

var LinkKind_name = map[int32]string{
	0: "UNKNOWN",
	1: "OBJECT",
	2: "SCHEMA",
	3: "BUCKET",
}

var LinkKind_value = map[string]int32{
	"UNKNOWN": 0,
	"OBJECT":  1,
	"SCHEMA":  2,
	"BUCKET":  3,
}

func (x LinkKind) String() string {
	return proto.EnumName(LinkKind_name, int32(x))
}

func (LinkKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a184c368e8c5a046, []int{0}
}

type SchemaKind int32

const (
	SchemaKind_LIST   SchemaKind = 0
	SchemaKind_BOOL   SchemaKind = 1
	SchemaKind_INT    SchemaKind = 2
	SchemaKind_FLOAT  SchemaKind = 3
	SchemaKind_STRING SchemaKind = 4
	SchemaKind_BYTES  SchemaKind = 5
	SchemaKind_LINK   SchemaKind = 6
	SchemaKind_ANY    SchemaKind = 7
)

var SchemaKind_name = map[int32]string{
	0: "LIST",
	1: "BOOL",
	2: "INT",
	3: "FLOAT",
	4: "STRING",
	5: "BYTES",
	6: "LINK",
	7: "ANY",
}

var SchemaKind_value = map[string]int32{
	"LIST":   0,
	"BOOL":   1,
	"INT":    2,
	"FLOAT":  3,
	"STRING": 4,
	"BYTES":  5,
	"LINK":   6,
	"ANY":    7,
}

func (x SchemaKind) String() string {
	return proto.EnumName(SchemaKind_name, int32(x))
}

func (SchemaKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a184c368e8c5a046, []int{1}
}

type MetadataDefintion struct {
	// key of the metadata
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// metadata
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *MetadataDefintion) Reset()         { *m = MetadataDefintion{} }
func (m *MetadataDefintion) String() string { return proto.CompactTextString(m) }
func (*MetadataDefintion) ProtoMessage()    {}
func (*MetadataDefintion) Descriptor() ([]byte, []int) {
	return fileDescriptor_a184c368e8c5a046, []int{0}
}
func (m *MetadataDefintion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetadataDefintion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetadataDefintion.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetadataDefintion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataDefintion.Merge(m, src)
}
func (m *MetadataDefintion) XXX_Size() int {
	return m.Size()
}
func (m *MetadataDefintion) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataDefintion.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataDefintion proto.InternalMessageInfo

func (m *MetadataDefintion) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MetadataDefintion) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SchemaKindDefinition struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type of a single schema property
	Field SchemaKind `protobuf:"varint,2,opt,name=field,proto3,enum=sonrio.sonr.schema.SchemaKind" json:"field,omitempty"`
	// Optional field for a link context if `SchemaKind` is of type `Link`
	LinkKind LinkKind `protobuf:"varint,3,opt,name=link_kind,json=linkKind,proto3,enum=sonrio.sonr.schema.LinkKind" json:"link_kind,omitempty"`
	Link     string   `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
}

func (m *SchemaKindDefinition) Reset()         { *m = SchemaKindDefinition{} }
func (m *SchemaKindDefinition) String() string { return proto.CompactTextString(m) }
func (*SchemaKindDefinition) ProtoMessage()    {}
func (*SchemaKindDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_a184c368e8c5a046, []int{1}
}
func (m *SchemaKindDefinition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SchemaKindDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SchemaKindDefinition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SchemaKindDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaKindDefinition.Merge(m, src)
}
func (m *SchemaKindDefinition) XXX_Size() int {
	return m.Size()
}
func (m *SchemaKindDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaKindDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaKindDefinition proto.InternalMessageInfo

func (m *SchemaKindDefinition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SchemaKindDefinition) GetField() SchemaKind {
	if m != nil {
		return m.Field
	}
	return SchemaKind_LIST
}

func (m *SchemaKindDefinition) GetLinkKind() LinkKind {
	if m != nil {
		return m.LinkKind
	}
	return LinkKind_UNKNOWN
}

func (m *SchemaKindDefinition) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

type SchemaDefinition struct {
	// the DID for this schema should not be populated by request
	Did     string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// an alternative reference point
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	// the properties of this schema
	Fields []*SchemaKindDefinition `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (m *SchemaDefinition) Reset()         { *m = SchemaDefinition{} }
func (m *SchemaDefinition) String() string { return proto.CompactTextString(m) }
func (*SchemaDefinition) ProtoMessage()    {}
func (*SchemaDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_a184c368e8c5a046, []int{2}
}
func (m *SchemaDefinition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SchemaDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SchemaDefinition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SchemaDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaDefinition.Merge(m, src)
}
func (m *SchemaDefinition) XXX_Size() int {
	return m.Size()
}
func (m *SchemaDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaDefinition proto.InternalMessageInfo

func (m *SchemaDefinition) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *SchemaDefinition) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SchemaDefinition) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *SchemaDefinition) GetFields() []*SchemaKindDefinition {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterEnum("sonrio.sonr.schema.LinkKind", LinkKind_name, LinkKind_value)
	proto.RegisterEnum("sonrio.sonr.schema.SchemaKind", SchemaKind_name, SchemaKind_value)
	proto.RegisterType((*MetadataDefintion)(nil), "sonrio.sonr.schema.MetadataDefintion")
	proto.RegisterType((*SchemaKindDefinition)(nil), "sonrio.sonr.schema.SchemaKindDefinition")
	proto.RegisterType((*SchemaDefinition)(nil), "sonrio.sonr.schema.SchemaDefinition")
}

func init() { proto.RegisterFile("schema/v1/schema.proto", fileDescriptor_a184c368e8c5a046) }

var fileDescriptor_a184c368e8c5a046 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xe3, 0x26, 0xfd, 0xf7, 0x4e, 0x42, 0xc6, 0xaa, 0x50, 0x0e, 0x28, 0x9a, 0x7a, 0x40,
	0xd5, 0x24, 0x12, 0x31, 0xb8, 0xa0, 0x5d, 0x68, 0x4a, 0x81, 0xd2, 0x2e, 0x41, 0x69, 0x26, 0x18,
	0x07, 0x50, 0x5a, 0x7b, 0xcc, 0x6a, 0x6a, 0x4f, 0x6d, 0x36, 0xb1, 0x6f, 0xc1, 0x89, 0xaf, 0xc1,
	0xd7, 0xe0, 0xd8, 0x23, 0x47, 0xd4, 0x7e, 0x11, 0x64, 0x3b, 0x51, 0x91, 0x40, 0x3b, 0xf9, 0xf7,
	0xda, 0x79, 0x1e, 0x3f, 0x4f, 0x64, 0x78, 0xb0, 0x9e, 0x5f, 0xb2, 0x65, 0x16, 0xdc, 0x3c, 0x09,
	0x0c, 0xf9, 0x57, 0x2b, 0x59, 0x48, 0x42, 0xd6, 0x52, 0xac, 0xb8, 0xf4, 0xd5, 0xe2, 0x9b, 0x93,
	0xee, 0x09, 0xdc, 0x3f, 0x65, 0x45, 0x46, 0xb3, 0x22, 0x7b, 0xc9, 0x2e, 0xb8, 0x28, 0xb8, 0x14,
	0x04, 0x83, 0xbd, 0x60, 0xb7, 0x2e, 0x3a, 0x44, 0xbd, 0x76, 0xa2, 0x90, 0x74, 0xa0, 0x7e, 0x93,
	0xe5, 0xd7, 0xcc, 0xad, 0xe9, 0x3d, 0x33, 0x74, 0x7f, 0x20, 0xe8, 0x4c, 0xb5, 0xcf, 0x98, 0x0b,
	0xaa, 0xf5, 0x5c, 0x1b, 0x10, 0x70, 0x44, 0xb6, 0x64, 0xa5, 0x83, 0x66, 0xf2, 0x0c, 0xea, 0x17,
	0x9c, 0xe5, 0x54, 0x5b, 0xdc, 0x3b, 0xf6, 0xfc, 0x7f, 0xd3, 0xf8, 0x7b, 0xb3, 0xc4, 0x7c, 0x4c,
	0x9e, 0x43, 0x3b, 0xe7, 0x62, 0xf1, 0x79, 0xc1, 0x05, 0x75, 0x6d, 0xad, 0x7c, 0xf8, 0x3f, 0xe5,
	0x84, 0x8b, 0x85, 0xd6, 0xb5, 0xf2, 0x92, 0x54, 0x08, 0xc5, 0xae, 0x63, 0x42, 0x28, 0xee, 0x7e,
	0x47, 0x80, 0xcd, 0x25, 0x7f, 0xa5, 0xc5, 0x60, 0x53, 0x4e, 0xab, 0xba, 0x94, 0x53, 0xe2, 0x42,
	0x73, 0xbe, 0x62, 0x59, 0x21, 0x57, 0x65, 0xe1, 0x6a, 0x54, 0x3f, 0x22, 0xcf, 0x66, 0x2c, 0xd7,
	0x59, 0xda, 0x89, 0x19, 0xc8, 0x0b, 0x68, 0xe8, 0xb8, 0x6b, 0xd7, 0x39, 0xb4, 0x7b, 0x07, 0xc7,
	0xbd, 0xbb, 0xcb, 0xed, 0xef, 0x4e, 0x4a, 0xdd, 0xd1, 0x09, 0xb4, 0xaa, 0x0a, 0xe4, 0x00, 0x9a,
	0x67, 0xd1, 0x38, 0x8a, 0xdf, 0x47, 0xd8, 0x22, 0x00, 0x8d, 0x38, 0x7c, 0x3b, 0x1c, 0xa4, 0x18,
	0x29, 0x9e, 0x0e, 0xde, 0x0c, 0x4f, 0xfb, 0xb8, 0xa6, 0x38, 0x3c, 0x1b, 0x8c, 0x87, 0x29, 0xb6,
	0x8f, 0x3e, 0x01, 0xec, 0xcd, 0x49, 0x0b, 0x9c, 0xc9, 0x68, 0x9a, 0x62, 0x4b, 0x51, 0x18, 0xc7,
	0x13, 0x8c, 0x48, 0x13, 0xec, 0x51, 0x94, 0xe2, 0x1a, 0x69, 0x43, 0xfd, 0xd5, 0x24, 0xee, 0xa7,
	0xd8, 0xd6, 0x6e, 0x69, 0x32, 0x8a, 0x5e, 0x63, 0x47, 0x6d, 0x87, 0xe7, 0xe9, 0x70, 0x8a, 0xeb,
	0x46, 0x1e, 0x8d, 0x71, 0x43, 0x89, 0xfa, 0xd1, 0x39, 0x6e, 0x86, 0x1f, 0x7e, 0x6e, 0x3d, 0xb4,
	0xd9, 0x7a, 0xe8, 0xf7, 0xd6, 0x43, 0xdf, 0x76, 0x9e, 0xb5, 0xd9, 0x79, 0xd6, 0xaf, 0x9d, 0x67,
	0x41, 0xa7, 0xea, 0x58, 0xdc, 0x5e, 0xb1, 0x75, 0xd9, 0xf4, 0x1d, 0xfa, 0xf8, 0xe8, 0x0b, 0x2f,
	0x2e, 0xaf, 0x67, 0xfe, 0x5c, 0x2e, 0x03, 0x75, 0xfe, 0x98, 0x4b, 0xbd, 0x06, 0x5f, 0xcb, 0x27,
	0x19, 0x68, 0xc1, 0xac, 0xa1, 0x5f, 0xe6, 0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x73, 0x14,
	0x8b, 0x00, 0xb3, 0x02, 0x00, 0x00,
}

func (m *MetadataDefintion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetadataDefintion) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetadataDefintion) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SchemaKindDefinition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SchemaKindDefinition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SchemaKindDefinition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Link) > 0 {
		i -= len(m.Link)
		copy(dAtA[i:], m.Link)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Link)))
		i--
		dAtA[i] = 0x22
	}
	if m.LinkKind != 0 {
		i = encodeVarintSchema(dAtA, i, uint64(m.LinkKind))
		i--
		dAtA[i] = 0x18
	}
	if m.Field != 0 {
		i = encodeVarintSchema(dAtA, i, uint64(m.Field))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SchemaDefinition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SchemaDefinition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SchemaDefinition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fields) > 0 {
		for iNdEx := len(m.Fields) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fields[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSchema(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSchema(dAtA []byte, offset int, v uint64) int {
	offset -= sovSchema(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MetadataDefintion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	return n
}

func (m *SchemaKindDefinition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	if m.Field != 0 {
		n += 1 + sovSchema(uint64(m.Field))
	}
	if m.LinkKind != 0 {
		n += 1 + sovSchema(uint64(m.LinkKind))
	}
	l = len(m.Link)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	return n
}

func (m *SchemaDefinition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	if len(m.Fields) > 0 {
		for _, e := range m.Fields {
			l = e.Size()
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	return n
}

func sovSchema(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSchema(x uint64) (n int) {
	return sovSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetadataDefintion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MetadataDefintion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetadataDefintion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SchemaKindDefinition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SchemaKindDefinition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SchemaKindDefinition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			m.Field = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Field |= SchemaKind(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LinkKind", wireType)
			}
			m.LinkKind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LinkKind |= LinkKind(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Link", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Link = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SchemaDefinition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SchemaDefinition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SchemaDefinition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fields", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fields = append(m.Fields, &SchemaKindDefinition{})
			if err := m.Fields[len(m.Fields)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSchema
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSchema
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSchema
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSchema        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchema          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSchema = fmt.Errorf("proto: unexpected end of group")
)
