package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/auth/repository"
	userRepo "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
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
}

func NewAuthService(ar repository.AuthRepository, ur userRepo.UserRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: ar,
		UserRepository: ur,
		Validate:       validate,
	}
}
