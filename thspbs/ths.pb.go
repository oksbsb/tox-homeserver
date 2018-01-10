// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ths.proto

/*
Package thspbs is a generated protocol buffer package.

It is generated from these files:
	ths.proto

It has these top-level messages:
	Event
	BaseInfo
	FriendInfo
	GroupInfo
	MemberInfo
	EmptyReq
	HelloReq
	HelloResp
*/
package thspbs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	Id    int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Mid   int64    `protobuf:"varint,2,opt,name=mid" json:"mid,omitempty"`
	Name  string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Args  []string `protobuf:"bytes,4,rep,name=args" json:"args,omitempty"`
	Margs []string `protobuf:"bytes,5,rep,name=margs" json:"margs,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Event) GetMargs() []string {
	if m != nil {
		return m.Margs
	}
	return nil
}

type BaseInfo struct {
	Id         string                 `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name       string                 `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Stmsg      string                 `protobuf:"bytes,3,opt,name=stmsg" json:"stmsg,omitempty"`
	Status     uint32                 `protobuf:"varint,4,opt,name=status" json:"status,omitempty"`
	Friends    map[uint32]*FriendInfo `protobuf:"bytes,5,rep,name=friends" json:"friends,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Groups     map[uint32]*GroupInfo  `protobuf:"bytes,6,rep,name=groups" json:"groups,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ConnStatus int32                  `protobuf:"varint,7,opt,name=connStatus" json:"connStatus,omitempty"`
}

func (m *BaseInfo) Reset()                    { *m = BaseInfo{} }
func (m *BaseInfo) String() string            { return proto.CompactTextString(m) }
func (*BaseInfo) ProtoMessage()               {}
func (*BaseInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BaseInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BaseInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BaseInfo) GetStmsg() string {
	if m != nil {
		return m.Stmsg
	}
	return ""
}

func (m *BaseInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *BaseInfo) GetFriends() map[uint32]*FriendInfo {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *BaseInfo) GetGroups() map[uint32]*GroupInfo {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *BaseInfo) GetConnStatus() int32 {
	if m != nil {
		return m.ConnStatus
	}
	return 0
}

