package repository

func (paymentRepository *PaymentRepositoryImpl) UpdateStatusPaymentHistory(status string, orderId string) error {
	result := paymentRepository.DB.Table("payment_histories").Where("id = ?", orderId).Update("status", status)
	if result.Error != nil {
		return result.Error
	}

	return nil
}