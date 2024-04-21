package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"graphql_api/graph/model"
	checkout "graphql_api/protos/checkoutpb"
	"graphql_api/protos/productspb"
	shoppingcart "graphql_api/protos/shoppingcartpb"
	"graphql_api/protos/usermanagementpb"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProductInput) (*checkout.Response, error) {
	fmt.Println("Mutation to create a product:", input)
	res, err := r.ProductsClient.CreateProduct(ctx, &productspb.CreateProductRequest{
		Name:        input.Name,
		UserId:      int32(input.UserID),
		Url:         input.URL,
		Price:       int32(input.Price),
		Description: input.Description,
		Title:       input.Title,
	})
	if err != nil {
		fmt.Printf("Error creating product: %v\n", err)
		return nil, handleError(ctx, err, "Error creating product")
	}
	if res == nil || res.Response == nil || !res.Response.Success {
		fmt.Printf("No product created\n")
		graphql.AddError(ctx, gqlerror.Errorf("No product created"))
		return nil, gqlerror.Errorf("No product found.")
	}
	fmt.Println("Product created successfully:", res.Response)
	return &checkout.Response{
		Message: res.Response.Message,
		Success: res.Response.Success,
	}, nil
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, input model.UpdateProductInput) (*checkout.Response, error) {
	fmt.Println("Mutation to update a product:", input)
	res, err := r.ProductsClient.UpdateProduct(ctx, &productspb.UpdateProductRequest{
		Product: &productspb.Product{
			ProductId:   int32(input.ProductID),
			Name:        input.Name,
			Url:         input.URL,
			Price:       int32(input.Price),
			Description: input.Description,
			Title:       input.Title,
		},
	})
	if err != nil {
		fmt.Printf("Error updating product: %v\n", err)
		return nil, handleError(ctx, err, "Error updating product")
	}
	if res == nil || res.Response == nil || !res.Response.Success {
		fmt.Printf("No product updated\n")
		graphql.AddError(ctx, gqlerror.Errorf("No product updated"))
		return nil, gqlerror.Errorf("No product found.")
	}
	fmt.Println("Product updated successfully:", res.Response)
	return &checkout.Response{
		Message: res.Response.Message,
		Success: res.Response.Success,
	}, nil
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, productID int) (*checkout.Response, error) {
	fmt.Println("Mutation to delete a product:", productID)
	res, err := r.ProductsClient.DeleteProduct(ctx, &productspb.DeleteProductRequest{
		ProductId: int32(productID),
	})

	if err != nil {
		fmt.Printf("Error deleting product: %v\n", err)
		return nil, handleError(ctx, err, "Error deleting product")
	}

	if res == nil || res.Response == nil || !res.Response.Success {
		fmt.Println("No product deleted")
		graphql.AddError(ctx, gqlerror.Errorf("No product deleted"))
		return nil, gqlerror.Errorf("No product found.")
	}

	fmt.Println("Product deleted successfully:", res.Response)
	return &checkout.Response{
		Message: res.Response.Message,
		Success: res.Response.Success,
	}, nil
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.CreateOrderInput) (*checkout.Response, error) {
	fmt.Println("Mutation to create a order:", input)

	convertedItems := make([]*checkout.OrderItem, len(input.OrderItems))
	for i, item := range input.OrderItems {
		convertedItems[i] = &checkout.OrderItem{
			ProductId: int32(item.ProductID),
			Quantity:  int32(item.Quantity),
		}
	}
	res, err := r.CheckoutClient.CreateOrder(ctx, &checkout.CreateOrderRequest{
		UserId:     int32(input.UserID),
		TotalPrice: int32(input.TotalPrice),
		OrderItems: convertedItems,
	})

	if err != nil {
		fmt.Printf("Error creating order: %v\n", err)
		return nil, handleError(ctx, err, "Error creating order")
	}

	if res == nil || res.Response == nil || !res.Response.Success {
		fmt.Println("No order created")
		graphql.AddError(ctx, gqlerror.Errorf("No order created"))
		return nil, gqlerror.Errorf("No order found.")
	}

	fmt.Println("order created successfully:", res.Response)
	return &checkout.Response{
		Message: res.Response.Message,
		Success: res.Response.Success,
	}, nil
}

