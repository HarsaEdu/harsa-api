package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/history_module/handler"
	"github.com/labstack/echo/v4"
)

type HistoryModuleRoutes interface {
	HistoryModuleWeb(apiGroup *echo.Group)
	HistoryModuleMobile(apiGroup *echo.Group)
}

type HistoryModuleRoutesImpl struct {
	Echo                 *echo.Echo
	HistoryModuleHandler handler.HistoryModuleHandler
}

func NewHistoryModuleRoutes(historyModuleHandler handler.HistoryModuleHandler) HistoryModuleRoutes {
	return &HistoryModuleRoutesImpl{
		HistoryModuleHandler: historyModuleHandler,
	}
}
