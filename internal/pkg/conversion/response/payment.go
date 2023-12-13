package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func PaymentHistoryDomainToPaymentHistoryResponse(paymentHistory *domain.PaymentHistory) *web.GetPaymentResponse {
	return &web.GetPaymentResponse{
		ID:              paymentHistory.ID,
		UserId:          paymentHistory.UserId,
		ItemId:          paymentHistory.ItemId,
		Status:          paymentHistory.Status,
		Method:          paymentHistory.Method,
		GrossAmount:     paymentHistory.GrossAmount,
		BankName:        paymentHistory.BankName,
		VaNumber:        paymentHistory.VaNumber,
		TransactionTime: paymentHistory.TransactionTime,
		ExpiryTime:      paymentHistory.ExpiryTime,
		SettlementTime:  paymentHistory.SettlementTime,
		Customer: web.PaymentCustomerResponse{
			ID:       paymentHistory.User.ID,
			Name:     paymentHistory.User.UserProfile.FirstName + " " + paymentHistory.User.UserProfile.LastName,
			Email:    paymentHistory.User.Email,
			Username: paymentHistory.User.Username,
		},
		Item: web.PaymentItemResponse{
			ID:   paymentHistory.Item.ID,
			Name: paymentHistory.Item.Title,
		},
	}
}

func PaymentHistoryDomainToPaymentHistoryResponses(paymentHistories []domain.PaymentHistory) []web.GetPaymentResponse {
	var paymentResponses []web.GetPaymentResponse
	for _, paymentHistory := range paymentHistories {
		paymentResponses = append(paymentResponses, *PaymentHistoryDomainToPaymentHistoryResponse(&paymentHistory))
	}
	return paymentResponses
}