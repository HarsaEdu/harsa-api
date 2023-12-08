package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/recommendations/handler"
	"github.com/labstack/echo/v4"
)

type RecommendationsRoutes interface {
	RecommendationsWeb(apiGroup *echo.Group) *echo.Group
	RecommendationsMobile(apiGroup *echo.Group) *echo.Group
}

type RecommendationsRoutesImpl struct {
	RecommendationsHandler handler.RecommendationsHandler
}

func NewRecommendationsRoutes(RecommendationsHandler handler.RecommendationsHandler) RecommendationsRoutes {
	return &RecommendationsRoutesImpl{
		RecommendationsHandler: RecommendationsHandler,
	}
}