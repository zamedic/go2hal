// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remote_command.proto

package remoteTelegramCommands

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

type RemoteCommandRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Group                int64    `protobuf:"varint,3,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteCommandRequest) Reset()         { *m = RemoteCommandRequest{} }
func (m *RemoteCommandRequest) String() string { return proto.CompactTextString(m) }
func (*RemoteCommandRequest) ProtoMessage()    {}
func (*RemoteCommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_command_388f4e5403ef75cf, []int{0}
}
func (m *RemoteCommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteCommandRequest.Unmarshal(m, b)
}
func (m *RemoteCommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteCommandRequest.Marshal(b, m, deterministic)
}
func (dst *RemoteCommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteCommandRequest.Merge(dst, src)
}
func (m *RemoteCommandRequest) XXX_Size() int {
	return xxx_messageInfo_RemoteCommandRequest.Size(m)
}
func (m *RemoteCommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteCommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteCommandRequest proto.InternalMessageInfo

func (m *RemoteCommandRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RemoteCommandRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RemoteCommandRequest) GetGroup() int64 {
	if m != nil {
		return m.Group
	}
	return 0
}

type RemoteRequest struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Chat                 string   `protobuf:"bytes,3,opt,name=chat,proto3" json:"chat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteRequest) Reset()         { *m = RemoteRequest{} }
func (m *RemoteRequest) String() string { return proto.CompactTextString(m) }
func (*RemoteRequest) ProtoMessage()    {}
func (*RemoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_command_388f4e5403ef75cf, []int{1}
}
func (m *RemoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteRequest.Unmarshal(m, b)
}
func (m *RemoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteRequest.Marshal(b, m, deterministic)
}
func (dst *RemoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteRequest.Merge(dst, src)
}
func (m *RemoteRequest) XXX_Size() int {
	return xxx_messageInfo_RemoteRequest.Size(m)
}
func (m *RemoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteRequest proto.InternalMessageInfo

func (m *RemoteRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *RemoteRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RemoteRequest) GetChat() string {
	if m != nil {
		return m.Chat
	}
	return ""
}

type Request struct {
	State                string   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Nextstate            string   `protobuf:"bytes,2,opt,name=nextstate,proto3" json:"nextstate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_command_388f4e5403ef75cf, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Request) GetNextstate() string {
	if m != nil {
		return m.Nextstate
	}
	return ""
}

type Response struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Fields               []string `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_command_388f4e5403ef75cf, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoteCommandRequest)(nil), "remoteTelegramCommands.RemoteCommandRequest")
	proto.RegisterType((*RemoteRequest)(nil), "remoteTelegramCommands.RemoteRequest")
	proto.RegisterType((*Request)(nil), "remoteTelegramCommands.Request")
	proto.RegisterType((*Response)(nil), "remoteTelegramCommands.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RemoteCommandClient is the client API for RemoteCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemoteCommandClient interface {
	RegisterCommand(ctx context.Context, in *RemoteCommandRequest, opts ...grpc.CallOption) (RemoteCommand_RegisterCommandClient, error)
	RegisterCommandLet(ctx context.Context, in *Request, opts ...grpc.CallOption) (RemoteCommand_RegisterCommandLetClient, error)
}

type remoteCommandClient struct {
	cc *grpc.ClientConn
}

func NewRemoteCommandClient(cc *grpc.ClientConn) RemoteCommandClient {
	return &remoteCommandClient{cc}
}

func (c *remoteCommandClient) RegisterCommand(ctx context.Context, in *RemoteCommandRequest, opts ...grpc.CallOption) (RemoteCommand_RegisterCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RemoteCommand_serviceDesc.Streams[0], "/remoteTelegramCommands.RemoteCommand/RegisterCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &remoteCommandRegisterCommandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RemoteCommand_RegisterCommandClient interface {
	Recv() (*RemoteRequest, error)
	grpc.ClientStream
}

type remoteCommandRegisterCommandClient struct {
	grpc.ClientStream
}

func (x *remoteCommandRegisterCommandClient) Recv() (*RemoteRequest, error) {
	m := new(RemoteRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *remoteCommandClient) RegisterCommandLet(ctx context.Context, in *Request, opts ...grpc.CallOption) (RemoteCommand_RegisterCommandLetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RemoteCommand_serviceDesc.Streams[1], "/remoteTelegramCommands.RemoteCommand/RegisterCommandLet", opts...)
	if err != nil {
		return nil, err
	}
	x := &remoteCommandRegisterCommandLetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RemoteCommand_RegisterCommandLetClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type remoteCommandRegisterCommandLetClient struct {
	grpc.ClientStream
}

func (x *remoteCommandRegisterCommandLetClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RemoteCommandServer is the server API for RemoteCommand service.
type RemoteCommandServer interface {
	RegisterCommand(*RemoteCommandRequest, RemoteCommand_RegisterCommandServer) error
	RegisterCommandLet(*Request, RemoteCommand_RegisterCommandLetServer) error
}

func RegisterRemoteCommandServer(s *grpc.Server, srv RemoteCommandServer) {
	s.RegisterService(&_RemoteCommand_serviceDesc, srv)
}

func _RemoteCommand_RegisterCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RemoteCommandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RemoteCommandServer).RegisterCommand(m, &remoteCommandRegisterCommandServer{stream})
}

