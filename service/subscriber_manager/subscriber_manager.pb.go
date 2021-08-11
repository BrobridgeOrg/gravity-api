// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: subscriber_manager.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubscriberType int32

const (
	SubscriberType_TRANSMITTER SubscriberType = 0
	SubscriberType_EXPORTER    SubscriberType = 1
)

// Enum value maps for SubscriberType.
var (
	SubscriberType_name = map[int32]string{
		0: "TRANSMITTER",
		1: "EXPORTER",
	}
	SubscriberType_value = map[string]int32{
		"TRANSMITTER": 0,
		"EXPORTER":    1,
	}
)

func (x SubscriberType) Enum() *SubscriberType {
	p := new(SubscriberType)
	*p = x
	return p
}

func (x SubscriberType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubscriberType) Descriptor() protoreflect.EnumDescriptor {
	return file_subscriber_manager_proto_enumTypes[0].Descriptor()
}

func (SubscriberType) Type() protoreflect.EnumType {
	return &file_subscriber_manager_proto_enumTypes[0]
}

func (x SubscriberType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubscriberType.Descriptor instead.
func (SubscriberType) EnumDescriptor() ([]byte, []int) {
	return file_subscriber_manager_proto_rawDescGZIP(), []int{0}
}

type RegisterSubscriberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberID string         `protobuf:"bytes,1,opt,name=subscriberID,proto3" json:"subscriberID,omitempty"`
	Name         string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type         SubscriberType `protobuf:"varint,3,opt,name=type,proto3,enum=gravity.api.subscriber_manager.SubscriberType" json:"type,omitempty"`
	Component    string         `protobuf:"bytes,4,opt,name=component,proto3" json:"component,omitempty"`
	AppID        string         `protobuf:"bytes,5,opt,name=appID,proto3" json:"appID,omitempty"`
	Token        []byte         `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RegisterSubscriberRequest) Reset() {
	*x = RegisterSubscriberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriber_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterSubscriberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterSubscriberRequest) ProtoMessage() {}

func (x *RegisterSubscriberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterSubscriberRequest.ProtoReflect.Descriptor instead.
func (*RegisterSubscriberRequest) Descriptor() ([]byte, []int) {
	return file_subscriber_manager_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterSubscriberRequest) GetSubscriberID() string {
	if x != nil {
		return x.SubscriberID
	}
	return ""
}

func (x *RegisterSubscriberRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterSubscriberRequest) GetType() SubscriberType {
	if x != nil {
		return x.Type
	}
	return SubscriberType_TRANSMITTER
}

func (x *RegisterSubscriberRequest) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *RegisterSubscriberRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *RegisterSubscriberRequest) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

type RegisterSubscriberReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason  string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RegisterSubscriberReply) Reset() {
	*x = RegisterSubscriberReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriber_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterSubscriberReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterSubscriberReply) ProtoMessage() {}

func (x *RegisterSubscriberReply) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterSubscriberReply.ProtoReflect.Descriptor instead.
func (*RegisterSubscriberReply) Descriptor() ([]byte, []int) {
	return file_subscriber_manager_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterSubscriberReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RegisterSubscriberReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *RegisterSubscriberReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UnregisterSubscriberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberID string `protobuf:"bytes,1,opt,name=subscriberID,proto3" json:"subscriberID,omitempty"`
}

func (x *UnregisterSubscriberRequest) Reset() {
	*x = UnregisterSubscriberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriber_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregisterSubscriberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterSubscriberRequest) ProtoMessage() {}

func (x *UnregisterSubscriberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterSubscriberRequest.ProtoReflect.Descriptor instead.
func (*UnregisterSubscriberRequest) Descriptor() ([]byte, []int) {
	return file_subscriber_manager_proto_rawDescGZIP(), []int{2}
}

func (x *UnregisterSubscriberRequest) GetSubscriberID() string {
	if x != nil {
		return x.SubscriberID
	}
	return ""
}

type UnregisterSubscriberReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason  string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *UnregisterSubscriberReply) Reset() {
	*x = UnregisterSubscriberReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriber_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregisterSubscriberReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterSubscriberReply) ProtoMessage() {}

func (x *UnregisterSubscriberReply) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterSubscriberReply.ProtoReflect.Descriptor instead.
func (*UnregisterSubscriberReply) Descriptor() ([]byte, []int) {
	return file_subscriber_manager_proto_rawDescGZIP(), []int{3}
}

func (x *UnregisterSubscriberReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UnregisterSubscriberReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberID string `protobuf:"bytes,1,opt,name=subscriberID,proto3" json:"subscriberID,omitempty"`
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriber_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_subscriber_manager_proto_rawDescGZIP(), []int{4}
}

func (x *HealthCheckRequest) GetSubscriberID() string {
	if x != nil {
		return x.SubscriberID
	}
	return ""
}

type HealthCheckReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason  string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *HealthCheckReply) Reset() {
	*x = HealthCheckReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriber_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckReply) ProtoMessage() {}

func (x *HealthCheckReply) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckReply.ProtoReflect.Descriptor instead.
func (*HealthCheckReply) Descriptor() ([]byte, []int) {
	return file_subscriber_manager_proto_rawDescGZIP(), []int{5}
}

func (x *HealthCheckReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *HealthCheckReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type Subscriber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberID string                 `protobuf:"bytes,1,opt,name=subscriberID,proto3" json:"subscriberID,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type         SubscriberType         `protobuf:"varint,3,opt,name=type,proto3,enum=gravity.api.subscriber_manager.SubscriberType" json:"type,omitempty"`
	Component    string                 `protobuf:"bytes,4,opt,name=component,proto3" json:"component,omitempty"`
	LastCheck    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=lastCheck,proto3" json:"lastCheck,omitempty"`
	AppID        string                 `protobuf:"bytes,6,opt,name=appID,proto3" json:"appID,omitempty"`
	AccessKey    string                 `protobuf:"bytes,7,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
	Permissions  []string               `protobuf:"bytes,8,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Collections  []string               `protobuf:"bytes,9,rep,name=collections,proto3" json:"collections,omitempty"`
}

