// Code generated by protoc-gen-go. DO NOT EDIT.
// source: punchh/your_service_api.proto

package examplepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// EchoRequest represents a request.
type EchoRequest struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_681954a240202d92, []int{0}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// EchoResponse represents a response.
type EchoResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_681954a240202d92, []int{1}
}

func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "example.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "example.EchoResponse")
}

func init() { proto.RegisterFile("punchh/your_service_api.proto", fileDescriptor_681954a240202d92) }

var fileDescriptor_681954a240202d92 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x28, 0xcd, 0x4b,
	0xce, 0xc8, 0xd0, 0xaf, 0xcc, 0x2f, 0x2d, 0x8a, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x8d,
	0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xad, 0x48, 0xcc, 0x2d,
	0xc8, 0x49, 0x95, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f,
	0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x53, 0x52, 0xe6, 0xe2,
	0x76, 0x4d, 0xce, 0xc8, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x2d,
	0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0x54, 0xb8,
	0x78, 0x20, 0x8a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xb1, 0xab, 0x32, 0x4a, 0xe2, 0xe2, 0x8b,
	0xcc, 0x2f, 0x2d, 0x0a, 0x86, 0x38, 0xc5, 0x31, 0xc0, 0x53, 0x28, 0x80, 0x8b, 0x05, 0xa4, 0x4f,
	0x48, 0x44, 0x0f, 0xea, 0x18, 0x3d, 0x24, 0xbb, 0xa4, 0x44, 0xd1, 0x44, 0x21, 0x86, 0x2b, 0x49,
	0x37, 0x5d, 0x7e, 0x32, 0x99, 0x49, 0x54, 0x49, 0x40, 0xbf, 0xcc, 0x50, 0x1f, 0xaa, 0x42, 0x3f,
	0x35, 0x39, 0x23, 0xdf, 0x8a, 0x51, 0xcb, 0x89, 0x3b, 0x8a, 0x13, 0x2a, 0x54, 0x90, 0x94, 0xc4,
	0x06, 0xf6, 0x82, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x83, 0x5d, 0xc0, 0x0a, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// YourServiceAPIClient is the client API for YourServiceAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type YourServiceAPIClient interface {
	// Echo sends an echo request.
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type yourServiceAPIClient struct {
	cc *grpc.ClientConn
}

func NewYourServiceAPIClient(cc *grpc.ClientConn) YourServiceAPIClient {
	return &yourServiceAPIClient{cc}
}

func (c *yourServiceAPIClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/example.YourServiceAPI/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YourServiceAPIServer is the server API for YourServiceAPI service.
type YourServiceAPIServer interface {
	// Echo sends an echo request.
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

// UnimplementedYourServiceAPIServer can be embedded to have forward compatible implementations.
type UnimplementedYourServiceAPIServer struct {
}

func (*UnimplementedYourServiceAPIServer) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterYourServiceAPIServer(s *grpc.Server, srv YourServiceAPIServer) {
	s.RegisterService(&_YourServiceAPI_serviceDesc, srv)
}

func _YourServiceAPI_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YourServiceAPIServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.YourServiceAPI/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YourServiceAPIServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _YourServiceAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.YourServiceAPI",
	HandlerType: (*YourServiceAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _YourServiceAPI_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "punchh/your_service_api.proto",
}
