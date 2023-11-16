package service

import (
	userRepo "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	UserCreate(ctx echo.Context, userRequest web.UserCreateRequest) error
	UserUpdate(ctx echo.Context, userRequest web.UserUpdateRequest) error
	UserProfileUpdate(ctx echo.Context, userRequest web.UserProfileUpdateRequest) error
	UserDelete(ctx echo.Context, userRequest web.UserDeleteRequest) error
	UserGetAll(ctx echo.Context) ([]domain.UserEntity, int64, error)
	GetUserDetail(ctx echo.Context, userRequest web.UserGetByIDRequest) (*domain.UserDetail, error)
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
