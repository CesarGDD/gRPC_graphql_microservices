package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"usermanagmentService/database"
	proto "usermanagmentService/proto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserManagementServer struct {
	proto.UnimplementedUserManagementServiceServer
	queries *db.Queries
}

func NewUserManagementServer() (*UserManagementServer, error) {
	// Read environment variables
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")

	// Create connection string
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s", user, password, host, dbName)

	// Connect to database using a connection pool
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Failed to parse pool configuration: %v", err)
		return nil, err
	}

	// Optionally, configure the pool here (e.g., MaxConns, MinConns)
	poolConfig.MaxConns = 10 // Example: setting max connections in the pool

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalf("Failed to establish connection pool: %v", err)
		return nil, err
	}

	queries := db.New(pool) // Assuming `db.New` accepts a pgxpool.Pool

	server := &UserManagementServer{
		queries: queries,
	}

	return server, nil
}
