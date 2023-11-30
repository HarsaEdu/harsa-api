package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (paymentHandler *PaymentHandlerImpl) GetPaymentHistoryById(ctx echo.Context) error {
	id := ctx.Param("id")

	response, err := paymentHandler.PaymentService.GetPaymentHistoryById(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "payment history not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get payment history, something happen", err)
	}

	return res.StatusOK(ctx, "success to get payment history", response, nil)
}

func (paymentHandler *PaymentHandlerImpl) GetAllPaymentHistory(ctx echo.Context) error {
	response, err := paymentHandler.PaymentService.GetAllPaymentHistory()
	if err != nil {
		if response == nil {
			return res.StatusNotFound(ctx, "payment history not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get payment history, something happen", err)
	}

	return res.StatusOK(ctx, "success to get payment history", response, nil)
}

func (paymentHandler *PaymentHandlerImpl) GetAllPaymentHistoryByUserId(ctx echo.Context) error {
	userId := ctx.Get("user_id").(uint)

	response, err := paymentHandler.PaymentService.GetAllPaymentHistoryByUserId(userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "payment history not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get payment history, something happen", err)
	}

	return res.StatusOK(ctx, "success to get payment history", response, nil)
}