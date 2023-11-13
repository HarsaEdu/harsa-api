package routes

import "github.com/labstack/echo/v4"

func (authRoutes *AuthRoutesImpl) Auth(apiGroup *echo.Group) {
	authGroup := apiGroup.Group("/auth")

	authGroup.POST("/register", authRoutes.AuthHandler.RegisterUser)
	authGroup.POST("/login", authRoutes.AuthHandler.LoginUser)
}
