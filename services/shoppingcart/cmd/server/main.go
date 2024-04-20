package main

import (
	"log"
	"net"
	"shoppingcart/server"
	proto "shoppingcart/proto"
	"google.golang.org/grpc"
)

func main() {
	// Initialize the gRPC server
	grpcServer := grpc.NewServer()

	// Register the ShoppingCart service with the server
	shoppingCartServer, err := server.NewShoppingCartServer()
	if err != nil {
		log.Fatalf("Failed to initialize ShoppingCart server: %v", err)
	}
	proto.RegisterShoppingCartServiceServer(grpcServer, shoppingCartServer)

	// Start the server
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Server started on :50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
