package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreatePaymentHistory(paymentHistory *domain.PaymentHistory) error
	GetPaymentHistoryById(paymentHistoryId string) (*domain.PaymentHistory, error)
	GetAllPaymentHistory(offset, limit int, search string, status string) ([]domain.PaymentHistory, int64, error)
	GetAllPaymentHistoryByUserId(userId uint, offset, limit int, search string, status string) ([]domain.PaymentHistory, int64, error)
	GetPaymentHistoryByUserIdAndPaymentId(userId uint, paymentId string) (*domain.PaymentHistory, error)
	UpdateStatusPaymentHistory(status string, transactionResult *domain.PaymentTransactionStatus) error
}

type PaymentRepositoryImpl struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{
		DB: db,
	}
}