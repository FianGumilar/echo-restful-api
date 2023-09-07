package handler

import (
	"net/http"
	"strconv"

	"github.com/FianGumilar/restful-api-echo/interfaces"
	"github.com/FianGumilar/restful-api-echo/models/dto"
	echo "github.com/labstack/echo/v4"
)

type handler struct {
	categoryServices interfaces.CategoryService
}

func NewCategoryHandler(categoryServices interfaces.CategoryService) interfaces.CategoryHandler {
	return &handler{
		categoryServices: categoryServices,
	}
}

// Create implements interfaces.CategoryHandler.
func (h handler) Create(ctx echo.Context) error {
	requestCategory := new(dto.ReqCreateCategory)
	ctx.Bind(requestCategory)

	_, err := h.categoryServices.Create(ctx.Request().Context(), requestCategory)
	if err != nil {
		panic(err.Error())
	}

	return ctx.JSON(http.StatusOK, dto.ApiResponse{
		Code:   201,
		Status: "OK",
	})
}

// Delete implements interfaces.CategoryHandler.
func (h handler) Delete(ctx echo.Context) error {
	requestCategory := new(dto.ReqDeleteCategory)
	err := ctx.Bind(requestCategory)
	if err != nil {
		return err
	}

	h.categoryServices.Delete(ctx.Request().Context(), requestCategory)

	return ctx.JSON(http.StatusOK, dto.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   requestCategory,
	})
}

// FindAll implements interfaces.CategoryHandler.
func (h handler) FindAll(ctx echo.Context) error {
	requestCategories, err := h.categoryServices.FindAll(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, dto.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   requestCategories,
	})
}

// FindByID implements interfaces.CategoryHandler.
func (h handler) FindByID(ctx echo.Context) error {
	idS := ctx.Param("id")
	id, err := strconv.Atoi(idS)

	if err != nil {
		return err
	}

	responseCategory, err := h.categoryServices.FindByID(ctx.Request().Context(), int64(id))

	return ctx.JSON(http.StatusOK, dto.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseCategory,
	})
}

// Update implements interfaces.CategoryHandler.
func (h handler) Update(ctx echo.Context) error {
	requestCategory := new(dto.ReqUpdateCategory)
	err := ctx.Bind(requestCategory)
	if err != nil {
		return err
	}

	h.categoryServices.Update(ctx.Request().Context(), requestCategory)

	return ctx.JSON(http.StatusOK, dto.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   requestCategory,
	})
}
