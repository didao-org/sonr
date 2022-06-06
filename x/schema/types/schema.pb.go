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

// Represents the types of fields a schema can have
type SchemaKind int32

const (
	SchemaKind_SCHEMA_KIND_INVALID SchemaKind = 0
	SchemaKind_SCHEMA_KIND_MAP     SchemaKind = 1
	SchemaKind_SCHEMA_KIND_LIST    SchemaKind = 2
	SchemaKind_SCHEMA_KIND_UNIT    SchemaKind = 3
	SchemaKind_SCHEMA_KIND_BOOL    SchemaKind = 4
	SchemaKind_SCHEMA_KIND_INT     SchemaKind = 5
	SchemaKind_SCHEMA_KIND_FLOAT   SchemaKind = 6
	SchemaKind_SCHEMA_KIND_STRING  SchemaKind = 7
	SchemaKind_SCHEMA_KIND_BYTES   SchemaKind = 8
	SchemaKind_SCHEMA_KIND_LINK    SchemaKind = 9
	SchemaKind_SCHEMA_KIND_STRUCT  SchemaKind = 10
	SchemaKind_SCHEMA_KIND_UNION   SchemaKind = 11
	SchemaKind_SCHEMA_KIND_ENUM    SchemaKind = 12
	SchemaKind_SCHEMA_KIND_ANY     SchemaKind = 13
)

var SchemaKind_name = map[int32]string{
	0:  "SCHEMA_KIND_INVALID",
	1:  "SCHEMA_KIND_MAP",
	2:  "SCHEMA_KIND_LIST",
	3:  "SCHEMA_KIND_UNIT",
	4:  "SCHEMA_KIND_BOOL",
	5:  "SCHEMA_KIND_INT",
	6:  "SCHEMA_KIND_FLOAT",
	7:  "SCHEMA_KIND_STRING",
	8:  "SCHEMA_KIND_BYTES",
	9:  "SCHEMA_KIND_LINK",
	10: "SCHEMA_KIND_STRUCT",
	11: "SCHEMA_KIND_UNION",
	12: "SCHEMA_KIND_ENUM",
	13: "SCHEMA_KIND_ANY",
}

var SchemaKind_value = map[string]int32{
	"SCHEMA_KIND_INVALID": 0,
	"SCHEMA_KIND_MAP":     1,
	"SCHEMA_KIND_LIST":    2,
	"SCHEMA_KIND_UNIT":    3,
	"SCHEMA_KIND_BOOL":    4,
	"SCHEMA_KIND_INT":     5,
	"SCHEMA_KIND_FLOAT":   6,
	"SCHEMA_KIND_STRING":  7,
	"SCHEMA_KIND_BYTES":   8,
	"SCHEMA_KIND_LINK":    9,
	"SCHEMA_KIND_STRUCT":  10,
	"SCHEMA_KIND_UNION":   11,
	"SCHEMA_KIND_ENUM":    12,
	"SCHEMA_KIND_ANY":     13,
}

func (x SchemaKind) String() string {
	return proto.EnumName(SchemaKind_name, int32(x))
}

func (SchemaKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a184c368e8c5a046, []int{0}
}

// Schema defines the shapes of schemas on Sonr
type Schema struct {
	// the DID for this schema
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// an alternative reference point
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// the properties of this schema
	Fields map[string]SchemaKind `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=sonrio.sonr.schema.SchemaKind"`
}

func (m *Schema) Reset()         { *m = Schema{} }
func (m *Schema) String() string { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()    {}
func (*Schema) Descriptor() ([]byte, []int) {
	return fileDescriptor_a184c368e8c5a046, []int{0}
}
func (m *Schema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Schema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Schema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Schema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schema.Merge(m, src)
}
func (m *Schema) XXX_Size() int {
	return m.Size()
}
func (m *Schema) XXX_DiscardUnknown() {
	xxx_messageInfo_Schema.DiscardUnknown(m)
}

var xxx_messageInfo_Schema proto.InternalMessageInfo

func (m *Schema) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *Schema) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Schema) GetFields() map[string]SchemaKind {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterEnum("sonrio.sonr.schema.SchemaKind", SchemaKind_name, SchemaKind_value)
	proto.RegisterType((*Schema)(nil), "sonrio.sonr.schema.Schema")
	proto.RegisterMapType((map[string]SchemaKind)(nil), "sonrio.sonr.schema.Schema.FieldsEntry")
}

func init() { proto.RegisterFile("schema/v1/schema.proto", fileDescriptor_a184c368e8c5a046) }

