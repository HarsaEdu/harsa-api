package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
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

	chargeRequest := conversion.CreatePaymentSubscriptionRequestToMidtransChargeRequest(existingSubsPlan, existingUser, request)

	chargeRequestResponse, err := paymentService.MidtransCoreApi.ChargeTransaction(chargeRequest)
	if err != nil {
		return nil, fmt.Errorf("error when create payment : %s", err.Error())
	}

	paymentHistory := conversion.ChargeResponseToPaymentHistoryDomain(chargeRequestResponse, existingUser, existingSubsPlan.ID)

	err = paymentService.PaymentRepository.CreatePaymentHistory(paymentHistory)
	if err != nil {
		return nil, fmt.Errorf("error when create payment history : %s", err.Error())
	}

	paymentHistoryResponse, err := paymentService.PaymentRepository.GetPaymentHistory(paymentHistory.ID)
	if err != nil {
		return nil, fmt.Errorf("error when get payment history : %s", err.Error())
	}

	response := conversion.PaymentHistoryDomainToPaymentHistoryResponse(paymentHistoryResponse)

	return response, nil
}