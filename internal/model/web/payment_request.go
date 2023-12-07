package web

type CreatePaymentSubscriptionRequest struct {
	PlanId   int    `json:"plan_id" form:"plan_id" validate:"required"`
	BankName string `json:"bank_name" form:"bank_name" validate:"required"`
}