// ProcessPayment is the resolver for the processPayment field.
func (r *mutationResolver) ProcessPayment(ctx context.Context, input model.ProcessPaymentInput) (*checkout.Response, error) {
	fmt.Println("Mutation to process payment:", input)
	res, err := r.CheckoutClient.ProcessPayment(ctx, &checkout.ProcessPaymentRequest{
		OrderId:       int32(input.OrderID),
		PaymentMethod: input.PaymentMethod,
	})

	if err != nil {
		fmt.Printf("Error processing payment: %v\n", err)
		return nil, handleError(ctx, err, "Error processing payment")
	}

	if res == nil || res.Response == nil || !res.Response.Success {
		fmt.Println("No payment processed")
		graphql.AddError(ctx, gqlerror.Errorf("No payment processed"))
		return nil, gqlerror.Errorf("No product found.")
	}

	fmt.Println("payment processed successfully:", res.Response)
	return &checkout.Response{
		Message: res.Response.Message,
		Success: res.Response.Success,
	}, nil
}

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUserInput) (*checkout.Response, error) {
	fmt.Println("Mutation to register an user:", input)
	res, err := r.UserManagementClient.Register(ctx, &usermanagementpb.RegisterRequest{
		Username: input.Username,
		Password: input.Password,
		Role:     input.Role.String(),
	})

	if err != nil {
		fmt.Printf("Error registering user: %v\n", err)
		return nil, handleError(ctx, err, "Error registering user")
	}

	if res == nil || res.Response == nil || !res.Response.Success {
		fmt.Println("No user registered")
		graphql.AddError(ctx, gqlerror.Errorf("No user registered"))
		return nil, gqlerror.Errorf("No user found.")
	}

	fmt.Println("user registered successfully:", res.Response)
	return &checkout.Response{
		Message: res.Response.Message,
		Success: res.Response.Success,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*checkout.Response, error) {
	fmt.Println("Mutation to update an user:", input)
	res, err := r.UserManagementClient.UpdateUser(ctx, &usermanagementpb.UpdateUserRequest{
		User: &usermanagementpb.User{
			UserId:       int32(input.UserID),
			Username:     input.Username,
			PasswordHash: input.Password,
			Role:         input.Role.String(),
		},
	})

	if err != nil {
		fmt.Printf("Error updating user: %v\n", err)
		return nil, handleError(ctx, err, "Error updating user")
	}

	if res == nil || res.Response == nil || !res.Response.Success {
		fmt.Println("No user updated")
		graphql.AddError(ctx, gqlerror.Errorf("No user updated"))
		return nil, gqlerror.Errorf("No user found.")
	}

	fmt.Println("user updated successfully:", res.Response)
	return &checkout.Response{
		Message: res.Response.Message,
		Success: res.Response.Success,
	}, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, userID int) (*checkout.Response, error) {
	fmt.Println("Mutation to delete an user:", userID)
	res, err := r.UserManagementClient.DeleteUser(ctx, &usermanagementpb.DeleteUserRequest{
		UserId: int32(userID),
	})

	if err != nil {
		fmt.Printf("Error deleting user: %v\n", err)
		return nil, handleError(ctx, err, "Error deleting user")
	}

	if res == nil || res.Response == nil || !res.Response.Success {
		fmt.Println("No user deleted")
		graphql.AddError(ctx, gqlerror.Errorf("No user deleted"))
		return nil, gqlerror.Errorf("No user found.")
	}

	fmt.Println("user deleted successfully:", res.Response)
	return &checkout.Response{
		Message: res.Response.Message,
		Success: res.Response.Success,
	}, nil
}

// AddItem is the resolver for the addItem field.
func (r *mutationResolver) AddItem(ctx context.Context, input model.AddItemInput) (*checkout.Response, error) {
	fmt.Println("Mutation to add an item:", input)
	res, err := r.ShoppingCartClient.AddItem(ctx, &shoppingcart.AddItemRequest{
		UserId: int32(input.UserID),
		Item: &shoppingcart.ProductItem{
			ProductId: int32(input.Item.ProductID),
			Quantity:  int32(input.Item.Quantity),
		},
	})

	if err != nil {
		fmt.Printf("Error adding item: %v\n", err)
		return nil, handleError(ctx, err, "Error adding item")
	}

	if res == nil || res.Response == nil || !res.Response.Success {
		fmt.Println("No item added")
		graphql.AddError(ctx, gqlerror.Errorf("No item added"))
		return nil, gqlerror.Errorf("No item found.")
	}

	fmt.Println("item added successfully:", res.Response)
	return &checkout.Response{
		Message: res.Response.Message,
		Success: res.Response.Success,
	}, nil
}

// UpdateItem is the resolver for the updateItem field.
func (r *mutationResolver) UpdateItem(ctx context.Context, input model.UpdateItemInput) (*checkout.Response, error) {
	fmt.Println("Mutation to update an item:", input)
	res, err := r.ShoppingCartClient.UpdateItem(ctx, &shoppingcart.UpdateItemRequest{
		UserId: int32(input.UserID),
		Item: &shoppingcart.ProductItem{
			ProductId: int32(input.Item.ProductID),
			Quantity:  int32(input.Item.Quantity),
		},
	})

	if err != nil {
		fmt.Printf("Error updating item: %v\n", err)
		return nil, handleError(ctx, err, "Error updating item")
	}

	if res == nil || res.Response == nil || !res.Response.Success {
		fmt.Println("No item updated")
		graphql.AddError(ctx, gqlerror.Errorf("No item updated"))
		return nil, gqlerror.Errorf("No item found.")
	}

	fmt.Println("item updated successfully:", res.Response)
	return &checkout.Response{
		Message: res.Response.Message,
		Success: res.Response.Success,
	}, nil
}

// RemoveItem is the resolver for the removeItem field.
func (r *mutationResolver) RemoveItem(ctx context.Context, input model.RemoveItemInput) (*checkout.Response, error) {
	fmt.Println("Mutation to remove an item:", input)
	res, err := r.ShoppingCartClient.RemoveItem(ctx, &shoppingcart.RemoveItemRequest{
		UserId:    int32(input.UserID),
		ProductId: int32(input.ProductID),
	})

	if err != nil {
		fmt.Printf("Error removing item: %v\n", err)
		return nil, handleError(ctx, err, "Error removing item")
	}

	if res == nil || res.Response == nil || !res.Response.Success {
		fmt.Println("No item removed")
		graphql.AddError(ctx, gqlerror.Errorf("No item removed"))
		return nil, gqlerror.Errorf("No item found.")
	}

	fmt.Println("item removed successfully:", res.Response)
	return &checkout.Response{
		Message: res.Response.Message,
		Success: res.Response.Success,
	}, nil
}

// ClearCart is the resolver for the clearCart field.
func (r *mutationResolver) ClearCart(ctx context.Context, input model.ClearCartInput) (*checkout.Response, error) {
	fmt.Println("Mutation to clear cart:", input)
	res, err := r.ShoppingCartClient.ClearCar(ctx, &shoppingcart.ClearCartRequest{
		UserId: int32(input.UserID),
	})

	if err != nil {
		fmt.Printf("Error clearing cart: %v\n", err)
		return nil, handleError(ctx, err, "Error clearing cart")
	}

	if res == nil || res.Response == nil || !res.Response.Success {
		fmt.Println("No cart cleared")
		graphql.AddError(ctx, gqlerror.Errorf("No cart cleared"))
		return nil, gqlerror.Errorf("No cart cleared.")
	}

	fmt.Println("car cleared successfully:", res.Response)
	return &checkout.Response{
		Message: res.Response.Message,
		Success: res.Response.Success,
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
