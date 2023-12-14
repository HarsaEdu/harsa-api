package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (notificationRepository *NotificationRepositoryImpl) FindById(id int) (*domain.Notification, error) {
	notification := domain.Notification{}
	result := notificationRepository.DB.First(&notification, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &notification, nil
}