func (x *Subscriber) Reset() {
	*x = Subscriber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriber_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscriber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscriber) ProtoMessage() {}

func (x *Subscriber) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscriber.ProtoReflect.Descriptor instead.
func (*Subscriber) Descriptor() ([]byte, []int) {
	return file_subscriber_manager_proto_rawDescGZIP(), []int{6}
}

func (x *Subscriber) GetSubscriberID() string {
	if x != nil {
		return x.SubscriberID
	}
	return ""
}

func (x *Subscriber) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Subscriber) GetType() SubscriberType {
	if x != nil {
		return x.Type
	}
	return SubscriberType_TRANSMITTER
}

func (x *Subscriber) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *Subscriber) GetLastCheck() *timestamppb.Timestamp {
	if x != nil {
		return x.LastCheck
	}
	return nil
}

func (x *Subscriber) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Subscriber) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Subscriber) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Subscriber) GetCollections() []string {
	if x != nil {
		return x.Collections
	}
	return nil
}

type GetSubscribersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSubscribersRequest) Reset() {
	*x = GetSubscribersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriber_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscribersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscribersRequest) ProtoMessage() {}

func (x *GetSubscribersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscribersRequest.ProtoReflect.Descriptor instead.
func (*GetSubscribersRequest) Descriptor() ([]byte, []int) {
	return file_subscriber_manager_proto_rawDescGZIP(), []int{7}
}

type GetSubscribersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason      string        `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Subscribers []*Subscriber `protobuf:"bytes,3,rep,name=subscribers,proto3" json:"subscribers,omitempty"`
}

func (x *GetSubscribersReply) Reset() {
	*x = GetSubscribersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriber_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscribersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscribersReply) ProtoMessage() {}

func (x *GetSubscribersReply) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscribersReply.ProtoReflect.Descriptor instead.
func (*GetSubscribersReply) Descriptor() ([]byte, []int) {
	return file_subscriber_manager_proto_rawDescGZIP(), []int{8}
}

func (x *GetSubscribersReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetSubscribersReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *GetSubscribersReply) GetSubscribers() []*Subscriber {
	if x != nil {
		return x.Subscribers
	}
	return nil
}

type SubscribeToCollectionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberID string   `protobuf:"bytes,1,opt,name=subscriberID,proto3" json:"subscriberID,omitempty"`
	Collections  []string `protobuf:"bytes,2,rep,name=collections,proto3" json:"collections,omitempty"`
}

func (x *SubscribeToCollectionsRequest) Reset() {
	*x = SubscribeToCollectionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriber_manager_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToCollectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToCollectionsRequest) ProtoMessage() {}

func (x *SubscribeToCollectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_manager_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToCollectionsRequest.ProtoReflect.Descriptor instead.
func (*SubscribeToCollectionsRequest) Descriptor() ([]byte, []int) {
	return file_subscriber_manager_proto_rawDescGZIP(), []int{9}
}

func (x *SubscribeToCollectionsRequest) GetSubscriberID() string {
	if x != nil {
		return x.SubscriberID
	}
	return ""
}

func (x *SubscribeToCollectionsRequest) GetCollections() []string {
	if x != nil {
		return x.Collections
	}
	return nil
}

type SubscribeToCollectionsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason      string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Collections []string `protobuf:"bytes,3,rep,name=collections,proto3" json:"collections,omitempty"`
}

func (x *SubscribeToCollectionsReply) Reset() {
	*x = SubscribeToCollectionsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriber_manager_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToCollectionsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToCollectionsReply) ProtoMessage() {}

func (x *SubscribeToCollectionsReply) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_manager_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToCollectionsReply.ProtoReflect.Descriptor instead.
func (*SubscribeToCollectionsReply) Descriptor() ([]byte, []int) {
	return file_subscriber_manager_proto_rawDescGZIP(), []int{10}
}

func (x *SubscribeToCollectionsReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SubscribeToCollectionsReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *SubscribeToCollectionsReply) GetCollections() []string {
	if x != nil {
		return x.Collections
	}
	return nil
}

var File_subscriber_manager_proto protoreflect.FileDescriptor

var file_subscriber_manager_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x72, 0x61, 0x76,
	0x69, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x19,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x42, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2e, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x61, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x41, 0x0a, 0x1b, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x49, 0x44, 0x22, 0x4d, 0x0a, 0x19, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x44, 0x22, 0x44,
	0x0a, 0x10, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x22, 0xd8, 0x02, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x67, 0x72, 0x61, 0x76,
	0x69, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x95, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73,
	0x22, 0x65, 0x0a, 0x1d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x71, 0x0a, 0x1b, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x2f, 0x0a, 0x0e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x45, 0x58, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x52, 0x10, 0x01, 0x42, 0x04, 0x5a, 0x02, 0x2e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subscriber_manager_proto_rawDescOnce sync.Once
	file_subscriber_manager_proto_rawDescData = file_subscriber_manager_proto_rawDesc
)

func file_subscriber_manager_proto_rawDescGZIP() []byte {
	file_subscriber_manager_proto_rawDescOnce.Do(func() {
		file_subscriber_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscriber_manager_proto_rawDescData)
	})
	return file_subscriber_manager_proto_rawDescData
}

var file_subscriber_manager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_subscriber_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_subscriber_manager_proto_goTypes = []interface{}{
	(SubscriberType)(0),                   // 0: gravity.api.subscriber_manager.SubscriberType
	(*RegisterSubscriberRequest)(nil),     // 1: gravity.api.subscriber_manager.RegisterSubscriberRequest
	(*RegisterSubscriberReply)(nil),       // 2: gravity.api.subscriber_manager.RegisterSubscriberReply
	(*UnregisterSubscriberRequest)(nil),   // 3: gravity.api.subscriber_manager.UnregisterSubscriberRequest
	(*UnregisterSubscriberReply)(nil),     // 4: gravity.api.subscriber_manager.UnregisterSubscriberReply
	(*HealthCheckRequest)(nil),            // 5: gravity.api.subscriber_manager.HealthCheckRequest
	(*HealthCheckReply)(nil),              // 6: gravity.api.subscriber_manager.HealthCheckReply
	(*Subscriber)(nil),                    // 7: gravity.api.subscriber_manager.Subscriber
	(*GetSubscribersRequest)(nil),         // 8: gravity.api.subscriber_manager.GetSubscribersRequest
	(*GetSubscribersReply)(nil),           // 9: gravity.api.subscriber_manager.GetSubscribersReply
	(*SubscribeToCollectionsRequest)(nil), // 10: gravity.api.subscriber_manager.SubscribeToCollectionsRequest
	(*SubscribeToCollectionsReply)(nil),   // 11: gravity.api.subscriber_manager.SubscribeToCollectionsReply
	(*timestamppb.Timestamp)(nil),         // 12: google.protobuf.Timestamp
}
var file_subscriber_manager_proto_depIdxs = []int32{
	0,  // 0: gravity.api.subscriber_manager.RegisterSubscriberRequest.type:type_name -> gravity.api.subscriber_manager.SubscriberType
	0,  // 1: gravity.api.subscriber_manager.Subscriber.type:type_name -> gravity.api.subscriber_manager.SubscriberType
	12, // 2: gravity.api.subscriber_manager.Subscriber.lastCheck:type_name -> google.protobuf.Timestamp
	7,  // 3: gravity.api.subscriber_manager.GetSubscribersReply.subscribers:type_name -> gravity.api.subscriber_manager.Subscriber
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_subscriber_manager_proto_init() }
func file_subscriber_manager_proto_init() {
	if File_subscriber_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_subscriber_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterSubscriberRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subscriber_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterSubscriberReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subscriber_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregisterSubscriberRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subscriber_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregisterSubscriberReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subscriber_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subscriber_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subscriber_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscriber); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subscriber_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscribersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subscriber_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscribersReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subscriber_manager_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToCollectionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subscriber_manager_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToCollectionsReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_subscriber_manager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_subscriber_manager_proto_goTypes,
		DependencyIndexes: file_subscriber_manager_proto_depIdxs,
		EnumInfos:         file_subscriber_manager_proto_enumTypes,
		MessageInfos:      file_subscriber_manager_proto_msgTypes,
	}.Build()
	File_subscriber_manager_proto = out.File
	file_subscriber_manager_proto_rawDesc = nil
	file_subscriber_manager_proto_goTypes = nil
	file_subscriber_manager_proto_depIdxs = nil
}
