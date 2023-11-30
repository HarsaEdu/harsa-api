package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (feedbackHandler *FeedbackHandlerImpl) FindById(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	result, err := feedbackHandler.FeedbackService.FindById(id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "feedback not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get feedback, something happen", err)
	}

	return res.StatusOK(ctx, "success to get feedback", result, nil)
}
