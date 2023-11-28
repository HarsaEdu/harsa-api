package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/subscription/handler"
	"github.com/labstack/echo/v4"
)

type SubscriptionRoutes interface {
	// SubscriptionWeb(apiGroup *echo.Group) *echo.Group
	SubscriptionMobile(apiGroup *echo.Group) *echo.Group
}

type SubscriptionRoutesImpl struct {
	SubscriptionHandler handler.SubscriptionHandler
}

func NewSubscriptionRoutes(SubscriptionHandler handler.SubscriptionHandler) SubscriptionRoutes {
	return &SubscriptionRoutesImpl{
		SubscriptionHandler: SubscriptionHandler,
	}
}