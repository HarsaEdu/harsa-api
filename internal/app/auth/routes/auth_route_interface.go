package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/auth/handler"
	"github.com/labstack/echo/v4"
)

type AuthRoutes interface {
	Auth(apiGroup *echo.Group)
}

type AuthRoutesImpl struct {
	Echo        *echo.Echo
	AuthHandler handler.AuthHandler
}

func NewAuthRoutes(e *echo.Echo, authHandler handler.AuthHandler) AuthRoutes {
	return &AuthRoutesImpl{
		Echo:        e,
		AuthHandler: authHandler,
	}
}