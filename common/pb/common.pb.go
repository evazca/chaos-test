// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NetworkOperate int32

const (
	NetworkOperate_Delay NetworkOperate = 0
	NetworkOperate_Loss  NetworkOperate = 1
)

var NetworkOperate_name = map[int32]string{
	0: "Delay",
	1: "Loss",
}

var NetworkOperate_value = map[string]int32{
	"Delay": 0,
	"Loss":  1,
}

func (x NetworkOperate) String() string {
	return proto.EnumName(NetworkOperate_name, int32(x))
}

func (NetworkOperate) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type ErrorType int32

const (
	ErrorType_UnknownError      ErrorType = 0
	ErrorType_InProcessError    ErrorType = 1
	ErrorType_NoError           ErrorType = 2
	ErrorType_NotInProcessError ErrorType = 3
	ErrorType_UnMatchIdError    ErrorType = 4
)

var ErrorType_name = map[int32]string{
	0: "UnknownError",
	1: "InProcessError",
	2: "NoError",
	3: "NotInProcessError",
	4: "UnMatchIdError",
}

var ErrorType_value = map[string]int32{
	"UnknownError":      0,
	"InProcessError":    1,
	"NoError":           2,
	"NotInProcessError": 3,
	"UnMatchIdError":    4,
}

func (x ErrorType) String() string {
	return proto.EnumName(ErrorType_name, int32(x))
}

func (ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

type Separation struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Flag                 string   `protobuf:"bytes,3,opt,name=flag,proto3" json:"flag,omitempty"`
	Protocol             string   `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Separation) Reset()         { *m = Separation{} }
func (m *Separation) String() string { return proto.CompactTextString(m) }
func (*Separation) ProtoMessage()    {}
func (*Separation) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Separation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Separation.Unmarshal(m, b)
}
func (m *Separation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Separation.Marshal(b, m, deterministic)
}
func (m *Separation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Separation.Merge(m, src)
}
func (m *Separation) XXX_Size() int {
	return xxx_messageInfo_Separation.Size(m)
}
func (m *Separation) XXX_DiscardUnknown() {
	xxx_messageInfo_Separation.DiscardUnknown(m)
}

var xxx_messageInfo_Separation proto.InternalMessageInfo

func (m *Separation) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Separation) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Separation) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

func (m *Separation) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type NetworkOperator struct {
	Operate              NetworkOperate `protobuf:"varint,1,opt,name=operate,proto3,enum=pb.NetworkOperate" json:"operate,omitempty"`
	Probability          int32          `protobuf:"varint,2,opt,name=probability,proto3" json:"probability,omitempty"`
	Delay                int32          `protobuf:"varint,3,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NetworkOperator) Reset()         { *m = NetworkOperator{} }
func (m *NetworkOperator) String() string { return proto.CompactTextString(m) }
func (*NetworkOperator) ProtoMessage()    {}
func (*NetworkOperator) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *NetworkOperator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkOperator.Unmarshal(m, b)
}
func (m *NetworkOperator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkOperator.Marshal(b, m, deterministic)
}
func (m *NetworkOperator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkOperator.Merge(m, src)
}
func (m *NetworkOperator) XXX_Size() int {
	return xxx_messageInfo_NetworkOperator.Size(m)
}
func (m *NetworkOperator) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkOperator.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkOperator proto.InternalMessageInfo

func (m *NetworkOperator) GetOperate() NetworkOperate {
	if m != nil {
		return m.Operate
	}
	return NetworkOperate_Delay
}

func (m *NetworkOperator) GetProbability() int32 {
	if m != nil {
		return m.Probability
	}
	return 0
}

func (m *NetworkOperator) GetDelay() int32 {
	if m != nil {
		return m.Delay
	}
	return 0
}

