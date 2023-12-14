package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/notification/handler"
	"github.com/labstack/echo/v4"
)

type NotificationRoutes interface {
	NotificationWeb(apiGroup *echo.Group)
	NotificationMobile(apiGroup *echo.Group)
}

type NotificationRoutesImpl struct {
	NotificationHandler handler.NotificationHandler
}

func NewNotificationRoutes(NotificationHandler handler.NotificationHandler) NotificationRoutes {
	return &NotificationRoutesImpl{
		NotificationHandler: NotificationHandler,
	}
}
