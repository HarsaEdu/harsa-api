package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (feedbackHandler *FeedbackHandlerImpl) GetAll(ctx echo.Context) error {
	params := ctx.QueryParams()
	courseid, err := strconv.Atoi(params.Get("course_id"))

	if err != nil {
		return res.StatusBadRequest(ctx, "params course not valid", err)
	}

	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return res.StatusBadRequest(ctx, "params limit not valid", err)
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return res.StatusBadRequest(ctx, "params offset not valid", err)
	}

	response, pagination, err := feedbackHandler.FeedbackService.GetAll(courseid, offset, limit)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "feedback not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get all feedback, something happen", err)
	}

	return res.StatusOK(ctx, "success get feedback", response, pagination)

}
