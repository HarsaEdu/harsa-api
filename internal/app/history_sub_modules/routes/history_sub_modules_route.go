package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (routes *HistorySubModuleRoutesImpl) MobileHistorySubModule(apiGroup *echo.Group) {
	historySubModule := apiGroup.Group("/users/course/module/sub_module")

	historySubModule.POST("", routes.Handler.CreateHistoryModule, middleware.StudentMiddleare)
}
