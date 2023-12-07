package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/auth/handler"
	"github.com/labstack/echo/v4"
)

type AuthRoutes interface {
	AuthWeb(apiGroup *echo.Group)
	AuthMobile(apiGroup *echo.Group)
}

type AuthRoutesImpl struct {
	AuthHandler handler.AuthHandler
}

func NewAuthRoutes(authHandler handler.AuthHandler) AuthRoutes {
	return &AuthRoutesImpl{
		AuthHandler: authHandler,
	}
}