package web

import "time"

type Role string

const (
	Student    Role = "student"
	Instructor Role = "instructor"
	Admin      Role = "admin"
)

type AuthResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	RoleName  Role      `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	RoleName Role   `json:"role_name"`
	Token    string `json:"token"`
}
