package web

type CreateSubscribeRequest struct {
	PlanId int `json:"plan_id" form:"plan_id" validate:"required"`
}