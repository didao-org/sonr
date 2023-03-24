// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sonr/crypto/coin.proto

package crypto

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

// CoinType is the BIP-0044 coin type for each supported coin.
type CoinType int32

const (
	// Bitcoins coin type is 0
	CoinType_CoinType_BITCOIN CoinType = 0
	// Testnet coin type is 1
	CoinType_CoinType_TESTNET CoinType = 1
	// Litecoin coin type is 2
	CoinType_CoinType_LITECOIN CoinType = 2
	// Dogecoin coin type is 3
	CoinType_CoinType_DOGE CoinType = 3
	// Ethereum coin type is 60
	CoinType_CoinType_ETHEREUM CoinType = 4
	// Sonr coin type is 703
	CoinType_CoinType_SONR CoinType = 5
	// Cosmos coin type is 118
	CoinType_CoinType_COSMOS CoinType = 6
	// Filecoin coin type is 461
	CoinType_CoinType_FILECOIN CoinType = 7
	// Handshake coin type is 5353
	CoinType_CoinType_HNS CoinType = 8
	// Solana coin type is 501
	CoinType_CoinType_SOLANA CoinType = 9
	// Ripple coin type is 144
	CoinType_CoinType_XRP CoinType = 10
)

var CoinType_name = map[int32]string{
	0:  "CoinType_BITCOIN",
	1:  "CoinType_TESTNET",
	2:  "CoinType_LITECOIN",
	3:  "CoinType_DOGE",
	4:  "CoinType_ETHEREUM",
	5:  "CoinType_SONR",
	6:  "CoinType_COSMOS",
	7:  "CoinType_FILECOIN",
	8:  "CoinType_HNS",
	9:  "CoinType_SOLANA",
	10: "CoinType_XRP",
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
	"CoinType_SOLANA":   9,
	"CoinType_XRP":      10,
}

func (x CoinType) String() string {
	return proto.EnumName(CoinType_name, int32(x))
}

func (CoinType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_60f7eb46e7085759, []int{0}
}

func init() {
	proto.RegisterEnum("sonrhq.sonr.crypto.CoinType", CoinType_name, CoinType_value)
}

func init() { proto.RegisterFile("sonr/crypto/coin.proto", fileDescriptor_60f7eb46e7085759) }

var fileDescriptor_60f7eb46e7085759 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0xce, 0xcf, 0x2b,
	0xd2, 0x4f, 0x2e, 0xaa, 0x2c, 0x28, 0xc9, 0xd7, 0x4f, 0xce, 0xcf, 0xcc, 0xd3, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x02, 0x89, 0x67, 0x14, 0xea, 0x81, 0x28, 0x3d, 0x88, 0xb4, 0xd6, 0x7b,
	0x46, 0x2e, 0x0e, 0xe7, 0xfc, 0xcc, 0xbc, 0x90, 0xca, 0x82, 0x54, 0x21, 0x11, 0x2e, 0x01, 0x18,
	0x3b, 0xde, 0xc9, 0x33, 0xc4, 0xd9, 0xdf, 0xd3, 0x4f, 0x80, 0x01, 0x45, 0x34, 0xc4, 0x35, 0x38,
	0xc4, 0xcf, 0x35, 0x44, 0x80, 0x51, 0x48, 0x94, 0x4b, 0x10, 0x2e, 0xea, 0xe3, 0x19, 0xe2, 0x0a,
	0x56, 0xcc, 0x24, 0x24, 0xc8, 0xc5, 0x0b, 0x17, 0x76, 0xf1, 0x77, 0x77, 0x15, 0x60, 0x46, 0x51,
	0xe9, 0x1a, 0xe2, 0xe1, 0x1a, 0xe4, 0x1a, 0xea, 0x2b, 0xc0, 0x82, 0xa2, 0x32, 0xd8, 0xdf, 0x2f,
	0x48, 0x80, 0x55, 0x48, 0x98, 0x8b, 0x1f, 0x2e, 0xe4, 0xec, 0x1f, 0xec, 0xeb, 0x1f, 0x2c, 0xc0,
	0x86, 0xa2, 0xdd, 0xcd, 0xd3, 0x07, 0x62, 0x11, 0xbb, 0x90, 0x00, 0x17, 0x0f, 0x5c, 0xd8, 0xc3,
	0x2f, 0x58, 0x80, 0x03, 0x45, 0x77, 0xb0, 0xbf, 0x8f, 0xa3, 0x9f, 0xa3, 0x00, 0x27, 0x8a, 0xb2,
	0x88, 0xa0, 0x00, 0x01, 0x2e, 0x27, 0xdb, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63,
	0x88, 0x52, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x04, 0x95,
	0x7e, 0x72, 0x7e, 0x51, 0xaa, 0x7e, 0x49, 0x65, 0x41, 0x6a, 0x31, 0x34, 0x3c, 0x93, 0xd8, 0xc0,
	0x61, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x55, 0x41, 0xe0, 0x65, 0x01, 0x00, 0x00,
}