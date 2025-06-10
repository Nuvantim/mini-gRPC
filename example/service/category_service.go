package service

import (
	"connectrpc.com/connect"
	"context"
	"example/gen/category/v1/categoryv1connect"
	"example/helper"
	"example/repository"

	"example/gen/category/v1"
)

type (
	CreateCategory = connect.Request[categoryv1.CreateCategoryRequest]
	CreateCategoryResponse = connect.Response[categoryv1.CreateCategoryResponse]
	GetCategoryRequest = connect.Request[categoryv1.GetCategoryRequest]
	GetCategoryResponse = connect.Response[categoryv1.GetCategoryResponse]
	ListCategoriesRequest = connect.Request[categoryv1.ListCategoriesRequest]
	ListCategoriesResponse = connect.Response[categoryv1.ListCategoriesResponse]
	UpdateCategoryRequest = connect.Request[categoryv1.UpdateCategoryRequest]
	UpdateCategoryResponse = connect.Response[categoryv1.UpdateCategoryResponse]
	DeleteCategoryRequest = connect.Request[categoryv1.DeleteCategoryRequest]
	DeleteCategoryResponse = connect.Response[categoryv1.DeleteCategoryResponse]
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(db *sqlx.DB) *CategoryService {
	return &CategoryService{
		repo: repository.NewCategoryRepository(db),
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context,req *CreateCategory,) (*CreateCategoryResponse, error) {
	category, err := s.repo.Create(ctx, req.Msg.Name, req.Msg.Description)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&categoryv1.CreateCategoryResponse{
		Category: helper.CategoryToProto(category),
	}), nil
}

func (s *CategoryService) GetCategory(ctx context.Context,req *GetCategory,) (*GetCategoryResponse, error) {
	category, err := s.repo.GetByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if category == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("category not found"))
	}

	return connect.NewResponse(&categoryv1.GetCategoryResponse{
		Category: helper.CategoryToProto(category),
	}), nil
}

func (s *CategoryService) ListCategories(ctx context.Context,req *ListCategoriesRequest,) (*ListCategoriesResponse, error) {
	categories, err := s.repo.List(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	protoCategories := make([]*categoryv1.Category, 0, len(categories))
	for _, cat := range categories {
		protoCategories = append(protoCategories, helper.CategoryToProto(&cat))
	}

	return connect.NewResponse(&categoryv1.ListCategoriesResponse{
		Categories: protoCategories,
	}), nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context,req *UpdateCategoryRequest,) (*UpdateCategoryResponse, error) {
	category, err := s.repo.Update(ctx, req.Msg.Id, req.Msg.Name, req.Msg.Description)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&categoryv1.UpdateCategoryResponse{
		Category: helper.CategoryToProto(category),
	}), nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context,req *DeleteCategoryRequest,) (*DeleteCategoryResponse, error) {
	success, err := s.repo.Delete(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&categoryv1.DeleteCategoryResponse{
		Success: success,
	}), nil
}
