package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/subscription/repository"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type SubscriptionService interface {
	SubscriptionAdd(user_id uint, days uint) error
	IsSubscription(ctx echo.Context, user_id uint) (bool, error)
}

type SubscriptionServiceImpl struct {
	SubscriptionRepository repository.SubscriptionRepository
	Validate               *validator.Validate
}

func NewSubscriptionService(SubscriptionRepository repository.SubscriptionRepository, validate *validator.Validate) SubscriptionService {
	return &SubscriptionServiceImpl{
		SubscriptionRepository: SubscriptionRepository,
		Validate:               validate,
	}
}
