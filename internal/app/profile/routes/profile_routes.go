package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (profileRoutes *ProfileRoutesImpl) ProfileMobile(apiGroup *echo.Group) {
	profilesGroup := apiGroup.Group("/users/profile")

	profilesGroup.POST("", profileRoutes.ProfileHandler.CreateProfile, middleware.AllUserMiddleare)
	profilesGroup.PUT("/myprofile", profileRoutes.ProfileHandler.UpdateMyProfile, middleware.AllUserMiddleare)
	profilesGroup.GET("/myprofile", profileRoutes.ProfileHandler.MyProfile, middleware.AllUserMiddleare)
}
