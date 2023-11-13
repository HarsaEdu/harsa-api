package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/categories/handler"
	"github.com/labstack/echo/v4"
)

type CategoryRoutes interface {
	Category(apiGroup *echo.Group)
}

type CategoryRoutesImpl struct {
	Echo            *echo.Echo
	CategoryHandler handler.CategoryHandler
}

func NewCategoryRoutes(e *echo.Echo, categoryHandler handler.CategoryHandler) CategoryRoutes {
	return &CategoryRoutesImpl{
		Echo:            e,
		CategoryHandler: categoryHandler,
	}
}
