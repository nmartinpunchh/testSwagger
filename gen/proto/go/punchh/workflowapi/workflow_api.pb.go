// Code generated by protoc-gen-go. DO NOT EDIT.
// source: punchh/workflowapi/workflow_api.proto

package workflowapipb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	workflow "github.com/nmartinpunchh/testSwagger/master/gen/proto/go/punchh/workflow"
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

// CreateWorkflowRequest represents a workflow requests
type CreateWorkflowRequest struct {
	Workflow             *workflow.Workflow `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateWorkflowRequest) Reset()         { *m = CreateWorkflowRequest{} }
func (m *CreateWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWorkflowRequest) ProtoMessage()    {}
func (*CreateWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7198b79946a0984, []int{0}
}

func (m *CreateWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWorkflowRequest.Unmarshal(m, b)
}
func (m *CreateWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWorkflowRequest.Marshal(b, m, deterministic)
}
func (m *CreateWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWorkflowRequest.Merge(m, src)
}
func (m *CreateWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWorkflowRequest.Size(m)
}
func (m *CreateWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWorkflowRequest proto.InternalMessageInfo

func (m *CreateWorkflowRequest) GetWorkflow() *workflow.Workflow {
	if m != nil {
		return m.Workflow
	}
	return nil
}

// EchoResponse represents a response.
type CreateWorkflowResponse struct {
	Workflow             *workflow.Workflow `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateWorkflowResponse) Reset()         { *m = CreateWorkflowResponse{} }
func (m *CreateWorkflowResponse) String() string { return proto.CompactTextString(m) }
func (*CreateWorkflowResponse) ProtoMessage()    {}
func (*CreateWorkflowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7198b79946a0984, []int{1}
}

func (m *CreateWorkflowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWorkflowResponse.Unmarshal(m, b)
}
func (m *CreateWorkflowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWorkflowResponse.Marshal(b, m, deterministic)
}
func (m *CreateWorkflowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWorkflowResponse.Merge(m, src)
}
func (m *CreateWorkflowResponse) XXX_Size() int {
	return xxx_messageInfo_CreateWorkflowResponse.Size(m)
}
func (m *CreateWorkflowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWorkflowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWorkflowResponse proto.InternalMessageInfo

func (m *CreateWorkflowResponse) GetWorkflow() *workflow.Workflow {
	if m != nil {
		return m.Workflow
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateWorkflowRequest)(nil), "workflowapi.CreateWorkflowRequest")
	proto.RegisterType((*CreateWorkflowResponse)(nil), "workflowapi.CreateWorkflowResponse")
}

func init() {
	proto.RegisterFile("punchh/workflowapi/workflow_api.proto", fileDescriptor_b7198b79946a0984)
}

var fileDescriptor_b7198b79946a0984 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x28, 0xcd, 0x4b,
	0xce, 0xc8, 0xd0, 0x2f, 0xcf, 0x2f, 0xca, 0x4e, 0xcb, 0xc9, 0x2f, 0x4f, 0x2c, 0xc8, 0x84, 0xb3,
	0xe3, 0x13, 0x0b, 0x32, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb8, 0x91, 0xe4, 0xa5, 0x64,
	0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x41, 0x6a, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b,
	0x32, 0xf3, 0xf3, 0x8a, 0x21, 0x4a, 0xa5, 0xe4, 0xd0, 0x4c, 0x84, 0x33, 0x20, 0xf2, 0x4a, 0xee,
	0x5c, 0xa2, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0xe1, 0x50, 0xf1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4,
	0xe2, 0x12, 0x21, 0x3d, 0x2e, 0x0e, 0x98, 0x52, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x21,
	0x3d, 0xb8, 0x5e, 0xb8, 0x62, 0xb8, 0x1a, 0x25, 0x0f, 0x2e, 0x31, 0x74, 0x83, 0x8a, 0x0b, 0xf2,
	0xf3, 0x8a, 0x53, 0x49, 0x35, 0xc9, 0xa8, 0x95, 0x91, 0x8b, 0x1b, 0x26, 0xec, 0x18, 0xe0, 0x29,
	0x54, 0xc6, 0xc5, 0x87, 0x6a, 0xb2, 0x90, 0x92, 0x1e, 0x52, 0x00, 0xe8, 0x61, 0x75, 0xbf, 0x94,
	0x32, 0x5e, 0x35, 0x10, 0xa7, 0x29, 0x49, 0x37, 0x5d, 0x7e, 0x32, 0x99, 0x49, 0x54, 0x49, 0x40,
	0xbf, 0xcc, 0x50, 0x3f, 0xb5, 0x22, 0x31, 0xb7, 0x20, 0x27, 0x55, 0x3f, 0x35, 0x39, 0x23, 0xdf,
	0x8a, 0x51, 0xcb, 0x89, 0x3f, 0x8a, 0x17, 0xc9, 0x88, 0x82, 0xa4, 0x24, 0x36, 0x70, 0x90, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x2b, 0xd9, 0x6d, 0xa6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkflowAPIClient is the client API for WorkflowAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkflowAPIClient interface {
	// CreateWorkflow creates a workflow.
	CreateWorkflow(ctx context.Context, in *CreateWorkflowRequest, opts ...grpc.CallOption) (*CreateWorkflowResponse, error)
}

type workflowAPIClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowAPIClient(cc *grpc.ClientConn) WorkflowAPIClient {
	return &workflowAPIClient{cc}
}

func (c *workflowAPIClient) CreateWorkflow(ctx context.Context, in *CreateWorkflowRequest, opts ...grpc.CallOption) (*CreateWorkflowResponse, error) {
	out := new(CreateWorkflowResponse)
	err := c.cc.Invoke(ctx, "/workflowapi.WorkflowAPI/CreateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowAPIServer is the server API for WorkflowAPI service.
type WorkflowAPIServer interface {
	// CreateWorkflow creates a workflow.
	CreateWorkflow(context.Context, *CreateWorkflowRequest) (*CreateWorkflowResponse, error)
}

// UnimplementedWorkflowAPIServer can be embedded to have forward compatible implementations.
type UnimplementedWorkflowAPIServer struct {
}

func (*UnimplementedWorkflowAPIServer) CreateWorkflow(ctx context.Context, req *CreateWorkflowRequest) (*CreateWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflow not implemented")
}

func RegisterWorkflowAPIServer(s *grpc.Server, srv WorkflowAPIServer) {
	s.RegisterService(&_WorkflowAPI_serviceDesc, srv)
}

func _WorkflowAPI_CreateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).CreateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowapi.WorkflowAPI/CreateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).CreateWorkflow(ctx, req.(*CreateWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "workflowapi.WorkflowAPI",
	HandlerType: (*WorkflowAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkflow",
			Handler:    _WorkflowAPI_CreateWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "punchh/workflowapi/workflow_api.proto",
}