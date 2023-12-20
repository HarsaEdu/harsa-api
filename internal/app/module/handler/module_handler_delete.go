package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (moduleHandler *ModuleHandlerImpl) DeleteModule(ctx echo.Context) error {

	idParam := ctx.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid module id", err)
	}

	user_id := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	err = moduleHandler.ModuleService.DeleteModule(uint(id), uint(user_id), roleString)
	if err != nil {
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot delete this module" ,err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "module not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to delete module, something happen", err)
	}

	return res.StatusOK(ctx, "success to delete module", nil, nil)
}

func (moduleHandler *ModuleHandlerImpl) DeleteSection(ctx echo.Context) error {

	idParam := ctx.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid section id", err)
	}

	user_id := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	err = moduleHandler.ModuleService.DeleteSection(uint(id), uint(user_id), roleString)
	if err != nil {
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot delete this section" ,err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "section not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to delete section, something happen", err)
	}

	return res.StatusOK(ctx, "success to delete section", nil, nil)
}

func (moduleHandler *ModuleHandlerImpl) DeleteSubModule(ctx echo.Context) error {

	idParam := ctx.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid subModule id", err)
	}

	user_id := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	err = moduleHandler.ModuleService.DeleteSubModule(uint(id), uint(user_id), roleString)
	if err != nil {
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot delete this sub module" ,err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "sub module not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to delete sub module, something happen", err)
	}

	return res.StatusOK(ctx, "success to delete sub module", nil, nil)
}

