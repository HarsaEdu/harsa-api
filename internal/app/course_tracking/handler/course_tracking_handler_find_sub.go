package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (courseTrackingHandler *CourseTrackingHandlerImpl) FindModuleHistory(ctx echo.Context) error {
	moduleID, _ := strconv.Atoi(ctx.Param("module-id"))
	id := ctx.Get("user_id").(uint)

	result, err := courseTrackingHandler.CourseTrackingService.FindModuleHistory(uint(moduleID), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "module not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get module tracking, something happen", err)
	}
	return res.StatusOK(ctx, "success to get module tracking", result, nil)
}

func (courseTrackingHandler *CourseTrackingHandlerImpl) FindSubModuleByID(ctx echo.Context) error {
	moduleID, _ := strconv.Atoi(ctx.Param("module-id"))
	subModuleID, _ := strconv.Atoi(ctx.Param("sub-module-id"))
	id := ctx.Get("user_id").(uint)

	result, err := courseTrackingHandler.CourseTrackingService.FindSubModuleByID(uint(moduleID), uint(subModuleID), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "sub module not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get course tracking, something happen", err)
	}
	return res.StatusOK(ctx, "success to get sub module tracking", result, nil)
}

func (courseTrackingHandler *CourseTrackingHandlerImpl) FindSubmissionByID(ctx echo.Context) error {
	moduleID, _ := strconv.Atoi(ctx.Param("module-id"))
	submissionID, _ := strconv.Atoi(ctx.Param("submission-id"))
	id := ctx.Get("user_id").(uint)

	result, err := courseTrackingHandler.CourseTrackingService.FindSubmissionByID(uint(moduleID), id, uint(submissionID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "submission tracking not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get submission tracking, something happen", err)
	}
	return res.StatusOK(ctx, "success to get submission answer tracking", result, nil)
}

func (courseTrackingHandler *CourseTrackingHandlerImpl) FindQuizzByID(ctx echo.Context) error {
	moduleID, _ := strconv.Atoi(ctx.Param("module-id"))
	quizzID, _ := strconv.Atoi(ctx.Param("quizz-id"))
	id := ctx.Get("user_id").(uint)

	result, err := courseTrackingHandler.CourseTrackingService.FindQuizzByID(uint(moduleID), id, uint(quizzID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "quizz tracking not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get quizz tracking, something happen", err)
	}
	return res.StatusOK(ctx, "success to get quizz tracking", result, nil)
}
