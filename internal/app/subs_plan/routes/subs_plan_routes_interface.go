package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/subs_plan/handler"
	"github.com/labstack/echo/v4"
)

type SubsPlanRoutes interface {
	SubsPlanWeb(apiGroup *echo.Group)
	SubsPlanMobile(apiGroup *echo.Group)
}

type SubsPlanRoutesImpl struct {
	subsPlanHandler handler.SubsPlanHandler
}

func NewSubsPlanRoutes(subsPlanHandler handler.SubsPlanHandler) SubsPlanRoutes {
	return &SubsPlanRoutesImpl{
		subsPlanHandler: subsPlanHandler,
	}
}
