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

func (subsPlanHandler *SubsPlanHandlerImpl) SetStatusDelete(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusInternalServerError(ctx, fmt.Sprintf("failed to convert param id to int: %s", err.Error()), err)
	}
	
	request := web.SubsPlanUpdateStatus{
		IsActive: false,
	}

	err = subsPlanHandler.SubsPlanService.SetStatus(&request, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "subs plan not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to set status to delete, something happen", err)
	}

	return res.StatusOK(ctx, "succes to set status to delete", nil, nil)
}

func (subsPlanHandler *SubsPlanHandlerImpl) SetStatusActive(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusInternalServerError(ctx, fmt.Sprintf("failed to convert param id to int: %s", err.Error()), err)
	}
	
	request := web.SubsPlanUpdateStatus{
		IsActive: true,
	}

	err = subsPlanHandler.SubsPlanService.SetStatus(&request, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "subs plan not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to set status to active, something happen", err)
	}

	return res.StatusOK(ctx, "succes to set status to active", nil, nil)
}

