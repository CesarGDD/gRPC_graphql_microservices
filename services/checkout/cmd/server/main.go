package main

import (
	"log"
	"net"
	"checkout/server"
	proto "checkout/proto"
	"google.golang.org/grpc"
)

func main() {
	// Initialize the gRPC server
	grpcServer := grpc.NewServer()

	// Register the Checkout service with the server
	checkoutServer, err := server.NewCheckoutServer()
	if err != nil {
		log.Fatalf("Failed to initialize ShoppingCart server: %v", err)
	}
	proto.RegisterCheckoutServiceServer(grpcServer, checkoutServer)

	// Start the server
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Server started on :50053")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
