package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/midtrans/midtrans-go/coreapi"
)

func (subscriptionService *SubscriptionServiceImpl) Subscribe(request *web.CreateSubscribeRequest, userId uint) (*coreapi.ChargeResponse, error) {
	err := subscriptionService.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	existingUser, err := subscriptionService.UserRepository.GetUserByID(userId)
	if err != nil {
		return nil, fmt.Errorf("error when get user : %s", err.Error())
	}

	existingSubsPlan, err := subscriptionService.SubsPlanRepository.FindById(request.PlanId)
	if err != nil {
		return nil, fmt.Errorf("error when get subscription plan : %s", err.Error())
	}

	response, err := subscriptionService.MidtransCoreApi.CreateTransaction(existingSubsPlan, existingUser)
	if err != nil {
		return nil, fmt.Errorf("error when subscribe : %s", err.Error())
	}

	return response, nil
}