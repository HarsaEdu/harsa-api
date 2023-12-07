package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (userRoutes *UserRoutesImpl) UserWeb(apiGroup *echo.Group) {
	userGroup := apiGroup.Group("/users")

	userGroup.POST("", userRoutes.UserHandler.UserCreate, middleware.AdminMiddleware)
	userGroup.PUT("", userRoutes.UserHandler.UserUpdate, middleware.AdminMiddleware)
	userGroup.DELETE("", userRoutes.UserHandler.UserDelete, middleware.AdminMiddleware)
	userGroup.PUT("/profile", userRoutes.UserHandler.UserProfileUpdate, middleware.InstructorMiddleware)
	userGroup.GET("", userRoutes.UserHandler.GetAllUsers, middleware.AdminMiddleware)
	userGroup.GET("/:id", userRoutes.UserHandler.GetUserDetailByID, middleware.AdminMiddleware)
}

func (userRoutes *UserRoutesImpl) UserMobile(apiGroup *echo.Group) {
	userGroup := apiGroup.Group("/users")

	userGroup.PUT("/profile", userRoutes.UserHandler.UserProfileUpdate, middleware.StudentMiddleare)

}
