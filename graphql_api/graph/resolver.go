package graph

import (
	"context"
	"fmt"
	"graphql_api/graph/model"
	checkoutpb "graphql_api/protos/checkoutpb"
	productspb "graphql_api/protos/productspb"
	shoppingcartpb "graphql_api/protos/shoppingcartpb"
	usermanagementpb "graphql_api/protos/usermanagementpb"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/status"
)

type Resolver struct {
	CheckoutClient        checkoutpb.CheckoutServiceClient
	UserManagementClient  usermanagementpb.UserManagementServiceClient
	ProductsClient        productspb.ProductServiceClient
	ShoppingCartClient    shoppingcartpb.ShoppingCartServiceClient
}

func mapStringToRole(roleStr string) (model.Role, error) {
    switch roleStr {
    case "ADMIN":
        return model.RoleAdmin, nil
    case "CUSTOMER":
        return model.RoleCustomer, nil
    default:
        return "", fmt.Errorf("invalid role %s", roleStr)
    }
}

func handleError(ctx context.Context, err error, message string) error {
    fmt.Printf("Error %s: %v\n", message, err)
    e, ok := status.FromError(err)
    if ok {
        // gRPC specific error handling
        fmt.Printf("gRPC error status: %v\n", e.Message())
        graphql.AddError(ctx, gqlerror.Errorf("gRPC error: %s", e.Message()))
    } else {
        // General error handling
        fmt.Printf("Non-gRPC error: %v\n", err)
        graphql.AddError(ctx, gqlerror.Errorf("Internal server error: %v", err))
    }
    return gqlerror.Errorf("Failed to %s", message)
}
