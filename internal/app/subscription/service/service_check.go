package service

import (
	"time"

	"github.com/labstack/echo/v4"
)

func (subscriptionService *SubscriptionServiceImpl) IsSubscription(ctx echo.Context, user_id uint) (bool, error) {
	userCreateDate := ctx.Get("create_user").(int64)
	now := time.Now().Unix()
	if userCreateDate > (now - (7 * 24 * 60 * 60)) {
		return true, nil
	}
	// check subscription
	subscription, err := subscriptionService.SubscriptionRepository.FindSubscription(user_id)

	if err != nil {
		return false, err
	}

	if subscription.ID == 0 {
		return false, nil
	}

	endDate := subscription.EndDate.Unix()

	if endDate < now {
		return false, nil
	}

	return true, nil
}


func (subscriptionService *SubscriptionServiceImpl) IsSubscriptionWeb(createdAt int64, user_id uint) (bool, error) {
	
	now := time.Now().Unix()
	if createdAt > (now - (7 * 24 * 60 * 60)) {
		return true, nil
	}
	// check subscription
	subscription, err := subscriptionService.SubscriptionRepository.FindSubscription(user_id)

	if err != nil {
		return false, err
	}

	if subscription.ID == 0 {
		return false, nil
	}

	endDate := subscription.EndDate.Unix()

	if endDate < now {
		return false, nil
	}

	return true, nil
}