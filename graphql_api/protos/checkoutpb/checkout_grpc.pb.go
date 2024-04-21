// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: checkoutpb/checkout.proto

package checkout

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
	CheckoutService_CreateOrder_FullMethodName              = "/checkout.CheckoutService/CreateOrder"
	CheckoutService_GetOrderDetails_FullMethodName          = "/checkout.CheckoutService/GetOrderDetails"
	CheckoutService_GetOrdersDetailsByUserId_FullMethodName = "/checkout.CheckoutService/GetOrdersDetailsByUserId"
	CheckoutService_ProcessPayment_FullMethodName           = "/checkout.CheckoutService/ProcessPayment"
	CheckoutService_GetPaymentDetails_FullMethodName        = "/checkout.CheckoutService/GetPaymentDetails"
)

// CheckoutServiceClient is the client API for CheckoutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckoutServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	GetOrderDetails(ctx context.Context, in *GetOrderDetailsRequest, opts ...grpc.CallOption) (*GetOrderDetailsResponse, error)
	GetOrdersDetailsByUserId(ctx context.Context, in *GetOrdersDetailsByUserIdRequest, opts ...grpc.CallOption) (*GetOrdersDetailsByUserIdResponse, error)
	ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error)
	GetPaymentDetails(ctx context.Context, in *GetPaymentDetailsRequest, opts ...grpc.CallOption) (*GetPaymentDetailsResponse, error)
}

type checkoutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckoutServiceClient(cc grpc.ClientConnInterface) CheckoutServiceClient {
	return &checkoutServiceClient{cc}
}

func (c *checkoutServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, CheckoutService_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkoutServiceClient) GetOrderDetails(ctx context.Context, in *GetOrderDetailsRequest, opts ...grpc.CallOption) (*GetOrderDetailsResponse, error) {
	out := new(GetOrderDetailsResponse)
	err := c.cc.Invoke(ctx, CheckoutService_GetOrderDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkoutServiceClient) GetOrdersDetailsByUserId(ctx context.Context, in *GetOrdersDetailsByUserIdRequest, opts ...grpc.CallOption) (*GetOrdersDetailsByUserIdResponse, error) {
	out := new(GetOrdersDetailsByUserIdResponse)
	err := c.cc.Invoke(ctx, CheckoutService_GetOrdersDetailsByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkoutServiceClient) ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error) {
	out := new(ProcessPaymentResponse)
	err := c.cc.Invoke(ctx, CheckoutService_ProcessPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkoutServiceClient) GetPaymentDetails(ctx context.Context, in *GetPaymentDetailsRequest, opts ...grpc.CallOption) (*GetPaymentDetailsResponse, error) {
	out := new(GetPaymentDetailsResponse)
	err := c.cc.Invoke(ctx, CheckoutService_GetPaymentDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckoutServiceServer is the server API for CheckoutService service.
// All implementations must embed UnimplementedCheckoutServiceServer
// for forward compatibility
type CheckoutServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	GetOrderDetails(context.Context, *GetOrderDetailsRequest) (*GetOrderDetailsResponse, error)
	GetOrdersDetailsByUserId(context.Context, *GetOrdersDetailsByUserIdRequest) (*GetOrdersDetailsByUserIdResponse, error)
	ProcessPayment(context.Context, *ProcessPaymentRequest) (*ProcessPaymentResponse, error)
	GetPaymentDetails(context.Context, *GetPaymentDetailsRequest) (*GetPaymentDetailsResponse, error)
	mustEmbedUnimplementedCheckoutServiceServer()
}

// UnimplementedCheckoutServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCheckoutServiceServer struct {
}

func (UnimplementedCheckoutServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedCheckoutServiceServer) GetOrderDetails(context.Context, *GetOrderDetailsRequest) (*GetOrderDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderDetails not implemented")
}
func (UnimplementedCheckoutServiceServer) GetOrdersDetailsByUserId(context.Context, *GetOrdersDetailsByUserIdRequest) (*GetOrdersDetailsByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersDetailsByUserId not implemented")
}
func (UnimplementedCheckoutServiceServer) ProcessPayment(context.Context, *ProcessPaymentRequest) (*ProcessPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessPayment not implemented")
}
func (UnimplementedCheckoutServiceServer) GetPaymentDetails(context.Context, *GetPaymentDetailsRequest) (*GetPaymentDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentDetails not implemented")
}
func (UnimplementedCheckoutServiceServer) mustEmbedUnimplementedCheckoutServiceServer() {}

// UnsafeCheckoutServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckoutServiceServer will
// result in compilation errors.
type UnsafeCheckoutServiceServer interface {
	mustEmbedUnimplementedCheckoutServiceServer()
}

func RegisterCheckoutServiceServer(s grpc.ServiceRegistrar, srv CheckoutServiceServer) {
	s.RegisterService(&CheckoutService_ServiceDesc, srv)
}

func _CheckoutService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckoutService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckoutService_GetOrderDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).GetOrderDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckoutService_GetOrderDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).GetOrderDetails(ctx, req.(*GetOrderDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckoutService_GetOrdersDetailsByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersDetailsByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).GetOrdersDetailsByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckoutService_GetOrdersDetailsByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).GetOrdersDetailsByUserId(ctx, req.(*GetOrdersDetailsByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckoutService_ProcessPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).ProcessPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckoutService_ProcessPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).ProcessPayment(ctx, req.(*ProcessPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckoutService_GetPaymentDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).GetPaymentDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckoutService_GetPaymentDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).GetPaymentDetails(ctx, req.(*GetPaymentDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckoutService_ServiceDesc is the grpc.ServiceDesc for CheckoutService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckoutService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "checkout.CheckoutService",
	HandlerType: (*CheckoutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _CheckoutService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrderDetails",
			Handler:    _CheckoutService_GetOrderDetails_Handler,
		},
		{
			MethodName: "GetOrdersDetailsByUserId",
			Handler:    _CheckoutService_GetOrdersDetailsByUserId_Handler,
		},
		{
			MethodName: "ProcessPayment",
			Handler:    _CheckoutService_ProcessPayment_Handler,
		},
		{
			MethodName: "GetPaymentDetails",
			Handler:    _CheckoutService_GetPaymentDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checkoutpb/checkout.proto",
}
