package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/history_sub_modules/service"
	"github.com/labstack/echo/v4"
)

type HistorySubModuleHandler interface {
	CreateHistoryModule(ctx echo.Context) error
}

type HistorySubModuleHandlerImpl struct {
	Service service.HistorySubModuleService
}

func NewHistorySubModuleHandler(service service.HistorySubModuleService) HistorySubModuleHandler {
	return &HistorySubModuleHandlerImpl{
		Service: service,
	}
}
