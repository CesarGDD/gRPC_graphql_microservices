package main

import (
	"log"
	"net"
	"products-svc/server"
	proto "products-svc/proto"
	"google.golang.org/grpc"
)

func main() {
	// Initialize the gRPC server
	grpcServer := grpc.NewServer()

	// Register the Products service with the server
	productsServer, err := server.NewProductsServer()
	if err != nil {
		log.Fatalf("Failed to initialize Products server: %v", err)
	}
	proto.RegisterProductServiceServer(grpcServer, productsServer)

	// Start the server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Server started on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
