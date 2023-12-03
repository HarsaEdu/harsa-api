package service

import (
	"fmt"
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (subscriptionService *SubscriptionServiceImpl) SubscriptionAdd(user_id uint, days uint) error {
	subscription := domain.Subscription{}

	// check subscription
	subscriptionResponse, _ := subscriptionService.SubscriptionRepository.FindSubscription(user_id)

	if subscriptionResponse != nil {
		subscription = *subscriptionResponse
	}

	subscription.UserID = user_id

	endDate := subscription.EndDate.Unix()
	now := time.Now().Unix()

	if endDate >= now {
		endDate = endDate + (24 * 60 * 60 * int64(days))
		subscription.EndDate = time.Unix(endDate, 0).UTC()
	} else {
		endDate = now + (24 * 60 * 60 * int64(days))
		subscription.EndDate = time.Unix(endDate, 0).UTC()
	}

	if &subscriptionResponse == nil {
		err := subscriptionService.SubscriptionRepository.AddSubscription(&subscription)
		if err != nil {
			return fmt.Errorf("error when add subscription %s:", err.Error())
		}
	} else {
		err := subscriptionService.SubscriptionRepository.UpdateSubscription(&subscription)
		if err != nil {
			return fmt.Errorf("error when update subscription %s:", err.Error())
		}
	}

	return nil
}
