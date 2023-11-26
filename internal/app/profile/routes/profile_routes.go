package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (profileRoutes *ProfileRoutesImpl) ProfileMobile(apiGroup *echo.Group) {
	profilesGroup := apiGroup.Group("/users/profile")

	profilesGroup.POST("", profileRoutes.ProfileHandler.CreateProfile, middleware.AllUserMiddleare)
	profilesGroup.GET("/:profile_id", profileRoutes.ProfileHandler.GetProfileByID, middleware.AdminMiddleware)
	profilesGroup.PUT("/:profile_id", profileRoutes.ProfileHandler.UpdateProfile, middleware.AllUserMiddleare)
}
