package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (paymentRepository *PaymentRepositoryImpl) GetPaymentHistoryById(paymentHistoryId string) (*domain.PaymentHistory, error) {
	paymentHistory := domain.PaymentHistory{}
	result := paymentRepository.DB.Preload("User.UserProfile").Preload("Item").Where("id = ?", paymentHistoryId).Order("created_at DESC").First(&paymentHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return &paymentHistory, nil
}

func (paymentRepository *PaymentRepositoryImpl) GetAllPaymentHistory(offset, limit int, search string, status string) ([]domain.PaymentHistory, int64, error) {
	var count int64
	paymentHistory := []domain.PaymentHistory{}
	query := paymentRepository.DB.Preload("User.UserProfile").Preload("Item")
	
	if status != "" {
		query.Where("status = ?", status)
	}

	if search != "" {
		query.Where("item.name LIKE ?", "%"+search+"%")
	}

	query.Order("created_at DESC").Find(&paymentHistory).Count(&count)

	query = query.Offset(offset).Limit(limit)

	result := query.Find(&paymentHistory)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if offset >= int(count) {
		return paymentHistory, 0, nil
	}

	return paymentHistory, count, nil
}

func (paymentRepository *PaymentRepositoryImpl) GetAllPaymentHistoryByUserId(userId uint, offset, limit int, search string, status string) ([]domain.PaymentHistory, int64, error) {
	var count int64
	paymentHistory := []domain.PaymentHistory{}
	query := paymentRepository.DB.Preload("User.UserProfile").Preload("Item")
	query.Where("user_id = ?", userId)

	if status != "" {
		query.Where("status = ?", status)
	}

	if search != "" {
		query.Where("Item.name LIKE ?", "%"+search+"%")
	}

	query.Order("created_at DESC").Find(&paymentHistory).Count(&count)
	query = query.Offset(offset).Limit(limit)

	result := query.Find(&paymentHistory)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if offset >= int(count) {
		return paymentHistory, 0, nil
	}

	return paymentHistory, count, nil
}

func (paymentRepository *PaymentRepositoryImpl) GetPaymentHistoryByUserIdAndPaymentId(userId uint, paymentId string) (*domain.PaymentHistory, error) {
	paymentHistory := domain.PaymentHistory{}
	result := paymentRepository.DB.Preload("User.UserProfile").Preload("Item").Where("user_id = ? AND id = ?", userId, paymentId).Order("created_at DESC").First(&paymentHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return &paymentHistory, nil
}