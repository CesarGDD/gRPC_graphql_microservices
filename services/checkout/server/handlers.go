package server

import (
	"context"
	"fmt"
	"log"
	db "checkout/database"
	proto "checkout/proto"
)

func (s *CheckoutServer) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	fmt.Println("Create order invoked", req)
	orderId, err := s.queries.CreateOrder(ctx, db.CreateOrderParams{
		UserID: req.UserId,
		TotalPrice: req.TotalPrice,
	})
	if err != nil {
		return &proto.CreateOrderResponse{
			Response: &proto.Response{
				Success: false,
				Message: "Failed to create order",
			},
		}, err
	}

	if len(req.OrderItems) > 0 {
		for _, item := range req.OrderItems {
				err := s.queries.AddOrderItem(ctx, db.AddOrderItemParams{
					OrderID: orderId,
					ProductID: item.ProductId,
					Quantity: item.Quantity,
				})
				if err != nil {
					return &proto.CreateOrderResponse{
						Response: &proto.Response{
							Success: false,
							Message: "Failed to add item into order",
						},
					}, err
				}
		}
	}

	return &proto.CreateOrderResponse{
		Response: &proto.Response{
			Success: true,
			Message: "Order created successfully",
		},
	}, nil
}

func (s *CheckoutServer) GetOrderDetails(ctx context.Context, req *proto.GetOrderDetailsRequest) (*proto.GetOrderDetailsResponse, error) {
	log.Println("Get order details invoked", req)
	order, err := s.queries.GetOrderDetails(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	items, err := s.queries.GetOrderItems(ctx, order.OrderID)
	if err != nil {
		return nil, err
	}

	var orderItems []*proto.OrderItem
	for _, item := range items {
		orderItems = append(orderItems, &proto.OrderItem{
			ProductId: item.ProductID,
			Quantity: item.Quantity,
		})
	}

	return &proto.GetOrderDetailsResponse{
		OrderDetails: &proto.OrderDetails{
			OrderId: order.OrderID,
			UserId: order.UserID,
			TotalPrice: order.TotalPrice,
			Status: order.Status,
			OrderItems: orderItems,
		},
	}, nil
}

func (s *CheckoutServer) GetOrdersDetailsByUserId(ctx context.Context, req *proto.GetOrdersDetailsByUserIdRequest) (*proto.GetOrdersDetailsByUserIdResponse, error) {
	log.Println("Get order details invoked", req)
	orders, err := s.queries.GetOrdersDetailsByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	response := &proto.GetOrdersDetailsByUserIdResponse{}
	for _, order := range orders {
		// Fetch the items for the current order
		items, err := s.queries.GetOrderItems(ctx, order.OrderID)
		if err != nil {
			return nil, err
		}

		var orderItems []*proto.OrderItem
		for _, item := range items {
			orderItems = append(orderItems, &proto.OrderItem{
				ProductId: item.ProductID,
				Quantity: item.Quantity,
			})
		}

		response.OrdersDetails = append(response.OrdersDetails, &proto.OrderDetails{
			OrderId:    order.OrderID,
			UserId:     order.UserID,
			TotalPrice: order.TotalPrice,
			Status:     order.Status,
			OrderItems: orderItems,
		})
	}

	return response, nil
}

func (s *CheckoutServer) ProcessPayment(ctx context.Context, req *proto.ProcessPaymentRequest) (*proto.ProcessPaymentResponse, error) {
	fmt.Println("Process payment invoked", req)
	err := s.queries.RecordPaymentTransaction(ctx, db.RecordPaymentTransactionParams{
		OrderID: req.OrderId,
		PaymentMethod: req.PaymentMethod,
	})
	if err != nil {
		return &proto.ProcessPaymentResponse{
			Response: &proto.Response{
				Success: false,
				Message: "Failed to process payment",
			},
		}, err
	}

	return &proto.ProcessPaymentResponse{
		Response: &proto.Response{
			Success: true,
			Message: "Payment successfull",
		},
	}, nil
}

func (s *CheckoutServer) GetPaymentDetails(ctx context.Context, req *proto.GetPaymentDetailsRequest) (*proto.GetPaymentDetailsResponse, error) {
	fmt.Println("Get payment details invoked", req)
	payment, err := s.queries.GetPaymentDetails(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	return &proto.GetPaymentDetailsResponse{
		PaymentMethod: payment,
	}, err
}
