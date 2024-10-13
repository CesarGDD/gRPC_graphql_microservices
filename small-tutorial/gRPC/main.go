package main

import (
	"context"
	tutorial "grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	tutorial.UnimplementedPersonServiceServer
}

func (s *server) GetPerson(ctx context.Context, req *tutorial.PersonRequest) (*tutorial.PersonResponse, error) {
	// Lógica de búsqueda de persona (aquí devolvemos una persona fija)
	person := &tutorial.Person{
		Id:    req.Id,
		Name:  "Juan Perez",
		Email: "juanperez@example.com",
	}

	return &tutorial.PersonResponse{Person: person}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	tutorial.RegisterPersonServiceServer(s, &server{})

	log.Printf("gRPC server listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
