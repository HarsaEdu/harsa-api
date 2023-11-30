package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/midtrans/midtrans-go/coreapi"
)

func (paymentService *PaymentServiceImpl) CreatePaymentSubscription(request *web.CreatePaymentSubscriptionRequest, userId uint) (*coreapi.ChargeResponse, error) {
	err := paymentService.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	existingUser, err := paymentService.UserRepository.GetUserByID(userId)
	if err != nil {
		return nil, fmt.Errorf("error when get user : %s", err.Error())
	}

	existingSubsPlan, err := paymentService.SubsPlanRepository.FindById(request.PlanId)
	if err != nil {
		return nil, fmt.Errorf("error when get subscription plan : %s", err.Error())
	}

	chargeRequest := conversion.CreatePaymentSubscriptionRequestToMidtransChargeRequest(existingSubsPlan, existingUser, request)

	response, err := paymentService.MidtransCoreApi.ChargeTransaction(chargeRequest)
	if err != nil {
		return nil, fmt.Errorf("error when create payment : %s", err.Error())
	}

	return response, nil
}