var fileDescriptor_a184c368e8c5a046 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4d, 0xcf, 0xd2, 0x40,
	0x10, 0x80, 0xbb, 0xad, 0x54, 0xdf, 0xc5, 0x8f, 0x71, 0x5f, 0x44, 0xe2, 0xa1, 0x21, 0x1e, 0x08,
	0x31, 0xb1, 0x8d, 0xe8, 0xc1, 0x78, 0x30, 0x16, 0x28, 0xba, 0xa1, 0x6c, 0x4d, 0x3f, 0x4c, 0xf0,
	0x42, 0x80, 0x56, 0x69, 0x28, 0x94, 0xd0, 0x42, 0xec, 0xbf, 0xf0, 0x5f, 0x69, 0xe2, 0x85, 0xa3,
	0x47, 0x03, 0x7f, 0xc4, 0xb4, 0xc5, 0xd8, 0x14, 0x3d, 0xcd, 0xf4, 0x99, 0xe9, 0x33, 0x93, 0xec,
	0xe0, 0x7a, 0x34, 0x5f, 0x78, 0xab, 0xa9, 0xb2, 0x7f, 0xa6, 0xe4, 0x99, 0xbc, 0xd9, 0x86, 0x71,
	0x48, 0x48, 0x14, 0xae, 0xb7, 0x7e, 0x28, 0xa7, 0x41, 0xce, 0x2b, 0x8f, 0x7f, 0x20, 0x2c, 0x5a,
	0x59, 0x4a, 0x00, 0x0b, 0xae, 0xef, 0x36, 0x50, 0x13, 0xb5, 0xaf, 0xcc, 0x34, 0x25, 0x35, 0x5c,
	0x09, 0xa6, 0x33, 0x2f, 0x68, 0xf0, 0x19, 0xcb, 0x3f, 0xc8, 0x6b, 0x2c, 0x7e, 0xf2, 0xbd, 0xc0,
	0x8d, 0x1a, 0x42, 0x53, 0x68, 0x57, 0x3b, 0x2d, 0xf9, 0xd2, 0x2b, 0xe7, 0x4e, 0x79, 0x90, 0x35,
	0x6a, 0xeb, 0x78, 0x9b, 0x98, 0xe7, 0xbf, 0x1e, 0x8d, 0x71, 0xb5, 0x80, 0xd3, 0xb1, 0x4b, 0x2f,
	0xf9, 0x33, 0x76, 0xe9, 0x25, 0xe4, 0x05, 0xae, 0xec, 0xa7, 0xc1, 0xce, 0xcb, 0xc6, 0xde, 0xed,
	0x48, 0xff, 0xf7, 0x0f, 0xfd, 0xb5, 0x6b, 0xe6, 0xcd, 0xaf, 0xf8, 0x97, 0xe8, 0xc9, 0x37, 0x1e,
	0xe3, 0xbf, 0x15, 0xf2, 0x10, 0x5f, 0x5b, 0xbd, 0x77, 0xda, 0x48, 0x9d, 0x0c, 0x29, 0xeb, 0x4f,
	0x28, 0xfb, 0xa0, 0xea, 0xb4, 0x0f, 0x1c, 0xb9, 0xc6, 0xf7, 0x8a, 0x85, 0x91, 0xfa, 0x1e, 0x10,
	0xa9, 0x61, 0x28, 0x42, 0x9d, 0x5a, 0x36, 0xf0, 0x65, 0xea, 0x30, 0x6a, 0x83, 0x50, 0xa6, 0x5d,
	0xc3, 0xd0, 0xe1, 0x46, 0x59, 0x4b, 0x99, 0x0d, 0x15, 0xf2, 0x00, 0xdf, 0x2f, 0xc2, 0x81, 0x6e,
	0xa8, 0x36, 0x88, 0xa4, 0x8e, 0x49, 0x11, 0x5b, 0xb6, 0x49, 0xd9, 0x5b, 0xb8, 0x59, 0x6e, 0xef,
	0x8e, 0x6d, 0xcd, 0x82, 0x5b, 0x97, 0xcb, 0xb1, 0x21, 0x5c, 0xfd, 0x43, 0xe2, 0xf4, 0x6c, 0xc0,
	0x65, 0x89, 0xc3, 0xa8, 0xc1, 0xa0, 0x5a, 0x96, 0x68, 0xcc, 0x19, 0xc1, 0xed, 0xf2, 0xd6, 0x2a,
	0x1b, 0xc3, 0x9d, 0xee, 0x9b, 0xef, 0x47, 0x09, 0x1d, 0x8e, 0x12, 0xfa, 0x75, 0x94, 0xd0, 0xd7,
	0x93, 0xc4, 0x1d, 0x4e, 0x12, 0xf7, 0xf3, 0x24, 0x71, 0x1f, 0x5b, 0x9f, 0xfd, 0x78, 0xb1, 0x9b,
	0xc9, 0xf3, 0x70, 0xa5, 0xa4, 0x2f, 0xf2, 0xd4, 0x0f, 0xb3, 0xa8, 0x7c, 0x39, 0x5f, 0x9b, 0x12,
	0x27, 0x1b, 0x2f, 0x9a, 0x89, 0xd9, 0xd1, 0x3d, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xee, 0xe5,
	0x6e, 0x46, 0x8e, 0x02, 0x00, 0x00,
}

func (m *Schema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Schema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fields) > 0 {
		for k := range m.Fields {
			v := m.Fields[k]
			baseI := i
			i = encodeVarintSchema(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintSchema(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintSchema(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Label)))
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
func (m *Schema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	if len(m.Fields) > 0 {
		for k, v := range m.Fields {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovSchema(uint64(len(k))) + 1 + sovSchema(uint64(v))
			n += mapEntrySize + 1 + sovSchema(uint64(mapEntrySize))
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
func (m *Schema) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Schema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Schema: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 3:
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
			if m.Fields == nil {
				m.Fields = make(map[string]SchemaKind)
			}
			var mapkey string
			var mapvalue SchemaKind
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSchema
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthSchema
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthSchema
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSchema
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= SchemaKind(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipSchema(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthSchema
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Fields[mapkey] = mapvalue
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