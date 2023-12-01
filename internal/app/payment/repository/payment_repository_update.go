package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (paymentRepository *PaymentRepositoryImpl) UpdateStatusPaymentHistory(status string, transactionResult *domain.PaymentTransactionStatus) error {
	result := paymentRepository.DB.Table("payment_histories").Where("id = ?", transactionResult.OrderID).Updates(domain.PaymentHistory{
		Status: status,
		SettlementTime: transactionResult.SettlementTime,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}