type RemoteCommand_RegisterCommandServer interface {
	Send(*RemoteRequest) error
	grpc.ServerStream
}

type remoteCommandRegisterCommandServer struct {
	grpc.ServerStream
}

func (x *remoteCommandRegisterCommandServer) Send(m *RemoteRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _RemoteCommand_RegisterCommandLet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RemoteCommandServer).RegisterCommandLet(m, &remoteCommandRegisterCommandLetServer{stream})
}

type RemoteCommand_RegisterCommandLetServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type remoteCommandRegisterCommandLetServer struct {
	grpc.ServerStream
}

func (x *remoteCommandRegisterCommandLetServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _RemoteCommand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remoteTelegramCommands.RemoteCommand",
	HandlerType: (*RemoteCommandServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterCommand",
			Handler:       _RemoteCommand_RegisterCommand_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RegisterCommandLet",
			Handler:       _RemoteCommand_RegisterCommandLet_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "remote_command.proto",
}

func init() {
	proto.RegisterFile("remote_command.proto", fileDescriptor_remote_command_388f4e5403ef75cf)
}

var fileDescriptor_remote_command_388f4e5403ef75cf = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xad, 0x75, 0x9b, 0x7d, 0x22, 0xc2, 0xa3, 0x8c, 0x22, 0x82, 0xa5, 0x20, 0xf4, 0x20,
	0x45, 0xf4, 0xec, 0xc9, 0xab, 0x07, 0x0d, 0xde, 0x3c, 0x48, 0xd6, 0xbe, 0xd5, 0xca, 0xd2, 0xd4,
	0x24, 0x03, 0xff, 0x54, 0xff, 0x1c, 0x69, 0x7e, 0xe0, 0x90, 0x4d, 0xd8, 0xed, 0x7d, 0x7f, 0xf0,
	0x79, 0x69, 0x52, 0x48, 0x15, 0x09, 0x69, 0xe8, 0xad, 0x96, 0x42, 0xf0, 0xbe, 0xa9, 0x06, 0x25,
	0x8d, 0xc4, 0xb9, 0x73, 0x5f, 0x68, 0x45, 0xad, 0xe2, 0xe2, 0xc1, 0x85, 0xba, 0x58, 0x40, 0xca,
	0x6c, 0xe2, 0x1d, 0x46, 0x9f, 0x6b, 0xd2, 0x06, 0x11, 0x8e, 0x7a, 0x2e, 0x28, 0x8b, 0xf2, 0xa8,
	0x4c, 0x98, 0x9d, 0x31, 0x87, 0x93, 0x86, 0x74, 0xad, 0xba, 0xc1, 0x74, 0xb2, 0xcf, 0x0e, 0x6d,
	0xb4, 0x69, 0x61, 0x0a, 0x93, 0x56, 0xc9, 0xf5, 0x90, 0xc5, 0x79, 0x54, 0xc6, 0xcc, 0x89, 0xe2,
	0x19, 0x4e, 0xdd, 0x8e, 0x0d, 0xf8, 0x52, 0x49, 0x11, 0xe0, 0xe3, 0x8c, 0x19, 0xcc, 0x04, 0x69,
	0xcd, 0x5b, 0xf2, 0xe0, 0x20, 0xc7, 0x76, 0xfd, 0xce, 0x8d, 0x65, 0x26, 0xcc, 0xce, 0xc5, 0x3d,
	0xcc, 0x02, 0x2c, 0x85, 0x89, 0x36, 0xdc, 0x84, 0xa3, 0x3a, 0x81, 0x17, 0x90, 0xf4, 0xf4, 0x65,
	0x5c, 0xe2, 0x80, 0xbf, 0x46, 0xf1, 0x04, 0xc7, 0x8c, 0xf4, 0x20, 0x7b, 0x4d, 0x7b, 0x1e, 0x66,
	0x0e, 0xd3, 0x65, 0x47, 0xab, 0x46, 0x67, 0x71, 0x1e, 0x97, 0x09, 0xf3, 0xea, 0xf6, 0x3b, 0x0a,
	0x1f, 0xe9, 0x2f, 0x12, 0x3f, 0xe0, 0x8c, 0x51, 0xdb, 0x69, 0x43, 0x2a, 0x58, 0xd7, 0xd5, 0xf6,
	0x57, 0xa8, 0xb6, 0x3d, 0xc1, 0xf9, 0xd5, 0xff, 0x6d, 0x5f, 0x2b, 0x0e, 0x6e, 0x22, 0x7c, 0x05,
	0xfc, 0xb3, 0xeb, 0x91, 0x0c, 0x5e, 0xee, 0x06, 0xb8, 0x0d, 0xf9, 0xee, 0x82, 0xbb, 0x9c, 0x11,
	0xbe, 0x98, 0xda, 0x3f, 0xe8, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0x97, 0xfc, 0x7a, 0x2c, 0x59,
	0x02, 0x00, 0x00,
}
