package clients

import (
    "log"
    "os"
    "google.golang.org/grpc"
    checkoutpb "graphql_api/protos/checkoutpb"
    productspb "graphql_api/protos/productspb"
    shoppingcartpb "graphql_api/protos/shoppingcartpb"
    usermanagementpb "graphql_api/protos/usermanagementpb"
)

// CreateServiceClient handles the connection setup and returns a gRPC client
func CreateServiceClient(serviceURL string, creator func(*grpc.ClientConn) interface{}) interface{} {
    opts := grpc.WithInsecure()
    log.Printf("Attempting to connect to service at URL: %s", serviceURL)
    conn, err := grpc.Dial(serviceURL, opts)
    if err != nil {
        log.Fatalf("Could not connect to %s: %v", serviceURL, err)
    }
    log.Printf("Successfully connected to service at URL: %s", serviceURL)
    return creator(conn)
}

func CheckoutSvc() checkoutpb.CheckoutServiceClient {
    serviceURL := os.Getenv("CHECKOUT_URL")
    creator := func(conn *grpc.ClientConn) interface{} {
        return checkoutpb.NewCheckoutServiceClient(conn)
    }
    return CreateServiceClient(serviceURL, creator).(checkoutpb.CheckoutServiceClient)
}

func UsermanagmentSvc() usermanagementpb.UserManagementServiceClient {
    serviceURL := os.Getenv("USERMANAGEMENT_URL")
    creator := func(conn *grpc.ClientConn) interface{} {
        return usermanagementpb.NewUserManagementServiceClient(conn)
    }
    return CreateServiceClient(serviceURL, creator).(usermanagementpb.UserManagementServiceClient)
}

func ProductsSvc() productspb.ProductServiceClient {
    serviceURL := os.Getenv("PRODUCTS_URL")
    creator := func(conn *grpc.ClientConn) interface{} {
        return productspb.NewProductServiceClient(conn)
    }
    return CreateServiceClient(serviceURL, creator).(productspb.ProductServiceClient)
}

func ShoppingcartSvc() shoppingcartpb.ShoppingCartServiceClient {
    serviceURL := os.Getenv("SHOPPINGCART_URL")
    creator := func(conn *grpc.ClientConn) interface{} {
        return shoppingcartpb.NewShoppingCartServiceClient(conn)
    }
    return CreateServiceClient(serviceURL, creator).(shoppingcartpb.ShoppingCartServiceClient)
}
