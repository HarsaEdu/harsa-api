package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (moduleHandler *ModuleHandlerImpl) GetAllSection(ctx echo.Context) error {
	paramCourseId := ctx.Param("courseId")
	courseId, err := strconv.Atoi(paramCourseId)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid course id", err)
	}

	params := ctx.QueryParams()
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid offset", err)
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit", err)
	}

	response, pagination, err := moduleHandler.ModuleService.GetAllSectionByCourseId(offset, limit, params.Get("search"), uint(courseId))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "section not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get all section, something happen", err)
	}

	return res.StatusOK(ctx, "success to get section", response, pagination)
}

func (moduleHandler *ModuleHandlerImpl) GetAllModule(ctx echo.Context) error {
	paramSectionId := ctx.Param("id")
	sectionId, err := strconv.Atoi(paramSectionId)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid ection id", err)
	}

	response, err := moduleHandler.ModuleService.GetAllModuleBySectionId(uint(sectionId))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "module not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get all module, something happen", err)
	}

	return res.StatusOK(ctx, "success to get module", response, nil)
}

func (moduleHandler *ModuleHandlerImpl) GetModuleById(ctx echo.Context) error {
	paramCourseId := ctx.Param("id")
	courseId, err := strconv.Atoi(paramCourseId)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid ection id", err)
	}

	response, err := moduleHandler.ModuleService.GetModuleById(uint(courseId))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "module not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get all module, something happen", err)
	}

	return res.StatusOK(ctx, "success to get module", response, nil)
}