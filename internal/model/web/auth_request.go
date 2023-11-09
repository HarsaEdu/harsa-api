package web

type RegisterUserRequest struct {
	Email             string `json:"email" validate:"required,email"`
	Username          string `json:"username" validate:"required"`
	Password          string `json:"password" validate:"required,min=8,max=255"`
	RoleID            uint   `json:"role_id" validate:"required"`
	RegistrationToken string `json:"registration_token"`
}
