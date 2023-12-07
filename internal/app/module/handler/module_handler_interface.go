package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/module/service"
	"github.com/labstack/echo/v4"
)

type ModuleHandler interface {
	CreateSection(ctx echo.Context) error
	CreateModule(ctx echo.Context) error
	GetAllSection(ctx echo.Context) error
	GetAllModule(ctx echo.Context) error
	GetModuleById(ctx echo.Context) error
	UpdateModule(ctx echo.Context) error
	UpdateModuleOrder(ctx echo.Context) error
	UpdateSection(ctx echo.Context) error
	UpdateSectionOrder(ctx echo.Context) error
	DeleteModule(ctx echo.Context) error
	DeleteSection(ctx echo.Context) error
	DeleteSubModule(ctx echo.Context) error
}

type ModuleHandlerImpl struct {
	ModuleService service.ModuleService
}

func NewModuleHandler(ModuleService service.ModuleService) ModuleHandler {
	return &ModuleHandlerImpl{
		ModuleService: ModuleService,
	}
}