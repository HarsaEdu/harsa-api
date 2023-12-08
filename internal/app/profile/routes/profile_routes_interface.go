package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/profile/handler"
	"github.com/labstack/echo/v4"
)

type ProfileRoutes interface {
	ProfileMobile(apiGroup *echo.Group)
	ProfileWeb(apiGroup *echo.Group)
}

type ProfileRoutesImpl struct {
	ProfileHandler handler.ProfileHandler
}

func NewProfileRoutes(profileHandler handler.ProfileHandler) ProfileRoutes {
	return &ProfileRoutesImpl{
		ProfileHandler: profileHandler,
	}
}
