// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: checkoutpb/checkout.proto

package checkout

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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutpb_checkout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutpb_checkout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_checkoutpb_checkout_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The OrderItem message represents a single item in an order.
type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int32 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"` // Unique identifier for the product
	Quantity  int32 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`                    // Quantity of the product
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutpb_checkout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutpb_checkout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_checkoutpb_checkout_proto_rawDescGZIP(), []int{1}
}

func (x *OrderItem) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// The OrderDetails message represents the complete order.
type OrderDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    int32        `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`          // Unique identifier for the order
	UserId     int32        `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             // Unique identifier for the user
	TotalPrice int32        `protobuf:"varint,3,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"` // Total price of the order
	Status     string       `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`                            // Status of the order (e.g., "Pending", "Completed")
	OrderItems []*OrderItem `protobuf:"bytes,5,rep,name=orderItems,proto3" json:"orderItems,omitempty"`
}

func (x *OrderDetails) Reset() {
	*x = OrderDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutpb_checkout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDetails) ProtoMessage() {}

func (x *OrderDetails) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutpb_checkout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDetails.ProtoReflect.Descriptor instead.
func (*OrderDetails) Descriptor() ([]byte, []int) {
	return file_checkoutpb_checkout_proto_rawDescGZIP(), []int{2}
}

func (x *OrderDetails) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderDetails) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderDetails) GetTotalPrice() int32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *OrderDetails) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderDetails) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

// Request to create an order from the shopping cart.
type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int32        `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TotalPrice int32        `protobuf:"varint,2,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	OrderItems []*OrderItem `protobuf:"bytes,3,rep,name=orderItems,proto3" json:"orderItems,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutpb_checkout_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutpb_checkout_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_checkoutpb_checkout_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateOrderRequest) GetTotalPrice() int32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *CreateOrderRequest) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

// Response after creating an order.
type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutpb_checkout_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutpb_checkout_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_checkoutpb_checkout_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOrderResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

// Request to get order details.
type GetOrderDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int32 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *GetOrderDetailsRequest) Reset() {
	*x = GetOrderDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutpb_checkout_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderDetailsRequest) ProtoMessage() {}

func (x *GetOrderDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutpb_checkout_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetOrderDetailsRequest) Descriptor() ([]byte, []int) {
	return file_checkoutpb_checkout_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrderDetailsRequest) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

// Response to get order details.
type GetOrderDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderDetails *OrderDetails `protobuf:"bytes,1,opt,name=orderDetails,proto3" json:"orderDetails,omitempty"`
}

func (x *GetOrderDetailsResponse) Reset() {
	*x = GetOrderDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutpb_checkout_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderDetailsResponse) ProtoMessage() {}

func (x *GetOrderDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutpb_checkout_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetOrderDetailsResponse) Descriptor() ([]byte, []int) {
	return file_checkoutpb_checkout_proto_rawDescGZIP(), []int{6}
}

func (x *GetOrderDetailsResponse) GetOrderDetails() *OrderDetails {
	if x != nil {
		return x.OrderDetails
	}
	return nil
}

// Request to get order details.
type GetOrdersDetailsByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetOrdersDetailsByUserIdRequest) Reset() {
	*x = GetOrdersDetailsByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutpb_checkout_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersDetailsByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersDetailsByUserIdRequest) ProtoMessage() {}

func (x *GetOrdersDetailsByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutpb_checkout_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersDetailsByUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetOrdersDetailsByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_checkoutpb_checkout_proto_rawDescGZIP(), []int{7}
}

func (x *GetOrdersDetailsByUserIdRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// Response to get order details.
type GetOrdersDetailsByUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrdersDetails []*OrderDetails `protobuf:"bytes,1,rep,name=ordersDetails,proto3" json:"ordersDetails,omitempty"`
}

func (x *GetOrdersDetailsByUserIdResponse) Reset() {
	*x = GetOrdersDetailsByUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutpb_checkout_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersDetailsByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersDetailsByUserIdResponse) ProtoMessage() {}

