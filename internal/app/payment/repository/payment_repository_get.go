package repository

import (
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (paymentRepository *PaymentRepositoryImpl) GetPaymentHistoryById(paymentHistoryId string) (*domain.PaymentHistory, error) {
	paymentHistory := domain.PaymentHistory{}
	result := paymentRepository.DB.Preload("User.UserProfile").Preload("Item").Where("id = ?", paymentHistoryId).Order("transaction_time DESC").First(&paymentHistory)
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
    query.Joins("JOIN user_profiles ON user_profiles.id = payment_histories.user_id").
        Where("CONCAT(user_profiles.first_name, ' ', user_profiles.last_name) LIKE ?", "%"+search+"%")
	}

	query.Order("transaction_time DESC").Find(&paymentHistory).Count(&count)

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
		query.Joins("JOIN subs_plans ON subs_plans.id = payment_histories.item_id").Where("subs_plans.title LIKE ?", "%"+search+"%")
	}

	query.Order("transaction_time DESC").Find(&paymentHistory).Count(&count)
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
	result := paymentRepository.DB.Preload("User.UserProfile").Preload("Item").Where("user_id = ? AND id = ?", userId, paymentId).Order("transaction_time DESC").First(&paymentHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return &paymentHistory, nil
}

func (paymentRepository *PaymentRepositoryImpl) GetLastYearPaymentHistory(now time.Time) ([]domain.PaymentHistory, error) {
	paymentHistory := []domain.PaymentHistory{}
	result := paymentRepository.DB.Where("status = ?", "success").Where("settlement_time > ?", now.AddDate(-1, 0, 0)).Order("settlement_time DESC").Find(&paymentHistory)
	if result.Error != nil {
		return nil, result.Error
	}

	return paymentHistory, nil
}