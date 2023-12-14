package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/notification/service"
	"github.com/labstack/echo/v4"
)

type NotificationHandler interface {
	GetAll(ctx echo.Context) error
	GetById(ctx echo.Context) error
	Delete(ctx echo.Context) error
	ReadNotification(ctx echo.Context) error
	ArsipNotification(ctx echo.Context) error
}

type NotificationHandlerImpl struct {
	NotificationService service.NotificationService
}

func NewNotificationHandler(service service.NotificationService) NotificationHandler {
	return &NotificationHandlerImpl{NotificationService: service}
}
