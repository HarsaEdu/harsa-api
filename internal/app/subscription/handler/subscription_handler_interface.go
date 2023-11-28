package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/subscription/service"
	"github.com/labstack/echo/v4"
)

type SubscriptionHandler interface {
	Subscribe(ctx echo.Context) error
}

type SubscriptionHandlerImpl struct {
	SubscriptionService service.SubscriptionService
}

func NewSubscriptionHandler(subscriptionService service.SubscriptionService) SubscriptionHandler {
	return &SubscriptionHandlerImpl{
		SubscriptionService: subscriptionService,
	}
}