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

func (submissionHandler *SubmissionHandlerImpl) Update(ctx echo.Context) error {
	ParamId := ctx.Param("id")
	id, err := strconv.Atoi(ParamId)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid submission id", err)
	}

	req := web.SubmissionUpdateRequest{}
	err = ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind submission request", err)
	}

	userId := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	err = submissionHandler.SubmissionService.Update(ctx, &req, id, userId, roleString)
	if err != nil {
		if strings.Contains(err.Error(), "validate") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "submission not found", err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot update this submission" ,err)
		}
		return res.StatusInternalServerError(ctx, "failed update submission, something happen", err)

	}
	return res.StatusOK(ctx, "success to update submission", nil, nil)
}
