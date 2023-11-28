package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

// func (paymentRoutes *PaymentRoutesImpl) PaymentWeb(apiGroup *echo.Group) *echo.Group {
// 	paymentsGroup := apiGroup.Group("/payment")

// 	paymentsGroup.POST("", paymentRoutes.PaymentHandler.Create, middleware.InstructorMiddleware)

// 	return paymentsGroup
// }

func (paymentRoutes *PaymentRoutesImpl) PaymentMobile(apiGroup *echo.Group) *echo.Group {
	paymentsGroup := apiGroup.Group("/payments")

	paymentsGroup.POST("", paymentRoutes.PaymentHandler.CreatePayment, middleware.StudentMiddleare)

	return paymentsGroup
}