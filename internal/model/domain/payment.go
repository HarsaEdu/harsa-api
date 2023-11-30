package domain

import (
	"time"

	"gorm.io/gorm"
)

type PaymentHistory struct {
	ID             string         `gorm:"type:varchar(255);primarykey" json:"id"`
	UserId         uint           `gorm:"type:int" json:"user_id"`
	ItemId         uint           `gorm:"type:int" json:"item_id"`
	Status         string         `gorm:"type:varchar(255)" json:"status"`
	Method         string         `gorm:"type:varchar(255)" json:"method"`
	GrossAmount    string         `gorm:"type:string" json:"gross_amount"`
	BankName       string         `gorm:"type:varchar(255)" json:"bank_name"`
	VaNumber       string         `gorm:"type:varchar(255)" json:"va_number"`
	SettlementTime time.Time      `gorm:"type:date" json:"settlement_time"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"delete_at"`
	ExpiredAt      time.Time      `gorm:"type:date" json:"expired_at"`
	User           User           `gorm:"foreignKey:UserId;references:ID" json:"user"`
	Item           SubsPlan       `gorm:"foreignKey:ItemId;references:ID" json:"item"`
}

type PaymentTransactionStatus struct {
	OrderID           string    `json:"order_id"`
	TransactionStatus string    `json:"transaction_status"`
	FraudStatus       string    `json:"fraud_status"`
	SettlementTime    time.Time `json:"settlement_time"`
}