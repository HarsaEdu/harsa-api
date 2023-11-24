package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/module/handler"
	"github.com/labstack/echo/v4"
)

type ModuleRoutes interface {
	ModuleWeb(coursesGruop *echo.Group)
	ModuleMobile(coursesGruop *echo.Group)
}

type ModuleRoutesImpl struct {
	ModuleHandler handler.ModuleHandler
}

func NewModuleRoutes(ModuleHandler handler.ModuleHandler) ModuleRoutes {
	return &ModuleRoutesImpl{
		ModuleHandler: ModuleHandler,
	}
}