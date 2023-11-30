package handler

import (
	service "github.com/HarsaEdu/harsa-api/internal/app/subs_plan/service"
	"github.com/labstack/echo/v4"
)

type SubsPlanHandler interface {
	Create(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Update(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	UpdateImage(ctx echo.Context) error
}

type SubsPlanHandlerImpl struct {
	SubsPlanService service.SubsPlanService
}

func NewFaqsHandler(service service.SubsPlanService) SubsPlanHandler {
	return &SubsPlanHandlerImpl{SubsPlanService: service}
}
