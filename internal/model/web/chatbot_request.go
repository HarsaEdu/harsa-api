package web

type CreateThreadRequest struct {
	Topic string `json:"topic" form:"topic" validate:"required"`
}

type GetResponseRequest struct {
	Message string `json:"message" form:"message" validate:"required"`
}

type ChatWithAssistantRequest struct {
	Message string `json:"message" form:"message" validate:"required"`
}