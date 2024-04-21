package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"graphql_api/protos/productspb"
	"graphql_api/protos/usermanagementpb"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/status"
)

// UserOwner is the resolver for the userOwner field.
func (r *productResolver) UserOwner(ctx context.Context, obj *productspb.Product) (*usermanagementpb.User, error) {
	fmt.Println("Request to get user owner:", obj.UserId)

	res, err := r.UserManagementClient.GetUser(ctx, &usermanagementpb.GetUserRequest{
		UserId: int32(obj.UserId),
	})
	if err != nil {
		fmt.Printf("Error fetching user owner %d: %v\n", obj.UserId, err)
		// Convert gRPC error to GraphQL error
		e, ok := status.FromError(err)
		if ok {
			// gRPC specific error handling
			fmt.Printf("gRPC error status: %v\n", e.Message())
			graphql.AddError(ctx, gqlerror.Errorf("gRPC error: %s", e.Message()))
		} else {
			// General error handling
			fmt.Printf("Non-gRPC error: %v\n", err)
			graphql.AddError(ctx, gqlerror.Errorf("Internal server error: %v", err))
		}
		return nil, gqlerror.Errorf("Failed to fetch user owner details.")
	}
	if res == nil || res.User == nil {
		fmt.Printf("No user owner found for ID %d\n", obj.UserId)
		graphql.AddError(ctx, gqlerror.Errorf("No user owner found with ID %d", obj.UserId))
		return nil, gqlerror.Errorf("No user owner found.")
	}

	fmt.Println("user owner retrieved successfully:", res.User)
	return &usermanagementpb.User{
		UserId:   res.User.UserId,
		Username: res.User.Username,
	}, nil
}

// Product returns ProductResolver implementation.
func (r *Resolver) Product() ProductResolver { return &productResolver{r} }

type productResolver struct{ *Resolver }