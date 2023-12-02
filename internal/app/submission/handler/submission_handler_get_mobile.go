package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (submissionHandler *SubmissionHandlerImpl) GetAllMobile(ctx echo.Context) error {
	idParam := ctx.Param("moduleId")
	id, _ := strconv.Atoi(idParam)
	data, err := submissionHandler.SubmissionService.GetAllMobile(id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "submissions not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get submissions, something happen", err)
	}

	return res.StatusOK(ctx, "success to get all submissions", data, nil)
}
