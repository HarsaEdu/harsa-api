package subsplan

import (
	subsPlanHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/subs_plan/handler"
	subsPlanRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/subs_plan/repository"
	subsPlanRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/subs_plan/routes"
	subsPlanServicePkg "github.com/HarsaEdu/harsa-api/internal/app/subs_plan/service"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func SubsPlanSetup(db *gorm.DB, validate *validator.Validate, cloudinary cloudinary.CloudinaryUploader) (subsPlanRoutesPkg.SubsPlanRoutes, subsPlanRepositoryPkg.SubsPlanRepository) {
	subsPlanRepository := subsPlanRepositoryPkg.NewSubsPlanRepository(db)
	subsPlanService := subsPlanServicePkg.NewsubsPlanService(subsPlanRepository, validate, cloudinary)
	subsPlanHandler := subsPlanHandlerPkg.NewFaqsHandler(subsPlanService)
	subsPlanRoutes := subsPlanRoutesPkg.NewSubsPlanRoutes(subsPlanHandler)

	return subsPlanRoutes, subsPlanRepository
}
