package middleware

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/jwt"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc, condition string) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header["Authorization"]

		if len(authorization) <= 0 {
			return res.StatusUnauthorized(c, "missing or malformed jwt.", fmt.Errorf("unauthorized"))
		}

		userToken := strings.Split(authorization[0], " ")

		if len(userToken) <= 1 {
			return res.StatusBadRequest(c, "invalid token.", fmt.Errorf("invalid token"))
		}

		data, err := jwt.ExtractToken(userToken[1])

		if err != nil {
			return res.StatusBadRequest(c, "invalid token!", err)
		}

		if condition == "admin" {
			if data.RoleName != "admin" {
				return res.StatusForbidden(c, "you dont have access!", fmt.Errorf("access forbidden!"))
			}
		}

		if condition == "instructor" {
			if data.RoleName != "instructor" && data.RoleName != "admin" {
				return res.StatusForbidden(c, "you dont have access!", fmt.Errorf("access forbidden!"))
			}
		}

		if condition == "student" {
			if data.RoleName != "student" && data.RoleName != "admin" {
				return res.StatusForbidden(c, "you dont have access!", fmt.Errorf("access forbidden!"))
			}
		}

		c.Set("user_id", data.ID)
		c.Set("role_name", data.RoleName)
		c.Set("username", data.Username)

		return next(c)
	}
}

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return AuthMiddleware(next, "admin")
}

func InstructorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return AuthMiddleware(next, "instructor")
}
func StudentMiddleare(next echo.HandlerFunc) echo.HandlerFunc {
	return AuthMiddleware(next, "student")
}
