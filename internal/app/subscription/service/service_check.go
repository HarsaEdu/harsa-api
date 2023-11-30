package service

import (
	"time"

	"github.com/labstack/echo/v4"
)

func (subscriptionService *SubscriptionServiceImpl) IsSubscription(ctx echo.Context, user_id uint) (bool, error) {
	userCreateDate := ctx.Get("user_create").(int64)
	now := time.Now().Unix()
	if userCreateDate > (now - (7 * 60 * 60)) {
		return true, nil
	}
	// check subscription
	subscription, err := subscriptionService.SubscriptionRepository.FindSubscription(user_id)

	if err != nil {
		return false, err
	}

	if subscription == nil {
		return false, nil
	}

	endDate := subscription.EndDate.Unix()

	if endDate < now {
		return false, nil
	}

	return true, nil
}
