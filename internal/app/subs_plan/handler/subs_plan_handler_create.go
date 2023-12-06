package handler

import (
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
