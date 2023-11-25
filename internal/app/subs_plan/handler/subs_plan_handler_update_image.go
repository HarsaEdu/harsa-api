package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (subsPlanHandler *SubsPlanHandlerImpl) UpdateImage(ctx echo.Context) error {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to convert param id to int: ", err)
	}

	req := web.SubsPlanUpdateImage{}
	err = ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind subs plan update request", err)
	}

	err = subsPlanHandler.SubsPlanService.UpdateImage(ctx, &req, id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "subs plan not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed update image subs plan, something happen", err)

	}
	return res.StatusOK(ctx, "success to update image", nil, nil)
}
