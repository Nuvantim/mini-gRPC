package service

import (
	"connectrpc.com/connect"
	"context"
	"example/helper"
	"example/repository"

	"errors"
	"github.com/jmoiron/sqlx"
	pb "example/pb/proto/v1"
)

type (
	CreateCategoryRequest = connect.Request[pb.CreateCategoryRequest]
	CreateCategoryResponse = connect.Response[pb.CreateCategoryResponse]
	GetCategoryRequest = connect.Request[pb.GetCategoryRequest]
	GetCategoryResponse = connect.Response[pb.GetCategoryResponse]
	ListCategoriesRequest = connect.Request[pb.ListCategoriesRequest]
	ListCategoriesResponse = connect.Response[pb.ListCategoriesResponse]
	UpdateCategoryRequest = connect.Request[pb.UpdateCategoryRequest]
	UpdateCategoryResponse = connect.Response[pb.UpdateCategoryResponse]
	DeleteCategoryRequest = connect.Request[pb.DeleteCategoryRequest]
	DeleteCategoryResponse = connect.Response[pb.DeleteCategoryResponse]
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(db *sqlx.DB) *CategoryService {
	return &CategoryService{
		repo: repository.NewCategoryRepository(db),
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context,req *CreateCategoryRequest,) (*CreateCategoryResponse, error) {
	category, err := s.repo.Create(ctx, req.Msg.Name, req.Msg.Description)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.CreateCategoryResponse{
		Category: helper.CategoryToProto(category),
	}), nil
}

func (s *CategoryService) GetCategory(ctx context.Context,req *GetCategoryRequest,) (*GetCategoryResponse, error) {
	category, err := s.repo.GetByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if category == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("category not found"))
	}

	return connect.NewResponse(&pb.GetCategoryResponse{
		Category: helper.CategoryToProto(category),
	}), nil
}

func (s *CategoryService) ListCategories(ctx context.Context,req *ListCategoriesRequest,) (*ListCategoriesResponse, error) {
	categories, err := s.repo.List(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	protoCategories := make([]*pb.Category, 0, len(categories))
	for _, cat := range categories {
		protoCategories = append(protoCategories, helper.CategoryToProto(&cat))
	}

	return connect.NewResponse(&pb.ListCategoriesResponse{
		Categories: protoCategories,
	}), nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context,req *UpdateCategoryRequest,) (*UpdateCategoryResponse, error) {
	category, err := s.repo.Update(ctx, req.Msg.Id, req.Msg.Name, req.Msg.Description)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.UpdateCategoryResponse{
		Category: helper.CategoryToProto(category),
	}), nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context,req *DeleteCategoryRequest,) (*DeleteCategoryResponse, error) {
	success, err := s.repo.Delete(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.DeleteCategoryResponse{
		Success: success,
	}), nil
}
