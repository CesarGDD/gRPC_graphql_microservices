package server

import (
	"context"
	"fmt"
	"log"
	"os"
	db "checkout/database"
	proto "checkout/proto"

	"github.com/jackc/pgx/v5"
)

type CheckoutServer struct {
	proto.UnimplementedCheckoutServiceServer
	queries *db.Queries
}

func NewCheckoutServer() (*CheckoutServer, error) {
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

	server := &CheckoutServer{
		queries: queries,
	}

	return server, nil
}