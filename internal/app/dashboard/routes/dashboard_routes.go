package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (courseTrackingRoutes *DashboardRoutesImpl) DashboardWeb(apiGroup *echo.Group) {
	dashboardTrackingsGroup := apiGroup.Group("/admin")

	dashboardTrackingsGroup.GET("/dashboard/count", courseTrackingRoutes.DashboardHandler.CountAll, middleware.AdminMiddleware)

}