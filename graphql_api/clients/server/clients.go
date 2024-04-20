package clients

import (
	checkoutpb "graphql_api/protos/checkoutpb"
	"graphql_api/protos/productspb"
	shoppingcartpb "graphql_api/protos/shoppingcartpb"
	"graphql_api/protos/usermanagementpb"
	"log"
	"os"

	"google.golang.org/grpc"
)

// ServiceClientCreator is a function type that creates a new client from a gRPC connection
type ServiceClientCreator func(*grpc.ClientConn) interface{}

// CreateServiceClient handles the connection setup and invokes the client creator function
func CreateServiceClient(serviceURL string, creator ServiceClientCreator) interface{} {
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(os.Getenv(serviceURL), opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	return creator(conn)
}

// CheckoutSvc creates a client for the Checkout service
func CheckoutSvc() checkoutpb.CheckoutServiceClient {
	creator := func(conn *grpc.ClientConn) interface{} {
		return checkoutpb.NewCheckoutServiceClient(conn)
	}
	return CreateServiceClient("CHECKOUT_URL", creator).(checkoutpb.CheckoutServiceClient)
}

// UsermanagmentSvc creates a client for the Usermanagment service
func UsermanagmentSvc() usermanagementpb.UserManagementServiceClient {
	creator := func(conn *grpc.ClientConn) interface{} {
		return usermanagementpb.NewUserManagementServiceClient(conn)
	}
	return CreateServiceClient("USERMANAGEMENT_URL", creator).(usermanagementpb.UserManagementServiceClient)
}

// ProductsSvc creates a client for the Products service
func ProductsSvc() productspb.ProductServiceClient {
	creator := func(conn *grpc.ClientConn) interface{} {
		return productspb.NewProductServiceClient(conn)
	}
	return CreateServiceClient("PRODUCTS_URL", creator).(productspb.ProductServiceClient)
}

// ShoppingcartSvc creates a client for the Shoppingcart service
func ShoppingcartSvc() shoppingcartpb.ShoppingCartServiceClient {
	creator := func(conn *grpc.ClientConn) interface{} {
		return shoppingcartpb.NewShoppingCartServiceClient(conn)
	}
	return CreateServiceClient("SHOPPINGCART_URL", creator).(shoppingcartpb.ShoppingCartServiceClient)
}
