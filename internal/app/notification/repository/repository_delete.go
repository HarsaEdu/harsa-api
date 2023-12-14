package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (notificationRepository *NotificationRepositoryImpl) Delete(id int) error {

	result := notificationRepository.DB.Where("id = ?", id).Delete(&domain.Notification{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
