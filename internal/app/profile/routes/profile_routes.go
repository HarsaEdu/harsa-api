package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (profileRoutes *ProfileRoutesImpl) ProfileMobile(apiGroup *echo.Group) {
	profilesGroup := apiGroup.Group("/users/profile")

	profilesGroup.POST("/myprofile", profileRoutes.ProfileHandler.CreateProfile, middleware.AllUserMiddleare)
	profilesGroup.PUT("/myprofile", profileRoutes.ProfileHandler.UpdateMyProfile, middleware.AllUserMiddleare)
	profilesGroup.GET("/myprofile", profileRoutes.ProfileHandler.MyProfile, middleware.AllUserMiddleare)
}

func (profileRoutes *ProfileRoutesImpl) ProfileWeb(apiGroup *echo.Group) {
	profilesGroup := apiGroup.Group("/users/profile")

	profilesGroup.PUT("/:profile_id", profileRoutes.ProfileHandler.UpdateProfile, middleware.AdminMiddleware)
	profilesGroup.GET("/myprofile", profileRoutes.ProfileHandler.MyProfile, middleware.AllUserMiddleare)
	profilesGroup.PUT("/myprofile", profileRoutes.ProfileHandler.UpdateMyProfile, middleware.AllUserMiddleare)
}
