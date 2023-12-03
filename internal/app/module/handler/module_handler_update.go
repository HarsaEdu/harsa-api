package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (moduleHandler *ModuleHandlerImpl) UpdateModule(ctx echo.Context) error {
	idParam := ctx.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid module id", err)
	}

	user_id := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	request := web.ModuleRequest{}
	err = ctx.Bind(&request)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind request: ", err)
	}

	err = moduleHandler.ModuleService.UpdateModule(&request, uint(id), uint(user_id), roleString)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot update this module" ,err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "module not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to update module, something happen", err)
	}

	return res.StatusOK(ctx, "success to update module", nil, nil)
}

func (moduleHandler *ModuleHandlerImpl) UpdateModuleOrder(ctx echo.Context) error {
	idParam := ctx.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid module id", err)
	}

	user_id := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	request := web.ModuleOrderRequest{}
	err = ctx.Bind(&request)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind request: ", err)
	}

	err = moduleHandler.ModuleService.UpdateModuleOrder(&request, uint(id), uint(user_id), roleString)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot delete this module order" ,err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "module not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to update order module, something happen", err)
	}

	return res.StatusOK(ctx, "success to update order module", nil, nil)
}

func (moduleHandler *ModuleHandlerImpl) UpdateSection(ctx echo.Context) error {
	idParam := ctx.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid module id", err)
	}

	user_id := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	request := web.SectionRequest{}
	err = ctx.Bind(&request)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind request: ", err)
	}

	err = moduleHandler.ModuleService.UpdateSection(&request, uint(id), uint(user_id), roleString)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot update this section" ,err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "section not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to update section, something happen", err)
	}

	return res.StatusOK(ctx, "success to update section", nil, nil)
}

func (moduleHandler *ModuleHandlerImpl) UpdateSectionOrder(ctx echo.Context) error {
	idParam := ctx.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid module id", err)
	}

	user_id := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	request := web.SectionOrderRequest{}
	err = ctx.Bind(&request)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind request: ", err)
	}

	err = moduleHandler.ModuleService.UpdateSectionOrder(&request, uint(id), uint(user_id), roleString)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot update this section order" ,err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "section not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to update order section, something happen", err)
	}

	return res.StatusOK(ctx, "success to update order section", nil, nil)
}

