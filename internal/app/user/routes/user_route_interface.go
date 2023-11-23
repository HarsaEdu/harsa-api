package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/user/handler"
	"github.com/labstack/echo/v4"
)

type UserRoutes interface {
	UserWeb(apiGroup *echo.Group)
	UserMobile(apiGroup *echo.Group)
}

type UserRoutesImpl struct {
	UserHandler handler.UserHandler
}

func NewUserRoutes(userHandler handler.UserHandler) UserRoutes {
	return &UserRoutesImpl{
		UserHandler: userHandler,
	}
}
