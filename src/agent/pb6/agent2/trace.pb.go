// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent-v2/trace.proto

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

type SegmentObject struct {
	TraceSegmentId       *agent.UniqueId `protobuf:"bytes,1,opt,name=traceSegmentId,proto3" json:"traceSegmentId,omitempty"`
	Spans                []*SpanObjectV2 `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	ServiceId            int32           `protobuf:"varint,3,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	ServiceInstanceId    int32           `protobuf:"varint,4,opt,name=serviceInstanceId,proto3" json:"serviceInstanceId,omitempty"`
	IsSizeLimited        bool            `protobuf:"varint,5,opt,name=isSizeLimited,proto3" json:"isSizeLimited,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SegmentObject) Reset()         { *m = SegmentObject{} }
func (m *SegmentObject) String() string { return proto.CompactTextString(m) }
func (*SegmentObject) ProtoMessage()    {}
func (*SegmentObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_8124ab206744863a, []int{0}
}

func (m *SegmentObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentObject.Unmarshal(m, b)
}
func (m *SegmentObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentObject.Marshal(b, m, deterministic)
}
func (m *SegmentObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentObject.Merge(m, src)
}
func (m *SegmentObject) XXX_Size() int {
	return xxx_messageInfo_SegmentObject.Size(m)
}
func (m *SegmentObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentObject.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentObject proto.InternalMessageInfo

func (m *SegmentObject) GetTraceSegmentId() *agent.UniqueId {
	if m != nil {
		return m.TraceSegmentId
	}
	return nil
}

func (m *SegmentObject) GetSpans() []*SpanObjectV2 {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *SegmentObject) GetServiceId() int32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *SegmentObject) GetServiceInstanceId() int32 {
	if m != nil {
		return m.ServiceInstanceId
	}
	return 0
}

func (m *SegmentObject) GetIsSizeLimited() bool {
	if m != nil {
		return m.IsSizeLimited
	}
	return false
}

type SegmentReference struct {
	RefType                 agent.RefType   `protobuf:"varint,1,opt,name=refType,proto3,enum=RefType" json:"refType,omitempty"`
	ParentTraceSegmentId    *agent.UniqueId `protobuf:"bytes,2,opt,name=parentTraceSegmentId,proto3" json:"parentTraceSegmentId,omitempty"`
	ParentSpanId            int32           `protobuf:"varint,3,opt,name=parentSpanId,proto3" json:"parentSpanId,omitempty"`
	ParentServiceInstanceId int32           `protobuf:"varint,4,opt,name=parentServiceInstanceId,proto3" json:"parentServiceInstanceId,omitempty"`
	NetworkAddress          string          `protobuf:"bytes,5,opt,name=networkAddress,proto3" json:"networkAddress,omitempty"`
	NetworkAddressId        int32           `protobuf:"varint,6,opt,name=networkAddressId,proto3" json:"networkAddressId,omitempty"`
	EntryServiceInstanceId  int32           `protobuf:"varint,7,opt,name=entryServiceInstanceId,proto3" json:"entryServiceInstanceId,omitempty"`
	EntryEndpoint           string          `protobuf:"bytes,8,opt,name=entryEndpoint,proto3" json:"entryEndpoint,omitempty"`
	EntryEndpointId         int32           `protobuf:"varint,9,opt,name=entryEndpointId,proto3" json:"entryEndpointId,omitempty"`
	ParentEndpoint          string          `protobuf:"bytes,10,opt,name=parentEndpoint,proto3" json:"parentEndpoint,omitempty"`
	ParentEndpointId        int32           `protobuf:"varint,11,opt,name=parentEndpointId,proto3" json:"parentEndpointId,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}        `json:"-"`
	XXX_unrecognized        []byte          `json:"-"`
	XXX_sizecache           int32           `json:"-"`
}

func (m *SegmentReference) Reset()         { *m = SegmentReference{} }
func (m *SegmentReference) String() string { return proto.CompactTextString(m) }
func (*SegmentReference) ProtoMessage()    {}
func (*SegmentReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_8124ab206744863a, []int{1}
}

