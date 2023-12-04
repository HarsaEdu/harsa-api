package history_sub_modules

import (
	HandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/history_sub_modules/handler"
	RepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/history_sub_modules/repository"
	RoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/history_sub_modules/routes"
	ServicePkg "github.com/HarsaEdu/harsa-api/internal/app/history_sub_modules/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func HistorySubModuleSetup(db *gorm.DB, validate *validator.Validate) RoutesPkg.HistorySubModuleRoutes {
	Repository := RepositoryPkg.NewHistorySubModuleRepository(db)
	Service := ServicePkg.NewHistorySubModuleRepository(Repository, validate)
	Handler := HandlerPkg.NewHistorySubModuleHandler(Service)
	Route := RoutesPkg.NewHistorySubModukeRoutes(Handler)

	return Route
}
