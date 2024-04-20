package main

import (
	"graphql_api/graph"
	"graphql_api/clients/server"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
