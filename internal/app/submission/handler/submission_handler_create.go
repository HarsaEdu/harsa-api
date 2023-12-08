package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (submissionHandler *SubmissionHandlerImpl) Create(ctx echo.Context) error {

	paramId := ctx.Param("moduleId")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid module id", err)
	}
	req := web.SubmissionRequest{}
	err = ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind submission request", err)
	}

	req.ModuleID = uint(id)
	err = submissionHandler.SubmissionService.Create(ctx, req)
	if err != nil {
		if strings.Contains(err.Error(), "validate") {
			return validation.ValidationError(ctx, err)
		}
		return res.StatusInternalServerError(ctx, "failed to create submission, something happen", err)

	}
	return res.StatusCreated(ctx, "success to create submission", nil, nil)
}
