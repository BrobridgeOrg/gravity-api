// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: authenticator.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID      string `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
	AppName    string `protobuf:"bytes,2,opt,name=appName,proto3" json:"appName,omitempty"`
	Token      string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Properties []byte `protobuf:"bytes,4,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{0}
}

func (x *Entity) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Entity) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *Entity) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Entity) GetProperties() []byte {
	if x != nil {
		return x.Properties
	}
	return nil
}

type AuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID string `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateRequest) ProtoMessage() {}

func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *AuthenticateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthenticateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason  string  `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Entity  *Entity `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *AuthenticateReply) Reset() {
	*x = AuthenticateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateReply) ProtoMessage() {}

func (x *AuthenticateReply) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateReply.ProtoReflect.Descriptor instead.
func (*AuthenticateReply) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticateReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AuthenticateReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *AuthenticateReply) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type CreateEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *CreateEntityRequest) Reset() {
	*x = CreateEntityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntityRequest) ProtoMessage() {}

func (x *CreateEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntityRequest.ProtoReflect.Descriptor instead.
func (*CreateEntityRequest) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{3}
}

func (x *CreateEntityRequest) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type CreateEntityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason  string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *CreateEntityReply) Reset() {
	*x = CreateEntityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntityReply) ProtoMessage() {}

func (x *CreateEntityReply) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntityReply.ProtoReflect.Descriptor instead.
func (*CreateEntityReply) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{4}
}

func (x *CreateEntityReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateEntityReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type UpdateEntityTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID string `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UpdateEntityTokenRequest) Reset() {
	*x = UpdateEntityTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntityTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntityTokenRequest) ProtoMessage() {}

func (x *UpdateEntityTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntityTokenRequest.ProtoReflect.Descriptor instead.
func (*UpdateEntityTokenRequest) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateEntityTokenRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *UpdateEntityTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpdateEntityTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason  string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *UpdateEntityTokenReply) Reset() {
	*x = UpdateEntityTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntityTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntityTokenReply) ProtoMessage() {}

func (x *UpdateEntityTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntityTokenReply.ProtoReflect.Descriptor instead.
func (*UpdateEntityTokenReply) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateEntityTokenReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateEntityTokenReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type UpdateEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string  `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
	Entity *Entity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *UpdateEntityRequest) Reset() {
	*x = UpdateEntityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntityRequest) ProtoMessage() {}

func (x *UpdateEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntityRequest.ProtoReflect.Descriptor instead.
func (*UpdateEntityRequest) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateEntityRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *UpdateEntityRequest) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type UpdateEntityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason  string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *UpdateEntityReply) Reset() {
	*x = UpdateEntityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntityReply) ProtoMessage() {}

func (x *UpdateEntityReply) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntityReply.ProtoReflect.Descriptor instead.
func (*UpdateEntityReply) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateEntityReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateEntityReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type GetEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID string `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
}

func (x *GetEntityRequest) Reset() {
	*x = GetEntityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntityRequest) ProtoMessage() {}

func (x *GetEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntityRequest.ProtoReflect.Descriptor instead.
func (*GetEntityRequest) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{9}
}

func (x *GetEntityRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

type GetEntityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason  string  `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Entity  *Entity `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *GetEntityReply) Reset() {
	*x = GetEntityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntityReply) ProtoMessage() {}

func (x *GetEntityReply) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntityReply.ProtoReflect.Descriptor instead.
func (*GetEntityReply) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{10}
}

func (x *GetEntityReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetEntityReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *GetEntityReply) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type DeleteEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID string `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
}

func (x *DeleteEntityRequest) Reset() {
	*x = DeleteEntityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEntityRequest) ProtoMessage() {}

func (x *DeleteEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEntityRequest.ProtoReflect.Descriptor instead.
func (*DeleteEntityRequest) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteEntityRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

type DeleteEntityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason  string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *DeleteEntityReply) Reset() {
	*x = DeleteEntityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEntityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEntityReply) ProtoMessage() {}

func (x *DeleteEntityReply) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEntityReply.ProtoReflect.Descriptor instead.
func (*DeleteEntityReply) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteEntityReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteEntityReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type GetEntitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetEntitiesRequest) Reset() {
	*x = GetEntitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntitiesRequest) ProtoMessage() {}

func (x *GetEntitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntitiesRequest.ProtoReflect.Descriptor instead.
func (*GetEntitiesRequest) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{13}
}

func (x *GetEntitiesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetEntitiesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetEntitiesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool      `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason   string    `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Total    int32     `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Entities []*Entity `protobuf:"bytes,4,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *GetEntitiesReply) Reset() {
	*x = GetEntitiesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntitiesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntitiesReply) ProtoMessage() {}

