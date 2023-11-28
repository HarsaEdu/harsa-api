package web

type CreatePaymentRequest struct {
	PlanId int `json:"plan_id" form:"plan_id" validate:"required"`
}