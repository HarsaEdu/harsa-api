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
	UserGetAll(offset int, limit int, search string, roleId int) ([]domain.UserEntity, *web.Pagination, error)
	GetUserDetail(userRequest web.UserGetByIDRequest) (*domain.UserDetail, error)
	GetUserAccount(userID uint) (*web.UserAccountResponse, error)
	UserUpdateMobile(userRequest web.UserUpdateRequestMobile) error
	UserGetAllStudentSubscribe(offset int, limit int, search string, courseId uint) ([]domain.UserEntity, *web.Pagination, error)
	UserUpdatePasswordMobile(userRequest web.UserUpdatePasswordRequestMobile) error
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
