package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/midtrans/midtrans-go/coreapi"
)

func (paymentService *PaymentServiceImpl) CreatePayment(request *web.CreatePaymentRequest, userId uint) (*coreapi.ChargeResponse, error) {
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

	response, err := paymentService.MidtransCoreApi.CreateTransaction(existingSubsPlan, existingUser)
	if err != nil {
		return nil, fmt.Errorf("error when create payment : %s", err.Error())
	}

	return response, nil
}