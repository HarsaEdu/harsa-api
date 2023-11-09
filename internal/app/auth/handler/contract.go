package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/auth/service"
	"github.com/labstack/echo/v4"
)

type AuthHanlder interface {
	RegisterUser(ctx echo.Context) error
}

type AuthHanlderImpl struct {
	AuthService service.AuthService
}

func NewAuthHanlder(as service.AuthService) AuthHanlder {
	return &AuthHanlderImpl{AuthService: as}
}
