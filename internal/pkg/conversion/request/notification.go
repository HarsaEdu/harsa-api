package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func NotificationCreateRequestToNotificationDomain(userID uint, title, message string) *domain.Notification {
	return &domain.Notification{
		UserID:  userID,
		Title:   title,
		Content: message,
	}
}

func NotificationMultiCastRequestToNotificationDomain(request web.NotificationMultiCast, userIDs []uint) []domain.Notification {
	notifications := []domain.Notification{}

	for _, userID := range userIDs {
		notifications = append(notifications, domain.Notification{
			UserID: userID,
			Title: request.Title,
			Content: request.Message,
		})
	}

	return notifications
}