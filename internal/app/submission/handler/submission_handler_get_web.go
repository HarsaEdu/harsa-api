package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (submissionHandler *SubmissionHandlerImpl) GetAllWeb(ctx echo.Context) error {
	data, err := submissionHandler.SubmissionService.GetAll()
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