type FriendInfo struct {
	Fnum       uint32 `protobuf:"varint,1,opt,name=fnum" json:"fnum,omitempty"`
	Status     uint32 `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	Pubkey     string `protobuf:"bytes,3,opt,name=pubkey" json:"pubkey,omitempty"`
	Name       string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Stmsg      string `protobuf:"bytes,5,opt,name=stmsg" json:"stmsg,omitempty"`
	Avatar     string `protobuf:"bytes,6,opt,name=avatar" json:"avatar,omitempty"`
	Seen       uint64 `protobuf:"varint,7,opt,name=seen" json:"seen,omitempty"`
	ConnStatus int32  `protobuf:"varint,8,opt,name=connStatus" json:"connStatus,omitempty"`
}

func (m *FriendInfo) Reset()                    { *m = FriendInfo{} }
func (m *FriendInfo) String() string            { return proto.CompactTextString(m) }
func (*FriendInfo) ProtoMessage()               {}
func (*FriendInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FriendInfo) GetFnum() uint32 {
	if m != nil {
		return m.Fnum
	}
	return 0
}

func (m *FriendInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *FriendInfo) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *FriendInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FriendInfo) GetStmsg() string {
	if m != nil {
		return m.Stmsg
	}
	return ""
}

func (m *FriendInfo) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *FriendInfo) GetSeen() uint64 {
	if m != nil {
		return m.Seen
	}
	return 0
}

func (m *FriendInfo) GetConnStatus() int32 {
	if m != nil {
		return m.ConnStatus
	}
	return 0
}

type GroupInfo struct {
	Gnum    uint32                 `protobuf:"varint,1,opt,name=gnum" json:"gnum,omitempty"`
	Mtype   uint32                 `protobuf:"varint,2,opt,name=mtype" json:"mtype,omitempty"`
	GroupId string                 `protobuf:"bytes,3,opt,name=groupId" json:"groupId,omitempty"`
	Title   string                 `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Stmsg   string                 `protobuf:"bytes,5,opt,name=stmsg" json:"stmsg,omitempty"`
	Ours    bool                   `protobuf:"varint,6,opt,name=ours" json:"ours,omitempty"`
	Members map[uint32]*MemberInfo `protobuf:"bytes,7,rep,name=members" json:"members,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GroupInfo) Reset()                    { *m = GroupInfo{} }
func (m *GroupInfo) String() string            { return proto.CompactTextString(m) }
func (*GroupInfo) ProtoMessage()               {}
func (*GroupInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GroupInfo) GetGnum() uint32 {
	if m != nil {
		return m.Gnum
	}
	return 0
}

func (m *GroupInfo) GetMtype() uint32 {
	if m != nil {
		return m.Mtype
	}
	return 0
}

func (m *GroupInfo) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *GroupInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GroupInfo) GetStmsg() string {
	if m != nil {
		return m.Stmsg
	}
	return ""
}

func (m *GroupInfo) GetOurs() bool {
	if m != nil {
		return m.Ours
	}
	return false
}

func (m *GroupInfo) GetMembers() map[uint32]*MemberInfo {
	if m != nil {
		return m.Members
	}
	return nil
}

type MemberInfo struct {
	Pnum   uint32 `protobuf:"varint,1,opt,name=pnum" json:"pnum,omitempty"`
	Pubkey string `protobuf:"bytes,2,opt,name=pubkey" json:"pubkey,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *MemberInfo) Reset()                    { *m = MemberInfo{} }
func (m *MemberInfo) String() string            { return proto.CompactTextString(m) }
func (*MemberInfo) ProtoMessage()               {}
func (*MemberInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MemberInfo) GetPnum() uint32 {
	if m != nil {
		return m.Pnum
	}
	return 0
}

func (m *MemberInfo) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *MemberInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type EmptyReq struct {
}

func (m *EmptyReq) Reset()                    { *m = EmptyReq{} }
func (m *EmptyReq) String() string            { return proto.CompactTextString(m) }
func (*EmptyReq) ProtoMessage()               {}
func (*EmptyReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type HelloReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *HelloReq) Reset()                    { *m = HelloReq{} }
func (m *HelloReq) String() string            { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()               {}
func (*HelloReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *HelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type HelloResp struct {
	Code   int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Status int64 `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
}

func (m *HelloResp) Reset()                    { *m = HelloResp{} }
func (m *HelloResp) String() string            { return proto.CompactTextString(m) }
func (*HelloResp) ProtoMessage()               {}
func (*HelloResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *HelloResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HelloResp) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "thspbs.Event")
	proto.RegisterType((*BaseInfo)(nil), "thspbs.BaseInfo")
	proto.RegisterType((*FriendInfo)(nil), "thspbs.FriendInfo")
	proto.RegisterType((*GroupInfo)(nil), "thspbs.GroupInfo")
	proto.RegisterType((*MemberInfo)(nil), "thspbs.MemberInfo")
	proto.RegisterType((*EmptyReq)(nil), "thspbs.EmptyReq")
	proto.RegisterType((*HelloReq)(nil), "thspbs.HelloReq")
	proto.RegisterType((*HelloResp)(nil), "thspbs.HelloResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Toxhs service

type ToxhsClient interface {
	PollCallback(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (Toxhs_PollCallbackClient, error)
	GetBaseInfo(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BaseInfo, error)
	RmtCall(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	Ping(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyReq, error)
}

type toxhsClient struct {
	cc *grpc.ClientConn
}

func NewToxhsClient(cc *grpc.ClientConn) ToxhsClient {
	return &toxhsClient{cc}
}

func (c *toxhsClient) PollCallback(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (Toxhs_PollCallbackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Toxhs_serviceDesc.Streams[0], c.cc, "/thspbs.Toxhs/PollCallback", opts...)
	if err != nil {
		return nil, err
	}
	x := &toxhsPollCallbackClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Toxhs_PollCallbackClient interface {
	Recv() (*EmptyReq, error)
	grpc.ClientStream
}

type toxhsPollCallbackClient struct {
	grpc.ClientStream
}

func (x *toxhsPollCallbackClient) Recv() (*EmptyReq, error) {
	m := new(EmptyReq)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *toxhsClient) GetBaseInfo(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BaseInfo, error) {
	out := new(BaseInfo)
	err := grpc.Invoke(ctx, "/thspbs.Toxhs/GetBaseInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toxhsClient) RmtCall(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := grpc.Invoke(ctx, "/thspbs.Toxhs/RmtCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toxhsClient) Ping(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyReq, error) {
	out := new(EmptyReq)
	err := grpc.Invoke(ctx, "/thspbs.Toxhs/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Toxhs service

type ToxhsServer interface {
	PollCallback(*EmptyReq, Toxhs_PollCallbackServer) error
	GetBaseInfo(context.Context, *EmptyReq) (*BaseInfo, error)
	RmtCall(context.Context, *Event) (*Event, error)
	Ping(context.Context, *EmptyReq) (*EmptyReq, error)
}

func RegisterToxhsServer(s *grpc.Server, srv ToxhsServer) {
	s.RegisterService(&_Toxhs_serviceDesc, srv)
}

func _Toxhs_PollCallback_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ToxhsServer).PollCallback(m, &toxhsPollCallbackServer{stream})
}

type Toxhs_PollCallbackServer interface {
	Send(*EmptyReq) error
	grpc.ServerStream
}

type toxhsPollCallbackServer struct {
	grpc.ServerStream
}

func (x *toxhsPollCallbackServer) Send(m *EmptyReq) error {
	return x.ServerStream.SendMsg(m)
}

func _Toxhs_GetBaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToxhsServer).GetBaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thspbs.Toxhs/GetBaseInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToxhsServer).GetBaseInfo(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Toxhs_RmtCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToxhsServer).RmtCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thspbs.Toxhs/RmtCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToxhsServer).RmtCall(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Toxhs_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToxhsServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thspbs.Toxhs/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToxhsServer).Ping(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Toxhs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thspbs.Toxhs",
	HandlerType: (*ToxhsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBaseInfo",
			Handler:    _Toxhs_GetBaseInfo_Handler,
		},
		{
			MethodName: "RmtCall",
			Handler:    _Toxhs_RmtCall_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Toxhs_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PollCallback",
			Handler:       _Toxhs_PollCallback_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ths.proto",
}

// Client API for Greeter service

type GreeterClient interface {
	SayHello(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*HelloReq, error)
	// 测试带参数的hello
	SayHellox(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloReq, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*HelloReq, error) {
	out := new(HelloReq)
	err := grpc.Invoke(ctx, "/thspbs.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHellox(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloReq, error) {
	out := new(HelloReq)
	err := grpc.Invoke(ctx, "/thspbs.Greeter/SayHellox", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SayHello(context.Context, *EmptyReq) (*HelloReq, error)
	// 测试带参数的hello
	SayHellox(context.Context, *HelloReq) (*HelloReq, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thspbs.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHellox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHellox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thspbs.Greeter/SayHellox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHellox(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thspbs.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayHellox",
			Handler:    _Greeter_SayHellox_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ths.proto",
}

func init() { proto.RegisterFile("ths.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xef, 0x4e, 0xd4, 0x40,
	0x10, 0x4f, 0x7b, 0xd7, 0x7f, 0x73, 0x60, 0x70, 0x43, 0x48, 0x73, 0x51, 0x72, 0xe9, 0x17, 0x2e,
	0x7e, 0xa8, 0x78, 0x9a, 0x40, 0xfc, 0xa8, 0x41, 0x24, 0x41, 0x43, 0x16, 0x5f, 0xa0, 0x47, 0x97,
	0x52, 0xe9, 0x3f, 0xba, 0x5b, 0xc2, 0x3d, 0x82, 0x8f, 0xe0, 0x03, 0xe9, 0x73, 0x99, 0xd9, 0xdd,
	0x96, 0xd2, 0x43, 0xa2, 0xdf, 0x66, 0x66, 0x67, 0xe6, 0xb7, 0xf3, 0xfb, 0xcd, 0x2e, 0x78, 0xe2,
	0x8a, 0x87, 0x55, 0x5d, 0x8a, 0x92, 0xd8, 0xe2, 0x8a, 0x57, 0x4b, 0x1e, 0x24, 0x60, 0x1d, 0xdd,
	0xb2, 0x42, 0x90, 0x67, 0x60, 0xa6, 0xb1, 0x6f, 0xcc, 0x8c, 0xb9, 0x45, 0xcd, 0x34, 0x26, 0x5b,
	0x30, 0xca, 0xd3, 0xd8, 0x37, 0x67, 0xc6, 0x7c, 0x44, 0xd1, 0x24, 0x04, 0xc6, 0x45, 0x94, 0x33,
	0x7f, 0x34, 0x33, 0xe6, 0x1e, 0x95, 0x36, 0xc6, 0xa2, 0x3a, 0xe1, 0xfe, 0x78, 0x36, 0xc2, 0x18,
	0xda, 0x64, 0x1b, 0xac, 0x5c, 0x06, 0x2d, 0x19, 0x54, 0x4e, 0xf0, 0x63, 0x04, 0xee, 0x87, 0x88,
	0xb3, 0x93, 0xe2, 0xb2, 0xec, 0x81, 0x79, 0x12, 0xac, 0x6d, 0x6d, 0xf6, 0x5a, 0x6f, 0x83, 0xc5,
	0x45, 0xce, 0x13, 0x8d, 0xa7, 0x1c, 0xb2, 0x03, 0x36, 0x17, 0x91, 0x68, 0x10, 0xd2, 0x98, 0x6f,
	0x52, 0xed, 0x91, 0x03, 0x70, 0x2e, 0xeb, 0x94, 0x15, 0xb1, 0x82, 0x9d, 0x2c, 0x5e, 0x86, 0x6a,
	0xc2, 0xb0, 0x05, 0x0d, 0x3f, 0xa9, 0xf3, 0xa3, 0x42, 0xd4, 0x2b, 0xda, 0x66, 0x93, 0x77, 0x60,
	0x27, 0x75, 0xd9, 0x54, 0xdc, 0xb7, 0x65, 0xdd, 0x8b, 0xb5, 0xba, 0x63, 0x79, 0xac, 0xca, 0x74,
	0x2e, 0xd9, 0x05, 0xb8, 0x28, 0x8b, 0xe2, 0x5c, 0x5d, 0xc5, 0x91, 0xac, 0xf5, 0x22, 0xd3, 0xaf,
	0xb0, 0xd1, 0x87, 0x43, 0x36, 0xaf, 0xd9, 0x4a, 0x4e, 0xbc, 0x49, 0xd1, 0x24, 0x73, 0xb0, 0x6e,
	0xa3, 0xac, 0x51, 0x33, 0x4f, 0x16, 0xa4, 0x85, 0x55, 0x65, 0x08, 0x4c, 0x55, 0xc2, 0x7b, 0xf3,
	0xd0, 0x98, 0x9e, 0xc2, 0xa4, 0x77, 0x8d, 0x47, 0xda, 0xed, 0x3d, 0x6c, 0xf7, 0xbc, 0x6d, 0x27,
	0xab, 0x06, 0xdd, 0x82, 0xdf, 0x06, 0xc0, 0x3d, 0x0e, 0xb2, 0x7f, 0x59, 0x34, 0xb9, 0x6e, 0x27,
	0xed, 0x1e, 0xcf, 0xe6, 0x03, 0x9e, 0x77, 0xc0, 0xae, 0x9a, 0x25, 0x82, 0x2b, 0x59, 0xb4, 0xd7,
	0x29, 0x38, 0x7e, 0x4c, 0x41, 0x6b, 0xa0, 0x60, 0x74, 0x1b, 0x89, 0xa8, 0xf6, 0x6d, 0xd5, 0x41,
	0x79, 0xd8, 0x81, 0x33, 0x56, 0x48, 0x32, 0xc7, 0x54, 0xda, 0x03, 0x9a, 0xdd, 0x21, 0xcd, 0xc1,
	0x4f, 0x13, 0xbc, 0x6e, 0x42, 0xec, 0x90, 0xf4, 0xe6, 0x40, 0x5b, 0x2e, 0xa3, 0x58, 0x55, 0x4c,
	0x8f, 0xa1, 0x1c, 0xe2, 0x83, 0x23, 0x85, 0x3c, 0x89, 0xf5, 0x18, 0xad, 0x8b, 0xf9, 0x22, 0x15,
	0x59, 0x3b, 0x88, 0x72, 0xfe, 0x32, 0x09, 0x81, 0x71, 0xd9, 0xd4, 0x5c, 0xce, 0xe1, 0x52, 0x69,
	0x93, 0x43, 0x70, 0x72, 0x96, 0x2f, 0x59, 0x8d, 0x5b, 0x81, 0xfb, 0xb4, 0xbb, 0xa6, 0x44, 0xf8,
	0x45, 0x25, 0xe8, 0x45, 0xd4, 0xe9, 0xb8, 0x32, 0xfd, 0x83, 0xff, 0x58, 0x19, 0x55, 0x36, 0x14,
	0xf9, 0x14, 0xe0, 0xfe, 0x00, 0xef, 0x5a, 0xf5, 0xb8, 0xa9, 0xb4, 0xc6, 0x5a, 0x4b, 0xf3, 0x51,
	0x2d, 0x7b, 0x0f, 0x3d, 0x00, 0x70, 0x8f, 0xf2, 0x4a, 0xac, 0x28, 0xbb, 0x09, 0xf6, 0xc1, 0xfd,
	0xcc, 0xb2, 0xac, 0xa4, 0xec, 0xa6, 0xcb, 0x35, 0x7a, 0xba, 0xe3, 0xd7, 0xc1, 0x13, 0xdd, 0x14,
	0xcd, 0xe0, 0x00, 0x3c, 0x5d, 0xc1, 0x2b, 0x2c, 0xb9, 0x28, 0x63, 0xa6, 0xff, 0x1a, 0x69, 0x0f,
	0xd6, 0x6d, 0xd4, 0xae, 0xdb, 0xe2, 0x97, 0x01, 0xd6, 0xb7, 0xf2, 0xee, 0x0a, 0xdf, 0xe9, 0xc6,
	0x59, 0x99, 0x65, 0x1f, 0xa3, 0x2c, 0x5b, 0x46, 0x17, 0xd7, 0x64, 0xab, 0x9d, 0xbe, 0xbd, 0xd6,
	0x74, 0x2d, 0xb2, 0x6f, 0x90, 0x37, 0x30, 0x39, 0x66, 0xa2, 0xfb, 0x77, 0x9e, 0x28, 0xea, 0x72,
	0xf6, 0xc0, 0xa1, 0xb9, 0x40, 0x1c, 0xb2, 0xd9, 0xa5, 0xe3, 0x17, 0x39, 0x7d, 0xe8, 0x92, 0x57,
	0x30, 0x3e, 0x4b, 0x8b, 0xe4, 0x5f, 0x6e, 0xb2, 0xf8, 0x0e, 0xce, 0x71, 0xcd, 0x98, 0x60, 0x35,
	0x09, 0xc1, 0x3d, 0x8f, 0x56, 0x92, 0x8e, 0xa7, 0x4a, 0x3b, 0x86, 0x5f, 0x83, 0xd7, 0xe6, 0xdf,
	0x91, 0xb5, 0xe3, 0xf5, 0x82, 0xa5, 0x2d, 0x7f, 0xf8, 0xb7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xe8, 0x3b, 0x73, 0x70, 0xee, 0x05, 0x00, 0x00,
}
