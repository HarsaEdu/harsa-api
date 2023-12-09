package web

import "time"

type UserAccountResponse struct {
	ID        uint      `json:"id" `
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRole struct {
	ID   uint   `json:"id" `
	Name string `json:"name"`
}