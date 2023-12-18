package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (notificationRoutes *NotificationRoutesImpl) NotificationWeb(apiGroup *echo.Group) {
	notificationGroup := apiGroup.Group("/notification", middleware.AllUserMiddleare)

	notificationGroup.GET("", notificationRoutes.NotificationHandler.GetAll)
	notificationGroup.GET("/:id", notificationRoutes.NotificationHandler.GetById)
	notificationGroup.DELETE("/:id", notificationRoutes.NotificationHandler.Delete)
	notificationGroup.PATCH("/read/:id", notificationRoutes.NotificationHandler.ReadNotification)
	notificationGroup.PATCH("/arsip/:id", notificationRoutes.NotificationHandler.ArsipNotification)

}

func (notificationRoutes *NotificationRoutesImpl) NotificationMobile(apiGroup *echo.Group) {
	notificationGroup := apiGroup.Group("/notification", middleware.AllUserMiddleare)

	notificationGroup.GET("", notificationRoutes.NotificationHandler.GetAll)
	notificationGroup.GET("/:id", notificationRoutes.NotificationHandler.GetById)
	notificationGroup.DELETE("/:id", notificationRoutes.NotificationHandler.Delete)
	notificationGroup.PATCH("/read/:id", notificationRoutes.NotificationHandler.ReadNotification)
	notificationGroup.PATCH("/arsip/:id", notificationRoutes.NotificationHandler.ArsipNotification)

}
