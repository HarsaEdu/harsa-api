package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/user/service"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	UserCreate(ctx echo.Context) error
	UserUpdate(ctx echo.Context) error
	UserProfileUpdate(ctx echo.Context) error
	UserDelete(ctx echo.Context) error
	GetAllUsers(ctx echo.Context) error
	GetUserDetailByID(ctx echo.Context) error
}

type UserHandlerImpl struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &UserHandlerImpl{UserService: userService}
}
