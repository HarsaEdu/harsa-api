package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/faqs/service"
	"github.com/labstack/echo/v4"
)

type FaqsHandler interface {
	Create(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	GetById(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Update(ctx echo.Context) error
}

type FaqsHandlerImpl struct {
	FaqsService service.FaqsService
}

func NewFaqsHandler(service service.FaqsService) FaqsHandler {
	return &FaqsHandlerImpl{FaqsService: service}
}
