package routes

import "github.com/labstack/echo/v4"

func (profileRoutes *ProfileRoutesImpl) Profile(apiGroup *echo.Group) {
	profilesGroup := apiGroup.Group("/user_profile")

	profilesGroup.POST("", profileRoutes.ProfileHandler.CreateProfile)
	profilesGroup.GET("/:profile_id", profileRoutes.ProfileHandler.GetProfileByID)
	profilesGroup.PUT("/:profile_id", profileRoutes.ProfileHandler.UpdateProfile)
}
