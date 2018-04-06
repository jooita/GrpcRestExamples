// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

/*
Package echopb is a generated protocol buffer package.

It is generated from these files:
	echo.proto

It has these top-level messages:
	StringMessage
*/
package echopb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type StringMessage struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StringMessage) Reset()                    { *m = StringMessage{} }
func (m *StringMessage) String() string            { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()               {}
func (*StringMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StringMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StringMessage)(nil), "echopb.StringMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EchoService service

type EchoServiceClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := grpc.Invoke(ctx, "/echopb.EchoService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	Echo(context.Context, *StringMessage) (*StringMessage, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echopb.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echopb.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0xb1, 0x0b, 0x92, 0xa4, 0x64, 0xd2, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b,
	0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xaa, 0x94, 0x54, 0xb9, 0x78, 0x83, 0x4b, 0x8a, 0x32, 0xf3, 0xd2,
	0x7d, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0xa3, 0x08, 0x2e, 0x6e, 0xd7, 0xe4, 0x8c,
	0xfc, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x4f, 0x2e, 0x16, 0x10, 0x57, 0x48, 0x54,
	0x0f, 0x62, 0x89, 0x1e, 0x8a, 0x19, 0x52, 0xd8, 0x85, 0x95, 0x84, 0x9b, 0x2e, 0x3f, 0x99, 0xcc,
	0xc4, 0xab, 0xc4, 0xa1, 0x5f, 0x66, 0xa8, 0x0f, 0x52, 0x61, 0xc5, 0xa8, 0x95, 0xc4, 0x06, 0x76,
	0x87, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x79, 0x09, 0xa0, 0xc6, 0xbb, 0x00, 0x00, 0x00,
}