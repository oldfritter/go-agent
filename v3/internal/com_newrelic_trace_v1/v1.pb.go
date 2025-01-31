// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// +build go1.9
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v3/internal/com_oldfritter_trace_v1/v1.proto

package com_oldfritter_trace_v1

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

type SpanBatch struct {
	Spans                []*Span  `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpanBatch) Reset()         { *m = SpanBatch{} }
func (m *SpanBatch) String() string { return proto.CompactTextString(m) }
func (*SpanBatch) ProtoMessage()    {}
func (*SpanBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_10a7bb7b83f0c5c3, []int{0}
}

func (m *SpanBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanBatch.Unmarshal(m, b)
}
func (m *SpanBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanBatch.Marshal(b, m, deterministic)
}
func (m *SpanBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanBatch.Merge(m, src)
}
func (m *SpanBatch) XXX_Size() int {
	return xxx_messageInfo_SpanBatch.Size(m)
}
func (m *SpanBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanBatch.DiscardUnknown(m)
}

var xxx_messageInfo_SpanBatch proto.InternalMessageInfo

func (m *SpanBatch) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

type Span struct {
	TraceId              string                     `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Intrinsics           map[string]*AttributeValue `protobuf:"bytes,2,rep,name=intrinsics,proto3" json:"intrinsics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UserAttributes       map[string]*AttributeValue `protobuf:"bytes,3,rep,name=user_attributes,json=userAttributes,proto3" json:"user_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AgentAttributes      map[string]*AttributeValue `protobuf:"bytes,4,rep,name=agent_attributes,json=agentAttributes,proto3" json:"agent_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_10a7bb7b83f0c5c3, []int{1}
}

func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func (m *Span) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Span) GetIntrinsics() map[string]*AttributeValue {
	if m != nil {
		return m.Intrinsics
	}
	return nil
}

func (m *Span) GetUserAttributes() map[string]*AttributeValue {
	if m != nil {
		return m.UserAttributes
	}
	return nil
}

func (m *Span) GetAgentAttributes() map[string]*AttributeValue {
	if m != nil {
		return m.AgentAttributes
	}
	return nil
}

type AttributeValue struct {
	// Types that are valid to be assigned to Value:
	//	*AttributeValue_StringValue
	//	*AttributeValue_BoolValue
	//	*AttributeValue_IntValue
	//	*AttributeValue_DoubleValue
	Value                isAttributeValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AttributeValue) Reset()         { *m = AttributeValue{} }
func (m *AttributeValue) String() string { return proto.CompactTextString(m) }
func (*AttributeValue) ProtoMessage()    {}
func (*AttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_10a7bb7b83f0c5c3, []int{2}
}

func (m *AttributeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeValue.Unmarshal(m, b)
}
func (m *AttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeValue.Marshal(b, m, deterministic)
}
func (m *AttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeValue.Merge(m, src)
}
func (m *AttributeValue) XXX_Size() int {
	return xxx_messageInfo_AttributeValue.Size(m)
}
func (m *AttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeValue proto.InternalMessageInfo

type isAttributeValue_Value interface {
	isAttributeValue_Value()
}

type AttributeValue_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type AttributeValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type AttributeValue_IntValue struct {
	IntValue int64 `protobuf:"varint,3,opt,name=int_value,json=intValue,proto3,oneof"`
}

type AttributeValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

func (*AttributeValue_StringValue) isAttributeValue_Value() {}

func (*AttributeValue_BoolValue) isAttributeValue_Value() {}

func (*AttributeValue_IntValue) isAttributeValue_Value() {}

func (*AttributeValue_DoubleValue) isAttributeValue_Value() {}

func (m *AttributeValue) GetValue() isAttributeValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *AttributeValue) GetStringValue() string {
	if x, ok := m.GetValue().(*AttributeValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *AttributeValue) GetBoolValue() bool {
	if x, ok := m.GetValue().(*AttributeValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *AttributeValue) GetIntValue() int64 {
	if x, ok := m.GetValue().(*AttributeValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *AttributeValue) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*AttributeValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AttributeValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AttributeValue_StringValue)(nil),
		(*AttributeValue_BoolValue)(nil),
		(*AttributeValue_IntValue)(nil),
		(*AttributeValue_DoubleValue)(nil),
	}
}

