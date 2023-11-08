package web

type Role string

const (
	Student    Role = "student"
	Instructor Role = "instructor"
	Admin      Role = "admin"
)

type UserLoginResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     Role   `json:"role"`
}
