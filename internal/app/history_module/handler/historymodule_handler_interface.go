package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/history_module/service"
	"github.com/labstack/echo/v4"
)

type HistoryModuleHandler interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	FindById(ctx echo.Context) error
	GetAll(ctx echo.Context) error
}

type HistoryModuleHandlerImpl struct {
	HistoryModuleService service.HistoryModuleService
}

func NewHistoryModuleHandler(service service.HistoryModuleService) HistoryModuleHandler {
	return &HistoryModuleHandlerImpl{HistoryModuleService: service}
}