func (x *GetOrdersDetailsByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutpb_checkout_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersDetailsByUserIdResponse.ProtoReflect.Descriptor instead.
func (*GetOrdersDetailsByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_checkoutpb_checkout_proto_rawDescGZIP(), []int{8}
}

func (x *GetOrdersDetailsByUserIdResponse) GetOrdersDetails() []*OrderDetails {
	if x != nil {
		return x.OrdersDetails
	}
	return nil
}

// Request for payment processing.
type ProcessPaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId       int32  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	PaymentMethod string `protobuf:"bytes,2,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"` // Details about the payment method (e.g., credit card, PayPal)
}

func (x *ProcessPaymentRequest) Reset() {
	*x = ProcessPaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutpb_checkout_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPaymentRequest) ProtoMessage() {}

func (x *ProcessPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutpb_checkout_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPaymentRequest.ProtoReflect.Descriptor instead.
func (*ProcessPaymentRequest) Descriptor() ([]byte, []int) {
	return file_checkoutpb_checkout_proto_rawDescGZIP(), []int{9}
}

func (x *ProcessPaymentRequest) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *ProcessPaymentRequest) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

// Response after processing payment.
type ProcessPaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ProcessPaymentResponse) Reset() {
	*x = ProcessPaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutpb_checkout_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPaymentResponse) ProtoMessage() {}

