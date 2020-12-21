// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data_source_adapter.proto

package gravity_api_dsa

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

type PublishRequest struct {
	EventName            string            `protobuf:"bytes,1,opt,name=eventName,proto3" json:"eventName,omitempty"`
	Payload              []byte            `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Meta                 map[string][]byte `protobuf:"bytes,3,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_102a39ffe4b3d003, []int{0}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *PublishRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PublishRequest) GetMeta() map[string][]byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

type PublishReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishReply) Reset()         { *m = PublishReply{} }
func (m *PublishReply) String() string { return proto.CompactTextString(m) }
func (*PublishReply) ProtoMessage()    {}
func (*PublishReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_102a39ffe4b3d003, []int{1}
}

func (m *PublishReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishReply.Unmarshal(m, b)
}
func (m *PublishReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishReply.Marshal(b, m, deterministic)
}
func (m *PublishReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishReply.Merge(m, src)
}
func (m *PublishReply) XXX_Size() int {
	return xxx_messageInfo_PublishReply.Size(m)
}
func (m *PublishReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishReply.DiscardUnknown(m)
}

var xxx_messageInfo_PublishReply proto.InternalMessageInfo

func (m *PublishReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PublishReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*PublishRequest)(nil), "gravity.api.dsa.PublishRequest")
	proto.RegisterMapType((map[string][]byte)(nil), "gravity.api.dsa.PublishRequest.MetaEntry")
	proto.RegisterType((*PublishReply)(nil), "gravity.api.dsa.PublishReply")
}

func init() { proto.RegisterFile("data_source_adapter.proto", fileDescriptor_102a39ffe4b3d003) }

var fileDescriptor_102a39ffe4b3d003 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xcd, 0xa6, 0x9b, 0xbd, 0xce, 0x7f, 0x41, 0xa4, 0x0e, 0xc5, 0xd2, 0xa7, 0xfa, 0x92,
	0x87, 0xf9, 0xa0, 0x08, 0x82, 0x82, 0x7b, 0x12, 0x45, 0xb2, 0x0f, 0x30, 0xee, 0xda, 0x8b, 0x16,
	0xbb, 0x36, 0x26, 0x69, 0x21, 0x1f, 0xcc, 0x77, 0x3f, 0x9a, 0xb4, 0xeb, 0x26, 0x0a, 0xe2, 0x83,
	0x6f, 0x39, 0x87, 0xdc, 0xf3, 0x4b, 0xce, 0x85, 0xa3, 0x04, 0x2d, 0x4e, 0x4d, 0x51, 0xea, 0x98,
	0xa6, 0x98, 0xa0, 0xb2, 0xa4, 0x85, 0xd2, 0x85, 0x2d, 0xf8, 0xee, 0xb3, 0xc6, 0x2a, 0xb5, 0x4e,
	0xa0, 0x4a, 0x45, 0x62, 0x30, 0xfc, 0x60, 0xb0, 0xf3, 0x54, 0xce, 0xb2, 0xd4, 0xbc, 0x48, 0x7a,
	0x2b, 0xc9, 0x58, 0x7e, 0x0c, 0x1e, 0x55, 0x94, 0xdb, 0x47, 0x9c, 0x93, 0xcf, 0x02, 0x16, 0x79,
	0xf2, 0xcb, 0xe0, 0x3e, 0xf4, 0x15, 0xba, 0xac, 0xc0, 0xc4, 0xef, 0x04, 0x2c, 0x1a, 0xc8, 0xa5,
	0xe4, 0xd7, 0xb0, 0x3e, 0x27, 0x8b, 0x7e, 0x37, 0xe8, 0x46, 0x5b, 0xa3, 0x33, 0xf1, 0x03, 0x25,
	0xbe, 0x63, 0xc4, 0x03, 0x59, 0x1c, 0xe7, 0x56, 0x3b, 0xd9, 0x8c, 0x0d, 0x2f, 0xc0, 0x5b, 0x59,
	0x7c, 0x0f, 0xba, 0xaf, 0xe4, 0x5a, 0x7a, 0x7d, 0xe4, 0x07, 0xb0, 0x51, 0x61, 0x56, 0x52, 0x4b,
	0x5d, 0x88, 0xab, 0xce, 0x25, 0x0b, 0x6f, 0x60, 0xb0, 0x8a, 0x56, 0x99, 0xab, 0x5f, 0x68, 0xca,
	0x38, 0x26, 0x63, 0x9a, 0xf9, 0x4d, 0xb9, 0x94, 0xfc, 0x10, 0x7a, 0x9a, 0xd0, 0x14, 0x79, 0x13,
	0xe2, 0xc9, 0x56, 0x8d, 0xde, 0x19, 0xec, 0xdf, 0xa1, 0xc5, 0x49, 0x53, 0xd9, 0xed, 0xa2, 0x31,
	0x7e, 0x0f, 0xfd, 0x36, 0x97, 0x9f, 0xfe, 0xf1, 0x99, 0xe1, 0xc9, 0xef, 0x17, 0x54, 0xe6, 0xc2,
	0x35, 0x3e, 0x81, 0xed, 0xd6, 0x19, 0xd7, 0x55, 0x9a, 0xff, 0x47, 0x46, 0x6c, 0xd6, 0x6b, 0x96,
	0x7a, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x0a, 0x22, 0x24, 0xf1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataSourceAdapterClient is the client API for DataSourceAdapter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataSourceAdapterClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishReply, error)
	PublishEvents(ctx context.Context, opts ...grpc.CallOption) (DataSourceAdapter_PublishEventsClient, error)
}