type NetworkOperateMark struct {
	Mark                 int32            `protobuf:"varint,1,opt,name=mark,proto3" json:"mark,omitempty"`
	Ip                   string           `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Separation           []*Separation    `protobuf:"bytes,3,rep,name=separation,proto3" json:"separation,omitempty"`
	NetworkOperator      *NetworkOperator `protobuf:"bytes,4,opt,name=networkOperator,proto3" json:"networkOperator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NetworkOperateMark) Reset()         { *m = NetworkOperateMark{} }
func (m *NetworkOperateMark) String() string { return proto.CompactTextString(m) }
func (*NetworkOperateMark) ProtoMessage()    {}
func (*NetworkOperateMark) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *NetworkOperateMark) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkOperateMark.Unmarshal(m, b)
}
func (m *NetworkOperateMark) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkOperateMark.Marshal(b, m, deterministic)
}
func (m *NetworkOperateMark) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkOperateMark.Merge(m, src)
}
func (m *NetworkOperateMark) XXX_Size() int {
	return xxx_messageInfo_NetworkOperateMark.Size(m)
}
func (m *NetworkOperateMark) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkOperateMark.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkOperateMark proto.InternalMessageInfo

func (m *NetworkOperateMark) GetMark() int32 {
	if m != nil {
		return m.Mark
	}
	return 0
}

func (m *NetworkOperateMark) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *NetworkOperateMark) GetSeparation() []*Separation {
	if m != nil {
		return m.Separation
	}
	return nil
}

func (m *NetworkOperateMark) GetNetworkOperator() *NetworkOperator {
	if m != nil {
		return m.NetworkOperator
	}
	return nil
}

type CommonResp struct {
	Result               bool      `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	ErrorType            ErrorType `protobuf:"varint,2,opt,name=errorType,proto3,enum=pb.ErrorType" json:"errorType,omitempty"`
	Id                   string    `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CommonResp) Reset()         { *m = CommonResp{} }
func (m *CommonResp) String() string { return proto.CompactTextString(m) }
func (*CommonResp) ProtoMessage()    {}
func (*CommonResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *CommonResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonResp.Unmarshal(m, b)
}
func (m *CommonResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonResp.Marshal(b, m, deterministic)
}
func (m *CommonResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResp.Merge(m, src)
}
func (m *CommonResp) XXX_Size() int {
	return xxx_messageInfo_CommonResp.Size(m)
}
func (m *CommonResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResp.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResp proto.InternalMessageInfo

func (m *CommonResp) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *CommonResp) GetErrorType() ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return ErrorType_UnknownError
}

