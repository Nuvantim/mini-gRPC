package service

import (
	"connectrpc.com/connect"
	"context"
	"example/internal/helper"
	"example/internal/repository"

	"errors"
	pb "example/pb/proto/category/v1"
)

type (
	CreateCategoryRequest  = connect.Request[pb.CreateCategoryRequest]
	CreateCategoryResponse = connect.Response[pb.CreateCategoryResponse]
	GetCategoryRequest     = connect.Request[pb.GetCategoryRequest]
	GetCategoryResponse    = connect.Response[pb.GetCategoryResponse]
	ListCategoriesRequest  = connect.Request[pb.ListCategoriesRequest]
	ListCategoriesResponse = connect.Response[pb.ListCategoriesResponse]
	UpdateCategoryRequest  = connect.Request[pb.UpdateCategoryRequest]
	UpdateCategoryResponse = connect.Response[pb.UpdateCategoryResponse]
	DeleteCategoryRequest  = connect.Request[pb.DeleteCategoryRequest]
	DeleteCategoryResponse = connect.Response[pb.DeleteCategoryResponse]
)

type CategoryService struct {
	queries *repository.Queries
}

func NewCategoryService(queries *repository.Queries) *CategoryService {
	return &CategoryService{queries: queries}
}

// Create Category
func (s *CategoryService) CreateCategory(ctx context.Context, req *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	category, err := s.queries.CreateCategory(context.Background(), req.Msg.Name)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&pb.CreateCategoryResponse{
		Category: helper.CategoryToProto(category),
	}), nil
}

// Get Category
func (s *CategoryService) GetCategory(ctx context.Context, req *GetCategoryRequest) (*GetCategoryResponse, error) {
	category, err := s.queries.GetCategory(context.Background(), req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.GetCategoryResponse{
		Category: helper.CategoryToProto(category),
	}), nil
}

// List Category
func (s *CategoryService) ListCategories(ctx context.Context, req *ListCategoriesRequest) (*ListCategoriesResponse, error) {
	categories, err := s.queries.ListCategory(context.Background())
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	protoCategories := make([]*pb.Category, len(categories))
	for i, ctg := range categories {
		protoCategories[i] = helper.CategoryToProto(ctg)
	}

	return connect.NewResponse(&pb.ListCategoriesResponse{
		Categories: protoCategories,
	}), nil
}

// Update Category
func (s *CategoryService) UpdateCategory(ctx context.Context, req *UpdateCategoryRequest) (*UpdateCategoryResponse, error) {
	if req.Msg.Name == "" {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("name is required"))
	}
	data := repository.UpdateCategoryParams{
		ID:   req.Msg.Id,
		Name: req.Msg.Name,
	}

	category, err := s.queries.UpdateCategory(context.Background(), data)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.UpdateCategoryResponse{
		Category: helper.CategoryToProto(category),
	}), nil
}

// Delete Category
func (s *CategoryService) DeleteCategory(ctx context.Context, req *DeleteCategoryRequest) (*DeleteCategoryResponse, error) {
	if err := s.queries.DeleteCategory(context.Background(), req.Msg.Id); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.DeleteCategoryResponse{
		Success: true,
	}), nil
}
