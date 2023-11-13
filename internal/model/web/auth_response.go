package web

type AuthResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleName Role   `json:"role_name"`
}
