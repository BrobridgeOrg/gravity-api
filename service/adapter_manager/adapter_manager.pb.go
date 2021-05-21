// Code generated by protoc-gen-go. DO NOT EDIT.
// source: adapter_manager.proto

package gravity_api_adapter_manager

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RegisterAdapterRequest struct {
	AdapterID            string   `protobuf:"bytes,1,opt,name=adapterID,proto3" json:"adapterID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Component            string   `protobuf:"bytes,3,opt,name=component,proto3" json:"component,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterAdapterRequest) Reset()         { *m = RegisterAdapterRequest{} }
func (m *RegisterAdapterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterAdapterRequest) ProtoMessage()    {}
func (*RegisterAdapterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3481fc891e3ad44, []int{0}
}

func (m *RegisterAdapterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterAdapterRequest.Unmarshal(m, b)
}
func (m *RegisterAdapterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterAdapterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterAdapterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAdapterRequest.Merge(m, src)
}
func (m *RegisterAdapterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterAdapterRequest.Size(m)
}
func (m *RegisterAdapterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAdapterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAdapterRequest proto.InternalMessageInfo

func (m *RegisterAdapterRequest) GetAdapterID() string {
	if m != nil {
		return m.AdapterID
	}
	return ""
}

func (m *RegisterAdapterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterAdapterRequest) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

type RegisterAdapterReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterAdapterReply) Reset()         { *m = RegisterAdapterReply{} }
func (m *RegisterAdapterReply) String() string { return proto.CompactTextString(m) }
func (*RegisterAdapterReply) ProtoMessage()    {}
func (*RegisterAdapterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3481fc891e3ad44, []int{1}
}

func (m *RegisterAdapterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterAdapterReply.Unmarshal(m, b)
}
func (m *RegisterAdapterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterAdapterReply.Marshal(b, m, deterministic)
}
func (m *RegisterAdapterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAdapterReply.Merge(m, src)
}
func (m *RegisterAdapterReply) XXX_Size() int {
	return xxx_messageInfo_RegisterAdapterReply.Size(m)
}
func (m *RegisterAdapterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAdapterReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAdapterReply proto.InternalMessageInfo

func (m *RegisterAdapterReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RegisterAdapterReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type UnregisterAdapterRequest struct {
	AdapterID            string   `protobuf:"bytes,1,opt,name=adapterID,proto3" json:"adapterID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterAdapterRequest) Reset()         { *m = UnregisterAdapterRequest{} }
func (m *UnregisterAdapterRequest) String() string { return proto.CompactTextString(m) }
func (*UnregisterAdapterRequest) ProtoMessage()    {}
func (*UnregisterAdapterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3481fc891e3ad44, []int{2}
}

func (m *UnregisterAdapterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterAdapterRequest.Unmarshal(m, b)
}
func (m *UnregisterAdapterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterAdapterRequest.Marshal(b, m, deterministic)
}
func (m *UnregisterAdapterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterAdapterRequest.Merge(m, src)
}
func (m *UnregisterAdapterRequest) XXX_Size() int {
	return xxx_messageInfo_UnregisterAdapterRequest.Size(m)
}
func (m *UnregisterAdapterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterAdapterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterAdapterRequest proto.InternalMessageInfo

func (m *UnregisterAdapterRequest) GetAdapterID() string {
	if m != nil {
		return m.AdapterID
	}
	return ""
}

type UnregisterAdapterReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterAdapterReply) Reset()         { *m = UnregisterAdapterReply{} }
func (m *UnregisterAdapterReply) String() string { return proto.CompactTextString(m) }
func (*UnregisterAdapterReply) ProtoMessage()    {}
func (*UnregisterAdapterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3481fc891e3ad44, []int{3}
}

func (m *UnregisterAdapterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterAdapterReply.Unmarshal(m, b)
}
func (m *UnregisterAdapterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterAdapterReply.Marshal(b, m, deterministic)
}
func (m *UnregisterAdapterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterAdapterReply.Merge(m, src)
}
func (m *UnregisterAdapterReply) XXX_Size() int {
	return xxx_messageInfo_UnregisterAdapterReply.Size(m)
}
func (m *UnregisterAdapterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterAdapterReply.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterAdapterReply proto.InternalMessageInfo

func (m *UnregisterAdapterReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UnregisterAdapterReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type Adapter struct {
	AdapterID            string   `protobuf:"bytes,1,opt,name=adapterID,proto3" json:"adapterID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Component            string   `protobuf:"bytes,3,opt,name=component,proto3" json:"component,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Adapter) Reset()         { *m = Adapter{} }
func (m *Adapter) String() string { return proto.CompactTextString(m) }
func (*Adapter) ProtoMessage()    {}
func (*Adapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3481fc891e3ad44, []int{4}
}

func (m *Adapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Adapter.Unmarshal(m, b)
}
func (m *Adapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Adapter.Marshal(b, m, deterministic)
}
func (m *Adapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Adapter.Merge(m, src)
}
func (m *Adapter) XXX_Size() int {
	return xxx_messageInfo_Adapter.Size(m)
}
func (m *Adapter) XXX_DiscardUnknown() {
	xxx_messageInfo_Adapter.DiscardUnknown(m)
}

var xxx_messageInfo_Adapter proto.InternalMessageInfo

func (m *Adapter) GetAdapterID() string {
	if m != nil {
		return m.AdapterID
	}
	return ""
}

func (m *Adapter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Adapter) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

type GetAdaptersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdaptersRequest) Reset()         { *m = GetAdaptersRequest{} }
func (m *GetAdaptersRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdaptersRequest) ProtoMessage()    {}
func (*GetAdaptersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3481fc891e3ad44, []int{5}
}

func (m *GetAdaptersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdaptersRequest.Unmarshal(m, b)
}
func (m *GetAdaptersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdaptersRequest.Marshal(b, m, deterministic)
}
func (m *GetAdaptersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdaptersRequest.Merge(m, src)
}
func (m *GetAdaptersRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdaptersRequest.Size(m)
}
func (m *GetAdaptersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdaptersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdaptersRequest proto.InternalMessageInfo

type GetAdaptersReply struct {
	Success              bool       `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string     `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Adapters             []*Adapter `protobuf:"bytes,3,rep,name=adapters,proto3" json:"adapters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAdaptersReply) Reset()         { *m = GetAdaptersReply{} }
func (m *GetAdaptersReply) String() string { return proto.CompactTextString(m) }
func (*GetAdaptersReply) ProtoMessage()    {}
func (*GetAdaptersReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3481fc891e3ad44, []int{6}
}

func (m *GetAdaptersReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdaptersReply.Unmarshal(m, b)
}
func (m *GetAdaptersReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdaptersReply.Marshal(b, m, deterministic)
}
func (m *GetAdaptersReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdaptersReply.Merge(m, src)
}
func (m *GetAdaptersReply) XXX_Size() int {
	return xxx_messageInfo_GetAdaptersReply.Size(m)
}
func (m *GetAdaptersReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdaptersReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdaptersReply proto.InternalMessageInfo

func (m *GetAdaptersReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetAdaptersReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *GetAdaptersReply) GetAdapters() []*Adapter {
	if m != nil {
		return m.Adapters
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterAdapterRequest)(nil), "gravity.api.adapter_manager.RegisterAdapterRequest")
	proto.RegisterType((*RegisterAdapterReply)(nil), "gravity.api.adapter_manager.RegisterAdapterReply")
	proto.RegisterType((*UnregisterAdapterRequest)(nil), "gravity.api.adapter_manager.UnregisterAdapterRequest")
	proto.RegisterType((*UnregisterAdapterReply)(nil), "gravity.api.adapter_manager.UnregisterAdapterReply")
	proto.RegisterType((*Adapter)(nil), "gravity.api.adapter_manager.Adapter")
	proto.RegisterType((*GetAdaptersRequest)(nil), "gravity.api.adapter_manager.GetAdaptersRequest")
	proto.RegisterType((*GetAdaptersReply)(nil), "gravity.api.adapter_manager.GetAdaptersReply")
}

func init() { proto.RegisterFile("adapter_manager.proto", fileDescriptor_b3481fc891e3ad44) }

var fileDescriptor_b3481fc891e3ad44 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x4c, 0x49, 0x2c,
	0x28, 0x49, 0x2d, 0x8a, 0xcf, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x4e, 0x2f, 0x4a, 0x2c, 0xcb, 0x2c, 0xa9, 0xd4, 0x4b, 0x2c, 0xc8, 0xd4, 0x43,
	0x53, 0xa2, 0x94, 0xc1, 0x25, 0x16, 0x94, 0x9a, 0x9e, 0x59, 0x5c, 0x92, 0x5a, 0xe4, 0x08, 0x91,
	0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe1, 0xe2, 0x84, 0x2a, 0xf6, 0x74, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6,
	0x4a, 0x30, 0x81, 0x25, 0xc0, 0x6c, 0x90, 0x8e, 0xe4, 0xfc, 0xdc, 0x82, 0xfc, 0xbc, 0xd4, 0xbc,
	0x12, 0x09, 0x66, 0x88, 0x0e, 0xb8, 0x80, 0x92, 0x07, 0x97, 0x08, 0x86, 0x4d, 0x05, 0x39, 0x95,
	0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa5, 0xc9, 0xc9, 0xa9, 0xc5, 0xc5, 0x60, 0x5b, 0x38, 0x82, 0x60,
	0x5c, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xc4, 0xe2, 0xfc, 0x3c, 0xa8, 0x2d, 0x50, 0x9e, 0x92,
	0x05, 0x97, 0x44, 0x68, 0x5e, 0x11, 0x19, 0xae, 0x56, 0xf2, 0xe2, 0x12, 0xc3, 0xa2, 0x93, 0x3c,
	0x57, 0x44, 0x72, 0xb1, 0x43, 0x4d, 0xa0, 0x7a, 0x50, 0x89, 0x70, 0x09, 0xb9, 0xa7, 0x96, 0x40,
	0x4d, 0x2f, 0x86, 0x7a, 0x4d, 0xa9, 0x8d, 0x91, 0x4b, 0x00, 0x45, 0x98, 0x2c, 0x77, 0x0b, 0x39,
	0x70, 0x71, 0x40, 0xdd, 0x56, 0x2c, 0xc1, 0xac, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xa2, 0x87, 0x27,
	0x85, 0xe8, 0xc1, 0x82, 0x09, 0xae, 0x2b, 0x89, 0x0d, 0x9c, 0xae, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x58, 0xae, 0x36, 0x81, 0x70, 0x02, 0x00, 0x00,
}
