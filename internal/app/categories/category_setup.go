package category

import (
	categoryHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/categories/handler"
	categoryRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/categories/repository"
	categoryRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/categories/routes"
	categoryServicePkg "github.com/HarsaEdu/harsa-api/internal/app/categories/service"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func CategorySetup(db *gorm.DB, validate *validator.Validate,cloudinary cloudinary.CloudinaryUploader) categoryRoutesPkg.CategoryRoutes{
	categoryRepository := categoryRepositoryPkg.NewCategoryRepository(db)
	categoryService := categoryServicePkg.NewCategoryService(categoryRepository,validate,cloudinary)
	categoryHandler := categoryHandlerPkg.NewCategoryHandler(categoryService)
	categoryRoute := categoryRoutesPkg.NewCategoryRoutes(categoryHandler)

	return categoryRoute
}
