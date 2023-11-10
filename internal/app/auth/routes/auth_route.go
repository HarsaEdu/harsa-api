package routes

import "github.com/labstack/echo/v4"

func (ar *AuthRoutesImpl) Auth(apiGroup *echo.Group) {
	authGroup := apiGroup.Group("/auth")

	authGroup.POST("/register", ar.AuthHandler.RegisterUser)
}