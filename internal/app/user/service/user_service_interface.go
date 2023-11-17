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
	UserUpdate(userRequest web.UserUpdateRequest) error
	UserProfileUpdate(userRequest web.UserProfileUpdateRequest) error
	UserDelete(userRequest web.UserDeleteRequest) error
	UserGetAll(offset int, limit int, search string) ([]domain.UserEntity, int64, error)
	GetUserDetail(userRequest web.UserGetByIDRequest) (*domain.UserDetail, error)
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