type dataSourceAdapterClient struct {
	cc *grpc.ClientConn
}

func NewDataSourceAdapterClient(cc *grpc.ClientConn) DataSourceAdapterClient {
	return &dataSourceAdapterClient{cc}
}

func (c *dataSourceAdapterClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishReply, error) {
	out := new(PublishReply)
	err := c.cc.Invoke(ctx, "/gravity.api.dsa.DataSourceAdapter/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceAdapterClient) PublishEvents(ctx context.Context, opts ...grpc.CallOption) (DataSourceAdapter_PublishEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataSourceAdapter_serviceDesc.Streams[0], "/gravity.api.dsa.DataSourceAdapter/PublishEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataSourceAdapterPublishEventsClient{stream}
	return x, nil
}

type DataSourceAdapter_PublishEventsClient interface {
	Send(*PublishRequest) error
	CloseAndRecv() (*PublishReply, error)
	grpc.ClientStream
}

type dataSourceAdapterPublishEventsClient struct {
	grpc.ClientStream
}

func (x *dataSourceAdapterPublishEventsClient) Send(m *PublishRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataSourceAdapterPublishEventsClient) CloseAndRecv() (*PublishReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PublishReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataSourceAdapterServer is the server API for DataSourceAdapter service.
type DataSourceAdapterServer interface {
	Publish(context.Context, *PublishRequest) (*PublishReply, error)
	PublishEvents(DataSourceAdapter_PublishEventsServer) error
}

// UnimplementedDataSourceAdapterServer can be embedded to have forward compatible implementations.
type UnimplementedDataSourceAdapterServer struct {
}

func (*UnimplementedDataSourceAdapterServer) Publish(ctx context.Context, req *PublishRequest) (*PublishReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedDataSourceAdapterServer) PublishEvents(srv DataSourceAdapter_PublishEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method PublishEvents not implemented")
}

func RegisterDataSourceAdapterServer(s *grpc.Server, srv DataSourceAdapterServer) {
	s.RegisterService(&_DataSourceAdapter_serviceDesc, srv)
}

func _DataSourceAdapter_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceAdapterServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gravity.api.dsa.DataSourceAdapter/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceAdapterServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSourceAdapter_PublishEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataSourceAdapterServer).PublishEvents(&dataSourceAdapterPublishEventsServer{stream})
}

type DataSourceAdapter_PublishEventsServer interface {
	SendAndClose(*PublishReply) error
	Recv() (*PublishRequest, error)
	grpc.ServerStream
}

type dataSourceAdapterPublishEventsServer struct {
	grpc.ServerStream
}

func (x *dataSourceAdapterPublishEventsServer) SendAndClose(m *PublishReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataSourceAdapterPublishEventsServer) Recv() (*PublishRequest, error) {
	m := new(PublishRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DataSourceAdapter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gravity.api.dsa.DataSourceAdapter",
	HandlerType: (*DataSourceAdapterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _DataSourceAdapter_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PublishEvents",
			Handler:       _DataSourceAdapter_PublishEvents_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "data_source_adapter.proto",
}
