// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aexp.proto

package aexp

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

// request Hello
type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_80368e8b1937549a, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// response Hello AEXP
type Response struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_80368e8b1937549a, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "aexp.Request")
	proto.RegisterType((*Response)(nil), "aexp.Response")
}

func init() { proto.RegisterFile("aexp.proto", fileDescriptor_80368e8b1937549a) }

var fileDescriptor_80368e8b1937549a = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4c, 0xad, 0x28,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x64, 0xb9, 0xd8, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x15, 0x2e, 0x8e, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc,
	0xe2, 0x54, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0x74, 0x98, 0x12, 0x18, 0xd7,
	0xc8, 0x8c, 0x8b, 0xc5, 0xd1, 0x35, 0x22, 0x40, 0x48, 0x8f, 0x8b, 0x3b, 0x38, 0x35, 0x2f, 0x05,
	0x66, 0x20, 0xaf, 0x1e, 0xd8, 0x3a, 0x28, 0x57, 0x8a, 0x0f, 0xc6, 0x85, 0x98, 0xa7, 0xc4, 0x90,
	0xc4, 0x06, 0x76, 0x89, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xce, 0x73, 0x10, 0x20, 0x97, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AEXPClient is the client API for AEXP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AEXPClient interface {
	SendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type aEXPClient struct {
	cc *grpc.ClientConn
}

func NewAEXPClient(cc *grpc.ClientConn) AEXPClient {
	return &aEXPClient{cc}
}

func (c *aEXPClient) SendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/aexp.AEXP/SendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AEXPServer is the server API for AEXP service.
type AEXPServer interface {
	SendRequest(context.Context, *Request) (*Response, error)
}

// UnimplementedAEXPServer can be embedded to have forward compatible implementations.
type UnimplementedAEXPServer struct {
}

func (*UnimplementedAEXPServer) SendRequest(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRequest not implemented")
}

func RegisterAEXPServer(s *grpc.Server, srv AEXPServer) {
	s.RegisterService(&_AEXP_serviceDesc, srv)
}

func _AEXP_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AEXPServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aexp.AEXP/SendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AEXPServer).SendRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _AEXP_serviceDesc = grpc.ServiceDesc{
	ServiceName: "aexp.AEXP",
	HandlerType: (*AEXPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRequest",
			Handler:    _AEXP_SendRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aexp.proto",
}
