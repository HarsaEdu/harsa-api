package routes

import "github.com/labstack/echo/v4"

func (categoryRoutes *CategoryRoutesImpl) Category(apiGroup *echo.Group) {
	categoriesGroup := apiGroup.Group("/categories")

	categoriesGroup.POST("", categoryRoutes.CategoryHandler.Create)
	categoriesGroup.GET("", categoryRoutes.CategoryHandler.GetAll)
	categoriesGroup.PUT("/:id", categoryRoutes.CategoryHandler.Update)
	categoriesGroup.PATCH("/:id", categoryRoutes.CategoryHandler.UploadImage)
	categoriesGroup.DELETE("/:id", categoryRoutes.CategoryHandler.Delete)
	categoriesGroup.GET("/:id", categoryRoutes.CategoryHandler.FindById)

}
