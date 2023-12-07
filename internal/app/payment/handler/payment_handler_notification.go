package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (paymentHandler *PaymentHandlerImpl) NotificationPayment(ctx echo.Context) error {
	var notificationPayload map[string]interface{}
	err := ctx.Bind(&notificationPayload)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind request", err)
	}

	err = paymentHandler.PaymentService.NotificationPayment(notificationPayload)
	if err != nil {
		if strings.Contains(err.Error(), "error when get order id") {
			return res.StatusAlreadyExist(ctx, "order id not found", err)
		}

		res.StatusInternalServerError(ctx, "failed to get notification payment", err)
	}

	return res.StatusOK(ctx, "ok", nil, nil)
}