package Dashboard

import (
	DashboardHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/dashboard/handler"
	DashboardRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/dashboard/repository"
	DashboardRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/dashboard/routes"
	DashboardServicePkg "github.com/HarsaEdu/harsa-api/internal/app/dashboard/service"
	"gorm.io/gorm"
)

func DashboardSetup(db *gorm.DB) DashboardRoutesPkg.DashboardRoutes{
	DashboardRepository := DashboardRepositoryPkg.NewDashboardRepository(db)
	DashboardService := DashboardServicePkg.NewDashboardService(DashboardRepository)
	DashboardHandler := DashboardHandlerPkg.NewDashboardHandler(DashboardService)
	DashboardRoute := DashboardRoutesPkg.NewDashboardRoutes(DashboardHandler)

	return DashboardRoute
}

