// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: testpb/v1/test.proto

package testpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TestService_GetTestItem_FullMethodName    = "/testpb.v1.TestService/GetTestItem"
	TestService_CreateTestItem_FullMethodName = "/testpb.v1.TestService/CreateTestItem"
)

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	GetTestItem(ctx context.Context, in *GetTestItemRequest, opts ...grpc.CallOption) (*GetTestItemResponse, error)
	CreateTestItem(ctx context.Context, in *CreateTestItemRequest, opts ...grpc.CallOption) (*CreateTestItemResponse, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) GetTestItem(ctx context.Context, in *GetTestItemRequest, opts ...grpc.CallOption) (*GetTestItemResponse, error) {
	out := new(GetTestItemResponse)
	err := c.cc.Invoke(ctx, TestService_GetTestItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) CreateTestItem(ctx context.Context, in *CreateTestItemRequest, opts ...grpc.CallOption) (*CreateTestItemResponse, error) {
	out := new(CreateTestItemResponse)
	err := c.cc.Invoke(ctx, TestService_CreateTestItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	GetTestItem(context.Context, *GetTestItemRequest) (*GetTestItemResponse, error)
	CreateTestItem(context.Context, *CreateTestItemRequest) (*CreateTestItemResponse, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) GetTestItem(context.Context, *GetTestItemRequest) (*GetTestItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestItem not implemented")
}
func (UnimplementedTestServiceServer) CreateTestItem(context.Context, *CreateTestItemRequest) (*CreateTestItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTestItem not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_GetTestItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).GetTestItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_GetTestItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).GetTestItem(ctx, req.(*GetTestItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_CreateTestItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTestItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).CreateTestItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_CreateTestItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).CreateTestItem(ctx, req.(*CreateTestItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testpb.v1.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTestItem",
			Handler:    _TestService_GetTestItem_Handler,
		},
		{
			MethodName: "CreateTestItem",
			Handler:    _TestService_CreateTestItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testpb/v1/test.proto",
}