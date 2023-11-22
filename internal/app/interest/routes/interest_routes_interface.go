package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/interest/handler"
	"github.com/labstack/echo/v4"
)

type InterestRoutes interface {
	Interest(apiGroup *echo.Group)
}

type InterestRoutesImpl struct {
	Handler handler.InterestHandler
}

func NewInterestRoutes(handler handler.InterestHandler) InterestRoutes {
	return &InterestRoutesImpl{
		Handler: handler,
	}
}
