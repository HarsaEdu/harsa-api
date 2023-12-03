package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (feedbackHandler *FeedbackHandlerImpl) UpdateByUserAndCourseId(ctx echo.Context) error {
	courseIdParam := ctx.Param("courseId")
	courseId, _ := strconv.Atoi(courseIdParam)

	userId := ctx.Get("user_id").(uint)

	req := web.FeedbackUpdateRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind feedback request", err)
	}

	err = feedbackHandler.FeedbackService.UpdateByUserAndCourseId(req, userId, uint(courseId))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "feedback not found", err)
		}

		if strings.Contains(err.Error(), "already exist") {
			return res.StatusAlreadyExist(ctx, "feedback name already exist", err)
		}

		return res.StatusInternalServerError(ctx, "failed to update feedback, something happen", err)

	}

	return res.StatusOK(ctx, "success to update feedback", nil, nil)

}
