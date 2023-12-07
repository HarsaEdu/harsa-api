package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/submission/service"
	"github.com/labstack/echo/v4"
)

type SubmissionHandler interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	GetAllWeb(ctx echo.Context) error
	GetAllMobile(ctx echo.Context) error
	FindById(ctx echo.Context) error
}

type SubmissionHandlerImpl struct {
	SubmissionService service.SubmissionService
}

func NewSubmissionHandler(service service.SubmissionService) SubmissionHandler {
	return &SubmissionHandlerImpl{SubmissionService: service}
}
