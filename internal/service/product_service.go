package service

import (
	"connectrpc.com/connect"
	"context"
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
func (s *ProductService) CreateProduct(ctx context.Context, req *CreateProductRequest) (*CreateProductResponse, error) {
	// Search Category
	category, _ := s.queries.CountCategory(context.Background(), req.Msg.CategoryId)

	if category == 0 {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("category not found"))
	}
	// input data from message protobuf
	var data = repository.CreateProductParams{
		Name:        req.Msg.Name,
		Description: req.Msg.Description,
		CategoryID:  req.Msg.CategoryId,
		Price:       req.Msg.Price,
	}
	// execution queries
	product, err := s.queries.CreateProduct(context.Background(), data)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// return to protobuf message
	return connect.NewResponse(&pb.CreateProductResponse{
		Product: helper.ProductToProto(product),
	}), nil

}

// Get Product
func (s *ProductService) GetProduct(ctx context.Context, req *GetProductRequest) (*GetProductResponse, error) {
	// Input id form protobuf message
	product, err := s.queries.GetProduct(context.Background(), req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if product.Product.ID == 0 {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("product is not found"))
	}

	// return to protobuf message
	return connect.NewResponse(&pb.GetProductResponse{
		Products: helper.ProductsToProto(product),
	}), nil

}

// ListProduct
func (s *ProductService) ListProduct(ctx context.Context, req *ListProductRequest) (*ListProductResponse, error) {
	products, err := s.queries.ListProduct(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var protoProducts []*pb.ProductWithCategory
	for _, prd := range products {
		protoProducts = append(protoProducts, helper.ListToProto(prd)...)
	}

	return connect.NewResponse(&pb.ListProductResponse{
		Products: protoProducts,
	}), nil
}

// UpdateProduct
func (s *ProductService) UpdateProduct(ctx context.Context, req *UpdateProductRequest) (*UpdateProductResponse, error) {
	// input data from message protobuf
	var data = repository.UpdateProductParams{
		ID:          req.Msg.Id,
		Name:        req.Msg.Name,
		Description: req.Msg.Description,
		CategoryID:  req.Msg.CategoryId,
		Price:       req.Msg.Price,
	}
	// execution queries
	product, err := s.queries.UpdateProduct(context.Background(), data)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// return to protobuf message
	return connect.NewResponse(&pb.UpdateProductResponse{
		Product: helper.ProductToProto(product),
	}), nil
}

// DeleteProduct
func (s *ProductService) DeleteProduct(ctx context.Context, req *DeleteProductRequest) (*DeleteProductResponse, error) {
	// Input id form protobuf message
	if err := s.queries.DeleteProduct(context.Background(), req.Msg.Id); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// return to protobuf message
	return connect.NewResponse(&pb.DeleteProductResponse{
		Success: true,
	}), nil
}
