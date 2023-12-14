package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/dashboard/handler"
	"github.com/labstack/echo/v4"
)

type DashboardRoutes interface {
	DashboardWeb(apiGroup *echo.Group)
}

type DashboardRoutesImpl struct {
	DashboardHandler handler.DashboardHandler
}

func NewDashboardRoutes(DashboardHandler handler.DashboardHandler) DashboardRoutes {
	return &DashboardRoutesImpl{
		DashboardHandler: DashboardHandler,
	}
}