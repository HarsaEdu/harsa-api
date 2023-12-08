package history_sub_modules

import (
	HandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/history_sub_modules/handler"
	RepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/history_sub_modules/repository"
	RoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/history_sub_modules/routes"
	ServicePkg "github.com/HarsaEdu/harsa-api/internal/app/history_sub_modules/service"
	subscriptionServicePkg "github.com/HarsaEdu/harsa-api/internal/app/subscription/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func HistorySubModuleSetup(db *gorm.DB, validate *validator.Validate, subcriptionService subscriptionServicePkg.SubscriptionService) RoutesPkg.HistorySubModuleRoutes {
	Repository := RepositoryPkg.NewHistorySubModuleRepository(db)
	Service := ServicePkg.NewHistorySubModuleRepository(Repository, validate)
	Handler := HandlerPkg.NewHistorySubModuleHandler(Service, subcriptionService)
	Route := RoutesPkg.NewHistorySubModukeRoutes(Handler)

	return Route
}
