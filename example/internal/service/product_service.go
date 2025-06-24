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
func (s *ProductService) CreateProduct(ctx context.Context, req *CreateProductRequest)(*CreateProductResponse, error){
        tx,err := database.DB.Begin(ctx.Background())
        err != nil {
                return nil, connect.NewError(connect.CodeNotFound)
        }
}
// Get Product
func (s *ProductService) GetProduct(ctx context.Context, req *GetProductRequest)(*GetProductResponse, error){}
// ListProduct
func (s *ProductService) ListProduct(ctx context.Context, req *ListProductRequest)(*ListProductResponse, error){}
// UpdateProduct
func (s *ProductService) UpdateProduct(ctx context.Context, req *UpdateProductRequest)(*UpdateProductResponse, error){}
// DeleteProduct
func (s *ProductService) DeleteProduct(ctx context.Context, req *DeleteProductRequest)(*DeleteProductResponse, error){}
