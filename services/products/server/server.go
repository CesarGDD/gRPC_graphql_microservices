package server

import (
	"context"
	"fmt"
	"log"
	"os"
	db "products-svc/database"
	proto "products-svc/proto"

	"github.com/jackc/pgx/v5"
)

type ProductsServer struct {
	proto.UnimplementedProductServiceServer
	queries *db.Queries
}

func NewProductsServer() (*ProductsServer, error) {
	// Read environment variables
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")

	// Create connection string
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s", user, password, host, dbName)

	// Connect to database
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	queries := db.New(conn)

	server := &ProductsServer{
		queries: queries,
	}

	return server, nil
}