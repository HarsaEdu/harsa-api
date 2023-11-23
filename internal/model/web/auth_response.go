package web

type Role string

const (
	Student    Role = "student"
	Instructor Role = "instructor"
	Admin      Role = "admin"
)

type AuthResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleName Role   `json:"role_name"`
}

type UserLoginResponse struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	RoleName     Role   `json:"role_name"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
