// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: panel.proto

package panelv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Panel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Panel) Reset() {
	*x = Panel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Panel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Panel) ProtoMessage() {}

func (x *Panel) ProtoReflect() protoreflect.Message {
	mi := &file_panel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Panel.ProtoReflect.Descriptor instead.
func (*Panel) Descriptor() ([]byte, []int) {
	return file_panel_proto_rawDescGZIP(), []int{0}
}

func (x *Panel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Panel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Panel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Panel) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Panel) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type PanelMutable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Description *string `protobuf:"bytes,2,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *PanelMutable) Reset() {
	*x = PanelMutable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PanelMutable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PanelMutable) ProtoMessage() {}

func (x *PanelMutable) ProtoReflect() protoreflect.Message {
	mi := &file_panel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PanelMutable.ProtoReflect.Descriptor instead.
func (*PanelMutable) Descriptor() ([]byte, []int) {
	return file_panel_proto_rawDescGZIP(), []int{1}
}

func (x *PanelMutable) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *PanelMutable) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type CreatePanelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *PanelMutable `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreatePanelRequest) Reset() {
	*x = CreatePanelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePanelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePanelRequest) ProtoMessage() {}

func (x *CreatePanelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_panel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePanelRequest.ProtoReflect.Descriptor instead.
func (*CreatePanelRequest) Descriptor() ([]byte, []int) {
	return file_panel_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePanelRequest) GetData() *PanelMutable {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetPanelByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPanelByIdRequest) Reset() {
	*x = GetPanelByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPanelByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPanelByIdRequest) ProtoMessage() {}

func (x *GetPanelByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_panel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPanelByIdRequest.ProtoReflect.Descriptor instead.
func (*GetPanelByIdRequest) Descriptor() ([]byte, []int) {
	return file_panel_proto_rawDescGZIP(), []int{3}
}

func (x *GetPanelByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPanelByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPanelByNameRequest) Reset() {
	*x = GetPanelByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPanelByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPanelByNameRequest) ProtoMessage() {}

func (x *GetPanelByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_panel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPanelByNameRequest.ProtoReflect.Descriptor instead.
func (*GetPanelByNameRequest) Descriptor() ([]byte, []int) {
	return file_panel_proto_rawDescGZIP(), []int{4}
}

func (x *GetPanelByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdatePanelByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data *PanelMutable `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdatePanelByIdRequest) Reset() {
	*x = UpdatePanelByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panel_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePanelByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePanelByIdRequest) ProtoMessage() {}

func (x *UpdatePanelByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_panel_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePanelByIdRequest.ProtoReflect.Descriptor instead.
func (*UpdatePanelByIdRequest) Descriptor() ([]byte, []int) {
	return file_panel_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePanelByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePanelByIdRequest) GetData() *PanelMutable {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdatePanelByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data *PanelMutable `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdatePanelByNameRequest) Reset() {
	*x = UpdatePanelByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panel_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePanelByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePanelByNameRequest) ProtoMessage() {}

func (x *UpdatePanelByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_panel_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePanelByNameRequest.ProtoReflect.Descriptor instead.
func (*UpdatePanelByNameRequest) Descriptor() ([]byte, []int) {
	return file_panel_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePanelByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePanelByNameRequest) GetData() *PanelMutable {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeletePanelByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePanelByIdRequest) Reset() {
	*x = DeletePanelByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panel_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePanelByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePanelByIdRequest) ProtoMessage() {}

func (x *DeletePanelByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_panel_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePanelByIdRequest.ProtoReflect.Descriptor instead.
func (*DeletePanelByIdRequest) Descriptor() ([]byte, []int) {
	return file_panel_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePanelByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePanelByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeletePanelByNameRequest) Reset() {
	*x = DeletePanelByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panel_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePanelByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePanelByNameRequest) ProtoMessage() {}

func (x *DeletePanelByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_panel_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePanelByNameRequest.ProtoReflect.Descriptor instead.
func (*DeletePanelByNameRequest) Descriptor() ([]byte, []int) {
	return file_panel_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePanelByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Kafka Event Schema
type PanelEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Data *Panel `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PanelEvent) Reset() {
	*x = PanelEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panel_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PanelEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PanelEvent) ProtoMessage() {}

func (x *PanelEvent) ProtoReflect() protoreflect.Message {
	mi := &file_panel_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PanelEvent.ProtoReflect.Descriptor instead.
func (*PanelEvent) Descriptor() ([]byte, []int) {
	return file_panel_proto_rawDescGZIP(), []int{9}
}

func (x *PanelEvent) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PanelEvent) GetData() *Panel {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_panel_proto protoreflect.FileDescriptor

var file_panel_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70,
	0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a,
	0x05, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x67, 0x0a, 0x0c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x4d, 0x75, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x61, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x6e, 0x65, 0x6c, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x61, 0x6e,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x4d, 0x75, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2e, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x6e, 0x65,
	0x6c, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x0a, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x61, 0x6e, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0xd4, 0x04, 0x0a, 0x0c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x6e, 0x65, 0x6c,
	0x12, 0x23, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x22, 0x00, 0x12,
	0x4a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x12, 0x24, 0x2e, 0x70, 0x61,
	0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x2e,
	0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x22, 0x00, 0x12,
	0x50, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x12, 0x27,
	0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x22,
	0x00, 0x12, 0x58, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x6e, 0x65, 0x6c,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x6e, 0x65, 0x6c, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x12, 0x27, 0x2e, 0x70, 0x61, 0x6e,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x58, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x29, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x61, 0x6e, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x6e, 0x65, 0x6c,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_panel_proto_rawDescOnce sync.Once
	file_panel_proto_rawDescData = file_panel_proto_rawDesc
)

func file_panel_proto_rawDescGZIP() []byte {
	file_panel_proto_rawDescOnce.Do(func() {
		file_panel_proto_rawDescData = protoimpl.X.CompressGZIP(file_panel_proto_rawDescData)
	})
	return file_panel_proto_rawDescData
}

var file_panel_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_panel_proto_goTypes = []interface{}{
	(*Panel)(nil),                    // 0: panels.panel.v1.Panel
	(*PanelMutable)(nil),             // 1: panels.panel.v1.PanelMutable
	(*CreatePanelRequest)(nil),       // 2: panels.panel.v1.CreatePanelRequest
	(*GetPanelByIdRequest)(nil),      // 3: panels.panel.v1.GetPanelByIdRequest
	(*GetPanelByNameRequest)(nil),    // 4: panels.panel.v1.GetPanelByNameRequest
	(*UpdatePanelByIdRequest)(nil),   // 5: panels.panel.v1.UpdatePanelByIdRequest
	(*UpdatePanelByNameRequest)(nil), // 6: panels.panel.v1.UpdatePanelByNameRequest
	(*DeletePanelByIdRequest)(nil),   // 7: panels.panel.v1.DeletePanelByIdRequest
	(*DeletePanelByNameRequest)(nil), // 8: panels.panel.v1.DeletePanelByNameRequest
	(*PanelEvent)(nil),               // 9: panels.panel.v1.PanelEvent
	(*timestamppb.Timestamp)(nil),    // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),            // 11: google.protobuf.Empty
}
var file_panel_proto_depIdxs = []int32{
	10, // 0: panels.panel.v1.Panel.created_at:type_name -> google.protobuf.Timestamp
	10, // 1: panels.panel.v1.Panel.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 2: panels.panel.v1.CreatePanelRequest.data:type_name -> panels.panel.v1.PanelMutable
	1,  // 3: panels.panel.v1.UpdatePanelByIdRequest.data:type_name -> panels.panel.v1.PanelMutable
	1,  // 4: panels.panel.v1.UpdatePanelByNameRequest.data:type_name -> panels.panel.v1.PanelMutable
	0,  // 5: panels.panel.v1.PanelEvent.data:type_name -> panels.panel.v1.Panel
	2,  // 6: panels.panel.v1.PanelService.CreatePanel:input_type -> panels.panel.v1.CreatePanelRequest
	3,  // 7: panels.panel.v1.PanelService.GetPanel:input_type -> panels.panel.v1.GetPanelByIdRequest
	4,  // 8: panels.panel.v1.PanelService.GetPanelByName:input_type -> panels.panel.v1.GetPanelByNameRequest
	5,  // 9: panels.panel.v1.PanelService.UpdatePanel:input_type -> panels.panel.v1.UpdatePanelByIdRequest
	6,  // 10: panels.panel.v1.PanelService.UpdatePanelByName:input_type -> panels.panel.v1.UpdatePanelByNameRequest
	7,  // 11: panels.panel.v1.PanelService.DeletePanel:input_type -> panels.panel.v1.DeletePanelByIdRequest
	8,  // 12: panels.panel.v1.PanelService.DeletePanelByName:input_type -> panels.panel.v1.DeletePanelByNameRequest
	0,  // 13: panels.panel.v1.PanelService.CreatePanel:output_type -> panels.panel.v1.Panel
	0,  // 14: panels.panel.v1.PanelService.GetPanel:output_type -> panels.panel.v1.Panel
	0,  // 15: panels.panel.v1.PanelService.GetPanelByName:output_type -> panels.panel.v1.Panel
	0,  // 16: panels.panel.v1.PanelService.UpdatePanel:output_type -> panels.panel.v1.Panel
	0,  // 17: panels.panel.v1.PanelService.UpdatePanelByName:output_type -> panels.panel.v1.Panel
	11, // 18: panels.panel.v1.PanelService.DeletePanel:output_type -> google.protobuf.Empty
	11, // 19: panels.panel.v1.PanelService.DeletePanelByName:output_type -> google.protobuf.Empty
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_panel_proto_init() }
func file_panel_proto_init() {
	if File_panel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_panel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Panel); i {
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
		file_panel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PanelMutable); i {
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
		file_panel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePanelRequest); i {
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
		file_panel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPanelByIdRequest); i {
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
		file_panel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPanelByNameRequest); i {
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
		file_panel_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePanelByIdRequest); i {
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
		file_panel_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePanelByNameRequest); i {
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
		file_panel_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePanelByIdRequest); i {
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
		file_panel_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePanelByNameRequest); i {
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
		file_panel_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PanelEvent); i {
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
	file_panel_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_panel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_panel_proto_goTypes,
		DependencyIndexes: file_panel_proto_depIdxs,
		MessageInfos:      file_panel_proto_msgTypes,
	}.Build()
	File_panel_proto = out.File
	file_panel_proto_rawDesc = nil
	file_panel_proto_goTypes = nil
	file_panel_proto_depIdxs = nil
}