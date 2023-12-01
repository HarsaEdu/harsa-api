package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (subscriptionRepository *SubscriptionRepositoryImpl) AddSubscription(subscription *domain.Subscription) error {
	result := subscriptionRepository.DB.Create(&subscription)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
