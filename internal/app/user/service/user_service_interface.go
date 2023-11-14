package service

import (
	userRepo "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	UserCreate(ctx echo.Context, userRequest web.UserCreateRequest) error
}

type UserServiceImpl struct {
	UserRepository userRepo.UserRepository
	Validate       *validator.Validate
}

func NewUserService(userRepository userRepo.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}
