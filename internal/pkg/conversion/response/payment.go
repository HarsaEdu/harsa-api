package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func PaymentHistoryDomainToPaymentHistoryResponse(paymentHistory *domain.PaymentHistory) *web.GetPaymentResponse {
	return &web.GetPaymentResponse{
		ID:          paymentHistory.ID,
		UserId:      paymentHistory.UserId,
		ItemId:      paymentHistory.ItemId,
		Status:      paymentHistory.Status,
		Method:      paymentHistory.Method,
		GrossAmount: paymentHistory.GrossAmount,
		BankName:    paymentHistory.BankName,
		VaNumber:    paymentHistory.VaNumber,
		CreatedAt:   paymentHistory.CreatedAt,
		UpdatedAt:   paymentHistory.UpdatedAt,
		ExpiredAt:   paymentHistory.ExpiredAt,
		Customer: web.PaymentCustomerResponse{
			ID:   paymentHistory.User.ID,
			Name: paymentHistory.User.UserProfile.FirstName + " " + paymentHistory.User.UserProfile.LastName,
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