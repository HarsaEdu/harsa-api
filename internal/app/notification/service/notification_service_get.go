package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (notificationService *NotificationServiceImpl) GetAll(offset, limit int, userID uint) ([]domain.Notification, *web.Pagination, error) {
	result, total, err := notificationService.NotificationRepository.GetAll(offset, limit, userID)

	if total == 0 {
		return nil, nil, fmt.Errorf("notification not found")
	}

	if err != nil {
		return nil, nil, fmt.Errorf("internal Server Error")
	}

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return result, pagination, nil
}

func (notificationService *NotificationServiceImpl) GetById(id int) (*domain.Notification, error) {
	result, err := notificationService.NotificationRepository.FindById(id)
	if result == nil {
		return nil, fmt.Errorf("notification not found")
	}

	if err != nil {
		return nil, fmt.Errorf("internal Server Error")
	}

	return result, nil
}
