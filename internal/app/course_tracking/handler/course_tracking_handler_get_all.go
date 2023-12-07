package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (courseTrackingHandler *CourseTrackingHandlerImpl) GetAllTracking(ctx echo.Context) error {
	
	params := ctx.QueryParams()
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid offset", err)
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit", err)
	}
	user_id := ctx.Get("user_id").(uint)

	response, pagination, err := courseTrackingHandler.CourseTrackingService.GetAllCourseByUserIdMobile(offset, limit, params.Get("search"), user_id, params.Get("status"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get all course, something happen", err)
	}

	return res.StatusOK(ctx, "success get courses", response, pagination)
}

func (courseTrackingHandler *CourseTrackingHandlerImpl) GetAllTrackingWeb(ctx echo.Context) error {
	idParam := ctx.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid user id", err)
	}
	
	params := ctx.QueryParams()
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid offset", err)
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit", err)
	}

	response, pagination, err := courseTrackingHandler.CourseTrackingService.GetAllCourseByUserIdWeb(offset, limit, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get all course, something happen", err)
	}

	return res.StatusOK(ctx, "success get courses", response, pagination)
}

func (courseTrackingHandler *CourseTrackingHandlerImpl) GetAllTrackingUserWeb(ctx echo.Context) error {
	idParam := ctx.Param("course-id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid course id", err)
	}
	
	params := ctx.QueryParams()
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid offset", err)
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit", err)
	}

	response, pagination, err := courseTrackingHandler.CourseTrackingService.GetAllUserCourseWeb(offset,limit,uint(id),params.Get("search"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "user not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get all user, something happen", err)
	}

	return res.StatusOK(ctx, "success get user", response, pagination)
}