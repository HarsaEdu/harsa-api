package routes

import "github.com/labstack/echo/v4"

func (authRoutes *AuthRoutesImpl) AuthWeb(apiGroup *echo.Group) {
	authGroup := apiGroup.Group("/auth")

	authGroup.POST("/register", authRoutes.AuthHandler.RegisterUser)
	authGroup.POST("/login", authRoutes.AuthHandler.LoginUser)
}

func (authRoutes *AuthRoutesImpl) AuthMobile(apiGroup *echo.Group) {
	authGroup := apiGroup.Group("/auth")

	authGroup.POST("/register", authRoutes.AuthHandler.RegisterUser)
	authGroup.POST("/login", authRoutes.AuthHandler.LoginUserStudent)
	authGroup.POST("/access-token", authRoutes.AuthHandler.GetAccessToken)
}