func (x *GetEntitiesReply) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntitiesReply.ProtoReflect.Descriptor instead.
func (*GetEntitiesReply) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{14}
}

func (x *GetEntitiesReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetEntitiesReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *GetEntitiesReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetEntitiesReply) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

var File_authenticator_proto protoreflect.FileDescriptor

var file_authenticator_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72,
	0x22, 0x6e, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x22, 0x41, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72,
	0x61, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x50, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x45, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22,
	0x46, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4a, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x66, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44,
	0x12, 0x39, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x45, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x22, 0x28, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x22, 0x7d, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x39, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x2b, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x22, 0x45, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22,
	0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x3d, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42,
	0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authenticator_proto_rawDescOnce sync.Once
	file_authenticator_proto_rawDescData = file_authenticator_proto_rawDesc
)

func file_authenticator_proto_rawDescGZIP() []byte {
	file_authenticator_proto_rawDescOnce.Do(func() {
		file_authenticator_proto_rawDescData = protoimpl.X.CompressGZIP(file_authenticator_proto_rawDescData)
	})
	return file_authenticator_proto_rawDescData
}

var file_authenticator_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_authenticator_proto_goTypes = []interface{}{
	(*Entity)(nil),                   // 0: gravity.api.authenticator.Entity
	(*AuthenticateRequest)(nil),      // 1: gravity.api.authenticator.AuthenticateRequest
	(*AuthenticateReply)(nil),        // 2: gravity.api.authenticator.AuthenticateReply
	(*CreateEntityRequest)(nil),      // 3: gravity.api.authenticator.CreateEntityRequest
	(*CreateEntityReply)(nil),        // 4: gravity.api.authenticator.CreateEntityReply
	(*UpdateEntityTokenRequest)(nil), // 5: gravity.api.authenticator.UpdateEntityTokenRequest
	(*UpdateEntityTokenReply)(nil),   // 6: gravity.api.authenticator.UpdateEntityTokenReply
	(*UpdateEntityRequest)(nil),      // 7: gravity.api.authenticator.UpdateEntityRequest
	(*UpdateEntityReply)(nil),        // 8: gravity.api.authenticator.UpdateEntityReply
	(*GetEntityRequest)(nil),         // 9: gravity.api.authenticator.GetEntityRequest
	(*GetEntityReply)(nil),           // 10: gravity.api.authenticator.GetEntityReply
	(*DeleteEntityRequest)(nil),      // 11: gravity.api.authenticator.DeleteEntityRequest
	(*DeleteEntityReply)(nil),        // 12: gravity.api.authenticator.DeleteEntityReply
	(*GetEntitiesRequest)(nil),       // 13: gravity.api.authenticator.GetEntitiesRequest
	(*GetEntitiesReply)(nil),         // 14: gravity.api.authenticator.GetEntitiesReply
}
var file_authenticator_proto_depIdxs = []int32{
	0, // 0: gravity.api.authenticator.AuthenticateReply.entity:type_name -> gravity.api.authenticator.Entity
	0, // 1: gravity.api.authenticator.CreateEntityRequest.entity:type_name -> gravity.api.authenticator.Entity
	0, // 2: gravity.api.authenticator.UpdateEntityRequest.entity:type_name -> gravity.api.authenticator.Entity
	0, // 3: gravity.api.authenticator.GetEntityReply.entity:type_name -> gravity.api.authenticator.Entity
	0, // 4: gravity.api.authenticator.GetEntitiesReply.entities:type_name -> gravity.api.authenticator.Entity
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_authenticator_proto_init() }
func file_authenticator_proto_init() {
	if File_authenticator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authenticator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_authenticator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateRequest); i {
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
		file_authenticator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateReply); i {
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
		file_authenticator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntityRequest); i {
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
		file_authenticator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntityReply); i {
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
		file_authenticator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntityTokenRequest); i {
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
		file_authenticator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntityTokenReply); i {
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
		file_authenticator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntityRequest); i {
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
		file_authenticator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntityReply); i {
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
		file_authenticator_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntityRequest); i {
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
		file_authenticator_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntityReply); i {
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
		file_authenticator_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEntityRequest); i {
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
		file_authenticator_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEntityReply); i {
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
		file_authenticator_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntitiesRequest); i {
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
		file_authenticator_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntitiesReply); i {
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
			RawDescriptor: file_authenticator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_authenticator_proto_goTypes,
		DependencyIndexes: file_authenticator_proto_depIdxs,
		MessageInfos:      file_authenticator_proto_msgTypes,
	}.Build()
	File_authenticator_proto = out.File
	file_authenticator_proto_rawDesc = nil
	file_authenticator_proto_goTypes = nil
	file_authenticator_proto_depIdxs = nil
}
