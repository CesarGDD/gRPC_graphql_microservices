package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"graphql_api/protos/productspb"
	shoppingcart "graphql_api/protos/shoppingcartpb"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Product is the resolver for the product field.
func (r *productItemResolver) Product(ctx context.Context, obj *shoppingcart.ProductItem) (*productspb.Product, error) {
	fmt.Println("Request to get product from shopping cart:", obj.ProductId)

	res, err := r.ProductsClient.GetProduct(ctx, &productspb.GetProductRequest{
		ProductId: int32(obj.ProductId),
	})
	if err != nil {
		fmt.Printf("Error fetching product %d: %v\n", obj.ProductId, err)
		return nil, handleError(ctx, err, "fetch product")
	}
	if res == nil || res.Product == nil {
		fmt.Printf("No product found for ID %d\n", obj.ProductId)
		graphql.AddError(ctx, gqlerror.Errorf("No product found with ID %d", obj.ProductId))
		return nil, gqlerror.Errorf("No product found.")
	}

	fmt.Println("Product retrieved successfully:", res.Product.Product.Name)
	return res.Product.Product, nil
}

// ProductItem returns ProductItemResolver implementation.
func (r *Resolver) ProductItem() ProductItemResolver { return &productItemResolver{r} }

type productItemResolver struct{ *Resolver }