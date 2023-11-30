package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/faqs/handler"
	"github.com/labstack/echo/v4"
)

type FaqsRoutes interface {
	Faqs(apiGroup *echo.Group)
}

type FaqsRoutesImpl struct {
	Echo        *echo.Echo
	FaqsHandler handler.FaqsHandler
}

func NewFaqsRoutes(e *echo.Echo, FaqsHandler handler.FaqsHandler) FaqsRoutes {
	return &FaqsRoutesImpl{
		Echo:        e,
		FaqsHandler: FaqsHandler,
	}
}
