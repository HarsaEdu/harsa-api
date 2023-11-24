package options

import (
	optionsHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/options/handler"
	optionsRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/options/repository"
	optionsRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/options/routes"
	optionsServicePkg "github.com/HarsaEdu/harsa-api/internal/app/options/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func OptionsSetup(db *gorm.DB, validate *validator.Validate)optionsRoutesPkg.OptionsRoutes {
	optionsRepository := optionsRepositoryPkg.NewOptionsRepository(db)
	optionsService := optionsServicePkg.NewOptionsService(optionsRepository, validate)
	optionsHandler := optionsHandlerPkg.NewOptionsHandler(optionsService)
	optionsRoute := optionsRoutesPkg.NewOptionsRoutes(optionsHandler)

	return optionsRoute
}