type RecordStatus struct {
	MessagesSeen         uint64   `protobuf:"varint,1,opt,name=messages_seen,json=messagesSeen,proto3" json:"messages_seen,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordStatus) Reset()         { *m = RecordStatus{} }
func (m *RecordStatus) String() string { return proto.CompactTextString(m) }
func (*RecordStatus) ProtoMessage()    {}
func (*RecordStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_10a7bb7b83f0c5c3, []int{3}
}

func (m *RecordStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordStatus.Unmarshal(m, b)
}
func (m *RecordStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordStatus.Marshal(b, m, deterministic)
}
func (m *RecordStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordStatus.Merge(m, src)
}
func (m *RecordStatus) XXX_Size() int {
	return xxx_messageInfo_RecordStatus.Size(m)
}
func (m *RecordStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RecordStatus proto.InternalMessageInfo

func (m *RecordStatus) GetMessagesSeen() uint64 {
	if m != nil {
		return m.MessagesSeen
	}
	return 0
}

func init() {
	proto.RegisterType((*SpanBatch)(nil), "com.oldfritter.trace.v1.SpanBatch")
	proto.RegisterType((*Span)(nil), "com.oldfritter.trace.v1.Span")
	proto.RegisterMapType((map[string]*AttributeValue)(nil), "com.oldfritter.trace.v1.Span.AgentAttributesEntry")
	proto.RegisterMapType((map[string]*AttributeValue)(nil), "com.oldfritter.trace.v1.Span.IntrinsicsEntry")
	proto.RegisterMapType((map[string]*AttributeValue)(nil), "com.oldfritter.trace.v1.Span.UserAttributesEntry")
	proto.RegisterType((*AttributeValue)(nil), "com.oldfritter.trace.v1.AttributeValue")
	proto.RegisterType((*RecordStatus)(nil), "com.oldfritter.trace.v1.RecordStatus")
}

func init() {
	proto.RegisterFile("v3/internal/com_oldfritter_trace_v1/v1.proto", fileDescriptor_10a7bb7b83f0c5c3)
}

var fileDescriptor_10a7bb7b83f0c5c3 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x61, 0x8b, 0x12, 0x41,
	0x18, 0xc7, 0x1d, 0xf5, 0x3a, 0x7d, 0xf4, 0xce, 0x63, 0x2a, 0x30, 0x23, 0x5a, 0x94, 0x60, 0x29,
	0xda, 0x3d, 0xf5, 0x4d, 0x14, 0x1c, 0x9d, 0x10, 0x28, 0xbd, 0x5b, 0x2b, 0xa2, 0xa0, 0x65, 0x5c,
	0x1f, 0xd6, 0x21, 0x9d, 0x95, 0x99, 0xd9, 0x8d, 0xfb, 0x3c, 0x7d, 0x96, 0xbe, 0x47, 0x1f, 0x25,
	0x76, 0x46, 0x3d, 0x3d, 0x3c, 0xe3, 0x5e, 0xdc, 0xbb, 0xdd, 0xe7, 0xf9, 0xff, 0x7f, 0xff, 0x79,
	0x06, 0x9e, 0x81, 0x97, 0x59, 0xdf, 0xe7, 0x42, 0xa3, 0x14, 0x6c, 0xee, 0x47, 0xc9, 0x22, 0x14,
	0xf8, 0x4b, 0xe2, 0x9c, 0x47, 0xa1, 0x96, 0x2c, 0xc2, 0x30, 0xeb, 0xfa, 0x59, 0xd7, 0x5b, 0xca,
	0x44, 0x27, 0xf4, 0x71, 0x94, 0x2c, 0xbc, 0x75, 0xdf, 0x33, 0x7d, 0x2f, 0xeb, 0xb6, 0x2f, 0xa0,
	0x3a, 0x5e, 0x32, 0x31, 0x60, 0x3a, 0x9a, 0xd1, 0x2e, 0x1c, 0xa9, 0x25, 0x13, 0xaa, 0x49, 0x9c,
	0x92, 0x5b, 0xeb, 0x3d, 0xf5, 0xf6, 0x7a, 0xbc, 0xdc, 0x10, 0x58, 0x65, 0xfb, 0x6f, 0x19, 0xca,
	0xf9, 0x3f, 0x7d, 0x02, 0x15, 0x1b, 0xca, 0xa7, 0x4d, 0xe2, 0x10, 0xb7, 0x1a, 0x1c, 0x9b, 0xff,
	0xd1, 0x94, 0x7e, 0x04, 0xe0, 0x42, 0x4b, 0x2e, 0x14, 0x8f, 0x54, 0xb3, 0x68, 0xd8, 0xaf, 0x0e,
	0xb0, 0xbd, 0xd1, 0x46, 0xfd, 0x41, 0x68, 0x79, 0x15, 0x6c, 0xd9, 0xe9, 0x57, 0x68, 0xa4, 0x0a,
	0x65, 0xc8, 0xb4, 0x96, 0x7c, 0x92, 0x6a, 0x54, 0xcd, 0x92, 0x21, 0xfa, 0x87, 0x88, 0x9f, 0x15,
	0xca, 0xcb, 0x8d, 0xc3, 0x52, 0x4f, 0xd3, 0x9d, 0x22, 0xfd, 0x0e, 0x67, 0x2c, 0x46, 0xa1, 0xb7,
	0xd1, 0x65, 0x83, 0x3e, 0x3f, 0x84, 0xbe, 0xcc, 0x3d, 0x37, 0xd9, 0x0d, 0xb6, 0x5b, 0x6d, 0x4d,
	0xa1, 0x71, 0x63, 0x2a, 0x7a, 0x06, 0xa5, 0x9f, 0x78, 0xb5, 0xba, 0xac, 0xfc, 0x93, 0xbe, 0x83,
	0xa3, 0x8c, 0xcd, 0x53, 0x6c, 0x16, 0x1d, 0xe2, 0xd6, 0x7a, 0x2f, 0x6e, 0x89, 0xdd, 0x60, 0xbf,
	0xe4, 0xe2, 0xc0, 0x7a, 0xde, 0x16, 0xdf, 0x90, 0xd6, 0x0c, 0x1e, 0xee, 0x99, 0xf4, 0x3e, 0x92,
	0x38, 0x3c, 0xda, 0x37, 0xf8, 0x3d, 0x44, 0xb5, 0x7f, 0x13, 0x38, 0xdd, 0xed, 0xd2, 0x0e, 0xd4,
	0x55, 0x7e, 0x99, 0x71, 0x68, 0xd1, 0x26, 0x6e, 0x58, 0x08, 0x6a, 0xb6, 0x6a, 0x45, 0xcf, 0x01,
	0x26, 0x49, 0x32, 0x0f, 0xaf, 0xd3, 0x2b, 0xc3, 0x42, 0x50, 0xcd, 0x6b, 0x56, 0xf0, 0x0c, 0xaa,
	0x5c, 0xe8, 0x55, 0xbf, 0xe4, 0x10, 0xb7, 0x34, 0x2c, 0x04, 0x15, 0x2e, 0xf4, 0x26, 0x64, 0x9a,
	0xa4, 0x93, 0x39, 0xae, 0x14, 0x65, 0x87, 0xb8, 0x24, 0x0f, 0xb1, 0x55, 0x23, 0x1a, 0x1c, 0xaf,
	0xa6, 0x6b, 0xf7, 0xa1, 0x1e, 0x60, 0x94, 0xc8, 0xe9, 0x58, 0x33, 0x9d, 0x2a, 0xda, 0x81, 0x93,
	0x05, 0x2a, 0xc5, 0x62, 0x54, 0xa1, 0x42, 0x14, 0xe6, 0x8c, 0xe5, 0xa0, 0xbe, 0x2e, 0x8e, 0x11,
	0x45, 0xef, 0x0f, 0x81, 0x93, 0x91, 0x88, 0x51, 0xe9, 0x31, 0xca, 0x8c, 0x47, 0x48, 0x3f, 0x01,
	0xac, 0x30, 0xf9, 0x52, 0x1d, 0xda, 0xc0, 0x56, 0xe7, 0x96, 0xe6, 0xf6, 0x31, 0xda, 0x05, 0x97,
	0x9c, 0x13, 0xfa, 0x03, 0x1a, 0xd7, 0x54, 0xbb, 0xeb, 0xce, 0x01, 0xb4, 0x51, 0xdc, 0x81, 0x3f,
	0x78, 0xff, 0xed, 0x22, 0xe6, 0x7a, 0x96, 0x4e, 0x72, 0x8b, 0xbf, 0xb6, 0xf8, 0x71, 0xf2, 0xda,
	0xec, 0x81, 0xff, 0xdf, 0x77, 0x6a, 0xf2, 0xc0, 0xbc, 0x52, 0xfd, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x38, 0x06, 0x7c, 0x3d, 0xd3, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IngestServiceClient is the client API for IngestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IngestServiceClient interface {
	// Accepts a stream of Span messages, and returns an irregular stream of
	// RecordStatus messages.
	RecordSpan(ctx context.Context, opts ...grpc.CallOption) (IngestService_RecordSpanClient, error)
	// Accepts a stream of SpanBatch messages, and returns an irregular
	// stream of RecordStatus messages. This endpoint can be used to improve
	// throughput when Span messages are small
	RecordSpanBatch(ctx context.Context, opts ...grpc.CallOption) (IngestService_RecordSpanBatchClient, error)
}

type ingestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIngestServiceClient(cc grpc.ClientConnInterface) IngestServiceClient {
	return &ingestServiceClient{cc}
}

func (c *ingestServiceClient) RecordSpan(ctx context.Context, opts ...grpc.CallOption) (IngestService_RecordSpanClient, error) {
	stream, err := c.cc.NewStream(ctx, &_IngestService_serviceDesc.Streams[0], "/com.oldfritter.trace.v1.IngestService/RecordSpan", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingestServiceRecordSpanClient{stream}
	return x, nil
}

type IngestService_RecordSpanClient interface {
	Send(*Span) error
	Recv() (*RecordStatus, error)
	grpc.ClientStream
}

type ingestServiceRecordSpanClient struct {
	grpc.ClientStream
}

func (x *ingestServiceRecordSpanClient) Send(m *Span) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ingestServiceRecordSpanClient) Recv() (*RecordStatus, error) {
	m := new(RecordStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ingestServiceClient) RecordSpanBatch(ctx context.Context, opts ...grpc.CallOption) (IngestService_RecordSpanBatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_IngestService_serviceDesc.Streams[1], "/com.oldfritter.trace.v1.IngestService/RecordSpanBatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingestServiceRecordSpanBatchClient{stream}
	return x, nil
}

type IngestService_RecordSpanBatchClient interface {
	Send(*SpanBatch) error
	Recv() (*RecordStatus, error)
	grpc.ClientStream
}

type ingestServiceRecordSpanBatchClient struct {
	grpc.ClientStream
}

func (x *ingestServiceRecordSpanBatchClient) Send(m *SpanBatch) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ingestServiceRecordSpanBatchClient) Recv() (*RecordStatus, error) {
	m := new(RecordStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IngestServiceServer is the server API for IngestService service.
type IngestServiceServer interface {
	// Accepts a stream of Span messages, and returns an irregular stream of
	// RecordStatus messages.
	RecordSpan(IngestService_RecordSpanServer) error
	// Accepts a stream of SpanBatch messages, and returns an irregular
	// stream of RecordStatus messages. This endpoint can be used to improve
	// throughput when Span messages are small
	RecordSpanBatch(IngestService_RecordSpanBatchServer) error
}

// UnimplementedIngestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIngestServiceServer struct {
}

func (*UnimplementedIngestServiceServer) RecordSpan(srv IngestService_RecordSpanServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordSpan not implemented")
}
func (*UnimplementedIngestServiceServer) RecordSpanBatch(srv IngestService_RecordSpanBatchServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordSpanBatch not implemented")
}

func RegisterIngestServiceServer(s *grpc.Server, srv IngestServiceServer) {
	s.RegisterService(&_IngestService_serviceDesc, srv)
}

func _IngestService_RecordSpan_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngestServiceServer).RecordSpan(&ingestServiceRecordSpanServer{stream})
}

type IngestService_RecordSpanServer interface {
	Send(*RecordStatus) error
	Recv() (*Span, error)
	grpc.ServerStream
}

type ingestServiceRecordSpanServer struct {
	grpc.ServerStream
}

func (x *ingestServiceRecordSpanServer) Send(m *RecordStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ingestServiceRecordSpanServer) Recv() (*Span, error) {
	m := new(Span)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _IngestService_RecordSpanBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngestServiceServer).RecordSpanBatch(&ingestServiceRecordSpanBatchServer{stream})
}

type IngestService_RecordSpanBatchServer interface {
	Send(*RecordStatus) error
	Recv() (*SpanBatch, error)
	grpc.ServerStream
}

type ingestServiceRecordSpanBatchServer struct {
	grpc.ServerStream
}

func (x *ingestServiceRecordSpanBatchServer) Send(m *RecordStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ingestServiceRecordSpanBatchServer) Recv() (*SpanBatch, error) {
	m := new(SpanBatch)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _IngestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.oldfritter.trace.v1.IngestService",
	HandlerType: (*IngestServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecordSpan",
			Handler:       _IngestService_RecordSpan_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RecordSpanBatch",
			Handler:       _IngestService_RecordSpanBatch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "v3/internal/com_oldfritter_trace_v1/v1.proto",
}
