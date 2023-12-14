package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (notificationRepository *NotificationRepositoryImpl) ReadNotification(id int) error {
	result := notificationRepository.DB.Model(domain.Notification{}).Where("id=?", id).Update("is_read", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (notificationRepository *NotificationRepositoryImpl) ArsipNotification(arsip web.ArsipNotification) error {
	result := notificationRepository.DB.Model(domain.Notification{}).Updates(arsip)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
