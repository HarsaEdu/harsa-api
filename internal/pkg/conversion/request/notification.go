package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func NotificationCreateRequestToNotificationDomain(userID uint, title, message string) *domain.Notification {
	return &domain.Notification{
		UserID:  userID,
		Title:   title,
		Content: message,
	}
}
