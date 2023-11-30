package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/profile/service"
	"github.com/labstack/echo/v4"
)

type ProfileHandler interface {
	CreateProfile(ctx echo.Context) error
	GetProfileByID(ctx echo.Context) error
	UpdateProfile(ctx echo.Context) error
	MyProfile(ctx echo.Context) error
}

type ProfileHandlerImpl struct {
	ProfileService service.ProfileService
}

func NewProfileHandler(ps service.ProfileService) *ProfileHandlerImpl {
	return &ProfileHandlerImpl{ProfileService: ps}
}
