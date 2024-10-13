package graph

import (
	"context"
	tutorial "grpc/proto"
	"log"

	"google.golang.org/grpc"
)

type queryResolver struct{}

// Resolver de GraphQL que llama a gRPC
func (r *queryResolver) Person(ctx context.Context, id int) (*model.Person, error) {
	// Conectar con el servidor gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := tutorial.NewPersonServiceClient(conn)

	// Hacer la solicitud gRPC
	req := &tutorial.PersonRequest{Id: int32(id)}
	res, err := client.GetPerson(ctx, req)
	if err != nil {
		return nil, err
	}

	// Convertir la respuesta gRPC a un formato GraphQL
	person := &model.Person{
		ID:    int(res.Person.Id),
		Name:  res.Person.Name,
		Email: res.Person.Email,
	}

	return person, nil
}
