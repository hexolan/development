// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: payment/v1/events.proto

package payment_v1

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

type OrderPaymentStatus int32

const (
	OrderPaymentStatus_FAILURE OrderPaymentStatus = 0
	OrderPaymentStatus_SUCCESS OrderPaymentStatus = 1
)

// Enum value maps for OrderPaymentStatus.
var (
	OrderPaymentStatus_name = map[int32]string{
		0: "FAILURE",
		1: "SUCCESS",
	}
	OrderPaymentStatus_value = map[string]int32{
		"FAILURE": 0,
		"SUCCESS": 1,
	}
)

func (x OrderPaymentStatus) Enum() *OrderPaymentStatus {
	p := new(OrderPaymentStatus)
	*p = x
	return p
}

func (x OrderPaymentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderPaymentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_payment_v1_events_proto_enumTypes[0].Descriptor()
}

func (OrderPaymentStatus) Type() protoreflect.EnumType {
	return &file_payment_v1_events_proto_enumTypes[0]
}

func (x OrderPaymentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderPaymentStatus.Descriptor instead.
func (OrderPaymentStatus) EnumDescriptor() ([]byte, []int) {
	return file_payment_v1_events_proto_rawDescGZIP(), []int{0}
}

type PaymentCreditEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string  `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	CustomerId    string  `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Amount        float32 `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PaymentCreditEvent) Reset() {
	*x = PaymentCreditEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentCreditEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentCreditEvent) ProtoMessage() {}

func (x *PaymentCreditEvent) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentCreditEvent.ProtoReflect.Descriptor instead.
func (*PaymentCreditEvent) Descriptor() ([]byte, []int) {
	return file_payment_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentCreditEvent) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *PaymentCreditEvent) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *PaymentCreditEvent) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type PaymentDebitEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string  `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	CustomerId    string  `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Amount        float32 `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PaymentDebitEvent) Reset() {
	*x = PaymentDebitEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentDebitEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentDebitEvent) ProtoMessage() {}

func (x *PaymentDebitEvent) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentDebitEvent.ProtoReflect.Descriptor instead.
func (*PaymentDebitEvent) Descriptor() ([]byte, []int) {
	return file_payment_v1_events_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentDebitEvent) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *PaymentDebitEvent) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *PaymentDebitEvent) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type OrderPaymentEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        OrderPaymentStatus `protobuf:"varint,1,opt,name=status,proto3,enum=stocklet.payment.v1.OrderPaymentStatus" json:"status,omitempty"`
	PaymentAmount float32            `protobuf:"fixed32,2,opt,name=payment_amount,json=paymentAmount,proto3" json:"payment_amount,omitempty"`
	// todo: fix import pattern - pass along order saga details or some generic
	OrderId    string `protobuf:"bytes,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CustomerId string `protobuf:"bytes,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Products   string `protobuf:"bytes,5,opt,name=products,proto3" json:"products,omitempty"` // should not be a string
}

func (x *OrderPaymentEvent) Reset() {
	*x = OrderPaymentEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPaymentEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPaymentEvent) ProtoMessage() {}

func (x *OrderPaymentEvent) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPaymentEvent.ProtoReflect.Descriptor instead.
func (*OrderPaymentEvent) Descriptor() ([]byte, []int) {
	return file_payment_v1_events_proto_rawDescGZIP(), []int{2}
}

func (x *OrderPaymentEvent) GetStatus() OrderPaymentStatus {
	if x != nil {
		return x.Status
	}
	return OrderPaymentStatus_FAILURE
}

func (x *OrderPaymentEvent) GetPaymentAmount() float32 {
	if x != nil {
		return x.PaymentAmount
	}
	return 0
}

func (x *OrderPaymentEvent) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderPaymentEvent) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *OrderPaymentEvent) GetProducts() string {
	if x != nil {
		return x.Products
	}
	return ""
}

var File_payment_v1_events_proto protoreflect.FileDescriptor

var file_payment_v1_events_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x6c, 0x65, 0x74, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x74,
	0x0a, 0x12, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x73, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x62, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd3, 0x01, 0x0a, 0x11, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2a,
	0x2e, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x42,
	0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65,
	0x78, 0x6f, 0x6c, 0x61, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_payment_v1_events_proto_rawDescOnce sync.Once
	file_payment_v1_events_proto_rawDescData = file_payment_v1_events_proto_rawDesc
)

func file_payment_v1_events_proto_rawDescGZIP() []byte {
	file_payment_v1_events_proto_rawDescOnce.Do(func() {
		file_payment_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_v1_events_proto_rawDescData)
	})
	return file_payment_v1_events_proto_rawDescData
}

var file_payment_v1_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_payment_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_payment_v1_events_proto_goTypes = []interface{}{
	(OrderPaymentStatus)(0),    // 0: stocklet.payment.v1.OrderPaymentStatus
	(*PaymentCreditEvent)(nil), // 1: stocklet.payment.v1.PaymentCreditEvent
	(*PaymentDebitEvent)(nil),  // 2: stocklet.payment.v1.PaymentDebitEvent
	(*OrderPaymentEvent)(nil),  // 3: stocklet.payment.v1.OrderPaymentEvent
}
var file_payment_v1_events_proto_depIdxs = []int32{
	0, // 0: stocklet.payment.v1.OrderPaymentEvent.status:type_name -> stocklet.payment.v1.OrderPaymentStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payment_v1_events_proto_init() }
func file_payment_v1_events_proto_init() {
	if File_payment_v1_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_v1_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentCreditEvent); i {
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
		file_payment_v1_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentDebitEvent); i {
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
		file_payment_v1_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPaymentEvent); i {
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
			RawDescriptor: file_payment_v1_events_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payment_v1_events_proto_goTypes,
		DependencyIndexes: file_payment_v1_events_proto_depIdxs,
		EnumInfos:         file_payment_v1_events_proto_enumTypes,
		MessageInfos:      file_payment_v1_events_proto_msgTypes,
	}.Build()
	File_payment_v1_events_proto = out.File
	file_payment_v1_events_proto_rawDesc = nil
	file_payment_v1_events_proto_goTypes = nil
	file_payment_v1_events_proto_depIdxs = nil
}