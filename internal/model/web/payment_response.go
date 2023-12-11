package web

import "time"

type PaymentCustomerResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`

}

type PaymentItemResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetPaymentResponse struct {
	ID          string                  `json:"id"`
	UserId      uint                    `json:"user_id"`
	ItemId      uint                    `json:"item_id"`
	Status      string                  `json:"status"`
	Method      string                  `json:"method"`
	GrossAmount string                     `json:"gross_amount"`
	BankName    string                  `json:"bank_name"`
	VaNumber    string                  `json:"va_number"`
	CreatedAt   time.Time               `json:"created_at"`
	UpdatedAt   time.Time               `json:"updated_at"`
	ExpiredAt   time.Time               `json:"expired_at"`
	Customer    PaymentCustomerResponse `json:"customer"`
	Item        PaymentItemResponse     `json:"item"`
}