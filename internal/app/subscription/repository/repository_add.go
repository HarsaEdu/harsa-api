package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (subscriptionRepository *SubscriptionRepositoryImpl) AddSubscription(userID uint, days uint) error {
	subscription := &domain.Subscription{}

	result := subscriptionRepository.DB.Create(&subscription)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
