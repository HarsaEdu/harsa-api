package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (subscriptionRepository *SubscriptionRepositoryImpl) FindSubscription(userID uint) (*domain.Subscription, error) {
	subscription := &domain.Subscription{}
	result := subscriptionRepository.DB.Where("user_id = ?", userID).First(&subscription)

	if result.Error != nil {
		return nil, result.Error
	}

	return subscription, nil
}
