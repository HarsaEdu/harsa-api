package web

type Role string

const (
	Student    Role = "student"
	Instructor Role = "instructor"
	Admin      Role = "admin"
)

type UserLoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	RoleName Role   `json:"role_name"`
	Token    string `json:"token"`
}
