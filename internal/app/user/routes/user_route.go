package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (userRoutes *UserRoutesImpl) UserWeb(apiGroup *echo.Group) {
	userGroup := apiGroup.Group("/users")
	userSubcribe  := apiGroup.Group("/course/:course-id/user")

	userGroup.POST("", userRoutes.UserHandler.UserCreate, middleware.AdminMiddleware)
	userGroup.PUT("/:id", userRoutes.UserHandler.UserUpdate, middleware.AdminMiddleware)
	userGroup.DELETE("/:id", userRoutes.UserHandler.UserDelete, middleware.AdminMiddleware)
	userGroup.PUT("/profile", userRoutes.UserHandler.UserProfileUpdate, middleware.InstructorMiddleware)
	userGroup.GET("", userRoutes.UserHandler.GetAllUsers, middleware.AdminMiddleware)
	userSubcribe.GET("/subscribe", userRoutes.UserHandler.GetAllStudentSubscribe, middleware.AdminMiddleware)
	userGroup.GET("/:id", userRoutes.UserHandler.GetUserDetailByID, middleware.AdminMiddleware)
	userGroup.GET("/account/:id", userRoutes.UserHandler.GetUserAccountByID, middleware.AdminMiddleware)
	userGroup.GET("/my-account", userRoutes.UserHandler.GetUserMyAccount, middleware.InstructorMiddleware)

}

func (userRoutes *UserRoutesImpl) UserMobile(apiGroup *echo.Group) {
	userGroup := apiGroup.Group("/users")

	userGroup.GET("/my-account", userRoutes.UserHandler.GetUserMyAccount, middleware.StudentMiddleare)
	userGroup.PUT("/profile", userRoutes.UserHandler.UserProfileUpdate, middleware.StudentMiddleare)
	userGroup.PUT("/my-account", userRoutes.UserHandler.UserUpdateMobile, middleware.StudentMiddleare)

}
