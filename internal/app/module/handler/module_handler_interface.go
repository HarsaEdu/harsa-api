package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/module/service"
	"github.com/labstack/echo/v4"
)

type ModuleHandler interface {
	Create(ctx echo.Context) error
	// GetAll(ctx echo.Context) error
	// GetById(ctx echo.Context) error
	// Update(ctx echo.Context) error
	// UpdateImage(ctx echo.Context) error
	// Delete(ctx echo.Context) error
}

type ModuleHandlerImpl struct {
	ModuleService service.ModuleService
}

func NewModuleHandler(ModuleService service.ModuleService) ModuleHandler {
	return &ModuleHandlerImpl{
		ModuleService: ModuleService,
	}
}