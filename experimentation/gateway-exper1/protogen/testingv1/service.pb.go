// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: testingv1/service.proto

package testingv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ItemStatus int32

const (
	ItemStatus_ITEM_STATUS_UNSPECIFIED ItemStatus = 0
	ItemStatus_ITEM_STATUS_PENDING     ItemStatus = 1
	ItemStatus_ITEM_STATUS_REJECTED    ItemStatus = 2
	ItemStatus_ITEM_STATUS_APPROVED    ItemStatus = 3
)

// Enum value maps for ItemStatus.
var (
	ItemStatus_name = map[int32]string{
		0: "ITEM_STATUS_UNSPECIFIED",
		1: "ITEM_STATUS_PENDING",
		2: "ITEM_STATUS_REJECTED",
		3: "ITEM_STATUS_APPROVED",
	}
	ItemStatus_value = map[string]int32{
		"ITEM_STATUS_UNSPECIFIED": 0,
		"ITEM_STATUS_PENDING":     1,
		"ITEM_STATUS_REJECTED":    2,
		"ITEM_STATUS_APPROVED":    3,
	}
)

func (x ItemStatus) Enum() *ItemStatus {
	p := new(ItemStatus)
	*p = x
	return p
}

func (x ItemStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ItemStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_testingv1_service_proto_enumTypes[0].Descriptor()
}

func (ItemStatus) Type() protoreflect.EnumType {
	return &file_testingv1_service_proto_enumTypes[0]
}

func (x ItemStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ItemStatus.Descriptor instead.
func (ItemStatus) EnumDescriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{0}
}

type ItemStateEvent_Type int32

const (
	ItemStateEvent_TYPE_UNSPECIFIED ItemStateEvent_Type = 0
	ItemStateEvent_TYPE_CREATED     ItemStateEvent_Type = 1
	ItemStateEvent_TYPE_UPDATED     ItemStateEvent_Type = 2
	ItemStateEvent_TYPE_DELETED     ItemStateEvent_Type = 3
)

// Enum value maps for ItemStateEvent_Type.
var (
	ItemStateEvent_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_CREATED",
		2: "TYPE_UPDATED",
		3: "TYPE_DELETED",
	}
	ItemStateEvent_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_CREATED":     1,
		"TYPE_UPDATED":     2,
		"TYPE_DELETED":     3,
	}
)

func (x ItemStateEvent_Type) Enum() *ItemStateEvent_Type {
	p := new(ItemStateEvent_Type)
	*p = x
	return p
}

func (x ItemStateEvent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ItemStateEvent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_testingv1_service_proto_enumTypes[1].Descriptor()
}

func (ItemStateEvent_Type) Type() protoreflect.EnumType {
	return &file_testingv1_service_proto_enumTypes[1]
}

func (x ItemStateEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ItemStateEvent_Type.Descriptor instead.
func (ItemStateEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{1, 0}
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status    ItemStatus `protobuf:"varint,2,opt,name=status,proto3,enum=testingv1.ItemStatus" json:"status,omitempty"`
	Title     string     `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	CreatedAt int64      `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingv1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_testingv1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetStatus() ItemStatus {
	if x != nil {
		return x.Status
	}
	return ItemStatus_ITEM_STATUS_UNSPECIFIED
}

func (x *Item) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Item) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type ItemStateEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    ItemStateEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=testingv1.ItemStateEvent_Type" json:"type,omitempty"`
	Payload *Item               `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ItemStateEvent) Reset() {
	*x = ItemStateEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingv1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemStateEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemStateEvent) ProtoMessage() {}

func (x *ItemStateEvent) ProtoReflect() protoreflect.Message {
	mi := &file_testingv1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemStateEvent.ProtoReflect.Descriptor instead.
func (*ItemStateEvent) Descriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ItemStateEvent) GetType() ItemStateEvent_Type {
	if x != nil {
		return x.Type
	}
	return ItemStateEvent_TYPE_UNSPECIFIED
}

func (x *ItemStateEvent) GetPayload() *Item {
	if x != nil {
		return x.Payload
	}
	return nil
}

type GetItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *GetItemRequest) Reset() {
	*x = GetItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingv1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemRequest) ProtoMessage() {}

func (x *GetItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testingv1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemRequest.ProtoReflect.Descriptor instead.
func (*GetItemRequest) Descriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetItemRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type GetItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Item `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetItemResponse) Reset() {
	*x = GetItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingv1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemResponse) ProtoMessage() {}