func (x *ProcessPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutpb_checkout_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPaymentResponse.ProtoReflect.Descriptor instead.
func (*ProcessPaymentResponse) Descriptor() ([]byte, []int) {
	return file_checkoutpb_checkout_proto_rawDescGZIP(), []int{10}
}

func (x *ProcessPaymentResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

// Request get payment.
type GetPaymentDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int32 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *GetPaymentDetailsRequest) Reset() {
	*x = GetPaymentDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutpb_checkout_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentDetailsRequest) ProtoMessage() {}

func (x *GetPaymentDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutpb_checkout_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentDetailsRequest) Descriptor() ([]byte, []int) {
	return file_checkoutpb_checkout_proto_rawDescGZIP(), []int{11}
}

func (x *GetPaymentDetailsRequest) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

// Response get payment.
type GetPaymentDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentMethod string `protobuf:"bytes,1,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
}

func (x *GetPaymentDetailsResponse) Reset() {
	*x = GetPaymentDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutpb_checkout_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentDetailsResponse) ProtoMessage() {}

func (x *GetPaymentDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutpb_checkout_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetPaymentDetailsResponse) Descriptor() ([]byte, []int) {
	return file_checkoutpb_checkout_proto_rawDescGZIP(), []int{12}
}

func (x *GetPaymentDetailsResponse) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

var File_checkoutpb_checkout_proto protoreflect.FileDescriptor

var file_checkoutpb_checkout_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x22, 0x3e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0xb0, 0x01,
	0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x83, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x33, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x45, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x55, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0c, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x3a, 0x0a, 0x1f, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x59, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x22, 0x48, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x32, 0xdb, 0x03, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x2e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x71, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x2e, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x22, 0x2e, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checkoutpb_checkout_proto_rawDescOnce sync.Once
	file_checkoutpb_checkout_proto_rawDescData = file_checkoutpb_checkout_proto_rawDesc
)

func file_checkoutpb_checkout_proto_rawDescGZIP() []byte {
	file_checkoutpb_checkout_proto_rawDescOnce.Do(func() {
		file_checkoutpb_checkout_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkoutpb_checkout_proto_rawDescData)
	})
	return file_checkoutpb_checkout_proto_rawDescData
}

var file_checkoutpb_checkout_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_checkoutpb_checkout_proto_goTypes = []interface{}{
	(*Response)(nil),                         // 0: checkout.Response
	(*OrderItem)(nil),                        // 1: checkout.OrderItem
	(*OrderDetails)(nil),                     // 2: checkout.OrderDetails
	(*CreateOrderRequest)(nil),               // 3: checkout.CreateOrderRequest
	(*CreateOrderResponse)(nil),              // 4: checkout.CreateOrderResponse
	(*GetOrderDetailsRequest)(nil),           // 5: checkout.GetOrderDetailsRequest
	(*GetOrderDetailsResponse)(nil),          // 6: checkout.GetOrderDetailsResponse
	(*GetOrdersDetailsByUserIdRequest)(nil),  // 7: checkout.GetOrdersDetailsByUserIdRequest
	(*GetOrdersDetailsByUserIdResponse)(nil), // 8: checkout.GetOrdersDetailsByUserIdResponse
	(*ProcessPaymentRequest)(nil),            // 9: checkout.ProcessPaymentRequest
	(*ProcessPaymentResponse)(nil),           // 10: checkout.ProcessPaymentResponse
	(*GetPaymentDetailsRequest)(nil),         // 11: checkout.GetPaymentDetailsRequest
	(*GetPaymentDetailsResponse)(nil),        // 12: checkout.GetPaymentDetailsResponse
}
var file_checkoutpb_checkout_proto_depIdxs = []int32{
	1,  // 0: checkout.OrderDetails.orderItems:type_name -> checkout.OrderItem
	1,  // 1: checkout.CreateOrderRequest.orderItems:type_name -> checkout.OrderItem
	0,  // 2: checkout.CreateOrderResponse.response:type_name -> checkout.Response
	2,  // 3: checkout.GetOrderDetailsResponse.orderDetails:type_name -> checkout.OrderDetails
	2,  // 4: checkout.GetOrdersDetailsByUserIdResponse.ordersDetails:type_name -> checkout.OrderDetails
	0,  // 5: checkout.ProcessPaymentResponse.response:type_name -> checkout.Response
	3,  // 6: checkout.CheckoutService.CreateOrder:input_type -> checkout.CreateOrderRequest
	5,  // 7: checkout.CheckoutService.GetOrderDetails:input_type -> checkout.GetOrderDetailsRequest
	7,  // 8: checkout.CheckoutService.GetOrdersDetailsByUserId:input_type -> checkout.GetOrdersDetailsByUserIdRequest
	9,  // 9: checkout.CheckoutService.ProcessPayment:input_type -> checkout.ProcessPaymentRequest
	11, // 10: checkout.CheckoutService.GetPaymentDetails:input_type -> checkout.GetPaymentDetailsRequest
	4,  // 11: checkout.CheckoutService.CreateOrder:output_type -> checkout.CreateOrderResponse
	6,  // 12: checkout.CheckoutService.GetOrderDetails:output_type -> checkout.GetOrderDetailsResponse
	8,  // 13: checkout.CheckoutService.GetOrdersDetailsByUserId:output_type -> checkout.GetOrdersDetailsByUserIdResponse
	10, // 14: checkout.CheckoutService.ProcessPayment:output_type -> checkout.ProcessPaymentResponse
	12, // 15: checkout.CheckoutService.GetPaymentDetails:output_type -> checkout.GetPaymentDetailsResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_checkoutpb_checkout_proto_init() }
func file_checkoutpb_checkout_proto_init() {
	if File_checkoutpb_checkout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_checkoutpb_checkout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_checkoutpb_checkout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
		file_checkoutpb_checkout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDetails); i {
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
		file_checkoutpb_checkout_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_checkoutpb_checkout_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
		file_checkoutpb_checkout_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderDetailsRequest); i {
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
		file_checkoutpb_checkout_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderDetailsResponse); i {
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
		file_checkoutpb_checkout_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersDetailsByUserIdRequest); i {
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
		file_checkoutpb_checkout_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersDetailsByUserIdResponse); i {
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
		file_checkoutpb_checkout_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessPaymentRequest); i {
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
		file_checkoutpb_checkout_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessPaymentResponse); i {
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
		file_checkoutpb_checkout_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentDetailsRequest); i {
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
		file_checkoutpb_checkout_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentDetailsResponse); i {
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
			RawDescriptor: file_checkoutpb_checkout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_checkoutpb_checkout_proto_goTypes,
		DependencyIndexes: file_checkoutpb_checkout_proto_depIdxs,
		MessageInfos:      file_checkoutpb_checkout_proto_msgTypes,
	}.Build()
	File_checkoutpb_checkout_proto = out.File
	file_checkoutpb_checkout_proto_rawDesc = nil
	file_checkoutpb_checkout_proto_goTypes = nil
	file_checkoutpb_checkout_proto_depIdxs = nil
}
