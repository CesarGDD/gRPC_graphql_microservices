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

// GetProduct is the resolver for the getProduct field.
func (r *queryResolver) GetProduct(ctx context.Context, productID int) (*productspb.Product, error) {
	fmt.Println("Request to get product:", productID)

	res, err := r.ProductsClient.GetProduct(ctx, &productspb.GetProductRequest{
		ProductId: int32(productID),
	})
	if err != nil {
		fmt.Printf("Error fetching product %d: %v\n", productID, err)
		return nil, handleError(ctx, err, "fetch product")
	}
	if res == nil || res.Product == nil {
		fmt.Printf("No product found for ID %d\n", productID)
		graphql.AddError(ctx, gqlerror.Errorf("No product found with ID %d", productID))
		return nil, gqlerror.Errorf("No product found.")
	}

	fmt.Println("Product retrieved successfully:", res.Product.Product.Name)
	return res.Product.Product, nil
}

// GetProductByName is the resolver for the getProductByName field.
func (r *queryResolver) GetProductByName(ctx context.Context, name string) (*productspb.Product, error) {
	fmt.Println("Request to get product by name:", name)

	res, err := r.ProductsClient.GetProductByName(ctx, &productspb.GetProductByNameRequest{
		Name: name,
	})
	if err != nil {
		fmt.Printf("Error fetching product %s: %v\n", name, err)
		return nil, handleError(ctx, err, "fetch product by name")
	}
	if res == nil || res.Product == nil {
		fmt.Printf("No product found for ID %s\n", name)
		graphql.AddError(ctx, gqlerror.Errorf("No product found with ID %s", name))
		return nil, gqlerror.Errorf("No product found.")
	}

	fmt.Println("Product retrieved successfully:", res.Product.Product.Name)
	return res.Product.Product, nil
}

// GetProducts is the resolver for the getProducts field.
func (r *queryResolver) GetProducts(ctx context.Context) ([]*productspb.Product, error) {
	fmt.Println("Request to get products:")

	res, err := r.ProductsClient.GetProducts(ctx, &productspb.GetProductsRequest{})
	if err != nil {
		fmt.Printf("Error fetching products %v\n", err)
		return nil, handleError(ctx, err, "fetch product details")
	}
	if res == nil || res.Products == nil {
		fmt.Printf("No products found\n")
		graphql.AddError(ctx, gqlerror.Errorf("No products found"))
		return nil, gqlerror.Errorf("No product found.")
	}

	var products []*productspb.Product
	for _, response := range res.Products {
		if response != nil && response.Product != nil {
			products = append(products, response.Product)
		}
	}
	fmt.Println("Products retrieved successfully:")
	return products, nil
}

// GetOrderDetails is the resolver for the getOrderDetails field.
func (r *queryResolver) GetOrderDetails(ctx context.Context, input model.GetOrderDetailsInput) (*model.Order, error) {
	fmt.Println("Request to get order details:", input.OrderID)

	res, err := r.CheckoutClient.GetOrderDetails(ctx, &checkout.GetOrderDetailsRequest{
		OrderId: int32(input.OrderID),
	})
	if err != nil {
		fmt.Printf("Error fetching order %v\n", err)
		return nil, handleError(ctx, err, "Error fetching order")
	}
	if res == nil || res.OrderDetails == nil {
		fmt.Printf("No order found\n")
		graphql.AddError(ctx, gqlerror.Errorf("No order found"))
		return nil, gqlerror.Errorf("No order found.")
	}
	fmt.Println("Order retrieved successfully:")
	return &model.Order{
		OrderDetails: res.OrderDetails,
	}, nil
}

// GetOrdersDetailsByUserID is the resolver for the GetOrdersDetailsByUserId field.
func (r *queryResolver) GetOrdersDetailsByUserID(ctx context.Context, input model.GetOrdersDetailsByUserIDInput) ([]*checkout.OrderDetails, error) {
	fmt.Println("Request to get order details by userId:", input.UserID)

	res, err := r.CheckoutClient.GetOrdersDetailsByUserId(ctx, &checkout.GetOrdersDetailsByUserIdRequest{
		UserId: int32(input.UserID),
	})
	if err != nil {
		fmt.Printf("Error fetching order %v\n", err)
		return nil, handleError(ctx, err, "Error fetching order by userId")
	}
	if res == nil || res.OrdersDetails == nil {
		fmt.Printf("No orders by userId found\n")
		graphql.AddError(ctx, gqlerror.Errorf("No orders  by userId found"))
		return nil, gqlerror.Errorf("No orders by userId found.")
	}
	fmt.Println("Order retrieved successfully:")
	return res.OrdersDetails, nil
}

