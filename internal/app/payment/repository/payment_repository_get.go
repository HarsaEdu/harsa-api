package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (paymentRepository *PaymentRepositoryImpl) GetPaymentHistoryById(paymentHistoryId string) (*domain.PaymentHistory, error) {
	paymentHistory := domain.PaymentHistory{}
	result := paymentRepository.DB.Preload("User.UserProfile").Preload("Item").Where("id = ?", paymentHistoryId).Order("created_at DESC").First(&paymentHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return &paymentHistory, nil
}

func (paymentRepository *PaymentRepositoryImpl) GetAllPaymentHistory() ([]domain.PaymentHistory, error) {
	paymentHistory := []domain.PaymentHistory{}
	result := paymentRepository.DB.Preload("User.UserProfile").Preload("Item").Order("created_at DESC").Find(&paymentHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return paymentHistory, nil
}

func (paymentRepository *PaymentRepositoryImpl) GetAllPaymentHistoryByUserId(userId uint) ([]domain.PaymentHistory, error) {
	paymentHistory := []domain.PaymentHistory{}
	result := paymentRepository.DB.Preload("User.UserProfile").Preload("Item").Where("user_id = ?", userId).Order("created_at DESC").Find(&paymentHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return paymentHistory, nil
}

func (paymentRepository *PaymentRepositoryImpl) GetPaymentHistoryByUserIdAndPaymentId(userId uint, paymentId string) (*domain.PaymentHistory, error) {
	paymentHistory := domain.PaymentHistory{}
	result := paymentRepository.DB.Preload("User.UserProfile").Preload("Item").Where("user_id = ? AND id = ?", userId, paymentId).Order("created_at DESC").First(&paymentHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return &paymentHistory, nil
}