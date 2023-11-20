package web

type GetResponseRequest struct {
	Message string `json:"message" form:"message" validate:"required"`
}