package service

import (
	"fmt"
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
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


func (paymentService *PaymentServiceImpl) GetAllPaymentHistory(offset, limit int, search string, status string) ([]web.GetPaymentResponse, *web.Pagination, error) {
	paymentHistory, total, err := paymentService.PaymentRepository.GetAllPaymentHistory(offset, limit, search, status)
	if err != nil {
		return nil, nil, fmt.Errorf("error when get payment history : %s", err.Error())
	}

	if paymentHistory == nil {
		return nil, nil, fmt.Errorf("payment history not found")
	}

	response := conversionResponse.PaymentHistoryDomainToPaymentHistoryResponses(paymentHistory)

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return response, pagination, nil
}

func (paymentService *PaymentServiceImpl) GetAllPaymentHistoryByUserId(userId uint, offset, limit int, search string, status string) ([]web.GetPaymentResponse, *web.Pagination, error) {
	paymentHistory, total, err := paymentService.PaymentRepository.GetAllPaymentHistoryByUserId(userId, offset, limit, search, status)
	if err != nil {
		return nil, nil, fmt.Errorf("error when get payment history by user id : %s", err.Error())
	}

	if paymentHistory == nil {
		return nil, nil, fmt.Errorf("payment history not found")
	}

	response := conversionResponse.PaymentHistoryDomainToPaymentHistoryResponses(paymentHistory)
	
	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return response, pagination, nil
}

func (paymentService *PaymentServiceImpl) GetPaymentHistoryByUserIdAndPaymentId(userId uint, paymentId string) (*web.GetPaymentResponse, error) {
	paymentHistory, err := paymentService.PaymentRepository.GetPaymentHistoryByUserIdAndPaymentId(userId, paymentId)
	if err != nil {
		return nil, fmt.Errorf("error when get payment history by user id and payment id : %s", err.Error())
	}

	if paymentHistory == nil {
		return nil, fmt.Errorf("payment history not found")
	}

	response := conversionResponse.PaymentHistoryDomainToPaymentHistoryResponse(paymentHistory)

	return response, nil

}

func (paymentService *PaymentServiceImpl) GetLastYearPaymentHistory() (*web.PaymentLastYearHistoryResponse, error) {
	paymentHistories, err := paymentService.PaymentRepository.GetLastYearPaymentHistory(time.Now())
	if paymentHistories == nil {
		return nil, fmt.Errorf("error when get payment history : payment history not found")
	}

	if err != nil {
		return nil, fmt.Errorf("error when get payment history : %s", err.Error())
	}

	getLastYearData := conversion.PaymentLastYearHistoryDomainToPaymentLastYearHistoryResponse(paymentHistories)

	return getLastYearData, nil
}