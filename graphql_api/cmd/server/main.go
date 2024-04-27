package main

import (
	clients "graphql_api/clients/server"
	"graphql_api/directives"
	"graphql_api/graph"
	"graphql_api/pkg/middlewares"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize gRPC clients
	checkoutClient := clients.CheckoutSvc()
	userManagementClient := clients.UsermanagmentSvc()
	productsClient := clients.ProductsSvc()
	shoppingCartClient := clients.ShoppingcartSvc()

	// Pass the clients to the resolver struct
	resolver := &graph.Resolver{
		CheckoutClient:        checkoutClient,
		UserManagementClient:  userManagementClient,
		ProductsClient:        productsClient,
		ShoppingCartClient:    shoppingCartClient,
	}

	// Create GraphQL handler

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver,
		Directives: graph.DirectiveRoot{
			Auth: directives.Auth,
		},
	}))

	// Configure chi router with CORS
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not necessary but useful for preflight requests
	}))
	router.Use(middlewares.AuthMiddleware)
	// Set up HTTP routes
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	// Start server
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

