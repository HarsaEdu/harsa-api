package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/options/handler"
	"github.com/labstack/echo/v4"
)

type OptionsRoutes interface {
	OptionsWeb(apiGroup *echo.Group)
}

type OptionsRoutesImpl struct {
	OptionsHandler   handler.OptionsHandler
}

func NewOptionsRoutes(optionsHandler handler.OptionsHandler) OptionsRoutes {
	return &OptionsRoutesImpl{
		OptionsHandler: optionsHandler,
	}
}