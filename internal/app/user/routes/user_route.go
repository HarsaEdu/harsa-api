package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (userRoutes *UserRoutesImpl) User(apiGroup *echo.Group) {
	userGroup := apiGroup.Group("/users")

	userGroup.POST("", userRoutes.UserHandler.UserCreate, middleware.AdminMiddleware)
}
