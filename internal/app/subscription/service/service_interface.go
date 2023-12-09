package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/subscription/repository"
	"github.com/labstack/echo/v4"
)

type SubscriptionService interface {
	SubscriptionAdd(user_id uint, days uint) error
	IsSubscription(ctx echo.Context, user_id uint) (bool, error)
	IsSubscriptionWeb(createdAt int64, user_id uint) (bool, error)
}

type SubscriptionServiceImpl struct {
	SubscriptionRepository repository.SubscriptionRepository
}

func NewSubscriptionService(SubscriptionRepository repository.SubscriptionRepository) SubscriptionService {
	return &SubscriptionServiceImpl{
		SubscriptionRepository: SubscriptionRepository,
	}
}
