// Code generated by protoc-gen-go. DO NOT EDIT.
// source: synchronizer.proto

package gravity_api_synchronizer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type AssignPipelineRequest struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	PipelineID           uint64   `protobuf:"varint,2,opt,name=pipelineID,proto3" json:"pipelineID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignPipelineRequest) Reset()         { *m = AssignPipelineRequest{} }
func (m *AssignPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*AssignPipelineRequest) ProtoMessage()    {}
func (*AssignPipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{0}
}

func (m *AssignPipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignPipelineRequest.Unmarshal(m, b)
}
func (m *AssignPipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignPipelineRequest.Marshal(b, m, deterministic)
}
func (m *AssignPipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignPipelineRequest.Merge(m, src)
}
func (m *AssignPipelineRequest) XXX_Size() int {
	return xxx_messageInfo_AssignPipelineRequest.Size(m)
}
func (m *AssignPipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignPipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AssignPipelineRequest proto.InternalMessageInfo

func (m *AssignPipelineRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *AssignPipelineRequest) GetPipelineID() uint64 {
	if m != nil {
		return m.PipelineID
	}
	return 0
}

type AssignPipelineReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignPipelineReply) Reset()         { *m = AssignPipelineReply{} }
func (m *AssignPipelineReply) String() string { return proto.CompactTextString(m) }
func (*AssignPipelineReply) ProtoMessage()    {}
func (*AssignPipelineReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{1}
}

func (m *AssignPipelineReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignPipelineReply.Unmarshal(m, b)
}
func (m *AssignPipelineReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignPipelineReply.Marshal(b, m, deterministic)
}
func (m *AssignPipelineReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignPipelineReply.Merge(m, src)
}
func (m *AssignPipelineReply) XXX_Size() int {
	return xxx_messageInfo_AssignPipelineReply.Size(m)
}
func (m *AssignPipelineReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignPipelineReply.DiscardUnknown(m)
}

var xxx_messageInfo_AssignPipelineReply proto.InternalMessageInfo

func (m *AssignPipelineReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *AssignPipelineReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type RevokePipelineRequest struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	PipelineID           uint64   `protobuf:"varint,2,opt,name=pipelineID,proto3" json:"pipelineID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokePipelineRequest) Reset()         { *m = RevokePipelineRequest{} }
func (m *RevokePipelineRequest) String() string { return proto.CompactTextString(m) }
func (*RevokePipelineRequest) ProtoMessage()    {}
func (*RevokePipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{2}
}

func (m *RevokePipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokePipelineRequest.Unmarshal(m, b)
}
func (m *RevokePipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokePipelineRequest.Marshal(b, m, deterministic)
}
func (m *RevokePipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokePipelineRequest.Merge(m, src)
}
func (m *RevokePipelineRequest) XXX_Size() int {
	return xxx_messageInfo_RevokePipelineRequest.Size(m)
}
func (m *RevokePipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokePipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokePipelineRequest proto.InternalMessageInfo

func (m *RevokePipelineRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *RevokePipelineRequest) GetPipelineID() uint64 {
	if m != nil {
		return m.PipelineID
	}
	return 0
}

type RevokePipelineReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokePipelineReply) Reset()         { *m = RevokePipelineReply{} }
func (m *RevokePipelineReply) String() string { return proto.CompactTextString(m) }
func (*RevokePipelineReply) ProtoMessage()    {}
func (*RevokePipelineReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{3}
}

func (m *RevokePipelineReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokePipelineReply.Unmarshal(m, b)
}
func (m *RevokePipelineReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokePipelineReply.Marshal(b, m, deterministic)
}
func (m *RevokePipelineReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokePipelineReply.Merge(m, src)
}
func (m *RevokePipelineReply) XXX_Size() int {
	return xxx_messageInfo_RevokePipelineReply.Size(m)
}
func (m *RevokePipelineReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokePipelineReply.DiscardUnknown(m)
}