func (m *SegmentReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentReference.Unmarshal(m, b)
}
func (m *SegmentReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentReference.Marshal(b, m, deterministic)
}
func (m *SegmentReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentReference.Merge(m, src)
}
func (m *SegmentReference) XXX_Size() int {
	return xxx_messageInfo_SegmentReference.Size(m)
}
func (m *SegmentReference) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentReference.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentReference proto.InternalMessageInfo

func (m *SegmentReference) GetRefType() agent.RefType {
	if m != nil {
		return m.RefType
	}
	return agent.RefType_CrossProcess
}

func (m *SegmentReference) GetParentTraceSegmentId() *agent.UniqueId {
	if m != nil {
		return m.ParentTraceSegmentId
	}
	return nil
}

func (m *SegmentReference) GetParentSpanId() int32 {
	if m != nil {
		return m.ParentSpanId
	}
	return 0
}

func (m *SegmentReference) GetParentServiceInstanceId() int32 {
	if m != nil {
		return m.ParentServiceInstanceId
	}
	return 0
}

func (m *SegmentReference) GetNetworkAddress() string {
	if m != nil {
		return m.NetworkAddress
	}
	return ""
}

func (m *SegmentReference) GetNetworkAddressId() int32 {
	if m != nil {
		return m.NetworkAddressId
	}
	return 0
}

func (m *SegmentReference) GetEntryServiceInstanceId() int32 {
	if m != nil {
		return m.EntryServiceInstanceId
	}
	return 0
}

func (m *SegmentReference) GetEntryEndpoint() string {
	if m != nil {
		return m.EntryEndpoint
	}
	return ""
}

func (m *SegmentReference) GetEntryEndpointId() int32 {
	if m != nil {
		return m.EntryEndpointId
	}
	return 0
}

func (m *SegmentReference) GetParentEndpoint() string {
	if m != nil {
		return m.ParentEndpoint
	}
	return ""
}

func (m *SegmentReference) GetParentEndpointId() int32 {
	if m != nil {
		return m.ParentEndpointId
	}
	return 0
}

