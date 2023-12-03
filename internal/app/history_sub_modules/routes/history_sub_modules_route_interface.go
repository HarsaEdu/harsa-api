package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/history_sub_modules/handler"
	"github.com/labstack/echo/v4"
)

type HistorySubModuleRoutes interface {
	MobileHistorySubModule(apiGroup *echo.Group)
}

type HistorySubModuleRoutesImpl struct {
	Handler handler.HistorySubModuleHandler
}

func NewHistorySubModukeRoutes(handler handler.HistorySubModuleHandler) HistorySubModuleRoutes {
	return &HistorySubModuleRoutesImpl{
		Handler: handler,
	}
}
