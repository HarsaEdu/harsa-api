package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/payment/handler"
	"github.com/labstack/echo/v4"
)

type PaymentRoutes interface {
	PaymentWeb(apiGroup *echo.Group) *echo.Group
	PaymentMobile(apiGroup *echo.Group) *echo.Group
	PaymentSubscriptionsMobile(apiGroup *echo.Group) *echo.Group
}

type PaymentRoutesImpl struct {
	PaymentHandler handler.PaymentHandler
}

func NewPaymentRoutes(PaymentHandler handler.PaymentHandler) PaymentRoutes {
	return &PaymentRoutesImpl{
		PaymentHandler: PaymentHandler,
	}
}