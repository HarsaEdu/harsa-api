package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/user/handler"
	"github.com/labstack/echo/v4"
)

type UserRoutes interface {
	User(apiGroup *echo.Group)
}

type UserRoutesImpl struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func NewUserRoutes(e *echo.Echo, userHandler handler.UserHandler) UserRoutes {
	return &UserRoutesImpl{
		Echo:        e,
		UserHandler: userHandler,
	}
}
