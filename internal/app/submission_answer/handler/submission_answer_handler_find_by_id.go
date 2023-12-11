package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (submissionAnswerHanlder *SubmissionAnswerHandlerImpl) FindById(ctx echo.Context) error {
	idParam := ctx.Param("subsAnsId")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to convert param id to int: ", err)
	}

	result, err := submissionAnswerHanlder.SubmissionAnswerservice.FindById(id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNoContent(ctx, "submission answer not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get submission answer, something happen", err)
	}

	return res.StatusOK(ctx, "success to get submission answer", result, nil)
}
