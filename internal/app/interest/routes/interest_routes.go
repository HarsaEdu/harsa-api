package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (routes *InterestRoutesImpl) WebInterest(apiGroup *echo.Group) {
	interestGroup := apiGroup.Group("/users/interest")

	interestGroup.POST("", routes.Handler.CreateInterest, middleware.AllUserMiddleare)
	interestGroup.GET("", routes.Handler.GetInterestRecommendation, middleware.AllUserMiddleare)
}

func (routes *InterestRoutesImpl) MobileInterest(apiGroup *echo.Group) {
	interestGroup := apiGroup.Group("/users/interest")

	interestGroup.POST("", routes.Handler.CreateInterest, middleware.AllUserMiddleare)
	interestGroup.GET("", routes.Handler.GetInterestRecommendation, middleware.AllUserMiddleare)
}