type SpanObjectV2 struct {
	SpanId               int32                        `protobuf:"varint,1,opt,name=spanId,proto3" json:"spanId,omitempty"`
	ParentSpanId         int32                        `protobuf:"varint,2,opt,name=parentSpanId,proto3" json:"parentSpanId,omitempty"`
	StartTime            int64                        `protobuf:"varint,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              int64                        `protobuf:"varint,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Refs                 []*SegmentReference          `protobuf:"bytes,5,rep,name=refs,proto3" json:"refs,omitempty"`
	OperationNameId      int32                        `protobuf:"varint,6,opt,name=operationNameId,proto3" json:"operationNameId,omitempty"`
	OperationName        string                       `protobuf:"bytes,7,opt,name=operationName,proto3" json:"operationName,omitempty"`
	PeerId               int32                        `protobuf:"varint,8,opt,name=peerId,proto3" json:"peerId,omitempty"`
	Peer                 string                       `protobuf:"bytes,9,opt,name=peer,proto3" json:"peer,omitempty"`
	SpanType             agent.SpanType               `protobuf:"varint,10,opt,name=spanType,proto3,enum=SpanType" json:"spanType,omitempty"`
	SpanLayer            agent.SpanLayer              `protobuf:"varint,11,opt,name=spanLayer,proto3,enum=SpanLayer" json:"spanLayer,omitempty"`
	ComponentId          int32                        `protobuf:"varint,12,opt,name=componentId,proto3" json:"componentId,omitempty"`
	Component            string                       `protobuf:"bytes,13,opt,name=component,proto3" json:"component,omitempty"`
	IsError              bool                         `protobuf:"varint,14,opt,name=isError,proto3" json:"isError,omitempty"`
	Tags                 []*common.KeyStringValuePair `protobuf:"bytes,15,rep,name=tags,proto3" json:"tags,omitempty"`
	Logs                 []*Log                       `protobuf:"bytes,16,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SpanObjectV2) Reset()         { *m = SpanObjectV2{} }
func (m *SpanObjectV2) String() string { return proto.CompactTextString(m) }
func (*SpanObjectV2) ProtoMessage()    {}
func (*SpanObjectV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_8124ab206744863a, []int{2}
}

func (m *SpanObjectV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanObjectV2.Unmarshal(m, b)
}
func (m *SpanObjectV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanObjectV2.Marshal(b, m, deterministic)
}
func (m *SpanObjectV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanObjectV2.Merge(m, src)
}
func (m *SpanObjectV2) XXX_Size() int {
	return xxx_messageInfo_SpanObjectV2.Size(m)
}
func (m *SpanObjectV2) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanObjectV2.DiscardUnknown(m)
}

var xxx_messageInfo_SpanObjectV2 proto.InternalMessageInfo

func (m *SpanObjectV2) GetSpanId() int32 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *SpanObjectV2) GetParentSpanId() int32 {
	if m != nil {
		return m.ParentSpanId
	}
	return 0
}

func (m *SpanObjectV2) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *SpanObjectV2) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *SpanObjectV2) GetRefs() []*SegmentReference {
	if m != nil {
		return m.Refs
	}
	return nil
}

func (m *SpanObjectV2) GetOperationNameId() int32 {
	if m != nil {
		return m.OperationNameId
	}
	return 0
}

func (m *SpanObjectV2) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *SpanObjectV2) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *SpanObjectV2) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *SpanObjectV2) GetSpanType() agent.SpanType {
	if m != nil {
		return m.SpanType
	}
	return agent.SpanType_Entry
}

func (m *SpanObjectV2) GetSpanLayer() agent.SpanLayer {
	if m != nil {
		return m.SpanLayer
	}
	return agent.SpanLayer_Unknown
}

func (m *SpanObjectV2) GetComponentId() int32 {
	if m != nil {
		return m.ComponentId
	}
	return 0
}

func (m *SpanObjectV2) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *SpanObjectV2) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

func (m *SpanObjectV2) GetTags() []*common.KeyStringValuePair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SpanObjectV2) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

type Log struct {
	Time                 int64                        `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Data                 []*common.KeyStringValuePair `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_8124ab206744863a, []int{3}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Log) GetData() []*common.KeyStringValuePair {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SegmentObject)(nil), "SegmentObject")
	proto.RegisterType((*SegmentReference)(nil), "SegmentReference")
	proto.RegisterType((*SpanObjectV2)(nil), "SpanObjectV2")
	proto.RegisterType((*Log)(nil), "Log")
}

func init() { proto.RegisterFile("language-agent-v2/trace.proto", fileDescriptor_8124ab206744863a) }

var fileDescriptor_8124ab206744863a = []byte{
	// 752 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x6f, 0xe3, 0x44,
	0x10, 0xc7, 0x71, 0xe2, 0x34, 0xf1, 0xa4, 0xc9, 0xe5, 0xf6, 0x50, 0xf1, 0x55, 0x20, 0x45, 0x81,
	0x83, 0xe8, 0x44, 0x5d, 0xe1, 0x93, 0x4e, 0xbc, 0xf0, 0xc0, 0xa1, 0x13, 0x8a, 0x88, 0x8e, 0x6a,
	0xd3, 0x16, 0x89, 0xb7, 0xad, 0x3d, 0x35, 0x26, 0xf6, 0xae, 0x59, 0x6f, 0x5b, 0x85, 0x57, 0xfe,
	0x04, 0xfe, 0x0b, 0xfe, 0x23, 0xfe, 0x9b, 0xd3, 0x8e, 0x9d, 0x1f, 0x4e, 0xda, 0x97, 0x64, 0xe7,
	0x33, 0xe3, 0xd1, 0xce, 0x7c, 0x77, 0x06, 0xbe, 0xc8, 0x84, 0x4c, 0xee, 0x44, 0x82, 0x67, 0x22,
	0x41, 0x69, 0xce, 0xee, 0xc3, 0x73, 0xa3, 0x45, 0x84, 0x41, 0xa1, 0x95, 0x51, 0xa7, 0x2f, 0x22,
	0x95, 0xe7, 0x4a, 0x9e, 0x57, 0x7f, 0x35, 0x7c, 0x59, 0x43, 0x0a, 0x3c, 0xdb, 0x75, 0x4d, 0xfe,
	0x77, 0x60, 0xb0, 0xc0, 0x24, 0x47, 0x69, 0x7e, 0xbd, 0xf9, 0x13, 0x23, 0xc3, 0xbe, 0x83, 0x21,
	0xc5, 0xd5, 0x74, 0x16, 0xfb, 0xce, 0xd8, 0x99, 0xf6, 0x43, 0x2f, 0xb8, 0x92, 0xe9, 0x5f, 0x77,
	0x38, 0x8b, 0xf9, 0x5e, 0x00, 0xfb, 0x12, 0x3a, 0x65, 0x21, 0x64, 0xe9, 0xb7, 0xc6, 0xed, 0x69,
	0x3f, 0x1c, 0x04, 0x8b, 0x42, 0xc8, 0x2a, 0xdd, 0x75, 0xc8, 0x2b, 0x1f, 0xfb, 0x1c, 0xbc, 0x12,
	0xf5, 0x7d, 0x1a, 0xe1, 0x2c, 0xf6, 0xdb, 0x63, 0x67, 0xda, 0xe1, 0x5b, 0xc0, 0xbe, 0x85, 0xe7,
	0x6b, 0x43, 0x96, 0x46, 0x48, 0x8a, 0x72, 0x29, 0xea, 0xd0, 0xc1, 0xbe, 0x82, 0x41, 0x5a, 0x2e,
	0xd2, 0xbf, 0x71, 0x9e, 0xe6, 0xa9, 0xc1, 0xd8, 0xef, 0x8c, 0x9d, 0x69, 0x8f, 0x37, 0xe1, 0xe4,
	0x1f, 0x17, 0x46, 0xf5, 0x25, 0x39, 0xde, 0xa2, 0x46, 0x19, 0x21, 0x9b, 0x40, 0x57, 0xe3, 0xed,
	0xe5, 0xaa, 0x40, 0xaa, 0x6b, 0x18, 0xf6, 0x02, 0x5e, 0xd9, 0x7c, 0xed, 0x60, 0x3f, 0xc0, 0xa7,
	0x85, 0xd0, 0x28, 0xcd, 0x65, 0xb3, 0x11, 0xad, 0xfd, 0x46, 0x3c, 0x1a, 0xc6, 0x26, 0x70, 0x5c,
	0x71, 0xdb, 0x86, 0x4d, 0xb1, 0x0d, 0xc6, 0xbe, 0x87, 0xcf, 0x6a, 0xfb, 0x89, 0xaa, 0x9f, 0x72,
	0xb3, 0xaf, 0x61, 0x28, 0xd1, 0x3c, 0x28, 0xbd, 0xfc, 0x31, 0x8e, 0x35, 0x96, 0x25, 0x15, 0xef,
	0xf1, 0x3d, 0xca, 0x5e, 0xc3, 0xa8, 0x49, 0x66, 0xb1, 0x7f, 0x44, 0xa9, 0x0f, 0x38, 0x7b, 0x0b,
	0x27, 0x28, 0x8d, 0x5e, 0x1d, 0x5e, 0xa6, 0x4b, 0x5f, 0x3c, 0xe1, 0xb5, 0x3a, 0x90, 0xe7, 0xbd,
	0x8c, 0x0b, 0x95, 0x4a, 0xe3, 0xf7, 0xe8, 0x2a, 0x4d, 0xc8, 0xa6, 0xf0, 0xac, 0x01, 0x66, 0xb1,
	0xef, 0x51, 0xda, 0x7d, 0x6c, 0x6b, 0xab, 0xca, 0xde, 0x24, 0x84, 0xaa, 0xb6, 0x26, 0xb5, 0xb5,
	0x35, 0xc9, 0x2c, 0xf6, 0xfb, 0x55, 0x6d, 0xfb, 0x7c, 0xf2, 0xaf, 0x0b, 0xc7, 0xbb, 0xef, 0x91,
	0x9d, 0xc0, 0x51, 0x59, 0x09, 0xe3, 0xd0, 0x27, 0xb5, 0x75, 0x20, 0x5b, 0xeb, 0x11, 0xd9, 0xec,
	0x23, 0x36, 0x42, 0x9b, 0xcb, 0x34, 0x47, 0xd2, 0xb5, 0xcd, 0xb7, 0x80, 0xf9, 0xd0, 0x45, 0x19,
	0x93, 0xcf, 0x25, 0xdf, 0xda, 0x64, 0xaf, 0xc0, 0xd5, 0x78, 0x6b, 0xa5, 0xb2, 0x03, 0xf2, 0x3c,
	0xd8, 0x7f, 0x96, 0x9c, 0xdc, 0xb6, 0x53, 0xaa, 0x40, 0x2d, 0x4c, 0xaa, 0xe4, 0x07, 0x91, 0xe3,
	0x46, 0xb2, 0x7d, 0x6c, 0x3b, 0xdf, 0x40, 0x24, 0x94, 0xc7, 0x9b, 0xd0, 0x96, 0x5a, 0x20, 0xea,
	0x59, 0x4c, 0xc2, 0x74, 0x78, 0x6d, 0x31, 0x06, 0xae, 0x3d, 0x91, 0x0c, 0x1e, 0xa7, 0x33, 0x7b,
	0x05, 0x3d, 0xdb, 0x08, 0x9a, 0x0c, 0xa0, 0xc9, 0xf0, 0x68, 0x8e, 0x69, 0x34, 0x36, 0x2e, 0x36,
	0x05, 0xcf, 0x9e, 0xe7, 0x62, 0x85, 0x9a, 0x7a, 0x3e, 0x0c, 0x81, 0xe2, 0x88, 0xf0, 0xad, 0x93,
	0x8d, 0xa1, 0x1f, 0xa9, 0xbc, 0x50, 0xb2, 0x1a, 0x9e, 0x63, 0xba, 0xc1, 0x2e, 0xb2, 0xdd, 0xdc,
	0x98, 0xfe, 0x80, 0xee, 0xb2, 0x05, 0xb6, 0x9b, 0x69, 0xf9, 0x5e, 0x6b, 0xa5, 0xfd, 0x21, 0x8d,
	0xf7, 0xda, 0x64, 0xdf, 0x80, 0x6b, 0x44, 0x52, 0xfa, 0xcf, 0xa8, 0x9b, 0x2f, 0x82, 0x5f, 0x70,
	0xb5, 0x30, 0x3a, 0x95, 0xc9, 0xb5, 0xc8, 0xee, 0xf0, 0x42, 0xa4, 0x9a, 0x53, 0x00, 0xf3, 0xc1,
	0xcd, 0x54, 0x52, 0xfa, 0x23, 0x0a, 0x74, 0x83, 0xb9, 0x4a, 0x38, 0x91, 0xc9, 0x3b, 0x68, 0xcf,
	0x55, 0x62, 0x1b, 0x61, 0xac, 0x5c, 0x0e, 0xc9, 0x45, 0x67, 0x9b, 0x3d, 0x16, 0x46, 0xd4, 0xcb,
	0xec, 0xf1, 0xec, 0x36, 0x20, 0xfc, 0x19, 0x5e, 0xee, 0x4e, 0x3e, 0xc7, 0x42, 0xe9, 0xf5, 0xc0,
	0xb2, 0xd7, 0xd0, 0x8d, 0x54, 0x96, 0xd9, 0x8d, 0x3a, 0x0a, 0xae, 0x8a, 0xd2, 0x68, 0x14, 0x79,
	0x1d, 0x79, 0xea, 0x05, 0x3f, 0xa9, 0x3c, 0x17, 0x32, 0x2e, 0x27, 0x9f, 0x4c, 0x9d, 0x77, 0x19,
	0xbc, 0x51, 0x3a, 0x09, 0x44, 0x21, 0xa2, 0x3f, 0x30, 0x28, 0x97, 0xab, 0x07, 0x91, 0x2d, 0x53,
	0x69, 0x49, 0x1e, 0xd4, 0xc3, 0x1a, 0xac, 0xf7, 0x7e, 0x40, 0x7b, 0x3f, 0xb8, 0x0f, 0x2f, 0x9c,
	0xdf, 0x4f, 0xe8, 0x7c, 0x5e, 0xfd, 0x16, 0x37, 0x6f, 0xab, 0x53, 0xf8, 0x5f, 0xeb, 0x74, 0xb1,
	0x5c, 0xfd, 0x56, 0x27, 0xf9, 0x50, 0x25, 0xb8, 0xb0, 0xfb, 0x3e, 0x52, 0xd9, 0xcd, 0x11, 0x6d,
	0xfe, 0x37, 0x1f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xba, 0x2d, 0x7a, 0x4a, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceSegmentReportServiceClient is the client API for TraceSegmentReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceSegmentReportServiceClient interface {
	Collect(ctx context.Context, opts ...grpc.CallOption) (TraceSegmentReportService_CollectClient, error)
}

type traceSegmentReportServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceSegmentReportServiceClient(cc *grpc.ClientConn) TraceSegmentReportServiceClient {
	return &traceSegmentReportServiceClient{cc}
}

func (c *traceSegmentReportServiceClient) Collect(ctx context.Context, opts ...grpc.CallOption) (TraceSegmentReportService_CollectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceSegmentReportService_serviceDesc.Streams[0], "/TraceSegmentReportService/collect", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceSegmentReportServiceCollectClient{stream}
	return x, nil
}

type TraceSegmentReportService_CollectClient interface {
	Send(*agent.UpstreamSegment) error
	CloseAndRecv() (*common.Commands, error)
	grpc.ClientStream
}

type traceSegmentReportServiceCollectClient struct {
	grpc.ClientStream
}

func (x *traceSegmentReportServiceCollectClient) Send(m *agent.UpstreamSegment) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceSegmentReportServiceCollectClient) CloseAndRecv() (*common.Commands, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(common.Commands)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceSegmentReportServiceServer is the server API for TraceSegmentReportService service.
type TraceSegmentReportServiceServer interface {
	Collect(TraceSegmentReportService_CollectServer) error
}

// UnimplementedTraceSegmentReportServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTraceSegmentReportServiceServer struct {
}

func (*UnimplementedTraceSegmentReportServiceServer) Collect(srv TraceSegmentReportService_CollectServer) error {
	return status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterTraceSegmentReportServiceServer(s *grpc.Server, srv TraceSegmentReportServiceServer) {
	s.RegisterService(&_TraceSegmentReportService_serviceDesc, srv)
}

func _TraceSegmentReportService_Collect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceSegmentReportServiceServer).Collect(&traceSegmentReportServiceCollectServer{stream})
}

type TraceSegmentReportService_CollectServer interface {
	SendAndClose(*common.Commands) error
	Recv() (*agent.UpstreamSegment, error)
	grpc.ServerStream
}

type traceSegmentReportServiceCollectServer struct {
	grpc.ServerStream
}

func (x *traceSegmentReportServiceCollectServer) SendAndClose(m *common.Commands) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceSegmentReportServiceCollectServer) Recv() (*agent.UpstreamSegment, error) {
	m := new(agent.UpstreamSegment)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TraceSegmentReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TraceSegmentReportService",
	HandlerType: (*TraceSegmentReportServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "collect",
			Handler:       _TraceSegmentReportService_Collect_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "language-agent-v2/trace.proto",
}
