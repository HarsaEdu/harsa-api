package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

// func (subscriptionRoutes *SubscriptionRoutesImpl) SubscriptionWeb(apiGroup *echo.Group) *echo.Group {
// 	subscriptionsGroup := apiGroup.Group("/subscription")

// 	subscriptionsGroup.POST("", subscriptionRoutes.SubscriptionHandler.Create, middleware.InstructorMiddleware)

// 	return subscriptionsGroup
// }

func (subscriptionRoutes *SubscriptionRoutesImpl) SubscriptionMobile(apiGroup *echo.Group) *echo.Group {
	subscriptionsGroup := apiGroup.Group("/subscriptions")

	subscriptionsGroup.POST("", subscriptionRoutes.SubscriptionHandler.Subscribe, middleware.StudentMiddleare)

	return subscriptionsGroup
}