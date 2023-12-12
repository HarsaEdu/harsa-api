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

func (subsPlanHandler *SubsPlanHandlerImpl) Create(ctx echo.Context) error {
	req := web.SubsPlanCreateRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind subs plan request", err)
	}

	err = subsPlanHandler.SubsPlanService.Create(ctx, &req)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "file format") {
			return res.StatusBadRequest(ctx, "please input jpg or png file", err)
		}
		return res.StatusInternalServerError(ctx, "failed to create subs plan, something happen", err)
	}
	return res.StatusCreated(ctx, "success to create subs plan", nil, nil)

}

func (subsPlanHandler *SubsPlanHandlerImpl) CreateFromExisting(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusInternalServerError(ctx, fmt.Sprintf("failed to convert param id to int: %s", err.Error()), err)
	}
	
	req := web.SubsPlanUpdateRequest{}
	err = ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind subs plan request", err)
	}

	err = subsPlanHandler.SubsPlanService.CreateFromExisting(ctx, &req, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "subs plan not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to update subs plan, something happen", err)
	}
	return res.StatusCreated(ctx, "success to update subs plan", nil, nil)
}