var xxx_messageInfo_RevokePipelineReply proto.InternalMessageInfo

func (m *RevokePipelineReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RevokePipelineReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type RegisterSubscriberRequest struct {
	SubscriberID         string   `protobuf:"bytes,1,opt,name=subscriberID,proto3" json:"subscriberID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterSubscriberRequest) Reset()         { *m = RegisterSubscriberRequest{} }
func (m *RegisterSubscriberRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterSubscriberRequest) ProtoMessage()    {}
func (*RegisterSubscriberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{4}
}

func (m *RegisterSubscriberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterSubscriberRequest.Unmarshal(m, b)
}
func (m *RegisterSubscriberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterSubscriberRequest.Marshal(b, m, deterministic)
}
func (m *RegisterSubscriberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterSubscriberRequest.Merge(m, src)
}
func (m *RegisterSubscriberRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterSubscriberRequest.Size(m)
}
func (m *RegisterSubscriberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterSubscriberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterSubscriberRequest proto.InternalMessageInfo

func (m *RegisterSubscriberRequest) GetSubscriberID() string {
	if m != nil {
		return m.SubscriberID
	}
	return ""
}

func (m *RegisterSubscriberRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RegisterSubscriberReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterSubscriberReply) Reset()         { *m = RegisterSubscriberReply{} }
func (m *RegisterSubscriberReply) String() string { return proto.CompactTextString(m) }
func (*RegisterSubscriberReply) ProtoMessage()    {}
func (*RegisterSubscriberReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{5}
}

func (m *RegisterSubscriberReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterSubscriberReply.Unmarshal(m, b)
}
func (m *RegisterSubscriberReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterSubscriberReply.Marshal(b, m, deterministic)
}
func (m *RegisterSubscriberReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterSubscriberReply.Merge(m, src)
}
func (m *RegisterSubscriberReply) XXX_Size() int {
	return xxx_messageInfo_RegisterSubscriberReply.Size(m)
}
func (m *RegisterSubscriberReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterSubscriberReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterSubscriberReply proto.InternalMessageInfo

func (m *RegisterSubscriberReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RegisterSubscriberReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type UnregisterSubscriberRequest struct {
	SubscriberID         string   `protobuf:"bytes,1,opt,name=subscriberID,proto3" json:"subscriberID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterSubscriberRequest) Reset()         { *m = UnregisterSubscriberRequest{} }
func (m *UnregisterSubscriberRequest) String() string { return proto.CompactTextString(m) }
func (*UnregisterSubscriberRequest) ProtoMessage()    {}
func (*UnregisterSubscriberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{6}
}

func (m *UnregisterSubscriberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterSubscriberRequest.Unmarshal(m, b)
}
func (m *UnregisterSubscriberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterSubscriberRequest.Marshal(b, m, deterministic)
}
func (m *UnregisterSubscriberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterSubscriberRequest.Merge(m, src)
}
func (m *UnregisterSubscriberRequest) XXX_Size() int {
	return xxx_messageInfo_UnregisterSubscriberRequest.Size(m)
}
func (m *UnregisterSubscriberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterSubscriberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterSubscriberRequest proto.InternalMessageInfo

func (m *UnregisterSubscriberRequest) GetSubscriberID() string {
	if m != nil {
		return m.SubscriberID
	}
	return ""
}

type UnregisterSubscriberReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterSubscriberReply) Reset()         { *m = UnregisterSubscriberReply{} }
func (m *UnregisterSubscriberReply) String() string { return proto.CompactTextString(m) }
func (*UnregisterSubscriberReply) ProtoMessage()    {}
func (*UnregisterSubscriberReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{7}
}

func (m *UnregisterSubscriberReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterSubscriberReply.Unmarshal(m, b)
}
func (m *UnregisterSubscriberReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterSubscriberReply.Marshal(b, m, deterministic)
}
func (m *UnregisterSubscriberReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterSubscriberReply.Merge(m, src)
}
func (m *UnregisterSubscriberReply) XXX_Size() int {
	return xxx_messageInfo_UnregisterSubscriberReply.Size(m)
}
func (m *UnregisterSubscriberReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterSubscriberReply.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterSubscriberReply proto.InternalMessageInfo

func (m *UnregisterSubscriberReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UnregisterSubscriberReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type SubscribeToCollectionsRequest struct {
	SubscriberID         string   `protobuf:"bytes,1,opt,name=subscriberID,proto3" json:"subscriberID,omitempty"`
	Collections          []string `protobuf:"bytes,2,rep,name=Collections,proto3" json:"Collections,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeToCollectionsRequest) Reset()         { *m = SubscribeToCollectionsRequest{} }
func (m *SubscribeToCollectionsRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeToCollectionsRequest) ProtoMessage()    {}
func (*SubscribeToCollectionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{8}
}

func (m *SubscribeToCollectionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeToCollectionsRequest.Unmarshal(m, b)
}
func (m *SubscribeToCollectionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeToCollectionsRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeToCollectionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeToCollectionsRequest.Merge(m, src)
}
func (m *SubscribeToCollectionsRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeToCollectionsRequest.Size(m)
}
func (m *SubscribeToCollectionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeToCollectionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeToCollectionsRequest proto.InternalMessageInfo

func (m *SubscribeToCollectionsRequest) GetSubscriberID() string {
	if m != nil {
		return m.SubscriberID
	}
	return ""
}

func (m *SubscribeToCollectionsRequest) GetCollections() []string {
	if m != nil {
		return m.Collections
	}
	return nil
}

type SubscribeToCollectionsReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Collections          []string `protobuf:"bytes,3,rep,name=Collections,proto3" json:"Collections,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeToCollectionsReply) Reset()         { *m = SubscribeToCollectionsReply{} }
func (m *SubscribeToCollectionsReply) String() string { return proto.CompactTextString(m) }
func (*SubscribeToCollectionsReply) ProtoMessage()    {}
func (*SubscribeToCollectionsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{9}
}

func (m *SubscribeToCollectionsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeToCollectionsReply.Unmarshal(m, b)
}
func (m *SubscribeToCollectionsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeToCollectionsReply.Marshal(b, m, deterministic)
}
func (m *SubscribeToCollectionsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeToCollectionsReply.Merge(m, src)
}
func (m *SubscribeToCollectionsReply) XXX_Size() int {
	return xxx_messageInfo_SubscribeToCollectionsReply.Size(m)
}
func (m *SubscribeToCollectionsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeToCollectionsReply.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeToCollectionsReply proto.InternalMessageInfo

func (m *SubscribeToCollectionsReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SubscribeToCollectionsReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *SubscribeToCollectionsReply) GetCollections() []string {
	if m != nil {
		return m.Collections
	}
	return nil
}

type UnsubscribeFromCollectionsRequest struct {
	SubscriberID         string   `protobuf:"bytes,1,opt,name=subscriberID,proto3" json:"subscriberID,omitempty"`
	Collections          []string `protobuf:"bytes,2,rep,name=collections,proto3" json:"collections,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnsubscribeFromCollectionsRequest) Reset()         { *m = UnsubscribeFromCollectionsRequest{} }
func (m *UnsubscribeFromCollectionsRequest) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeFromCollectionsRequest) ProtoMessage()    {}
func (*UnsubscribeFromCollectionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{10}
}

func (m *UnsubscribeFromCollectionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsubscribeFromCollectionsRequest.Unmarshal(m, b)
}
func (m *UnsubscribeFromCollectionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsubscribeFromCollectionsRequest.Marshal(b, m, deterministic)
}
func (m *UnsubscribeFromCollectionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubscribeFromCollectionsRequest.Merge(m, src)
}
func (m *UnsubscribeFromCollectionsRequest) XXX_Size() int {
	return xxx_messageInfo_UnsubscribeFromCollectionsRequest.Size(m)
}
func (m *UnsubscribeFromCollectionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubscribeFromCollectionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubscribeFromCollectionsRequest proto.InternalMessageInfo

func (m *UnsubscribeFromCollectionsRequest) GetSubscriberID() string {
	if m != nil {
		return m.SubscriberID
	}
	return ""
}

func (m *UnsubscribeFromCollectionsRequest) GetCollections() []string {
	if m != nil {
		return m.Collections
	}
	return nil
}

type UnsubscribeFromCollectionsReply struct {
	Collections          []string `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnsubscribeFromCollectionsReply) Reset()         { *m = UnsubscribeFromCollectionsReply{} }
func (m *UnsubscribeFromCollectionsReply) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeFromCollectionsReply) ProtoMessage()    {}
func (*UnsubscribeFromCollectionsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{11}
}

func (m *UnsubscribeFromCollectionsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsubscribeFromCollectionsReply.Unmarshal(m, b)
}
func (m *UnsubscribeFromCollectionsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsubscribeFromCollectionsReply.Marshal(b, m, deterministic)
}
func (m *UnsubscribeFromCollectionsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubscribeFromCollectionsReply.Merge(m, src)
}
func (m *UnsubscribeFromCollectionsReply) XXX_Size() int {
	return xxx_messageInfo_UnsubscribeFromCollectionsReply.Size(m)
}
func (m *UnsubscribeFromCollectionsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubscribeFromCollectionsReply.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubscribeFromCollectionsReply proto.InternalMessageInfo

func (m *UnsubscribeFromCollectionsReply) GetCollections() []string {
	if m != nil {
		return m.Collections
	}
	return nil
}

type GetPipelineStateRequest struct {
	PipelineID           string   `protobuf:"bytes,1,opt,name=pipelineID,proto3" json:"pipelineID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPipelineStateRequest) Reset()         { *m = GetPipelineStateRequest{} }
func (m *GetPipelineStateRequest) String() string { return proto.CompactTextString(m) }
func (*GetPipelineStateRequest) ProtoMessage()    {}
func (*GetPipelineStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{12}
}

func (m *GetPipelineStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPipelineStateRequest.Unmarshal(m, b)
}
func (m *GetPipelineStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPipelineStateRequest.Marshal(b, m, deterministic)
}
func (m *GetPipelineStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPipelineStateRequest.Merge(m, src)
}
func (m *GetPipelineStateRequest) XXX_Size() int {
	return xxx_messageInfo_GetPipelineStateRequest.Size(m)
}
func (m *GetPipelineStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPipelineStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPipelineStateRequest proto.InternalMessageInfo

func (m *GetPipelineStateRequest) GetPipelineID() string {
	if m != nil {
		return m.PipelineID
	}
	return ""
}

type GetPipelineStateReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	LastSeq              uint64   `protobuf:"varint,3,opt,name=lastSeq,proto3" json:"lastSeq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPipelineStateReply) Reset()         { *m = GetPipelineStateReply{} }
func (m *GetPipelineStateReply) String() string { return proto.CompactTextString(m) }
func (*GetPipelineStateReply) ProtoMessage()    {}
func (*GetPipelineStateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{13}
}

func (m *GetPipelineStateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPipelineStateReply.Unmarshal(m, b)
}
func (m *GetPipelineStateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPipelineStateReply.Marshal(b, m, deterministic)
}
func (m *GetPipelineStateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPipelineStateReply.Merge(m, src)
}
func (m *GetPipelineStateReply) XXX_Size() int {
	return xxx_messageInfo_GetPipelineStateReply.Size(m)
}
func (m *GetPipelineStateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPipelineStateReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetPipelineStateReply proto.InternalMessageInfo

func (m *GetPipelineStateReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetPipelineStateReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *GetPipelineStateReply) GetLastSeq() uint64 {
	if m != nil {
		return m.LastSeq
	}
	return 0
}

type PipelineFetchRequest struct {
	SubscriberID         string   `protobuf:"bytes,1,opt,name=subscriberID,proto3" json:"subscriberID,omitempty"`
	PipelineID           uint64   `protobuf:"varint,2,opt,name=pipelineID,proto3" json:"pipelineID,omitempty"`
	StartAt              uint64   `protobuf:"varint,3,opt,name=startAt,proto3" json:"startAt,omitempty"`
	Offset               uint64   `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                int64    `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PipelineFetchRequest) Reset()         { *m = PipelineFetchRequest{} }
func (m *PipelineFetchRequest) String() string { return proto.CompactTextString(m) }
func (*PipelineFetchRequest) ProtoMessage()    {}
func (*PipelineFetchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{14}
}

func (m *PipelineFetchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PipelineFetchRequest.Unmarshal(m, b)
}
func (m *PipelineFetchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PipelineFetchRequest.Marshal(b, m, deterministic)
}
func (m *PipelineFetchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipelineFetchRequest.Merge(m, src)
}
func (m *PipelineFetchRequest) XXX_Size() int {
	return xxx_messageInfo_PipelineFetchRequest.Size(m)
}
func (m *PipelineFetchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PipelineFetchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PipelineFetchRequest proto.InternalMessageInfo

func (m *PipelineFetchRequest) GetSubscriberID() string {
	if m != nil {
		return m.SubscriberID
	}
	return ""
}

func (m *PipelineFetchRequest) GetPipelineID() uint64 {
	if m != nil {
		return m.PipelineID
	}
	return 0
}

func (m *PipelineFetchRequest) GetStartAt() uint64 {
	if m != nil {
		return m.StartAt
	}
	return 0
}

func (m *PipelineFetchRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *PipelineFetchRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type PipelineFetchReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Count                uint64   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	LastSeq              uint64   `protobuf:"varint,4,opt,name=lastSeq,proto3" json:"lastSeq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PipelineFetchReply) Reset()         { *m = PipelineFetchReply{} }
func (m *PipelineFetchReply) String() string { return proto.CompactTextString(m) }
func (*PipelineFetchReply) ProtoMessage()    {}
func (*PipelineFetchReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{15}
}

func (m *PipelineFetchReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PipelineFetchReply.Unmarshal(m, b)
}
func (m *PipelineFetchReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PipelineFetchReply.Marshal(b, m, deterministic)
}
func (m *PipelineFetchReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipelineFetchReply.Merge(m, src)
}
func (m *PipelineFetchReply) XXX_Size() int {
	return xxx_messageInfo_PipelineFetchReply.Size(m)
}
func (m *PipelineFetchReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PipelineFetchReply.DiscardUnknown(m)
}

var xxx_messageInfo_PipelineFetchReply proto.InternalMessageInfo

func (m *PipelineFetchReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PipelineFetchReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *PipelineFetchReply) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *PipelineFetchReply) GetLastSeq() uint64 {
	if m != nil {
		return m.LastSeq
	}
	return 0
}

func init() {
	proto.RegisterType((*AssignPipelineRequest)(nil), "gravity.api.synchronizer.AssignPipelineRequest")
	proto.RegisterType((*AssignPipelineReply)(nil), "gravity.api.synchronizer.AssignPipelineReply")
	proto.RegisterType((*RevokePipelineRequest)(nil), "gravity.api.synchronizer.RevokePipelineRequest")
	proto.RegisterType((*RevokePipelineReply)(nil), "gravity.api.synchronizer.RevokePipelineReply")
	proto.RegisterType((*RegisterSubscriberRequest)(nil), "gravity.api.synchronizer.RegisterSubscriberRequest")
	proto.RegisterType((*RegisterSubscriberReply)(nil), "gravity.api.synchronizer.RegisterSubscriberReply")
	proto.RegisterType((*UnregisterSubscriberRequest)(nil), "gravity.api.synchronizer.UnregisterSubscriberRequest")
	proto.RegisterType((*UnregisterSubscriberReply)(nil), "gravity.api.synchronizer.UnregisterSubscriberReply")
	proto.RegisterType((*SubscribeToCollectionsRequest)(nil), "gravity.api.synchronizer.SubscribeToCollectionsRequest")
	proto.RegisterType((*SubscribeToCollectionsReply)(nil), "gravity.api.synchronizer.SubscribeToCollectionsReply")
	proto.RegisterType((*UnsubscribeFromCollectionsRequest)(nil), "gravity.api.synchronizer.UnsubscribeFromCollectionsRequest")
	proto.RegisterType((*UnsubscribeFromCollectionsReply)(nil), "gravity.api.synchronizer.UnsubscribeFromCollectionsReply")
	proto.RegisterType((*GetPipelineStateRequest)(nil), "gravity.api.synchronizer.GetPipelineStateRequest")
	proto.RegisterType((*GetPipelineStateReply)(nil), "gravity.api.synchronizer.GetPipelineStateReply")
	proto.RegisterType((*PipelineFetchRequest)(nil), "gravity.api.synchronizer.PipelineFetchRequest")
	proto.RegisterType((*PipelineFetchReply)(nil), "gravity.api.synchronizer.PipelineFetchReply")
}

func init() { proto.RegisterFile("synchronizer.proto", fileDescriptor_f98557a2a36b798c) }

var fileDescriptor_f98557a2a36b798c = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x4d, 0xfa, 0x27, 0x43, 0xc5, 0xc1, 0xb4, 0xd4, 0xa5, 0x02, 0x16, 0x9f, 0x72,
	0xca, 0x01, 0x4e, 0x1c, 0xa3, 0x56, 0xad, 0x2a, 0x04, 0x42, 0x5e, 0xfa, 0x00, 0x1b, 0x6b, 0x9a,
	0x5a, 0x6c, 0xed, 0x8d, 0x3d, 0x89, 0x14, 0x1e, 0x87, 0x27, 0x45, 0xd9, 0x7f, 0x72, 0x76, 0x0b,
	0x28, 0x0b, 0xb7, 0xfd, 0xc6, 0xeb, 0x6f, 0x7e, 0xfe, 0xa4, 0x19, 0x60, 0x7e, 0x6d, 0xd4, 0x83,
	0xb3, 0x46, 0xff, 0x40, 0x37, 0xc9, 0x9d, 0x25, 0xcb, 0xf8, 0xdc, 0xa5, 0x2b, 0x4d, 0xeb, 0x49,
	0x9a, 0xeb, 0x49, 0x78, 0x2e, 0x12, 0x38, 0x9d, 0x7a, 0xaf, 0xe7, 0xe6, 0xab, 0xce, 0x31, 0xd3,
	0x06, 0x25, 0x2e, 0x96, 0xe8, 0x89, 0xbd, 0x82, 0x23, 0x95, 0x69, 0x34, 0x74, 0x7b, 0xc5, 0xa3,
	0x38, 0x1a, 0x8f, 0x64, 0xa3, 0xd9, 0x1b, 0x80, 0xbc, 0xfa, 0xfd, 0xf6, 0x8a, 0xef, 0xc5, 0xd1,
	0x78, 0x28, 0x83, 0x8a, 0xb8, 0x81, 0x17, 0x6d, 0xd3, 0x3c, 0x5b, 0x33, 0x0e, 0x87, 0x7e, 0xa9,
	0x14, 0x7a, 0x5f, 0x38, 0x1e, 0xc9, 0x5a, 0xb2, 0x97, 0x70, 0xe0, 0x30, 0xf5, 0xd6, 0x14, 0x66,
	0x23, 0x59, 0xa9, 0x0d, 0x9d, 0xc4, 0x95, 0xfd, 0x8e, 0xff, 0x99, 0xae, 0x6d, 0xda, 0x97, 0xee,
	0x5c, 0xe2, 0x5c, 0x7b, 0x42, 0x97, 0x2c, 0x67, 0x5e, 0x39, 0x3d, 0x43, 0x57, 0x13, 0x0a, 0x38,
	0xf6, 0x4d, 0xb1, 0xa1, 0xdc, 0xaa, 0x31, 0x06, 0xc3, 0x2f, 0xe9, 0x23, 0x56, 0xb6, 0xc5, 0xb7,
	0xf8, 0x04, 0x67, 0x4f, 0x99, 0xf6, 0x23, 0x9c, 0xc2, 0xc5, 0x9d, 0x71, 0xff, 0xc2, 0x28, 0x3e,
	0xc3, 0xf9, 0xd3, 0x16, 0xfd, 0x88, 0x10, 0x5e, 0x37, 0x26, 0xdf, 0xec, 0xa5, 0xcd, 0x32, 0x54,
	0xa4, 0xad, 0xf1, 0xbb, 0xe4, 0x16, 0xc3, 0xb3, 0xe0, 0x26, 0xdf, 0x8b, 0x07, 0xe3, 0x91, 0x0c,
	0x4b, 0x62, 0x01, 0x17, 0xbf, 0x6b, 0xd3, 0x8b, 0xbb, 0xdd, 0x72, 0xd0, 0x6d, 0xa9, 0xe1, 0xdd,
	0x9d, 0x69, 0x30, 0xaf, 0x9d, 0x7d, 0xec, 0xff, 0x3a, 0xd5, 0x7d, 0x5d, 0x50, 0x12, 0x97, 0xf0,
	0xf6, 0x4f, 0xad, 0x36, 0x2f, 0x6c, 0x99, 0x44, 0x5d, 0x93, 0x8f, 0x70, 0x76, 0x83, 0x54, 0xcf,
	0x40, 0x42, 0x29, 0x35, 0xd3, 0xb5, 0x3d, 0x41, 0x25, 0x63, 0x38, 0x41, 0x0a, 0x4e, 0xbb, 0x57,
	0xfb, 0xe5, 0xca, 0xe1, 0x30, 0x4b, 0x3d, 0x25, 0xb8, 0xe0, 0x83, 0x62, 0x52, 0x6b, 0x29, 0x7e,
	0x46, 0x70, 0x52, 0xb7, 0xb8, 0x46, 0x52, 0x0f, 0xbb, 0x64, 0xf8, 0x97, 0x1d, 0x50, 0x80, 0x52,
	0xea, 0x68, 0x4a, 0x75, 0xdb, 0x4a, 0x6e, 0x40, 0xed, 0xfd, 0xbd, 0x47, 0xe2, 0xc3, 0xe2, 0xa0,
	0x52, 0xec, 0x04, 0xf6, 0x95, 0x5d, 0x1a, 0xe2, 0xfb, 0x71, 0x34, 0x1e, 0xc8, 0x52, 0x88, 0x15,
	0xb0, 0x16, 0x63, 0xbf, 0x18, 0x1a, 0xf7, 0x92, 0xa6, 0x14, 0x61, 0x38, 0xc3, 0xad, 0x70, 0xde,
	0x3f, 0x87, 0xe3, 0x24, 0x58, 0xe3, 0xb3, 0x83, 0x62, 0xcf, 0x7f, 0xf8, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x7d, 0x21, 0xde, 0x4d, 0xfd, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SynchronizerClient is the client API for Synchronizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SynchronizerClient interface {
}

type synchronizerClient struct {
	cc *grpc.ClientConn
}

func NewSynchronizerClient(cc *grpc.ClientConn) SynchronizerClient {
	return &synchronizerClient{cc}
}

// SynchronizerServer is the server API for Synchronizer service.
type SynchronizerServer interface {
}

// UnimplementedSynchronizerServer can be embedded to have forward compatible implementations.
type UnimplementedSynchronizerServer struct {
}

func RegisterSynchronizerServer(s *grpc.Server, srv SynchronizerServer) {
	s.RegisterService(&_Synchronizer_serviceDesc, srv)
}

var _Synchronizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gravity.api.synchronizer.Synchronizer",
	HandlerType: (*SynchronizerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "synchronizer.proto",
}
