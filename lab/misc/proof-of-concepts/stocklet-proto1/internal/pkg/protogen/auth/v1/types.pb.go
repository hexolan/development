// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: stocklet/auth/v1/types.proto

package auth_v1

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

type ECPublicJWK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kty string `protobuf:"bytes,1,opt,name=kty,proto3" json:"kty,omitempty"`
	Use string `protobuf:"bytes,2,opt,name=use,proto3" json:"use,omitempty"`
	Alg string `protobuf:"bytes,3,opt,name=alg,proto3" json:"alg,omitempty"`
	Crv string `protobuf:"bytes,4,opt,name=crv,proto3" json:"crv,omitempty"`
	X   string `protobuf:"bytes,5,opt,name=x,proto3" json:"x,omitempty"`
	Y   string `protobuf:"bytes,6,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *ECPublicJWK) Reset() {
	*x = ECPublicJWK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stocklet_auth_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ECPublicJWK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ECPublicJWK) ProtoMessage() {}

func (x *ECPublicJWK) ProtoReflect() protoreflect.Message {
	mi := &file_stocklet_auth_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ECPublicJWK.ProtoReflect.Descriptor instead.
func (*ECPublicJWK) Descriptor() ([]byte, []int) {
	return file_stocklet_auth_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *ECPublicJWK) GetKty() string {
	if x != nil {
		return x.Kty
	}
	return ""
}

func (x *ECPublicJWK) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

func (x *ECPublicJWK) GetAlg() string {
	if x != nil {
		return x.Alg
	}
	return ""
}

func (x *ECPublicJWK) GetCrv() string {
	if x != nil {
		return x.Crv
	}
	return ""
}

func (x *ECPublicJWK) GetX() string {
	if x != nil {
		return x.X
	}
	return ""
}

func (x *ECPublicJWK) GetY() string {
	if x != nil {
		return x.Y
	}
	return ""
}

type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenType   string `protobuf:"bytes,1,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ExpiresIn   int64  `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stocklet_auth_v1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_stocklet_auth_v1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_stocklet_auth_v1_types_proto_rawDescGZIP(), []int{1}
}

func (x *AuthToken) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *AuthToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthToken) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

var File_stocklet_auth_v1_types_proto protoreflect.FileDescriptor

var file_stocklet_auth_v1_types_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x22, 0x71, 0x0a, 0x0b, 0x45, 0x43, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4a, 0x57, 0x4b, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x74,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x6c, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x72, 0x76, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x63, 0x72, 0x76, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x79, 0x22, 0x6c, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49,
	0x6e, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x65, 0x78, 0x6f, 0x6c, 0x61, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stocklet_auth_v1_types_proto_rawDescOnce sync.Once
	file_stocklet_auth_v1_types_proto_rawDescData = file_stocklet_auth_v1_types_proto_rawDesc
)

func file_stocklet_auth_v1_types_proto_rawDescGZIP() []byte {
	file_stocklet_auth_v1_types_proto_rawDescOnce.Do(func() {
		file_stocklet_auth_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_stocklet_auth_v1_types_proto_rawDescData)
	})
	return file_stocklet_auth_v1_types_proto_rawDescData
}

var file_stocklet_auth_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stocklet_auth_v1_types_proto_goTypes = []interface{}{
	(*ECPublicJWK)(nil), // 0: stocklet.auth.v1.ECPublicJWK
	(*AuthToken)(nil),   // 1: stocklet.auth.v1.AuthToken
}
var file_stocklet_auth_v1_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stocklet_auth_v1_types_proto_init() }
func file_stocklet_auth_v1_types_proto_init() {
	if File_stocklet_auth_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stocklet_auth_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ECPublicJWK); i {
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
		file_stocklet_auth_v1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
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
			RawDescriptor: file_stocklet_auth_v1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stocklet_auth_v1_types_proto_goTypes,
		DependencyIndexes: file_stocklet_auth_v1_types_proto_depIdxs,
		MessageInfos:      file_stocklet_auth_v1_types_proto_msgTypes,
	}.Build()
	File_stocklet_auth_v1_types_proto = out.File
	file_stocklet_auth_v1_types_proto_rawDesc = nil
	file_stocklet_auth_v1_types_proto_goTypes = nil
	file_stocklet_auth_v1_types_proto_depIdxs = nil
}