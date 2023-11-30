package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (paymentService *PaymentServiceImpl) GetPaymentHistoryById(orderId string) (*web.GetPaymentResponse, error) {
	paymentHistory, err := paymentService.PaymentRepository.GetPaymentHistoryById(orderId)
	if err != nil {
		return nil, fmt.Errorf("error when get payment history by id : %s", err.Error())
	}

	if paymentHistory == nil {
		return nil, fmt.Errorf("payment history not found")
	}

	response := conversion.PaymentHistoryDomainToPaymentHistoryResponse(paymentHistory)

	return response, nil
}