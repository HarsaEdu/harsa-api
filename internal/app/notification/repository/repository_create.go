package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (NotificationRepository *NotificationRepositoryImpl) Create(notification *domain.Notification) error {
	result := NotificationRepository.DB.Create(&notification)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (NotificationRepository *NotificationRepositoryImpl) CreateMany(Notification []domain.Notification) error {
	result := NotificationRepository.DB.Create(&Notification)
	if result.Error != nil {
		return result.Error
	}

	return nil
}