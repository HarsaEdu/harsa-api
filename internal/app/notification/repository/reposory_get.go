package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (notificationRepository *NotificationRepositoryImpl) GetAll(offset, limit int, userID uint) ([]domain.Notification, int64, error) {
	if offset < 0 || limit < 0 {
		return nil, 0, nil
	}

	notification := []domain.Notification{}
	var total int64

	query := notificationRepository.DB.Model(&notification).Where("user_id = ?", userID)

	query.Find(&notification).Count(&total)

	query = query.Limit(limit).Offset(offset)

	result := query.Find(&notification)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 {
		return nil, 0, nil
	}

	if offset >= int(total) {
		return nil, 0, nil
	}

	return notification, total, nil
}
