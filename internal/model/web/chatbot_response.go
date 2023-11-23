package web

import "time"

type GetAllThreadByUserIdResponse struct {
	ID        string   `json:"id" form:"id"`
	UserId    int      `json:"user_id" form:"user_id"`
	Topic     string   `json:"topic" form:"topic"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type GetMessageInThreadResponse struct {
	ID        string `json:"id" form:"id"`
	Role      string `json:"role" form:"role"`
	Message   string `json:"message" form:"message"`
	CreatedAt int    `json:"created_at" form:"created_at"`
}