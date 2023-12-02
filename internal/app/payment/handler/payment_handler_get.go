package handler

import (
	"strconv"
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
	params := ctx.QueryParams()
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid offset", err)
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit", err)
	}

	response, pagination, err := paymentHandler.PaymentService.GetAllPaymentHistory(offset, limit, params.Get("search"), params.Get("status"))
	if err != nil {
		if response == nil {
			return res.StatusNotFound(ctx, "payment history not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get payment history, something happen", err)
	}

	return res.StatusOK(ctx, "success to get payment history", response, pagination)
}

func (paymentHandler *PaymentHandlerImpl) GetAllPaymentHistoryByUserId(ctx echo.Context) error {
	userId := ctx.Get("user_id").(uint)

	params := ctx.QueryParams()
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid offset", err)
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit", err)
	}
	response, pagination, err := paymentHandler.PaymentService.GetAllPaymentHistoryByUserId(userId, offset, limit, params.Get("search"), params.Get("status"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "payment history not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get payment history, something happen", err)
	}

	return res.StatusOK(ctx, "success to get payment history", response, pagination)
}

func (paymentHandler *PaymentHandlerImpl) GetPaymentHistoryByUserIdAndPaymentId(ctx echo.Context) error {
	userId := ctx.Get("user_id").(uint)
	paymentId := ctx.Param("id")

	response, err := paymentHandler.PaymentService.GetPaymentHistoryByUserIdAndPaymentId(userId, paymentId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "payment history not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get payment history, something happen", err)
	}

	return res.StatusOK(ctx, "success to get payment history", response, nil)
}