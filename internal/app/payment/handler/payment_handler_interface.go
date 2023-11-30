package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/payment/service"
	"github.com/labstack/echo/v4"
)

type PaymentHandler interface {
	CreatePaymentSubscription(ctx echo.Context) error
	NotificationPayment(ctx echo.Context) error
	GetPaymentHistoryById(ctx echo.Context) error
	GetAllPaymentHistory(ctx echo.Context) error
	GetAllPaymentHistoryByUserId(ctx echo.Context) error
}

type PaymentHandlerImpl struct {
	PaymentService service.PaymentService
}

func NewPaymentHandler(paymentService service.PaymentService) PaymentHandler {
	return &PaymentHandlerImpl{
		PaymentService: paymentService,
	}
}