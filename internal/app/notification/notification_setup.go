package notification

import (
	notificationHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/notification/handler"
	notificationRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/notification/repository"
	notificationRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/notification/routes"
	notificationServicePkg "github.com/HarsaEdu/harsa-api/internal/app/notification/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func NotificationSetup(db *gorm.DB, validate *validator.Validate) notificationRoutesPkg.NotificationRoutes {
	notificationRepository := notificationRepositoryPkg.NewNotificationRepository(db)
	notificationService := notificationServicePkg.NewNotificationService(notificationRepository, validate)
	notificationHandler := notificationHandlerPkg.NewNotificationHandler(notificationService)
	notificationRoute := notificationRoutesPkg.NewNotificationRoutes(notificationHandler)

	return notificationRoute
}
