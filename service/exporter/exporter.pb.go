// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exporter.proto

package gravity_api_exporter

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

type SendEventRequest struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendEventRequest) Reset()         { *m = SendEventRequest{} }
func (m *SendEventRequest) String() string { return proto.CompactTextString(m) }
func (*SendEventRequest) ProtoMessage()    {}
func (*SendEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8826adc5babc285, []int{0}
}

func (m *SendEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendEventRequest.Unmarshal(m, b)
}
func (m *SendEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendEventRequest.Marshal(b, m, deterministic)
}
func (m *SendEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendEventRequest.Merge(m, src)
}
func (m *SendEventRequest) XXX_Size() int {
	return xxx_messageInfo_SendEventRequest.Size(m)
}
func (m *SendEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendEventRequest proto.InternalMessageInfo

func (m *SendEventRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *SendEventRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type SendEventReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendEventReply) Reset()         { *m = SendEventReply{} }
func (m *SendEventReply) String() string { return proto.CompactTextString(m) }
func (*SendEventReply) ProtoMessage()    {}
func (*SendEventReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8826adc5babc285, []int{1}
}

func (m *SendEventReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendEventReply.Unmarshal(m, b)
}
func (m *SendEventReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendEventReply.Marshal(b, m, deterministic)
}
func (m *SendEventReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendEventReply.Merge(m, src)
}
func (m *SendEventReply) XXX_Size() int {
	return xxx_messageInfo_SendEventReply.Size(m)
}
func (m *SendEventReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SendEventReply.DiscardUnknown(m)
}

var xxx_messageInfo_SendEventReply proto.InternalMessageInfo

func (m *SendEventReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SendEventReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*SendEventRequest)(nil), "gravity.api.exporter.SendEventRequest")
	proto.RegisterType((*SendEventReply)(nil), "gravity.api.exporter.SendEventReply")
}

func init() { proto.RegisterFile("exporter.proto", fileDescriptor_a8826adc5babc285) }

var fileDescriptor_a8826adc5babc285 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xad, 0x28, 0xc8,
	0x2f, 0x2a, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x49, 0x2f, 0x4a, 0x2c,
	0xcb, 0x2c, 0xa9, 0xd4, 0x4b, 0x2c, 0xc8, 0xd4, 0x83, 0xc9, 0x29, 0xb9, 0x71, 0x09, 0x04, 0xa7,
	0xe6, 0xa5, 0xb8, 0x96, 0xa5, 0xe6, 0x95, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49,
	0x70, 0xb1, 0x27, 0x67, 0x24, 0xe6, 0xe5, 0xa5, 0xe6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xc1, 0xb8, 0x20, 0x99, 0x82, 0xc4, 0xca, 0x9c, 0xfc, 0xc4, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0x9e, 0x20, 0x18, 0x57, 0xc9, 0x89, 0x8b, 0x0f, 0xc9, 0x9c, 0x82, 0x9c, 0x4a, 0x90, 0xda, 0xe2,
	0xd2, 0xe4, 0xe4, 0xd4, 0xe2, 0x62, 0xb0, 0x29, 0x1c, 0x41, 0x30, 0xae, 0x90, 0x18, 0x17, 0x5b,
	0x51, 0x6a, 0x62, 0x71, 0x7e, 0x1e, 0xd8, 0x10, 0xce, 0x20, 0x28, 0xcf, 0xe8, 0x0c, 0x23, 0x17,
	0x87, 0x2b, 0xd4, 0x61, 0x42, 0xd1, 0x5c, 0x9c, 0x70, 0x03, 0x85, 0xd4, 0xf4, 0xb0, 0x39, 0x5e,
	0x0f, 0xdd, 0xe5, 0x52, 0x2a, 0x04, 0xd5, 0x15, 0xe4, 0x54, 0x2a, 0x31, 0x08, 0x25, 0x73, 0xf1,
	0xc3, 0xc5, 0x82, 0x4b, 0x8a, 0x52, 0x13, 0x73, 0xa9, 0x6d, 0x85, 0x06, 0x63, 0x12, 0x1b, 0x38,
	0xdc, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x65, 0x32, 0xd0, 0x89, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExporterClient is the client API for Exporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExporterClient interface {
	SendEvent(ctx context.Context, in *SendEventRequest, opts ...grpc.CallOption) (*SendEventReply, error)
	SendEventStream(ctx context.Context, opts ...grpc.CallOption) (Exporter_SendEventStreamClient, error)
}

type exporterClient struct {
	cc *grpc.ClientConn
}

func NewExporterClient(cc *grpc.ClientConn) ExporterClient {
	return &exporterClient{cc}
}

func (c *exporterClient) SendEvent(ctx context.Context, in *SendEventRequest, opts ...grpc.CallOption) (*SendEventReply, error) {
	out := new(SendEventReply)
	err := c.cc.Invoke(ctx, "/gravity.api.exporter.Exporter/SendEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exporterClient) SendEventStream(ctx context.Context, opts ...grpc.CallOption) (Exporter_SendEventStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Exporter_serviceDesc.Streams[0], "/gravity.api.exporter.Exporter/SendEventStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &exporterSendEventStreamClient{stream}
	return x, nil
}

type Exporter_SendEventStreamClient interface {
	Send(*SendEventRequest) error
	CloseAndRecv() (*SendEventReply, error)
	grpc.ClientStream
}

type exporterSendEventStreamClient struct {
	grpc.ClientStream
}

func (x *exporterSendEventStreamClient) Send(m *SendEventRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exporterSendEventStreamClient) CloseAndRecv() (*SendEventReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendEventReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExporterServer is the server API for Exporter service.
type ExporterServer interface {
	SendEvent(context.Context, *SendEventRequest) (*SendEventReply, error)
	SendEventStream(Exporter_SendEventStreamServer) error
}

// UnimplementedExporterServer can be embedded to have forward compatible implementations.
type UnimplementedExporterServer struct {
}

func (*UnimplementedExporterServer) SendEvent(ctx context.Context, req *SendEventRequest) (*SendEventReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEvent not implemented")
}
func (*UnimplementedExporterServer) SendEventStream(srv Exporter_SendEventStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SendEventStream not implemented")
}

func RegisterExporterServer(s *grpc.Server, srv ExporterServer) {
	s.RegisterService(&_Exporter_serviceDesc, srv)
}

func _Exporter_SendEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExporterServer).SendEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gravity.api.exporter.Exporter/SendEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExporterServer).SendEvent(ctx, req.(*SendEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exporter_SendEventStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExporterServer).SendEventStream(&exporterSendEventStreamServer{stream})
}

type Exporter_SendEventStreamServer interface {
	SendAndClose(*SendEventReply) error
	Recv() (*SendEventRequest, error)
	grpc.ServerStream
}

type exporterSendEventStreamServer struct {
	grpc.ServerStream
}

func (x *exporterSendEventStreamServer) SendAndClose(m *SendEventReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exporterSendEventStreamServer) Recv() (*SendEventRequest, error) {
	m := new(SendEventRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Exporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gravity.api.exporter.Exporter",
	HandlerType: (*ExporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEvent",
			Handler:    _Exporter_SendEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendEventStream",
			Handler:       _Exporter_SendEventStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "exporter.proto",
}
