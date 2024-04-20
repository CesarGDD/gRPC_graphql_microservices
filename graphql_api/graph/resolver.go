package graph

import (
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