package setup

import (
	categoryHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/categories/handler"
	categoryRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/categories/repository"
	categoryRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/categories/routes"
	categoryServicePkg "github.com/HarsaEdu/harsa-api/internal/app/categories/service"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CategorySetup(db *gorm.DB, validate *validator.Validate,cloudinary cloudinary.CloudinaryUploader ,e *echo.Echo) categoryRoutesPkg.CategoryRoutes{
	categoryRepository := categoryRepositoryPkg.NewCategoryRepository(db)
	categoryService := categoryServicePkg.NewCategoryService(categoryRepository,validate,cloudinary)
	categoryHandler := categoryHandlerPkg.NewCategoryHandler(categoryService)
	categoryRoute := categoryRoutesPkg.NewCategoryRoutes(e, categoryHandler)

	return categoryRoute
}
