package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (routes *InterestRoutesImpl) Interest(apiGroup *echo.Group) {
	interestGroup := apiGroup.Group("/:profile_id/interest")

	interestGroup.POST("", routes.Handler.CreateInterest, middleware.AllUserMiddleare)
}
