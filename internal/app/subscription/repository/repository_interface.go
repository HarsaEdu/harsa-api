package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	AddSubscription(*domain.Subscription) error
	UpdateSubscription(*domain.Subscription) error
	FindSubscription(userID uint) (*domain.Subscription, error)
}

type SubscriptionRepositoryImpl struct {
	DB *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &SubscriptionRepositoryImpl{
		DB: db,
	}
}
