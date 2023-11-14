package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/profile/handler"
	"github.com/labstack/echo/v4"
)

type ProfileRoutes interface {
	Profile(apiGroup *echo.Group)
}

type ProfileRoutesImpl struct {
	Echo           *echo.Echo
	ProfileHandler handler.ProfileHandler
}

func NewProfileRoutes(e *echo.Echo, profileHandler handler.ProfileHandler) ProfileRoutes {
	return &ProfileRoutesImpl{
		Echo:           e,
		ProfileHandler: profileHandler,
	}
}
