package main

import (
	"log"
	"net"
	proto "usermanagmentService/proto"
	"usermanagmentService/server"

	"google.golang.org/grpc"
)

func main() {
	// Initialize the gRPC server
	grpcServer := grpc.NewServer()

	// Register the UserManagement service with the server
	userManagementServer, err := server.NewUserManagementServer()
	if err != nil {
		log.Fatalf("Failed to initialize UserManagement server: %v", err)
	}
	proto.RegisterUserManagementServiceServer(grpcServer, userManagementServer)

	// Start the server
	listener, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Server started on :50050")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