func (x *GetItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testingv1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemResponse.ProtoReflect.Descriptor instead.
func (*GetItemResponse) Descriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetItemResponse) GetData() *Item {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetItemsRequest) Reset() {
	*x = GetItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingv1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsRequest) ProtoMessage() {}

func (x *GetItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testingv1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsRequest.ProtoReflect.Descriptor instead.
func (*GetItemsRequest) Descriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{4}
}

type GetItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Item `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetItemsResponse) Reset() {
	*x = GetItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingv1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsResponse) ProtoMessage() {}

func (x *GetItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testingv1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsResponse.ProtoReflect.Descriptor instead.
func (*GetItemsResponse) Descriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetItemsResponse) GetData() []*Item {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Item `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateItemRequest) Reset() {
	*x = UpdateItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingv1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemRequest) ProtoMessage() {}

func (x *UpdateItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testingv1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateItemRequest) Descriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateItemRequest) GetData() *Item {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Item `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateItemResponse) Reset() {
	*x = UpdateItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingv1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemResponse) ProtoMessage() {}

func (x *UpdateItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testingv1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemResponse.ProtoReflect.Descriptor instead.
func (*UpdateItemResponse) Descriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateItemResponse) GetData() *Item {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *DeleteItemRequest) Reset() {
	*x = DeleteItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingv1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemRequest) ProtoMessage() {}

func (x *DeleteItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testingv1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteItemRequest) Descriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteItemRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type DeleteItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteItemResponse) Reset() {
	*x = DeleteItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingv1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemResponse) ProtoMessage() {}

func (x *DeleteItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testingv1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemResponse.ProtoReflect.Descriptor instead.
func (*DeleteItemResponse) Descriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{9}
}

type CreateItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Item `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateItemRequest) Reset() {
	*x = CreateItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingv1_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateItemRequest) ProtoMessage() {}

func (x *CreateItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testingv1_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateItemRequest.ProtoReflect.Descriptor instead.
func (*CreateItemRequest) Descriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{10}
}

