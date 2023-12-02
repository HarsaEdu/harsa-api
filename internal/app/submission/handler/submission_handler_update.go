package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (submissionHandler *SubmissionHandlerImpl) Update(ctx echo.Context) error {
	ParamId := ctx.Param("id")
	id, _ := strconv.Atoi(ParamId)

	req := web.SubmissionUpdateRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind submission request", err)
	}

	err = submissionHandler.SubmissionService.Update(ctx, &req, id)
	if err != nil {
		if strings.Contains(err.Error(), "validate") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "submission not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed update submission, something happen", err)

	}
	return res.StatusOK(ctx, "success to update submission", nil, nil)
}
