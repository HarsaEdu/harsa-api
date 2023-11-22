package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/subs_plan/handler"
	"github.com/labstack/echo/v4"
)

type SubsPlanRoutes interface {
	SubsPlan(apiGroup *echo.Group)
}

type SubsPlanRoutesImpl struct {
	Echo            *echo.Echo
	subsPlanHandler handler.SubsPlanHandler
}

func NewSubsPlanRoutes(e *echo.Echo, subsPlanHandler handler.SubsPlanHandler) SubsPlanRoutes {
	return &SubsPlanRoutesImpl{
		Echo:            e,
		subsPlanHandler: subsPlanHandler,
	}
}
