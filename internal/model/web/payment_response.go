package web

import "time"

type PaymentCustomerResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type PaymentItemResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetPaymentResponse struct {
	ID              string                  `json:"id"`
	UserId          uint                    `json:"user_id"`
	ItemId          uint                    `json:"item_id"`
	Status          string                  `json:"status"`
	Method          string                  `json:"method"`
	GrossAmount     string                  `json:"gross_amount"`
	BankName        string                  `json:"bank_name"`
	VaNumber        string                  `json:"va_number"`
	TransactionTime time.Time               `json:"transaction_time"`
	SettlementTime  time.Time               `json:"settlement_time"`
	ExpiryTime      time.Time               `json:"expiry_time"`
	Customer        PaymentCustomerResponse `json:"customer"`
	Item            PaymentItemResponse     `json:"item"`
}

type PaymentMonthlyHistoryResponse struct {
	Month string  `json:"month"`
	Total float64 `json:"total"`
}

type PaymentLastYearHistoryResponse struct {
	Total  float64                         `json:"total"`
	Months []PaymentMonthlyHistoryResponse `json:"months"`
}