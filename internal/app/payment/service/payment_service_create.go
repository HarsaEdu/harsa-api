package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	conversionResponse "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (paymentService *PaymentServiceImpl) CreatePaymentSubscription(request *web.CreatePaymentSubscriptionRequest, userId uint) (*web.GetPaymentResponse, error) {
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

	chargeRequest := conversionRequest.CreatePaymentSubscriptionRequestToMidtransChargeRequest(existingSubsPlan, existingUser, request)

	chargeRequestResponse, err := paymentService.MidtransCoreApi.ChargeTransaction(chargeRequest)
	if err != nil {
		return nil, fmt.Errorf("error when create payment : %s", err.Error())
	}

	paymentHistory := conversionRequest.ChargeResponseToPaymentHistoryDomain(chargeRequestResponse, existingUser, existingSubsPlan.ID)

	err = paymentService.PaymentRepository.CreatePaymentHistory(paymentHistory)
	if err != nil {
		return nil, fmt.Errorf("error when create payment history : %s", err.Error())
	}

	paymentHistoryResponse, err := paymentService.PaymentRepository.GetPaymentHistoryById(paymentHistory.ID)
	if err != nil {
		return nil, fmt.Errorf("error when get payment history : %s", err.Error())
	}

	response := conversionResponse.PaymentHistoryDomainToPaymentHistoryResponse(paymentHistoryResponse)

	return response, nil
}