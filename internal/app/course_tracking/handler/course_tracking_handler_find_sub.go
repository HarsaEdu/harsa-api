package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (courseTrackingHandler *CourseTrackingHandlerImpl) FindModuleHistory(ctx echo.Context) error {
	id := ctx.Get("user_id").(uint)

	idParam := ctx.Param("module-id")
	moduleId, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid module id", err)
	}


	result, err := courseTrackingHandler.CourseTrackingService.FindModuleHistory(ctx, uint(moduleId), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "module not found", err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"subscribe to enrolled this course" ,err)
		}
		if strings.Contains(err.Error(), "not enrolled") {
			return res.StatusUnauthorized(ctx,"you not enrolled this course" ,err)
		}

		return res.StatusInternalServerError(ctx, "failed to get module tracking, something happen", err)
	}
	return res.StatusOK(ctx, "success to get module tracking", result, nil)
}

func (courseTrackingHandler *CourseTrackingHandlerImpl) FindSubModuleByID(ctx echo.Context) error {
	id := ctx.Get("user_id").(uint)

	idParam := ctx.Param("module-id")
	moduleId, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid module id", err)
	}


	subModuleID, err := strconv.Atoi(ctx.Param("sub-module-id"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid sub-module id", err)
	}

	result, err := courseTrackingHandler.CourseTrackingService.FindSubModuleByID(ctx, uint(moduleId), uint(subModuleID), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "sub module not found", err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"subscribe to enrolled this course" ,err)
		}
		if strings.Contains(err.Error(), "not enrolled") {
			return res.StatusUnauthorized(ctx,"you not enrolled this course" ,err)
		}

		return res.StatusInternalServerError(ctx, "failed to get course tracking, something happen", err)
	}
	return res.StatusOK(ctx, "success to get sub module tracking", result, nil)
}

func (courseTrackingHandler *CourseTrackingHandlerImpl) FindSubmissionByID(ctx echo.Context) error {
	id := ctx.Get("user_id").(uint)

	idParam := ctx.Param("module-id")
	moduleID, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid module id", err)
	}


	submissionID, err := strconv.Atoi(ctx.Param("submission-id"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid submmison id", err)
	}

	result, err := courseTrackingHandler.CourseTrackingService.FindSubmissionByID(ctx, uint(moduleID), id, uint(submissionID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "submission tracking not found", err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"subscribe to enrolled this course" ,err)
		}
		if strings.Contains(err.Error(), "not enrolled") {
			return res.StatusUnauthorized(ctx,"you not enrolled this course" ,err)
		}

		return res.StatusInternalServerError(ctx, "failed to get submission tracking, something happen", err)
	}
	return res.StatusOK(ctx, "success to get submission answer tracking", result, nil)
}

func (courseTrackingHandler *CourseTrackingHandlerImpl) FindQuizzByID(ctx echo.Context) error {
	id := ctx.Get("user_id").(uint)

	idParam := ctx.Param("module-id")
	moduleID, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid module id", err)
	}

	quizzID, err := strconv.Atoi(ctx.Param("quizz-id"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid quiz id", err)
	}

	result, err := courseTrackingHandler.CourseTrackingService.FindQuizzByID(ctx, uint(moduleID), id, uint(quizzID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "quizz tracking not found", err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"subscribe to enrolled this course" ,err)
		}
		if strings.Contains(err.Error(), "not enrolled") {
			return res.StatusUnauthorized(ctx,"you not enrolled this course" ,err)
		}

		return res.StatusInternalServerError(ctx, "failed to get quizz tracking, something happen", err)
	}
	return res.StatusOK(ctx, "success to get quizz tracking", result, nil)
}
