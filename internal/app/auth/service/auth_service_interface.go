package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/auth/repository"
	notifRepo "github.com/HarsaEdu/harsa-api/internal/app/notification/repository"
	userRepo "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
	"github.com/HarsaEdu/harsa-api/internal/pkg/firebase"
)

type AuthService interface {
	RegisterUser(userRequest web.RegisterUserRequest) (*web.AuthResponse, error)
	LoginUser(userRequest web.LoginUserRequest) (*web.AuthResponse, error)
	GetAccessToken(tokenRequest web.AccessTokenRequest) (*web.AuthResponse, error)
}

type AuthServiceImpl struct {
	AuthRepository         repository.AuthRepository
	UserRepository         userRepo.UserRepository
	NotificationRepository notifRepo.NotificationRepository
	Validate               *validator.Validate
	Firebase               firebase.Firebase
}

func NewAuthService(ar repository.AuthRepository, ur userRepo.UserRepository, validate *validator.Validate, notifRepository notifRepo.NotificationRepository, firebaseImpl firebase.Firebase) AuthService {
	return &AuthServiceImpl{
		AuthRepository:         ar,
		UserRepository:         ur,
		Validate:               validate,
		NotificationRepository: notifRepository,
		Firebase: firebaseImpl,
	}
}
