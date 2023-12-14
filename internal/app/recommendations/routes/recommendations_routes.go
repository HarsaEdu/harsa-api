package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (recommendationsRoutes *RecommendationsRoutesImpl) RecommendationsWeb(apiGroup *echo.Group) *echo.Group {
	recommendationsGroup := apiGroup.Group("/recommendations")

	recommendationsGroup.GET("", recommendationsRoutes.RecommendationsHandler.SendRecommendationsForInstructor, middleware.AdminMiddleware)
	return recommendationsGroup
}

func (recommendationsRoutes *RecommendationsRoutesImpl) RecommendationsMobile(apiGroup *echo.Group) *echo.Group {
	recommendationsGroup := apiGroup.Group("/recommendations")

	recommendationsGroup.POST("", recommendationsRoutes.RecommendationsHandler.GetRecommendations, middleware.StudentMiddleare)

	return recommendationsGroup
}