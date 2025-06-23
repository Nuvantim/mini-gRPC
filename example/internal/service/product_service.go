package service

import (
	"connectrpc.com/connect"
	"context"
	"example/database"
	"example/internal/helper"
	"example/internal/repository"

	"errors"
	pb "example/pb/proto/product/v1"
)

type (
	CreateProductRequest  = connect.Request[pb.CreateProductRequest]
	CreateProductResponse = connect.Response[pb.CreateProductResponse]
	GetProductRequest     = connect.Request[pb.GetProductRequest]
	GetProductResponse    = connect.Response[pb.GetProductResponse]
	ListProductRequest    = connect.Request[pb.ListProductRequest]
	ListProductResponse   = connect.Response[pb.ListProductResponse]
	UpdateProductRequest  = connect.Request[pb.UpdateProductRequest]
	UpdateProductResponse = connect.Response[pb.UpdateProductResponse]
	DeleteProductRequest  = connect.Request[pb.DeleteProductRequest]
	DeleteProductResponse = connect.Response[pb.DeleteProductResponse]
)

type ProductService struct {
	queries *repository.Queries
}

func NewProductService(queries *repository.Queries) *ProductService {
	return &ProductService{queries: queries}
}

// Create Product
// Get Product
// ListProduct
// UpdateProduct
// DeleteProduct
