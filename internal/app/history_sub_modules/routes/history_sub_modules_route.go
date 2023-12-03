package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (routes *HistorySubModuleRoutesImpl) MobileHistorySubModule(apiGroup *echo.Group) {
	historySubModule := apiGroup.Group("/users/course/module/sub_module")

	historySubModule.POST("/:sub-module-id", routes.Handler.CreateHistoryModule, middleware.StudentMiddleare)
	// historySubModule.GET("", routes.Handler.GetHistoryModule, middleware.StudentMiddleare)
	historySubModule.PUT("/:sub_module_id", routes.Handler.UpdateHistorySubModule, middleware.StudentMiddleare)
}
