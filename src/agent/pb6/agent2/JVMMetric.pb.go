// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent-v2/JVMMetric.proto

package agent2

import (
	agent "agent/agent/pb6/agent"
	common "agent/agent/pb6/common"
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

type JVMMetricCollection struct {
	Metrics              []*agent.JVMMetric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	ServiceInstanceId    int32              `protobuf:"varint,2,opt,name=serviceInstanceId,proto3" json:"serviceInstanceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *JVMMetricCollection) Reset()         { *m = JVMMetricCollection{} }
func (m *JVMMetricCollection) String() string { return proto.CompactTextString(m) }
func (*JVMMetricCollection) ProtoMessage()    {}
func (*JVMMetricCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5bd38fe036677f3, []int{0}
}

func (m *JVMMetricCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JVMMetricCollection.Unmarshal(m, b)
}
func (m *JVMMetricCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JVMMetricCollection.Marshal(b, m, deterministic)
}
func (m *JVMMetricCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JVMMetricCollection.Merge(m, src)
}
func (m *JVMMetricCollection) XXX_Size() int {
	return xxx_messageInfo_JVMMetricCollection.Size(m)
}
func (m *JVMMetricCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_JVMMetricCollection.DiscardUnknown(m)
}

var xxx_messageInfo_JVMMetricCollection proto.InternalMessageInfo

func (m *JVMMetricCollection) GetMetrics() []*agent.JVMMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *JVMMetricCollection) GetServiceInstanceId() int32 {
	if m != nil {
		return m.ServiceInstanceId
	}
	return 0
}

func init() {
	proto.RegisterType((*JVMMetricCollection)(nil), "JVMMetricCollection")
}

func init() { proto.RegisterFile("language-agent-v2/JVMMetric.proto", fileDescriptor_a5bd38fe036677f3) }

var fileDescriptor_a5bd38fe036677f3 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x6a, 0x02, 0x31,
	0x10, 0x86, 0xab, 0xa5, 0x95, 0xa6, 0x97, 0x76, 0x2d, 0x22, 0x7b, 0xb2, 0xd2, 0x83, 0x07, 0x9d,
	0x85, 0x15, 0xfa, 0x00, 0x15, 0x0a, 0x0a, 0x5b, 0x44, 0xc1, 0x42, 0x6f, 0x31, 0x0e, 0xdb, 0xb0,
	0x49, 0x66, 0x49, 0xd2, 0x15, 0x5f, 0xa9, 0x4f, 0x59, 0xdc, 0xac, 0xf6, 0xd0, 0x5e, 0x92, 0xe1,
	0x9b, 0xc9, 0x9f, 0xe1, 0x63, 0x8f, 0x8a, 0x9b, 0xfc, 0x8b, 0xe7, 0x38, 0xe1, 0x39, 0x1a, 0x3f,
	0xa9, 0xd2, 0x64, 0xb1, 0xc9, 0x32, 0xf4, 0x56, 0x0a, 0x28, 0x2d, 0x79, 0x8a, 0xbb, 0x82, 0xb4,
	0x26, 0x93, 0x84, 0xab, 0x81, 0x77, 0x0d, 0x5c, 0x6c, 0xb2, 0x40, 0x86, 0x92, 0x75, 0xcf, 0x2f,
	0x67, 0xa4, 0x14, 0x0a, 0x2f, 0xc9, 0x44, 0x4f, 0xac, 0xa3, 0x6b, 0xe6, 0xfa, 0xad, 0xc1, 0xe5,
	0xe8, 0x36, 0x65, 0x70, 0x1e, 0x5b, 0x9d, 0x5a, 0xd1, 0x98, 0xdd, 0x3b, 0xb4, 0x95, 0x14, 0x38,
	0x37, 0xce, 0x73, 0x23, 0x70, 0xbe, 0xeb, 0xb7, 0x07, 0xad, 0xd1, 0xd5, 0xea, 0x6f, 0x23, 0x7d,
	0x65, 0xbd, 0xdf, 0x0c, 0x2c, 0xc9, 0xfa, 0x75, 0x98, 0x89, 0xc6, 0xac, 0x23, 0xc2, 0xdf, 0xd1,
	0x03, 0xfc, 0xb3, 0x4e, 0x7c, 0x03, 0x33, 0xd2, 0x9a, 0x9b, 0x9d, 0x1b, 0x5e, 0xbc, 0x28, 0x36,
	0x25, 0x9b, 0x03, 0x2f, 0xb9, 0xf8, 0x44, 0x70, 0xc5, 0x61, 0xcf, 0x55, 0x21, 0xcd, 0x91, 0x68,
	0x30, 0xe8, 0xf7, 0x64, 0x0b, 0x38, 0x09, 0x82, 0x5a, 0x10, 0x54, 0xe9, 0xb2, 0xf5, 0xd1, 0xab,
	0xeb, 0x24, 0x9c, 0xe5, 0xf6, 0x39, 0x54, 0xe9, 0x77, 0x3b, 0x5e, 0x17, 0x87, 0xf7, 0x26, 0xe4,
	0x2d, 0x04, 0x2c, 0x8f, 0x76, 0x04, 0xa9, 0xed, 0x75, 0xed, 0x69, 0xfa, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x73, 0x25, 0xdc, 0x74, 0x73, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JVMMetricReportServiceClient is the client API for JVMMetricReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JVMMetricReportServiceClient interface {
	Collect(ctx context.Context, in *JVMMetricCollection, opts ...grpc.CallOption) (*common.Commands, error)
}

type jVMMetricReportServiceClient struct {
	cc *grpc.ClientConn
}

func NewJVMMetricReportServiceClient(cc *grpc.ClientConn) JVMMetricReportServiceClient {
	return &jVMMetricReportServiceClient{cc}
}

func (c *jVMMetricReportServiceClient) Collect(ctx context.Context, in *JVMMetricCollection, opts ...grpc.CallOption) (*common.Commands, error) {
	out := new(common.Commands)
	err := c.cc.Invoke(ctx, "/JVMMetricReportService/collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JVMMetricReportServiceServer is the server API for JVMMetricReportService service.
type JVMMetricReportServiceServer interface {
	Collect(context.Context, *JVMMetricCollection) (*common.Commands, error)
}

// UnimplementedJVMMetricReportServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJVMMetricReportServiceServer struct {
}

func (*UnimplementedJVMMetricReportServiceServer) Collect(ctx context.Context, req *JVMMetricCollection) (*common.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterJVMMetricReportServiceServer(s *grpc.Server, srv JVMMetricReportServiceServer) {
	s.RegisterService(&_JVMMetricReportService_serviceDesc, srv)
}

func _JVMMetricReportService_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JVMMetricCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JVMMetricReportServiceServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JVMMetricReportService/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JVMMetricReportServiceServer).Collect(ctx, req.(*JVMMetricCollection))
	}
	return interceptor(ctx, in, info, handler)
}

var _JVMMetricReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "JVMMetricReportService",
	HandlerType: (*JVMMetricReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "collect",
			Handler:    _JVMMetricReportService_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language-agent-v2/JVMMetric.proto",
}
