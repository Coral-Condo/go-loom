// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/config/config.proto

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/config/config.proto

It has these top-level messages:
	ConfigInitRequest
	GetSetting
	UpdateSetting
	Value
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ReceiptStorage int32

const (
	ReceiptStorage_CHAIN   ReceiptStorage = 0
	ReceiptStorage_LEVELDB ReceiptStorage = 1
)

var ReceiptStorage_name = map[int32]string{
	0: "CHAIN",
	1: "LEVELDB",
}
var ReceiptStorage_value = map[string]int32{
	"CHAIN":   0,
	"LEVELDB": 1,
}

func (x ReceiptStorage) String() string {
	return proto.EnumName(ReceiptStorage_name, int32(x))
}
func (ReceiptStorage) EnumDescriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

type ConfigInitRequest struct {
	Oracle   *types.Address   `protobuf:"bytes,1,opt,name=oracle" json:"oracle,omitempty"`
	Settings []*UpdateSetting `protobuf:"bytes,2,rep,name=settings" json:"settings,omitempty"`
}

func (m *ConfigInitRequest) Reset()                    { *m = ConfigInitRequest{} }
func (m *ConfigInitRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigInitRequest) ProtoMessage()               {}
func (*ConfigInitRequest) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *ConfigInitRequest) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

func (m *ConfigInitRequest) GetSettings() []*UpdateSetting {
	if m != nil {
		return m.Settings
	}
	return nil
}

type GetSetting struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *GetSetting) Reset()                    { *m = GetSetting{} }
func (m *GetSetting) String() string            { return proto.CompactTextString(m) }
func (*GetSetting) ProtoMessage()               {}
func (*GetSetting) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

func (m *GetSetting) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type UpdateSetting struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *Value `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *UpdateSetting) Reset()                    { *m = UpdateSetting{} }
func (m *UpdateSetting) String() string            { return proto.CompactTextString(m) }
func (*UpdateSetting) ProtoMessage()               {}
func (*UpdateSetting) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

func (m *UpdateSetting) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UpdateSetting) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type Value struct {
	// Types that are valid to be assigned to Data:
	//	*Value_Uint64Val
	//	*Value_ReceiptStorage
	//	*Value_Address
	Data isValue_Data `protobuf_oneof:"data"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{3} }

type isValue_Data interface {
	isValue_Data()
}

type Value_Uint64Val struct {
	Uint64Val uint64 `protobuf:"varint,1,opt,name=uint64_val,json=uint64Val,proto3,oneof"`
}
type Value_ReceiptStorage struct {
	ReceiptStorage ReceiptStorage `protobuf:"varint,2,opt,name=receipt_storage,json=receiptStorage,proto3,enum=ReceiptStorage,oneof"`
}
type Value_Address struct {
	Address *types.Address `protobuf:"bytes,3,opt,name=address,oneof"`
}

func (*Value_Uint64Val) isValue_Data()      {}
func (*Value_ReceiptStorage) isValue_Data() {}
func (*Value_Address) isValue_Data()        {}

func (m *Value) GetData() isValue_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Value) GetUint64Val() uint64 {
	if x, ok := m.GetData().(*Value_Uint64Val); ok {
		return x.Uint64Val
	}
	return 0
}

func (m *Value) GetReceiptStorage() ReceiptStorage {
	if x, ok := m.GetData().(*Value_ReceiptStorage); ok {
		return x.ReceiptStorage
	}
	return ReceiptStorage_CHAIN
}

func (m *Value) GetAddress() *types.Address {
	if x, ok := m.GetData().(*Value_Address); ok {
		return x.Address
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_Uint64Val)(nil),
		(*Value_ReceiptStorage)(nil),
		(*Value_Address)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// data
	switch x := m.Data.(type) {
	case *Value_Uint64Val:
		_ = b.EncodeVarint(1<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Uint64Val))
	case *Value_ReceiptStorage:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.ReceiptStorage))
	case *Value_Address:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Address); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Value.Data has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // data.uint64_val
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Data = &Value_Uint64Val{x}
		return true, err
	case 2: // data.receipt_storage
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Data = &Value_ReceiptStorage{ReceiptStorage(x)}
		return true, err
	case 3: // data.address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types.Address)
		err := b.DecodeMessage(msg)
		m.Data = &Value_Address{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// data
	switch x := m.Data.(type) {
	case *Value_Uint64Val:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Uint64Val))
	case *Value_ReceiptStorage:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.ReceiptStorage))
	case *Value_Address:
		s := proto.Size(x.Address)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ConfigInitRequest)(nil), "ConfigInitRequest")
	proto.RegisterType((*GetSetting)(nil), "GetSetting")
	proto.RegisterType((*UpdateSetting)(nil), "UpdateSetting")
	proto.RegisterType((*Value)(nil), "Value")
	proto.RegisterEnum("ReceiptStorage", ReceiptStorage_name, ReceiptStorage_value)
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/config/config.proto", fileDescriptorConfig)
}

