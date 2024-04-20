package graph

import (
	"fmt"
	"graphql_api/graph/model"
	checkoutpb "graphql_api/protos/checkoutpb"
	productspb "graphql_api/protos/productspb"
	shoppingcartpb "graphql_api/protos/shoppingcartpb"
	usermanagementpb "graphql_api/protos/usermanagementpb"
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