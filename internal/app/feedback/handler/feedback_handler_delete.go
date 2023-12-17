package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (feedbackHandler *FeedbackHandlerImpl) DeleteByUserAndCourseId(ctx echo.Context) error {
	courseIdParam := ctx.Param("courseId")
	courseId, _ := strconv.Atoi(courseIdParam)

	userId := ctx.Get("user_id").(uint)

	err := feedbackHandler.FeedbackService.DeleteByUserAndCourseId(userId, uint(courseId))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "feedback not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to delete feedback, something happen", err)
	}

	return res.StatusOK(ctx, "success to delete feedback", nil, nil)
}

func (feedbackHandler *FeedbackHandlerImpl) DeleteById(ctx echo.Context) error {
	IdParam := ctx.Param("id")
	Id, _ := strconv.Atoi(IdParam)

	err := feedbackHandler.FeedbackService.DeleteById(uint(Id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "feedback not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to delete feedback, something happen", err)
	}

	return res.StatusOK(ctx, "success to delete feedback", nil, nil)
}
