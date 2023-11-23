package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/interest/service"
	"github.com/labstack/echo/v4"
)

type InterestHandler interface {
	CreateInterest(ctx echo.Context) error
	GetInterestRecommendation(ctx echo.Context) error
}

type InterestHandlerImpl struct {
	Service service.InterestService
}

func NewInterestHandler(service service.InterestService) InterestHandler {
	return &InterestHandlerImpl{
		Service: service,
	}
}
