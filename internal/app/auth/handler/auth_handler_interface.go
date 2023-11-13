package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/auth/service"
	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	RegisterUser(ctx echo.Context) error
	LoginUser(ctx echo.Context) error
}

type AuthHandlerImpl struct {
	AuthService service.AuthService
}

func NewAuthHandler(as service.AuthService) AuthHandler {
	return &AuthHandlerImpl{AuthService: as}
}
