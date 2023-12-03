package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (paymentRepository *PaymentRepositoryImpl) CreatePaymentHistory(paymentHistory *domain.PaymentHistory) error {
	result := paymentRepository.DB.Create(paymentHistory)
	if result.Error != nil {
		return result.Error
	}

	return nil
}