func (m *CommonResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.NetworkOperate", NetworkOperate_name, NetworkOperate_value)
	proto.RegisterEnum("pb.ErrorType", ErrorType_name, ErrorType_value)
	proto.RegisterType((*Separation)(nil), "pb.separation")
	proto.RegisterType((*NetworkOperator)(nil), "pb.networkOperator")
	proto.RegisterType((*NetworkOperateMark)(nil), "pb.networkOperateMark")
	proto.RegisterType((*CommonResp)(nil), "pb.commonResp")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x5d, 0xd2, 0x66, 0x6b, 0x6f, 0x47, 0x28, 0x97, 0x0f, 0x45, 0x13, 0x0f, 0x55, 0x25, 0xa4,
	0x6a, 0x20, 0x3f, 0x94, 0x67, 0x5e, 0x10, 0x3c, 0x4c, 0x62, 0x03, 0x59, 0xec, 0x1d, 0x27, 0x35,
	0xab, 0xd5, 0xd4, 0xd7, 0xb2, 0x3d, 0xaa, 0xf2, 0x87, 0xf8, 0x9b, 0x28, 0x37, 0x59, 0xd7, 0xae,
	0x6f, 0xf7, 0x9e, 0x9c, 0x1c, 0x9f, 0x73, 0x6c, 0x38, 0xaf, 0x68, 0xbd, 0x26, 0x2b, 0x9c, 0xa7,
	0x48, 0x98, 0xba, 0x72, 0xfa, 0x0b, 0x20, 0x68, 0xa7, 0xbc, 0x8a, 0x86, 0x2c, 0xe6, 0x90, 0x1a,
	0x57, 0x24, 0x93, 0x64, 0x36, 0x94, 0xa9, 0x71, 0x88, 0xd0, 0x77, 0xe4, 0x63, 0x91, 0x4e, 0x92,
	0x59, 0x26, 0x79, 0x6e, 0xb0, 0xdf, 0xb5, 0xba, 0x2b, 0x7a, 0xcc, 0xe2, 0x19, 0x2f, 0x60, 0xc0,
	0x92, 0x15, 0xd5, 0x45, 0x9f, 0xf1, 0xdd, 0x3e, 0xdd, 0xc0, 0x73, 0xab, 0xe3, 0x86, 0xfc, 0xea,
	0xbb, 0xd3, 0x5e, 0x45, 0xf2, 0xf8, 0x01, 0xce, 0x88, 0x67, 0xcd, 0x67, 0xe5, 0x73, 0x14, 0xae,
	0x14, 0x07, 0x2c, 0x2d, 0x1f, 0x28, 0x38, 0x81, 0x91, 0xf3, 0x54, 0xaa, 0xd2, 0xd4, 0x26, 0x6e,
	0x3b, 0x2f, 0xfb, 0x10, 0xbe, 0x82, 0x6c, 0xa1, 0x6b, 0xb5, 0x65, 0x4f, 0x99, 0x6c, 0x97, 0xe9,
	0xbf, 0x04, 0xf0, 0x50, 0xf3, 0x5a, 0xf9, 0x55, 0xe3, 0x7f, 0xad, 0xfc, 0x8a, 0x4f, 0xce, 0x24,
	0xcf, 0x5d, 0xee, 0x74, 0x97, 0x5b, 0xec, 0xb7, 0x52, 0xf4, 0x26, 0xbd, 0xd9, 0x68, 0x9e, 0x37,
	0x1e, 0x1f, 0x51, 0xb9, 0xdf, 0xdb, 0xa7, 0xa3, 0x8c, 0x5c, 0xc3, 0x68, 0xfe, 0xf2, 0x28, 0x18,
	0x79, 0xf9, 0x94, 0x3b, 0x55, 0x00, 0xed, 0xc5, 0x48, 0x1d, 0x1c, 0xbe, 0x81, 0x53, 0xaf, 0xc3,
	0x7d, 0x1d, 0xd9, 0xe2, 0x40, 0x76, 0x1b, 0xbe, 0x87, 0xa1, 0xf6, 0x9e, 0xfc, 0xcf, 0xad, 0xd3,
	0xec, 0x35, 0x9f, 0x3f, 0x6b, 0xe4, 0xbf, 0x3e, 0x80, 0xf2, 0xf1, 0x3b, 0x27, 0x5a, 0x74, 0x77,
	0x94, 0x9a, 0xc5, 0xe5, 0x3b, 0xc8, 0x0f, 0xbb, 0xc0, 0x21, 0x64, 0x5f, 0x9a, 0x9e, 0xc6, 0x27,
	0x38, 0x80, 0xfe, 0x37, 0x0a, 0x61, 0x9c, 0x5c, 0x1a, 0x18, 0xee, 0xe4, 0x70, 0x0c, 0xe7, 0xb7,
	0x76, 0x65, 0x69, 0x63, 0x19, 0x1b, 0x9f, 0x20, 0x42, 0x7e, 0x65, 0x7f, 0x78, 0xaa, 0x74, 0x08,
	0x2d, 0x96, 0xe0, 0x08, 0xce, 0x6e, 0xa8, 0x5d, 0x52, 0x7c, 0x0d, 0x2f, 0x6e, 0x28, 0x3e, 0xe1,
	0xf4, 0x9a, 0xff, 0x6e, 0xed, 0xb5, 0x8a, 0xd5, 0xf2, 0x6a, 0xd1, 0x62, 0xfd, 0xcf, 0x6f, 0xe1,
	0xa2, 0xa2, 0xb5, 0xb8, 0x33, 0x71, 0x79, 0x5f, 0x0a, 0xfd, 0x47, 0xfd, 0xad, 0x94, 0xa8, 0x96,
	0x8a, 0x82, 0x70, 0x65, 0x79, 0xca, 0xef, 0xe7, 0xe3, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52,
	0xe3, 0xf5, 0xaa, 0xb2, 0x02, 0x00, 0x00,
}
