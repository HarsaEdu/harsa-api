package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (categoryRoutes *CategoryRoutesImpl) Category(apiGroup *echo.Group) {
	categoriesGroup := apiGroup.Group("/categories")

	categoriesGroup.POST("", categoryRoutes.CategoryHandler.Create, middleware.AdminMiddleware)
	categoriesGroup.GET("", categoryRoutes.CategoryHandler.GetAll, middleware.AllUserMiddleare)
	categoriesGroup.PUT("/:id", categoryRoutes.CategoryHandler.Update, middleware.AdminMiddleware)
	categoriesGroup.PATCH("/:id", categoryRoutes.CategoryHandler.UploadImage, middleware.AdminMiddleware)
	categoriesGroup.DELETE("/:id", categoryRoutes.CategoryHandler.Delete, middleware.AdminMiddleware)
	categoriesGroup.GET("/:id", categoryRoutes.CategoryHandler.FindById, middleware.AdminMiddleware)

}
