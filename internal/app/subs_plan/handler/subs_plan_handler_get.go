package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (subsPlanHandler *SubsPlanHandlerImpl) GetAll(ctx echo.Context) error {
	params := ctx.QueryParams()
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid offset", err)
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit", err)
	}

	response, pagination, err := subsPlanHandler.SubsPlanService.GetAll(offset, limit, params.Get("search"))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "subs plan not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get all subs plan, something happen", err)
	}

	return res.StatusOK(ctx, "succes to get subsplan", response, pagination)
}
