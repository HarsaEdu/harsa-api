package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/auth/repository"
	userRepo "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
	"github.com/HarsaEdu/harsa-api/internal/pkg/password"
)

type AuthService interface {
	RegisterUser(userRequest web.RegisterUserRequest) (*web.AuthResponse, error)
	LoginUser(userRequest web.LoginUserRequest) (*web.AuthResponse, error)
	GetAccessToken(tokenRequest web.AccessTokenRequest) (*web.AuthResponse, error)
}

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	UserRepository userRepo.UserRepository
	Validate       *validator.Validate
	Password       password.PasswordComparer
}

func NewAuthService(ar repository.AuthRepository, ur userRepo.UserRepository, validate *validator.Validate, passwordImpl password.PasswordComparer) AuthService {
	return &AuthServiceImpl{
		AuthRepository: ar,
		UserRepository: ur,
		Validate:       validate,
		Password: passwordImpl,
	}
}
