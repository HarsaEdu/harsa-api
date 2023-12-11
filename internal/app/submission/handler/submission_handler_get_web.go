package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (submissionHandler *SubmissionHandlerImpl) GetAllWeb(ctx echo.Context) error {

	idParam := ctx.Param("moduleId")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid module id", err)
	}
	data, err := submissionHandler.SubmissionService.GetAll(id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNoContent(ctx, "submissions not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get submissions, something happen", err)
	}

	return res.StatusOK(ctx, "success to get all submissions", data, nil)
}
