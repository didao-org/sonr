// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/crypto/coin.proto

// Package common defines commonly used types agnostic to the node role on the Sonr network.

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

type CoinType int32

const (
	CoinType_CoinType_BITCOIN  CoinType = 0
	CoinType_CoinType_TESTNET  CoinType = 1
	CoinType_CoinType_LITECOIN CoinType = 2
	CoinType_CoinType_DOGE     CoinType = 3
	CoinType_CoinType_ETHEREUM CoinType = 4
	CoinType_CoinType_SONR     CoinType = 5
	CoinType_CoinType_COSMOS   CoinType = 6
	CoinType_CoinType_FILECOIN CoinType = 7
	CoinType_CoinType_HNS      CoinType = 8
)

var CoinType_name = map[int32]string{
	0: "CoinType_BITCOIN",
	1: "CoinType_TESTNET",
	2: "CoinType_LITECOIN",
	3: "CoinType_DOGE",
	4: "CoinType_ETHEREUM",
	5: "CoinType_SONR",
	6: "CoinType_COSMOS",
	7: "CoinType_FILECOIN",
	8: "CoinType_HNS",
}

var CoinType_value = map[string]int32{
	"CoinType_BITCOIN":  0,
	"CoinType_TESTNET":  1,
	"CoinType_LITECOIN": 2,
	"CoinType_DOGE":     3,
	"CoinType_ETHEREUM": 4,
	"CoinType_SONR":     5,
	"CoinType_COSMOS":   6,
	"CoinType_FILECOIN": 7,
	"CoinType_HNS":      8,
}

func (x CoinType) String() string {
	return proto.EnumName(CoinType_name, int32(x))
}

func (CoinType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0aeeac340f2c4abc, []int{0}
}

func init() {
	proto.RegisterEnum("sonrhq.common.v1.CoinType", CoinType_name, CoinType_value)
}

func init() { proto.RegisterFile("common/crypto/coin.proto", fileDescriptor_0aeeac340f2c4abc) }

var fileDescriptor_0aeeac340f2c4abc = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x4f, 0x2e, 0xaa, 0x2c, 0x28, 0xc9, 0xd7, 0x4f, 0xce, 0xcf, 0xcc, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28, 0xce, 0xcf, 0x2b, 0xca, 0x28, 0xd4, 0x83, 0x28, 0xd0,
	0x2b, 0x33, 0xd4, 0x3a, 0xc1, 0xc8, 0xc5, 0xe1, 0x9c, 0x9f, 0x99, 0x17, 0x52, 0x59, 0x90, 0x2a,
	0x24, 0xc2, 0x25, 0x00, 0x63, 0xc7, 0x3b, 0x79, 0x86, 0x38, 0xfb, 0x7b, 0xfa, 0x09, 0x30, 0xa0,
	0x88, 0x86, 0xb8, 0x06, 0x87, 0xf8, 0xb9, 0x86, 0x08, 0x30, 0x0a, 0x89, 0x72, 0x09, 0xc2, 0x45,
	0x7d, 0x3c, 0x43, 0x5c, 0xc1, 0x8a, 0x99, 0x84, 0x04, 0xb9, 0x78, 0xe1, 0xc2, 0x2e, 0xfe, 0xee,
	0xae, 0x02, 0xcc, 0x28, 0x2a, 0x5d, 0x43, 0x3c, 0x5c, 0x83, 0x5c, 0x43, 0x7d, 0x05, 0x58, 0x50,
	0x54, 0x06, 0xfb, 0xfb, 0x05, 0x09, 0xb0, 0x0a, 0x09, 0x73, 0xf1, 0xc3, 0x85, 0x9c, 0xfd, 0x83,
	0x7d, 0xfd, 0x83, 0x05, 0xd8, 0x50, 0xb4, 0xbb, 0x79, 0xfa, 0x40, 0x2c, 0x62, 0x17, 0x12, 0xe0,
	0xe2, 0x81, 0x0b, 0x7b, 0xf8, 0x05, 0x0b, 0x70, 0x38, 0x39, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x7a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x12, 0xc8, 0xdb, 0xfa, 0x90,
	0x10, 0xd0, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0xc8, 0x4e, 0x87, 0x05, 0x52, 0x49, 0x65, 0x41,
	0x6a, 0x71, 0x12, 0x1b, 0x38, 0x98, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x25, 0x81,
	0x21, 0x42, 0x01, 0x00, 0x00,
}
