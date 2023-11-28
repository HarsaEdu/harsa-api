package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (subscriptionHandler *SubscriptionHandlerImpl) Subscribe(ctx echo.Context) error {
	request := web.CreateSubscribeRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind subscription request", err)
	}

	userId := ctx.Get("user_id").(uint)

	response, err := subscriptionHandler.SubscriptionService.Subscribe(&request, userId)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "error when get user") {
			return res.StatusAlreadyExist(ctx, "user not found", err)
		}
		if strings.Contains(err.Error(), "error when get subscription plan") {
			return res.StatusNotFound(ctx, "subscription plan not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to subscribe , something happen", err)
	}

	return res.StatusCreated(ctx, "success to subscribe", response, nil)
}