package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (paymentRepository *PaymentRepositoryImpl) GetPaymentHistoryById(paymentHistoryId string) (*domain.PaymentHistory, error) {
	paymentHistory := domain.PaymentHistory{}
	result := paymentRepository.DB.Preload("User.UserProfile").Preload("Item").Where("id = ?", paymentHistoryId).First(&paymentHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return &paymentHistory, nil
}

func (paymentRepository *PaymentRepositoryImpl) GetAllPaymentHistory() ([]domain.PaymentHistory, error) {
	paymentHistory := []domain.PaymentHistory{}
	result := paymentRepository.DB.Preload("User.UserProfile").Preload("Item").Find(&paymentHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return paymentHistory, nil
}

func (paymentRepository *PaymentRepositoryImpl) GetAllPaymentHistoryByUserId(userId uint) ([]domain.PaymentHistory, error) {
	paymentHistory := []domain.PaymentHistory{}
	result := paymentRepository.DB.Preload("User.UserProfile").Preload("Item").Where("user_id = ?", userId).Find(&paymentHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return paymentHistory, nil
}

func (paymentRepository *PaymentRepositoryImpl) GetPaymentHistoryByUserIdAndPaymentId(userId uint, paymentId string) (*domain.PaymentHistory, error) {
	paymentHistory := domain.PaymentHistory{}
	result := paymentRepository.DB.Preload("User.UserProfile").Preload("Item").Where("user_id = ? AND id = ?", userId, paymentId).First(&paymentHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return &paymentHistory, nil
}