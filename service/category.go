package service

import (
	"context"

	"github.com/FianGumilar/restful-api-echo/interfaces"
	"github.com/FianGumilar/restful-api-echo/models/dto"
)

type service struct {
	categoryRepository interfaces.CategoryRepository
}

func NewCategoryService(categoryRepository interfaces.CategoryRepository) interfaces.CategoryService {
	return &service{categoryRepository: categoryRepository}
}

// Create implements interfaces.CategoryService.
func (s service) Create(ctx context.Context, req *dto.ReqCreateCategory) (dto.ResponseCategory, error) {
	category := &dto.Category{Name: req.Name}
	createdCategory := s.categoryRepository.Create(ctx, category)

	responseCategory := dto.ResponseCategory{
		ID:   *createdCategory.GetID(),
		Name: *createdCategory.GetName(),
	}

	return responseCategory, nil
}

// Delete implements interfaces.CategoryService.
func (s service) Delete(ctx context.Context, req *dto.ReqDeleteCategory) error {
	category, err := s.categoryRepository.FindByID(ctx, req.ID)
	if err != nil {
		panic(err.Error())
	}

	return s.categoryRepository.Delete(ctx, &category)
}

// FindAll implements interfaces.CategoryService.
func (s service) FindAll(ctx context.Context) ([]dto.ResponseCategory, error) {
	categories, err := s.categoryRepository.FindAll(ctx)
	if err != nil {
		panic(err.Error())
	}

	var responseCategories []dto.ResponseCategory

	for _, category := range categories {
		responseCategory := dto.ResponseCategory{
			ID:   *category.GetID(),
			Name: *category.GetName(),
		}
		responseCategories = append(responseCategories, responseCategory)
	}
	return responseCategories, nil
}

// FindByID implements interfaces.CategoryService.
func (s service) FindByID(ctx context.Context, id int64) (dto.ResponseCategory, error) {
	category, err := s.categoryRepository.FindByID(ctx, id)
	if err != nil {
		panic(err.Error())
	}

	responseCategory := dto.ResponseCategory{
		ID:   *category.GetID(),
		Name: *category.GetName(),
	}

	return responseCategory, nil
}

// Update implements interfaces.CategoryService.
func (s service) Update(ctx context.Context, req *dto.ReqUpdateCategory) error {
	category, err := s.categoryRepository.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	categoryName := req.Name
	category.SetName(&categoryName)

	return s.categoryRepository.Update(ctx, &category)
}
