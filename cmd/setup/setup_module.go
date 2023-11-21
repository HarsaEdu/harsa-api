package setup

import (
	moduleHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/module/handler"
	moduleRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/module/repository"
	moduleRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/module/routes"
	moduleServicePkg "github.com/HarsaEdu/harsa-api/internal/app/module/service"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ModuleSetup(db *gorm.DB, validate *validator.Validate, e *echo.Echo) moduleRoutesPkg.ModuleRoutes {
	moduleRepository := moduleRepositoryPkg.NewModuleRepository(db)
	moduleService := moduleServicePkg.NewModuleService(moduleRepository, validate)
	moduleHandler := moduleHandlerPkg.NewModuleHandler(moduleService)
	moduleRoute := moduleRoutesPkg.NewModuleRoutes(moduleHandler)

	return moduleRoute
}
