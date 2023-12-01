package service

import (
	"fmt"
	"time"
)

func (subscriptionService *SubscriptionServiceImpl) SubscriptionAdd(user_id uint, days uint) error {

	// check subscription
	subscription, _ := subscriptionService.SubscriptionRepository.FindSubscription(user_id)

	endDate := subscription.EndDate.Unix()
	now := time.Now().Unix()

	if endDate >= now {
		endDate = endDate + (24 * 60 * 60)
		subscription.EndDate = time.Unix(endDate, 0).UTC()
	} else {
		endDate = now + (24 * 60 * 60)
		subscription.EndDate = time.Unix(endDate, 0).UTC()
	}

	if subscription == nil {
		return fmt.Errorf("subscription not found")
	}

	if subscription == nil {
		err := subscriptionService.SubscriptionRepository.AddSubscription(subscription)
		if err != nil {
			return fmt.Errorf("error when add subscription %s:", err.Error())
		}
	} else {
		err := subscriptionService.SubscriptionRepository.UpdateSubscription(subscription)
		if err != nil {
			return fmt.Errorf("error when update subscription %s:", err.Error())
		}
	}

	return nil
}