var fileDescriptorConfig = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x5f, 0x4f, 0xf2, 0x30,
	0x14, 0x87, 0x37, 0xfe, 0x0c, 0x38, 0xe4, 0x1d, 0xbc, 0xbd, 0x5a, 0x8c, 0x51, 0xb2, 0x78, 0x41,
	0x48, 0xdc, 0x0c, 0x1a, 0x2f, 0xbc, 0x51, 0x40, 0xe2, 0x48, 0x88, 0x17, 0x25, 0x72, 0x4b, 0xca,
	0x56, 0x67, 0xc3, 0x58, 0xe7, 0x76, 0x86, 0xe1, 0x6b, 0xf8, 0x89, 0x0d, 0x2b, 0x9a, 0x2c, 0x5e,
	0x78, 0xd3, 0xf6, 0x3c, 0xa7, 0xbf, 0x27, 0xcd, 0x29, 0x3c, 0x84, 0x02, 0xdf, 0xf2, 0xb5, 0xe3,
	0xcb, 0xad, 0x1b, 0x49, 0xb9, 0x8d, 0x39, 0x7e, 0xc8, 0x74, 0xe3, 0x86, 0xf2, 0xf2, 0x50, 0xba,
	0xeb, 0x5c, 0x44, 0x28, 0x62, 0x17, 0xf7, 0x09, 0xcf, 0x5c, 0x5f, 0xc6, 0xaf, 0x22, 0x3c, 0x6e,
	0x4e, 0x92, 0x4a, 0x94, 0x27, 0x57, 0x7f, 0x18, 0x54, 0xb2, 0x58, 0x55, 0xc2, 0x66, 0xf0, 0x7f,
	0x52, 0x18, 0x66, 0xb1, 0x40, 0xca, 0xdf, 0x73, 0x9e, 0x21, 0xe9, 0x81, 0x21, 0x53, 0xe6, 0x47,
	0xdc, 0xd2, 0x7b, 0x7a, 0xbf, 0x3d, 0x6c, 0x3a, 0xa3, 0x20, 0x48, 0x79, 0x96, 0xd1, 0x23, 0x27,
	0x03, 0x68, 0x66, 0x1c, 0x51, 0xc4, 0x61, 0x66, 0x55, 0x7a, 0xd5, 0x7e, 0x7b, 0x68, 0x3a, 0x2f,
	0x49, 0xc0, 0x90, 0x2f, 0x14, 0xa6, 0x3f, 0x7d, 0xfb, 0x0c, 0xe0, 0x89, 0xe3, 0x91, 0x93, 0x2e,
	0x54, 0x37, 0x7c, 0x5f, 0x88, 0x5b, 0xf4, 0x70, 0xb4, 0xef, 0xe1, 0x5f, 0x29, 0xfa, 0xfb, 0x0a,
	0x39, 0x85, 0xfa, 0x8e, 0x45, 0x39, 0xb7, 0x2a, 0xc5, 0x7b, 0x0c, 0x67, 0x79, 0xa8, 0xa8, 0x82,
	0xf6, 0xa7, 0x0e, 0xf5, 0x02, 0x90, 0x73, 0x80, 0x5c, 0xc4, 0x78, 0x7b, 0xb3, 0xda, 0xb1, 0xa8,
	0x10, 0xd4, 0x3c, 0x8d, 0xb6, 0x14, 0x5b, 0xb2, 0x88, 0xdc, 0x41, 0x27, 0xe5, 0x3e, 0x17, 0x09,
	0xae, 0x32, 0x94, 0x29, 0x0b, 0x95, 0xd2, 0x1c, 0x76, 0x1c, 0xaa, 0xf8, 0x42, 0x61, 0x4f, 0xa3,
	0x66, 0x5a, 0x22, 0xe4, 0x02, 0x1a, 0x4c, 0x8d, 0xc1, 0xaa, 0x96, 0xc7, 0xe2, 0x69, 0xf4, 0xbb,
	0x35, 0x36, 0xa0, 0x16, 0x30, 0x64, 0x83, 0x3e, 0x98, 0x65, 0x23, 0x69, 0x41, 0x7d, 0xe2, 0x8d,
	0x66, 0xcf, 0x5d, 0x8d, 0xb4, 0xa1, 0x31, 0x9f, 0x2e, 0xa7, 0xf3, 0xc7, 0x71, 0x57, 0x5f, 0x1b,
	0xc5, 0x4f, 0x5c, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x12, 0xe6, 0x63, 0xf1, 0xff, 0x01, 0x00,
	0x00,
}