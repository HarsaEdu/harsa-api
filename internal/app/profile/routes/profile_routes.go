package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (profileRoutes *ProfileRoutesImpl) ProfileWeb(apiGroup *echo.Group) {
	profilesGroup := apiGroup.Group("/profile")
	profilesGroup.GET("/:profile_id", profileRoutes.ProfileHandler.GetProfileByID, middleware.AllUserMiddleare)
	profilesGroup.PUT("/:profile_id", profileRoutes.ProfileHandler.UpdateProfile, middleware.AllUserMiddleare)
}

func (profileRoutes *ProfileRoutesImpl) ProfileMobile(apiGroup *echo.Group) {
	profilesGroup := apiGroup.Group("/profile")

	profilesGroup.POST("", profileRoutes.ProfileHandler.CreateProfile, middleware.AllUserMiddleare)
	profilesGroup.GET("", profileRoutes.ProfileHandler.GetProfileByID, middleware.AllUserMiddleare)
	profilesGroup.PUT("/:profile_id", profileRoutes.ProfileHandler.UpdateProfile, middleware.AllUserMiddleare)
}
