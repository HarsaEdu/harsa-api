package domain

import "time"

type PaymentHistory struct {
	ID              string    `gorm:"type:varchar(255);primarykey" json:"id"`
	UserId          uint      `gorm:"type:int" json:"user_id"`
	ItemId          uint      `gorm:"type:int" json:"item_id"`
	Status          string    `gorm:"type:varchar(255)" json:"status"`
	Method          string    `gorm:"type:varchar(255)" json:"method"`
	GrossAmount     string    `gorm:"type:string" json:"gross_amount"`
	BankName        string    `gorm:"type:varchar(255)" json:"bank_name"`
	VaNumber        string    `gorm:"type:varchar(255)" json:"va_number"`
	SettlementTime  time.Time `gorm:"type:datetime;null;default:null" json:"settlement_time"`
	TransactionTime time.Time `gorm:"type:datetime" json:"transaction_time"`
	ExpiryTime      time.Time `gorm:"type:datetime" json:"expiry_time"`
	User            User      `gorm:"foreignKey:UserId;references:ID" json:"user"`
	Item            SubsPlan  `gorm:"foreignKey:ItemId;references:ID" json:"item"`
}

type PaymentTransactionStatus struct {
	OrderID           string    `json:"order_id"`
	TransactionStatus string    `json:"transaction_status"`
	FraudStatus       string    `json:"fraud_status"`
	SettlementTime    time.Time `json:"settlement_time"`
}