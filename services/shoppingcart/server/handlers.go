package server

import (
	"context"
	"fmt"
	"log"
	db "shoppingcart/database"
	proto "shoppingcart/proto"
)

func (s *ShoppingCartServer) AddItem(ctx context.Context, req *proto.AddItemRequest) (*proto.AddItemResponse, error) {
	fmt.Println("Add item to cart invoked")
	err := s.queries.AddItemToCart(ctx, db.AddItemToCartParams{
		UserID: req.UserId,
		ProductID: req.Item.ProductId,
		Quantity: req.Item.Quantity,
	})
	if err != nil {
		return &proto.AddItemResponse{
			Response: &proto.Response{
				Success: false,
				Message: "Failed to add product to the shopping cart",
			},
		}, err
	}

	return &proto.AddItemResponse{
		Response: &proto.Response{
			Success: true,
			Message: "Product added successfully to the shopping cart",
		},
	}, nil
}
func (s *ShoppingCartServer) UpdateItem(ctx context.Context, req *proto.UpdateItemRequest) (*proto.UpdateItemResponse, error) {
	fmt.Println("Update item in cart invoked")
	err := s.queries.UpdateCartItem(ctx, db.UpdateCartItemParams{
		UserID: req.UserId,
		ProductID: req.Item.ProductId,
		Quantity: req.Item.Quantity,
	})
	if err != nil {
		return &proto.UpdateItemResponse{
			Response: &proto.Response{
				Success: false,
				Message: "Failed to update item in the shopping cart",
			},
		}, err
	}

	return &proto.UpdateItemResponse{
		Response: &proto.Response{
			Success: true,
			Message: "Item updateed successfully in the shopping cart",
		},
	}, nil
}
func (s *ShoppingCartServer) RemoveItem(ctx context.Context, req *proto.RemoveItemRequest) (*proto.RemoveItemResponse, error) {
	log.Println("Remove item from shoping cart invoked")
	err := s.queries.RemoveItemFromCart(ctx, db.RemoveItemFromCartParams{
		UserID: req.UserId,
		ProductID: req.ProductId,
	})
	if err != nil {
		return &proto.RemoveItemResponse{
			Response: &proto.Response{
				Success: false,
				Message: "Fail to remove item from shopping cart",
			},
		}, err
	}

	return &proto.RemoveItemResponse{
		Response: &proto.Response{
			Success: true,
			Message: "Item removed successfully from shopping cart",
		},
	}, nil
}

func (s *ShoppingCartServer) ClearCar(ctx context.Context, req *proto.ClearCartRequest) (*proto.ClearCartResponse, error) {
	log.Println("Clear cart invoked")
	err := s.queries.ClearCart(ctx, req.UserId)
	if err != nil {
		return &proto.ClearCartResponse{
			Response: &proto.Response{
				Success: false,
				Message: "Fail to clear cart",
			},
		}, err
	}

	return &proto.ClearCartResponse{
		Response: &proto.Response{
			Success: true,
			Message: "Cart successfully cleared",
		},
	}, nil
}

func (s *ShoppingCartServer) GetCart(ctx context.Context, req *proto.GetCartRequest) (*proto.GetCartResponse, error) {
	log.Println("Get cart invoked")
	cart, err := s.queries.GetCartItems(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var items []*proto.ProductItem
	for _, item := range cart {
		items = append(items, &proto.ProductItem{
			ProductId: item.ProductID,
			Quantity: item.Quantity,
		})
	}

	return &proto.GetCartResponse{
		Success: true,
		Cart: &proto.ShoppingCart{
			UserId: req.UserId,
			Items: items,
		},
	}, nil
}
