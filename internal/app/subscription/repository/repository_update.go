package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (subscriptionRepository *SubscriptionRepositoryImpl) UpdateSubscription(subscription *domain.Subscription) error {
	result := subscriptionRepository.DB.Save(subscription)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
