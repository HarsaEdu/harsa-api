package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/categories/service"
	"github.com/labstack/echo/v4"
)

type CategoryHandler interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	UploadImage(ctx echo.Context) error
	FindById(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type CategoryHandlerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) CategoryHandler {
	return &CategoryHandlerImpl{CategoryService: service}
}
