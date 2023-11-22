package web

type GetMessageInThreadResponse struct {
	ID        string `json:"id" form:"id"`
	Role      string `json:"role" form:"role"`
	Message   string `json:"message" form:"message"`
	CreatedAt int    `json:"created_at" form:"created_at"`
}