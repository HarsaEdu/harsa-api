package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/dashboard/service"
	"github.com/labstack/echo/v4"
)

type DashboardHandler interface {
	CountAll(ctx echo.Context) error
}

type DashboardHandlerImpl struct {
	DashboardService service.DashboardService
}

func NewDashboardHandler(service service.DashboardService) DashboardHandler {
	return &DashboardHandlerImpl{DashboardService: service}
}