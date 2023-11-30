package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (paymentRoutes *PaymentRoutesImpl) PaymentWeb(apiGroup *echo.Group) *echo.Group {
	paymentsGroup := apiGroup.Group("/payments")

	paymentsGroup.POST("/notifications", paymentRoutes.PaymentHandler.NotificationPayment)

	return paymentsGroup
}

func (paymentRoutes *PaymentRoutesImpl) PaymentSubscriptionsMobile(apiGroup *echo.Group) *echo.Group {
	paymentsGroup := apiGroup.Group("/payments")
	subscriptionsGroup := paymentsGroup.Group("/subscriptions")

	subscriptionsGroup.POST("", paymentRoutes.PaymentHandler.CreatePaymentSubscription, middleware.StudentMiddleare)

	return subscriptionsGroup
}