// GetPaymentDetails is the resolver for the getPaymentDetails field.
func (r *queryResolver) GetPaymentDetails(ctx context.Context, input model.GetPaymentDetailsInput) (*model.PaymentDetails, error) {
	fmt.Println("Request to get payment details:", input.OrderID)

	res, err := r.CheckoutClient.GetPaymentDetails(ctx, &checkout.GetPaymentDetailsRequest{
		OrderId: int32(input.OrderID),
	})
	if err != nil {
		fmt.Printf("Error fetching payment details %v\n", err)
		return nil, handleError(ctx, err, "Error fetching payment details")
	}
	if res == nil || res.PaymentMethod == "" {
		fmt.Printf("No payment details found\n")
		graphql.AddError(ctx, gqlerror.Errorf("No payment details found"))
		return nil, gqlerror.Errorf("No payment details found.")
	}

	fmt.Println("Products retrieved successfully:")
	return &model.PaymentDetails{
		PaymentMethod: res.PaymentMethod,
	}, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, userID int) (*usermanagementpb.User, error) {
	fmt.Println("Request to get user:", userID)

	res, err := r.UserManagementClient.GetUser(ctx, &usermanagementpb.GetUserRequest{
		UserId: int32(userID),
	})
	if err != nil {
		fmt.Printf("Error fetching user %d: %v\n", userID, err)
		return nil, handleError(ctx, err, "Error fetching user")
	}
	if res == nil || res.User == nil {
		fmt.Printf("No user found for ID %d\n", userID)
		graphql.AddError(ctx, gqlerror.Errorf("No user found with ID %d", userID))
		return nil, gqlerror.Errorf("No user found.")
	}

	fmt.Println("user retrieved successfully:", res.User)
	return &usermanagementpb.User{
		UserId:   res.User.UserId,
		Username: res.User.Username,
		Role:     res.User.Role,
	}, nil
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*usermanagementpb.User, error) {
	fmt.Println("Request to get users:")

	res, err := r.UserManagementClient.GetUsers(ctx, &usermanagementpb.GetUsersRequest{})
	if err != nil {
		fmt.Printf("Error fetching users: %v\n", err)
		return nil, handleError(ctx, err, "Error fetching users")
	}
	if res == nil || res.Users == nil {
		fmt.Printf("No users found\n")
		graphql.AddError(ctx, gqlerror.Errorf("No users found"))
		return nil, gqlerror.Errorf("No user found.")
	}

	var users []*usermanagementpb.User
	for _, user := range res.Users {
		users = append(users, &usermanagementpb.User{
			UserId:   user.UserId,
			Username: user.Username,
			Role:     user.Role,
		})
	}

	fmt.Println("user retrieved successfully:", res.Users)
	return users, nil
}

// GetUserByUsername is the resolver for the getUserByUsername field.
func (r *queryResolver) GetUserByUsername(ctx context.Context, username string) (*usermanagementpb.User, error) {
	fmt.Println("Request to get user:", username)

	res, err := r.UserManagementClient.GetUserByUsername(ctx, &usermanagementpb.GetUserByUsernameRequest{
		Username: username,
	})
	if err != nil {
		fmt.Printf("Error fetching user %s: %v\n", username, err)
		return nil, handleError(ctx, err, "Error fetching user by username")
	}
	if res == nil || res.User == nil {
		fmt.Printf("No user found for ID %s\n", username)
		graphql.AddError(ctx, gqlerror.Errorf("No user found with ID %s", username))
		return nil, gqlerror.Errorf("No user found.")
	}

	fmt.Println("user retrieved successfully:", res.User)
	return &usermanagementpb.User{
		UserId:   res.User.UserId,
		Username: res.User.Username,
		Role:     res.User.Role,
	}, nil
}

// GetCart is the resolver for the getCart field.
func (r *queryResolver) GetCart(ctx context.Context, input model.GetCartInput) (*shoppingcart.ShoppingCart, error) {
	fmt.Println("Request to get shopping cart:", input.UserID)

	res, err := r.ShoppingCartClient.GetCart(ctx, &shoppingcart.GetCartRequest{
		UserId: int32(input.UserID),
	})
	if err != nil {
		fmt.Printf("Error fetching shopping cart %d: %v\n", input.UserID, err)
		return nil, handleError(ctx, err, "Error fetching shopping cart")
	}
	if res == nil || res.Cart == nil {
		fmt.Printf("No shopping cart found for ID %d\n", input.UserID)
		graphql.AddError(ctx, gqlerror.Errorf("No shopping cart found with ID %d", input.UserID))
		return nil, gqlerror.Errorf("No shopping cart found.")
	}

	fmt.Println("user retrieved successfully:", res.Cart)
	return res.Cart, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
