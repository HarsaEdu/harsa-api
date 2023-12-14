package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/notification/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type NotificationService interface {
	GetAll(offset, limit int, userID uint) ([]domain.Notification, *web.Pagination, error)
	GetById(id int) (*domain.Notification, error)
	Delete(id int) error
	ReadNotification(id int) error
	ArsipNotification(arsip web.ArsipNotification) error
}

type NotificationServiceImpl struct {
	NotificationRepository repository.NotificationRepository
	Validator              *validator.Validate
}

func NewNotificationService(notificationRepository repository.NotificationRepository, validator *validator.Validate) NotificationService {
	return &NotificationServiceImpl{NotificationRepository: notificationRepository, Validator: validator}

}
