package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/faqs/handler"
	"github.com/labstack/echo/v4"
)

type FaqsRoutes interface {
	FaqsWeb(apiGroup *echo.Group)
	FaqsMobile(apiGroup *echo.Group)
}

type FaqsRoutesImpl struct {
	FaqsHandler handler.FaqsHandler
}

func NewFaqsRoutes(FaqsHandler handler.FaqsHandler) FaqsRoutes {
	return &FaqsRoutesImpl{
		FaqsHandler: FaqsHandler,
	}
}
