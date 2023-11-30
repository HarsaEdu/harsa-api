package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionResponse "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (paymentService *PaymentServiceImpl) GetPaymentHistoryById(orderId string) (*web.GetPaymentResponse, error) {
	paymentHistory, err := paymentService.PaymentRepository.GetPaymentHistoryById(orderId)
	if err != nil {
		return nil, fmt.Errorf("error when get payment history by id : %s", err.Error())
	}

	if paymentHistory == nil {
		return nil, fmt.Errorf("payment history not found")
	}

	response := conversionResponse.PaymentHistoryDomainToPaymentHistoryResponse(paymentHistory)

	return response, nil
}


func (paymentService *PaymentServiceImpl) GetAllPaymentHistory() ([]web.GetPaymentResponse, error) {
	paymentHistory, err := paymentService.PaymentRepository.GetAllPaymentHistory()
	if err != nil {
		return nil, fmt.Errorf("error when get payment history : %s", err.Error())
	}

	if paymentHistory == nil {
		return nil, fmt.Errorf("payment history not found")
	}

	response := conversionResponse.PaymentHistoryDomainToPaymentHistoryResponses(paymentHistory)

	return response, nil
}

func (paymentService *PaymentServiceImpl) GetAllPaymentHistoryByUserId(userId uint) ([]web.GetPaymentResponse, error) {
	paymentHistory, err := paymentService.PaymentRepository.GetAllPaymentHistoryByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("error when get payment history by user id : %s", err.Error())
	}

	if paymentHistory == nil {
		return nil, fmt.Errorf("payment history not found")
	}

	response := conversionResponse.PaymentHistoryDomainToPaymentHistoryResponses(paymentHistory)

	return response, nil
}