func (x *CreateItemRequest) GetData() *Item {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Item `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateItemResponse) Reset() {
	*x = CreateItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testingv1_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateItemResponse) ProtoMessage() {}

func (x *CreateItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testingv1_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateItemResponse.ProtoReflect.Descriptor instead.
func (*CreateItemResponse) Descriptor() ([]byte, []int) {
	return file_testingv1_service_proto_rawDescGZIP(), []int{11}
}

func (x *CreateItemResponse) GetData() *Item {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_testingv1_service_proto protoreflect.FileDescriptor

var file_testingv1_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc3,
	0x01, 0x0a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x52, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x45, 0x44, 0x10, 0x03, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22,
	0x36, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x38, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x2a, 0x76, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1b, 0x0a, 0x17, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13,
	0x49, 0x54, 0x45, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x18, 0x0a, 0x14, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41,
	0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x03, 0x32, 0xef, 0x03, 0x0a, 0x0b, 0x54, 0x65,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x53, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x1a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x08, 0x12, 0x06, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x69, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x10, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x69, 0x64, 0x7d, 0x12, 0x63, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x1c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x2f, 0x7b, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5f, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x06, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x5c, 0x5a, 0x5a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x78, 0x6f, 0x6c, 0x61,
	0x6e, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x65, 0x78, 0x70, 0x65, 0x72, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x3b,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_testingv1_service_proto_rawDescOnce sync.Once
	file_testingv1_service_proto_rawDescData = file_testingv1_service_proto_rawDesc
)

func file_testingv1_service_proto_rawDescGZIP() []byte {
	file_testingv1_service_proto_rawDescOnce.Do(func() {
		file_testingv1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_testingv1_service_proto_rawDescData)
	})
	return file_testingv1_service_proto_rawDescData
}

var file_testingv1_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_testingv1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_testingv1_service_proto_goTypes = []interface{}{
	(ItemStatus)(0),            // 0: testingv1.ItemStatus
	(ItemStateEvent_Type)(0),   // 1: testingv1.ItemStateEvent.Type
	(*Item)(nil),               // 2: testingv1.Item
	(*ItemStateEvent)(nil),     // 3: testingv1.ItemStateEvent
	(*GetItemRequest)(nil),     // 4: testingv1.GetItemRequest
	(*GetItemResponse)(nil),    // 5: testingv1.GetItemResponse
	(*GetItemsRequest)(nil),    // 6: testingv1.GetItemsRequest
	(*GetItemsResponse)(nil),   // 7: testingv1.GetItemsResponse
	(*UpdateItemRequest)(nil),  // 8: testingv1.UpdateItemRequest
	(*UpdateItemResponse)(nil), // 9: testingv1.UpdateItemResponse
	(*DeleteItemRequest)(nil),  // 10: testingv1.DeleteItemRequest
	(*DeleteItemResponse)(nil), // 11: testingv1.DeleteItemResponse
	(*CreateItemRequest)(nil),  // 12: testingv1.CreateItemRequest
	(*CreateItemResponse)(nil), // 13: testingv1.CreateItemResponse
}
var file_testingv1_service_proto_depIdxs = []int32{
	0,  // 0: testingv1.Item.status:type_name -> testingv1.ItemStatus
	1,  // 1: testingv1.ItemStateEvent.type:type_name -> testingv1.ItemStateEvent.Type
	2,  // 2: testingv1.ItemStateEvent.payload:type_name -> testingv1.Item
	2,  // 3: testingv1.GetItemResponse.data:type_name -> testingv1.Item
	2,  // 4: testingv1.GetItemsResponse.data:type_name -> testingv1.Item
	2,  // 5: testingv1.UpdateItemRequest.data:type_name -> testingv1.Item
	2,  // 6: testingv1.UpdateItemResponse.data:type_name -> testingv1.Item
	2,  // 7: testingv1.CreateItemRequest.data:type_name -> testingv1.Item
	2,  // 8: testingv1.CreateItemResponse.data:type_name -> testingv1.Item
	4,  // 9: testingv1.TestService.GetItem:input_type -> testingv1.GetItemRequest
	6,  // 10: testingv1.TestService.GetItems:input_type -> testingv1.GetItemsRequest
	8,  // 11: testingv1.TestService.UpdateItem:input_type -> testingv1.UpdateItemRequest
	10, // 12: testingv1.TestService.DeleteItem:input_type -> testingv1.DeleteItemRequest
	12, // 13: testingv1.TestService.CreateItem:input_type -> testingv1.CreateItemRequest
	5,  // 14: testingv1.TestService.GetItem:output_type -> testingv1.GetItemResponse
	7,  // 15: testingv1.TestService.GetItems:output_type -> testingv1.GetItemsResponse
	9,  // 16: testingv1.TestService.UpdateItem:output_type -> testingv1.UpdateItemResponse
	11, // 17: testingv1.TestService.DeleteItem:output_type -> testingv1.DeleteItemResponse
	13, // 18: testingv1.TestService.CreateItem:output_type -> testingv1.CreateItemResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_testingv1_service_proto_init() }
func file_testingv1_service_proto_init() {
	if File_testingv1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testingv1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_testingv1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemStateEvent); i {
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
		file_testingv1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemRequest); i {
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
		file_testingv1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemResponse); i {
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
		file_testingv1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemsRequest); i {
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
		file_testingv1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemsResponse); i {
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
		file_testingv1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemRequest); i {
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
		file_testingv1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemResponse); i {
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
		file_testingv1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemRequest); i {
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
		file_testingv1_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemResponse); i {
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
		file_testingv1_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateItemRequest); i {
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
		file_testingv1_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateItemResponse); i {
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
			RawDescriptor: file_testingv1_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_testingv1_service_proto_goTypes,
		DependencyIndexes: file_testingv1_service_proto_depIdxs,
		EnumInfos:         file_testingv1_service_proto_enumTypes,
		MessageInfos:      file_testingv1_service_proto_msgTypes,
	}.Build()
	File_testingv1_service_proto = out.File
	file_testingv1_service_proto_rawDesc = nil
	file_testingv1_service_proto_goTypes = nil
	file_testingv1_service_proto_depIdxs = nil
}