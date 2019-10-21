// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PrepareReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareReq) Reset()         { *m = PrepareReq{} }
func (m *PrepareReq) String() string { return proto.CompactTextString(m) }
func (*PrepareReq) ProtoMessage()    {}
func (*PrepareReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{0}
}

func (m *PrepareReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareReq.Unmarshal(m, b)
}
func (m *PrepareReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareReq.Marshal(b, m, deterministic)
}
func (m *PrepareReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareReq.Merge(m, src)
}
func (m *PrepareReq) XXX_Size() int {
	return xxx_messageInfo_PrepareReq.Size(m)
}
func (m *PrepareReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareReq.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareReq proto.InternalMessageInfo

func (m *PrepareReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PrepareResp struct {
	CommonResp           *CommonResp `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PrepareResp) Reset()         { *m = PrepareResp{} }
func (m *PrepareResp) String() string { return proto.CompactTextString(m) }
func (*PrepareResp) ProtoMessage()    {}
func (*PrepareResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{1}
}

func (m *PrepareResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareResp.Unmarshal(m, b)
}
func (m *PrepareResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareResp.Marshal(b, m, deterministic)
}
func (m *PrepareResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareResp.Merge(m, src)
}
func (m *PrepareResp) XXX_Size() int {
	return xxx_messageInfo_PrepareResp.Size(m)
}
func (m *PrepareResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareResp.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareResp proto.InternalMessageInfo

func (m *PrepareResp) GetCommonResp() *CommonResp {
	if m != nil {
		return m.CommonResp
	}
	return nil
}

type RevokeReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Force                bool     `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeReq) Reset()         { *m = RevokeReq{} }
func (m *RevokeReq) String() string { return proto.CompactTextString(m) }
func (*RevokeReq) ProtoMessage()    {}
func (*RevokeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{2}
}

func (m *RevokeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeReq.Unmarshal(m, b)
}
func (m *RevokeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeReq.Marshal(b, m, deterministic)
}
func (m *RevokeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeReq.Merge(m, src)
}
func (m *RevokeReq) XXX_Size() int {
	return xxx_messageInfo_RevokeReq.Size(m)
}
func (m *RevokeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeReq.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeReq proto.InternalMessageInfo

func (m *RevokeReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RevokeReq) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

type RevokeResp struct {
	CommonResp           *CommonResp `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RevokeResp) Reset()         { *m = RevokeResp{} }
func (m *RevokeResp) String() string { return proto.CompactTextString(m) }
func (*RevokeResp) ProtoMessage()    {}
func (*RevokeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{3}
}

func (m *RevokeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeResp.Unmarshal(m, b)
}
func (m *RevokeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeResp.Marshal(b, m, deterministic)
}
func (m *RevokeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeResp.Merge(m, src)
}
func (m *RevokeResp) XXX_Size() int {
	return xxx_messageInfo_RevokeResp.Size(m)
}
func (m *RevokeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeResp.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeResp proto.InternalMessageInfo

func (m *RevokeResp) GetCommonResp() *CommonResp {
	if m != nil {
		return m.CommonResp
	}
	return nil
}

type NetworkOperateReq struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NetworkOperateMark   *NetworkOperateMark `protobuf:"bytes,2,opt,name=networkOperateMark,proto3" json:"networkOperateMark,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NetworkOperateReq) Reset()         { *m = NetworkOperateReq{} }
func (m *NetworkOperateReq) String() string { return proto.CompactTextString(m) }
func (*NetworkOperateReq) ProtoMessage()    {}
func (*NetworkOperateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{4}
}

func (m *NetworkOperateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkOperateReq.Unmarshal(m, b)
}
func (m *NetworkOperateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkOperateReq.Marshal(b, m, deterministic)
}
func (m *NetworkOperateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkOperateReq.Merge(m, src)
}
func (m *NetworkOperateReq) XXX_Size() int {
	return xxx_messageInfo_NetworkOperateReq.Size(m)
}
func (m *NetworkOperateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkOperateReq.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkOperateReq proto.InternalMessageInfo

func (m *NetworkOperateReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NetworkOperateReq) GetNetworkOperateMark() *NetworkOperateMark {
	if m != nil {
		return m.NetworkOperateMark
	}
	return nil
}

type NetworkOperateResp struct {
	CommonResp           *CommonResp `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp,omitempty"`
	Mark                 int32       `protobuf:"varint,2,opt,name=mark,proto3" json:"mark,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NetworkOperateResp) Reset()         { *m = NetworkOperateResp{} }
func (m *NetworkOperateResp) String() string { return proto.CompactTextString(m) }
func (*NetworkOperateResp) ProtoMessage()    {}
func (*NetworkOperateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{5}
}

func (m *NetworkOperateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkOperateResp.Unmarshal(m, b)
}
func (m *NetworkOperateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkOperateResp.Marshal(b, m, deterministic)
}
func (m *NetworkOperateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkOperateResp.Merge(m, src)
}
func (m *NetworkOperateResp) XXX_Size() int {
	return xxx_messageInfo_NetworkOperateResp.Size(m)
}
func (m *NetworkOperateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkOperateResp.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkOperateResp proto.InternalMessageInfo

func (m *NetworkOperateResp) GetCommonResp() *CommonResp {
	if m != nil {
		return m.CommonResp
	}
	return nil
}

func (m *NetworkOperateResp) GetMark() int32 {
	if m != nil {
		return m.Mark
	}
	return 0
}

type RevokeNetworkOperateReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mark                 int32    `protobuf:"varint,2,opt,name=mark,proto3" json:"mark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeNetworkOperateReq) Reset()         { *m = RevokeNetworkOperateReq{} }
func (m *RevokeNetworkOperateReq) String() string { return proto.CompactTextString(m) }
func (*RevokeNetworkOperateReq) ProtoMessage()    {}
func (*RevokeNetworkOperateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{6}
}

func (m *RevokeNetworkOperateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeNetworkOperateReq.Unmarshal(m, b)
}
func (m *RevokeNetworkOperateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeNetworkOperateReq.Marshal(b, m, deterministic)
}
func (m *RevokeNetworkOperateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeNetworkOperateReq.Merge(m, src)
}
func (m *RevokeNetworkOperateReq) XXX_Size() int {
	return xxx_messageInfo_RevokeNetworkOperateReq.Size(m)
}
func (m *RevokeNetworkOperateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeNetworkOperateReq.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeNetworkOperateReq proto.InternalMessageInfo

func (m *RevokeNetworkOperateReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RevokeNetworkOperateReq) GetMark() int32 {
	if m != nil {
		return m.Mark
	}
	return 0
}

type RevokeNetworkOperateResp struct {
	CommonResp           *CommonResp `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RevokeNetworkOperateResp) Reset()         { *m = RevokeNetworkOperateResp{} }
func (m *RevokeNetworkOperateResp) String() string { return proto.CompactTextString(m) }
func (*RevokeNetworkOperateResp) ProtoMessage()    {}
func (*RevokeNetworkOperateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{7}
}

func (m *RevokeNetworkOperateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeNetworkOperateResp.Unmarshal(m, b)
}
func (m *RevokeNetworkOperateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeNetworkOperateResp.Marshal(b, m, deterministic)
}
func (m *RevokeNetworkOperateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeNetworkOperateResp.Merge(m, src)
}
func (m *RevokeNetworkOperateResp) XXX_Size() int {
	return xxx_messageInfo_RevokeNetworkOperateResp.Size(m)
}
func (m *RevokeNetworkOperateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeNetworkOperateResp.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeNetworkOperateResp proto.InternalMessageInfo

func (m *RevokeNetworkOperateResp) GetCommonResp() *CommonResp {
	if m != nil {
		return m.CommonResp
	}
	return nil
}

func init() {
	proto.RegisterType((*PrepareReq)(nil), "pb.prepareReq")
	proto.RegisterType((*PrepareResp)(nil), "pb.prepareResp")
	proto.RegisterType((*RevokeReq)(nil), "pb.revokeReq")
	proto.RegisterType((*RevokeResp)(nil), "pb.revokeResp")
	proto.RegisterType((*NetworkOperateReq)(nil), "pb.networkOperateReq")
	proto.RegisterType((*NetworkOperateResp)(nil), "pb.networkOperateResp")
	proto.RegisterType((*RevokeNetworkOperateReq)(nil), "pb.revokeNetworkOperateReq")
	proto.RegisterType((*RevokeNetworkOperateResp)(nil), "pb.revokeNetworkOperateResp")
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor_56ede974c0020f77) }

var fileDescriptor_56ede974c0020f77 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x25, 0x11, 0x2d, 0xf4, 0x02, 0x41, 0x9c, 0x4a, 0xa9, 0x42, 0x87, 0xca, 0x53, 0x91, 0x90,
	0x25, 0xca, 0x4a, 0x87, 0x32, 0x30, 0x20, 0x01, 0x92, 0x59, 0x58, 0x9d, 0xd4, 0xb4, 0x55, 0x94,
	0xda, 0xb8, 0xa1, 0x48, 0xfc, 0x72, 0x46, 0x64, 0x87, 0x7e, 0x84, 0x18, 0x21, 0xba, 0xe5, 0x3e,
	0xde, 0xbb, 0x77, 0x97, 0x67, 0x08, 0xf8, 0x58, 0xcc, 0x72, 0xaa, 0xb4, 0xcc, 0x25, 0xfa, 0x2a,
	0x8e, 0x0e, 0x12, 0x99, 0x65, 0x72, 0x56, 0x64, 0x48, 0x07, 0x40, 0x69, 0xa1, 0xb8, 0x16, 0x4c,
	0xbc, 0x62, 0x08, 0xfe, 0x74, 0xd4, 0xf6, 0xba, 0x5e, 0xaf, 0xc1, 0xfc, 0xe9, 0x88, 0x0c, 0x20,
	0x58, 0x55, 0xe7, 0x0a, 0x29, 0x40, 0x01, 0x36, 0x91, 0x6d, 0x0b, 0xfa, 0x21, 0x55, 0x31, 0x5d,
	0x67, 0xd9, 0x46, 0x07, 0xb9, 0x84, 0x86, 0x16, 0x0b, 0x99, 0xba, 0xb8, 0xb1, 0x09, 0xb5, 0x17,
	0xa9, 0x13, 0xd1, 0xf6, 0xbb, 0x5e, 0x6f, 0x9f, 0x15, 0x01, 0xb9, 0x06, 0x58, 0x42, 0xb6, 0x18,
	0x98, 0xc2, 0xf1, 0x4c, 0xe4, 0xef, 0x52, 0xa7, 0x8f, 0x4a, 0x68, 0x9e, 0x3b, 0x07, 0xdf, 0x02,
	0x96, 0x9b, 0xee, 0xb9, 0x4e, 0xad, 0x8a, 0xa0, 0xdf, 0x32, 0xe4, 0xd5, 0x2a, 0x73, 0x20, 0xc8,
	0xf3, 0x4f, 0x9e, 0x6d, 0x24, 0x23, 0xc2, 0x6e, 0xb6, 0x9c, 0x5f, 0x63, 0xf6, 0x9b, 0x0c, 0xe0,
	0xb4, 0x38, 0xc2, 0xc3, 0x9f, 0xcb, 0xb8, 0xe0, 0x77, 0xd0, 0x76, 0xc3, 0xff, 0x2f, 0xaf, 0xff,
	0xe9, 0x41, 0x6d, 0x68, 0x1c, 0x84, 0x17, 0xb0, 0xf7, 0xed, 0x05, 0xb4, 0x80, 0xb5, 0x6d, 0xa2,
	0xa3, 0x52, 0x3c, 0x57, 0x64, 0x07, 0xcf, 0xa1, 0x5e, 0x68, 0xc0, 0x43, 0x53, 0x5c, 0xd9, 0x20,
	0x0a, 0x37, 0x43, 0xdb, 0x3a, 0x84, 0xb0, 0x7c, 0x47, 0x3c, 0xa9, 0xfe, 0x05, 0x03, 0x6d, 0xb9,
	0xd2, 0x96, 0xe2, 0x09, 0x9a, 0xae, 0x8d, 0xf1, 0x6c, 0x3d, 0xac, 0x72, 0xca, 0xa8, 0xf3, 0x7b,
	0xd1, 0x90, 0xde, 0x74, 0x20, 0x4a, 0x64, 0x46, 0xc7, 0xd3, 0x7c, 0xf2, 0x16, 0x53, 0xb1, 0xe0,
	0x1f, 0x09, 0xa7, 0xc9, 0x84, 0xcb, 0x39, 0x55, 0x71, 0x5c, 0xb7, 0xef, 0xe7, 0xea, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0x69, 0xd3, 0xbc, 0x6c, 0x60, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	Prepare(ctx context.Context, in *PrepareReq, opts ...grpc.CallOption) (*PrepareResp, error)
	Revoke(ctx context.Context, in *RevokeReq, opts ...grpc.CallOption) (*RevokeResp, error)
	NetworkOperate(ctx context.Context, in *NetworkOperateReq, opts ...grpc.CallOption) (*NetworkOperateResp, error)
	RevokeNetworkOperate(ctx context.Context, in *RevokeNetworkOperateReq, opts ...grpc.CallOption) (*RevokeNetworkOperateResp, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Prepare(ctx context.Context, in *PrepareReq, opts ...grpc.CallOption) (*PrepareResp, error) {
	out := new(PrepareResp)
	err := c.cc.Invoke(ctx, "/pb.Agent/prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) Revoke(ctx context.Context, in *RevokeReq, opts ...grpc.CallOption) (*RevokeResp, error) {
	out := new(RevokeResp)
	err := c.cc.Invoke(ctx, "/pb.Agent/revoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) NetworkOperate(ctx context.Context, in *NetworkOperateReq, opts ...grpc.CallOption) (*NetworkOperateResp, error) {
	out := new(NetworkOperateResp)
	err := c.cc.Invoke(ctx, "/pb.Agent/networkOperate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) RevokeNetworkOperate(ctx context.Context, in *RevokeNetworkOperateReq, opts ...grpc.CallOption) (*RevokeNetworkOperateResp, error) {
	out := new(RevokeNetworkOperateResp)
	err := c.cc.Invoke(ctx, "/pb.Agent/revokeNetworkOperate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	Prepare(context.Context, *PrepareReq) (*PrepareResp, error)
	Revoke(context.Context, *RevokeReq) (*RevokeResp, error)
	NetworkOperate(context.Context, *NetworkOperateReq) (*NetworkOperateResp, error)
	RevokeNetworkOperate(context.Context, *RevokeNetworkOperateReq) (*RevokeNetworkOperateResp, error)
}

// UnimplementedAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (*UnimplementedAgentServer) Prepare(ctx context.Context, req *PrepareReq) (*PrepareResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (*UnimplementedAgentServer) Revoke(ctx context.Context, req *RevokeReq) (*RevokeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Revoke not implemented")
}
func (*UnimplementedAgentServer) NetworkOperate(ctx context.Context, req *NetworkOperateReq) (*NetworkOperateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NetworkOperate not implemented")
}
func (*UnimplementedAgentServer) RevokeNetworkOperate(ctx context.Context, req *RevokeNetworkOperateReq) (*RevokeNetworkOperateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeNetworkOperate not implemented")
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Agent/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Prepare(ctx, req.(*PrepareReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Agent/Revoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Revoke(ctx, req.(*RevokeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_NetworkOperate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkOperateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).NetworkOperate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Agent/NetworkOperate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).NetworkOperate(ctx, req.(*NetworkOperateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_RevokeNetworkOperate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeNetworkOperateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RevokeNetworkOperate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Agent/RevokeNetworkOperate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RevokeNetworkOperate(ctx, req.(*RevokeNetworkOperateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "prepare",
			Handler:    _Agent_Prepare_Handler,
		},
		{
			MethodName: "revoke",
			Handler:    _Agent_Revoke_Handler,
		},
		{
			MethodName: "networkOperate",
			Handler:    _Agent_NetworkOperate_Handler,
		},
		{
			MethodName: "revokeNetworkOperate",
			Handler:    _Agent_RevokeNetworkOperate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}
