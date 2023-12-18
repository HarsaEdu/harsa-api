package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/recommendations/service"
	"github.com/labstack/echo/v4"
)

type RecommendationsHandler interface {
	GetRecommendations(ctx echo.Context) error
	SendRecommendationsForInstructor(ctx echo.Context) error
}

type RecommendationsHandlerImpl struct {
	RecommendationsService service.RecommendationsService
}

func NewRecommendationsHandler(recommendationsService service.RecommendationsService) RecommendationsHandler {
	return &RecommendationsHandlerImpl{
		RecommendationsService: recommendationsService,
	}
}