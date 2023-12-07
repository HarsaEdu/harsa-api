package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/categories/handler"
	"github.com/labstack/echo/v4"
)

type CategoryRoutes interface {
	CategoryWeb(apiGroup *echo.Group)
	CategoryMobile(apiGroup *echo.Group)
}

type CategoryRoutesImpl struct {
	CategoryHandler handler.CategoryHandler
}

func NewCategoryRoutes(categoryHandler handler.CategoryHandler) CategoryRoutes {
	return &CategoryRoutesImpl{
		CategoryHandler: categoryHandler,
	}
}
