package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(Notification *domain.Notification) error
	CreateMany(Notification []domain.Notification) error
	GetAll(offset, limit int, userID uint) ([]domain.Notification, int64, error)
	ReadNotification(id int) error
	ArsipNotification(arsip web.ArsipNotification) error
	Delete(id int) error
	FindById(id int) (*domain.Notification, error)
}

type NotificationRepositoryImpl struct {
	DB *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &NotificationRepositoryImpl{DB: db}
}
