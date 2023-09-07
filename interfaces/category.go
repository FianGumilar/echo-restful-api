package interfaces

import (
	"context"

	"github.com/FianGumilar/restful-api-echo/models/dto"
	"github.com/labstack/echo/v4"
)

type CategoryRepository interface {
	FindAll(ctx context.Context) ([]dto.Category, error)
	FindByID(ctx context.Context, id int64) (dto.Category, error)
	Create(ctx context.Context, category *dto.Category) dto.Category
	Update(ctx context.Context, categry *dto.Category) error
	Delete(ctx context.Context, category *dto.Category) error
}

type CategoryService interface {
	FindAll(ctx context.Context) ([]dto.ResponseCategory, error)
	FindByID(ctx context.Context, id int64) (dto.ResponseCategory, error)
	Create(ctx context.Context, req *dto.ReqCreateCategory) (dto.ResponseCategory, error)
	Update(ctx context.Context, req *dto.ReqUpdateCategory) error
	Delete(ctx context.Context, re *dto.ReqDeleteCategory) error
}

type CategoryHandler interface {
	FindAll(ctx echo.Context) error
	FindByID(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
