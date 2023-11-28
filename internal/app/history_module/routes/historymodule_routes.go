package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (historyModuleRoutes *HistoryModuleRoutesImpl) HistoryModuleWeb(apiGroup *echo.Group) {
	historyModuleGroup := apiGroup.Group("/historymodule")

	historyModuleGroup.POST("", historyModuleRoutes.HistoryModuleHandler.Create, middleware.StudentMiddleare)
	historyModuleGroup.GET("", historyModuleRoutes.HistoryModuleHandler.GetAll)
	historyModuleGroup.GET("/:id", historyModuleRoutes.HistoryModuleHandler.FindById)
	historyModuleGroup.PUT("/:id", historyModuleRoutes.HistoryModuleHandler.Update, middleware.StudentMiddleare)
}

func (historyModuleRoutes *HistoryModuleRoutesImpl) HistoryModuleMobile(apiGroup *echo.Group) {
	historyModuleGroup := apiGroup.Group("/historymodule")

	historyModuleGroup.POST("", historyModuleRoutes.HistoryModuleHandler.Create, middleware.StudentMiddleare)
	historyModuleGroup.GET("", historyModuleRoutes.HistoryModuleHandler.GetAll)
	historyModuleGroup.GET("/:id", historyModuleRoutes.HistoryModuleHandler.FindById)
	historyModuleGroup.PUT("/:id", historyModuleRoutes.HistoryModuleHandler.Update, middleware.StudentMiddleare)
}
