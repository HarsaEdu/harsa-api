package web

type RegisterUserRequest struct {
	Email             string `json:"email" validate:"required,email"`
	Username          string `json:"username" validate:"required"`
	Password          string `json:"password" validate:"required,min=8,max=255"`
	RegistrationToken string `json:"registration_token"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
