package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (profileRoutes *ProfileRoutesImpl) Profile(apiGroup *echo.Group) {
	profilesGroup := apiGroup.Group("/user_profile")

	profilesGroup.POST("", profileRoutes.ProfileHandler.CreateProfile, middleware.AllUserMiddleare)
	profilesGroup.GET("", profileRoutes.ProfileHandler.GetProfileByID, middleware.AllUserMiddleare)
	profilesGroup.PUT("", profileRoutes.ProfileHandler.UpdateProfile, middleware.AllUserMiddleare)
}
