package server

import (
	"context"
	"fmt"
	"log"
	db "products-svc/database"
	proto "products-svc/proto"
)

func (s *ProductsServer) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	fmt.Println("Create product invoked")
	err := s.queries.CreateProduct(ctx, db.CreateProductParams{
		Name: req.Name,
		UserID: req.UserId,
		Url: req.Url,
		Title: req.Title,
		Description: req.Description,
		Price: req.Price,
	})
	if err != nil {
		return &proto.CreateProductResponse{
			Response: &proto.Response{
				Success: false,
				Message: "Failed to create product",
			},
		}, err
	}

	return &proto.CreateProductResponse{
		Response: &proto.Response{
			Success: true,
			Message: "Product created successfully",
		},
	}, nil
}

func (s *ProductsServer) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest) (*proto.UpdateProductResponse, error) {
	fmt.Println("Update product invoked")
	err := s.queries.UpdateProduct(ctx, db.UpdateProductParams{
		ProductID: req.Product.ProductId,
		Name: req.Product.Name,
		Url: req.Product.Url,
		Description: req.Product.Description,
		Price: req.Product.Price,
	})
	if err != nil {
		return &proto.UpdateProductResponse{
			Response: &proto.Response{
				Success: false,
				Message: "Failed to update product",
			},
		}, err
	}

	return &proto.UpdateProductResponse{
		Response: &proto.Response{
			Success: true,
			Message: "Product updated successfully",
		},
	}, nil
}
func (s *ProductsServer) GetProduct(ctx context.Context, req *proto.GetProductRequest) (*proto.GetProductResponse, error) {
	log.Println("Get product invoked")
	product, err := s.queries.GetProduct(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}

	return &proto.GetProductResponse{
		Product: &proto.ProductResponse{
			Product: &proto.Product{
				ProductId: product.ProductID,
				Name: product.Name,
				Url: product.Url,
				Price: product.Price,
				Description: product.Description,

			},
		},
	}, nil
}

func (s *ProductsServer) GetProductByName(ctx context.Context, req *proto.GetProductByNameRequest) (*proto.GetProductByNameResponse, error){
	log.Println("Get product by name invoked", req.Name)
	product, err := s.queries.GetProductByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &proto.GetProductByNameResponse{
		Product: &proto.ProductResponse{
			Product: &proto.Product{
				ProductId: product.ProductID,
				Name: product.Name,
				Url: product.Url,
				Price: product.Price,
				Description: product.Description,

			},
		},
	}, nil
}

func (s *ProductsServer) GetProductByUserId(ctx context.Context, req *proto.GetProductByUserIdRequest) (*proto.GetProductByUserIdResponse, error){
	log.Println("Get products by userId invoked", req.UserId)
	products, err := s.queries.GetProductByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var productsResponses []*proto.ProductResponse
	for _, product := range products {
		productsResponses = append(productsResponses, &proto.ProductResponse{
			Product: &proto.Product{
				ProductId: product.ProductID,
				Name: product.Name,
				Url: product.Url,
				Price: product.Price,
				Description: product.Description,
				Title: product.Title,
			},
		})
	}
	return &proto.GetProductByUserIdResponse{
		Product: productsResponses,
	}, nil
}

func (s *ProductsServer) GetProducts(ctx context.Context, req *proto.GetProductsRequest) (*proto.GetProductsResponse, error) {
	log.Println("Get products invoked")
	products, err := s.queries.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	var productsResponses []*proto.ProductResponse
	for _, product := range products {
		productsResponses = append(productsResponses, &proto.ProductResponse{
			Product: &proto.Product{
				ProductId: product.ProductID,
				Name: product.Name,
				Url: product.Url,
				Price: product.Price,
				Description: product.Description,
				Title: product.Title,
			},
		})
	}
	return &proto.GetProductsResponse{
		Products: productsResponses,
	}, nil
}

func (s *ProductsServer) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) (*proto.DeleteProductResponse, error) {
	log.Println("Delete product invoked")
	err := s.queries.DeleteProduct(ctx, req.ProductId)
	if err != nil {
		return &proto.DeleteProductResponse{
			Response: &proto.Response{
				Success: false,
				Message: "Failed to delete product",
			},
		}, err
	}

	return &proto.DeleteProductResponse{
		Response: &proto.Response{
			Success: true,
			Message: "Product deleted successfully",
		},
	}, nil
}

