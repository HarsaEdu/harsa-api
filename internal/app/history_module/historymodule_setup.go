package faqs

import (
	historyModuleHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/history_module/handler"
	historyModuleRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/history_module/repository"
	historyModuleRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/history_module/routes"
	historyModuleServicePkg "github.com/HarsaEdu/harsa-api/internal/app/history_module/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func HistoryModuleSetup(db *gorm.DB, validate *validator.Validate) historyModuleRoutesPkg.HistoryModuleRoutes {
	historyModuleRepository := historyModuleRepositoryPkg.NewHistoryModuleRepository(db)
	historyModuleService := historyModuleServicePkg.NewHistoryModuleService(historyModuleRepository, validate)
	historyModuleHandler := historyModuleHandlerPkg.NewHistoryModuleHandler(historyModuleService)
	historyModuleRoute := historyModuleRoutesPkg.NewHistoryModuleRoutes(historyModuleHandler)

	return historyModuleRoute
}
