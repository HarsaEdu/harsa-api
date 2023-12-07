package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/history_sub_modules/service"
	"github.com/labstack/echo/v4"
	subscriptionServicePkg "github.com/HarsaEdu/harsa-api/internal/app/subscription/service"
)

type HistorySubModuleHandler interface {
	CreateHistoryModule(ctx echo.Context) error
	// GetHistoryModule(ctx echo.Context) error
	// UpdateHistorySubModule(ctx echo.Context) error
}

type HistorySubModuleHandlerImpl struct {
	Service service.HistorySubModuleService
	SubcriptionService subscriptionServicePkg.SubscriptionService
}

func NewHistorySubModuleHandler(service service.HistorySubModuleService, subcriptionService subscriptionServicePkg.SubscriptionService) HistorySubModuleHandler {
	return &HistorySubModuleHandlerImpl{
		SubcriptionService: subcriptionService,
		Service: service,
	}
}
