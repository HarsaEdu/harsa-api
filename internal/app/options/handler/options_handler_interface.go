package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/options/service"
	"github.com/labstack/echo/v4"
)

type OptionsHandler interface {
	Delete(ctx echo.Context) error
}

type OptionsHandlereImpl struct {
	OptionsService service.OptionsService
}

func NewOptionsHandler(service service.OptionsService) OptionsHandler {
	return &OptionsHandlereImpl{
		OptionsService: service,
	}
}
