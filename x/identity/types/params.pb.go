// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sonr/identity/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Params defines the parameters for the module.
type Params struct {
	DidBaseContext   string   `protobuf:"bytes,1,opt,name=did_base_context,json=didBaseContext,proto3" json:"did_base_context,omitempty"`
	DidMethodContext string   `protobuf:"bytes,2,opt,name=did_method_context,json=didMethodContext,proto3" json:"did_method_context,omitempty"`
	DidMethodName    string   `protobuf:"bytes,3,opt,name=did_method_name,json=didMethodName,proto3" json:"did_method_name,omitempty"`
	DidMethodVersion string   `protobuf:"bytes,4,opt,name=did_method_version,json=didMethodVersion,proto3" json:"did_method_version,omitempty"`
	DidNetwork       string   `protobuf:"bytes,5,opt,name=did_network,json=didNetwork,proto3" json:"did_network,omitempty"`
	IpfsGateway      string   `protobuf:"bytes,6,opt,name=ipfs_gateway,json=ipfsGateway,proto3" json:"ipfs_gateway,omitempty"`
	IpfsApi          string   `protobuf:"bytes,7,opt,name=ipfs_api,json=ipfsApi,proto3" json:"ipfs_api,omitempty"`
	HnsTlds          []string `protobuf:"bytes,8,rep,name=hns_tlds,json=hnsTlds,proto3" json:"hns_tlds,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_795aedede39c29d8, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDidBaseContext() string {
	if m != nil {
		return m.DidBaseContext
	}
	return ""
}

func (m *Params) GetDidMethodContext() string {
	if m != nil {
		return m.DidMethodContext
	}
	return ""
}

func (m *Params) GetDidMethodName() string {
	if m != nil {
		return m.DidMethodName
	}
	return ""
}

func (m *Params) GetDidMethodVersion() string {
	if m != nil {
		return m.DidMethodVersion
	}
	return ""
}

func (m *Params) GetDidNetwork() string {
	if m != nil {
		return m.DidNetwork
	}
	return ""
}

func (m *Params) GetIpfsGateway() string {
	if m != nil {
		return m.IpfsGateway
	}
	return ""
}

func (m *Params) GetIpfsApi() string {
	if m != nil {
		return m.IpfsApi
	}
	return ""
}

func (m *Params) GetHnsTlds() []string {
	if m != nil {
		return m.HnsTlds
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "sonrio.sonr.identity.Params")
}

func init() { proto.RegisterFile("sonr/identity/params.proto", fileDescriptor_795aedede39c29d8) }

var fileDescriptor_795aedede39c29d8 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd1, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0x06, 0xf0, 0xa4, 0xed, 0x6d, 0x7b, 0x5d, 0xfe, 0xc9, 0xea, 0x60, 0x3a, 0xb8, 0x85, 0x01,
	0x75, 0x80, 0x64, 0x60, 0x63, 0xa3, 0x0c, 0x4c, 0x54, 0x08, 0x21, 0x06, 0x96, 0xc8, 0xad, 0x4d,
	0x62, 0xd1, 0xd8, 0x26, 0x36, 0xb4, 0x79, 0x0b, 0x46, 0x46, 0x9e, 0x82, 0x67, 0x60, 0xec, 0xc8,
	0x88, 0x92, 0x17, 0x41, 0x76, 0xd4, 0x82, 0x3a, 0x9d, 0xe4, 0xfb, 0x7e, 0x3e, 0xcb, 0x01, 0x3d,
	0x2d, 0x45, 0x16, 0x72, 0xca, 0x84, 0xe1, 0x26, 0x0f, 0x15, 0xc9, 0x48, 0xaa, 0x03, 0x95, 0x49,
	0x23, 0x61, 0xd7, 0x76, 0x5c, 0x06, 0x76, 0x04, 0x2b, 0xd2, 0xeb, 0xc6, 0x32, 0x96, 0x0e, 0x84,
	0xf6, 0xab, 0xb2, 0x87, 0x1f, 0x35, 0xd0, 0xbc, 0x76, 0x8f, 0xe1, 0x10, 0xec, 0x51, 0x4e, 0xa3,
	0x09, 0xd1, 0x2c, 0x9a, 0x4a, 0x61, 0xd8, 0xc2, 0x20, 0x7f, 0xe0, 0x0f, 0xff, 0xdf, 0xec, 0x50,
	0x4e, 0x47, 0x44, 0xb3, 0x8b, 0x2a, 0x85, 0xc7, 0x00, 0x5a, 0x99, 0x32, 0x93, 0x48, 0xba, 0xb6,
	0x35, 0x67, 0xed, 0x8e, 0x2b, 0x57, 0xac, 0xf4, 0x11, 0xd8, 0xfd, 0xa3, 0x05, 0x49, 0x19, 0xaa,
	0x3b, 0xba, 0xbd, 0xa6, 0x63, 0x92, 0xb2, 0x8d, 0xad, 0x2f, 0x2c, 0xd3, 0x5c, 0x0a, 0xd4, 0xd8,
	0xd8, 0x7a, 0x57, 0xe5, 0xb0, 0x0f, 0x3a, 0x56, 0x0b, 0x66, 0xe6, 0x32, 0x7b, 0x44, 0xff, 0x1c,
	0x03, 0x94, 0xd3, 0x71, 0x95, 0xc0, 0x03, 0xb0, 0xc5, 0xd5, 0x83, 0x8e, 0x62, 0x62, 0xd8, 0x9c,
	0xe4, 0xa8, 0xe9, 0x44, 0xc7, 0x66, 0x97, 0x55, 0x04, 0xf7, 0x41, 0xdb, 0x11, 0xa2, 0x38, 0x6a,
	0xb9, 0xba, 0x65, 0xff, 0xcf, 0x15, 0xb7, 0x55, 0x22, 0x74, 0x64, 0x66, 0x54, 0xa3, 0xf6, 0xa0,
	0x6e, 0xab, 0x44, 0xe8, 0xdb, 0x19, 0xd5, 0x67, 0x8d, 0xb7, 0xf7, 0xbe, 0x37, 0x1a, 0x7d, 0x16,
	0xd8, 0x5f, 0x16, 0xd8, 0xff, 0x2e, 0xb0, 0xff, 0x5a, 0x62, 0x6f, 0x59, 0x62, 0xef, 0xab, 0xc4,
	0xde, 0xfd, 0x30, 0xe6, 0x26, 0x79, 0x9e, 0x04, 0x53, 0x99, 0x86, 0xf6, 0x04, 0x27, 0xc9, 0x93,
	0x9b, 0xe1, 0xe2, 0xf7, 0x5e, 0x26, 0x57, 0x4c, 0x4f, 0x9a, 0xee, 0x06, 0xa7, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xc0, 0x7c, 0x08, 0x7a, 0xcd, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HnsTlds) > 0 {
		for iNdEx := len(m.HnsTlds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.HnsTlds[iNdEx])
			copy(dAtA[i:], m.HnsTlds[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.HnsTlds[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.IpfsApi) > 0 {
		i -= len(m.IpfsApi)
		copy(dAtA[i:], m.IpfsApi)
		i = encodeVarintParams(dAtA, i, uint64(len(m.IpfsApi)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.IpfsGateway) > 0 {
		i -= len(m.IpfsGateway)
		copy(dAtA[i:], m.IpfsGateway)
		i = encodeVarintParams(dAtA, i, uint64(len(m.IpfsGateway)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DidNetwork) > 0 {
		i -= len(m.DidNetwork)
		copy(dAtA[i:], m.DidNetwork)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DidNetwork)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DidMethodVersion) > 0 {
		i -= len(m.DidMethodVersion)
		copy(dAtA[i:], m.DidMethodVersion)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DidMethodVersion)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DidMethodName) > 0 {
		i -= len(m.DidMethodName)
		copy(dAtA[i:], m.DidMethodName)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DidMethodName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DidMethodContext) > 0 {
		i -= len(m.DidMethodContext)
		copy(dAtA[i:], m.DidMethodContext)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DidMethodContext)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DidBaseContext) > 0 {
		i -= len(m.DidBaseContext)
		copy(dAtA[i:], m.DidBaseContext)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DidBaseContext)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DidBaseContext)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DidMethodContext)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DidMethodName)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DidMethodVersion)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DidNetwork)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.IpfsGateway)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.IpfsApi)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.HnsTlds) > 0 {
		for _, s := range m.HnsTlds {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidBaseContext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidBaseContext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidMethodContext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidMethodContext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidMethodName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidMethodName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidMethodVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidMethodVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidNetwork", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidNetwork = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpfsGateway", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpfsGateway = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpfsApi", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpfsApi = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HnsTlds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HnsTlds = append(m.HnsTlds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
