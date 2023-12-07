package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

// func (courseTrackingHandler *CourseTrackingHandlerImpl) GetById(ctx echo.Context) error {
// 	idParam := ctx.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		return res.StatusInternalServerError(ctx, "failed to convert param id to int: ", err)
// 	}

// 	response, err := courseTrackingHandler.CourseTrackingService.FindByIdMobile(uint(id))
// 	if err != nil {

// 		if strings.Contains(err.Error(), "not found") {
// 			return res.StatusNotFound(ctx, "course not found", err)
// 		}

// 		return res.StatusInternalServerError(ctx, "failed to get course tracking, something happen", err)
// 	}

// 	return res.StatusOK(ctx, "success to get course tracking", response, nil)
// }

func (courseTrackingHandler *CourseTrackingHandlerImpl) GetByUserIdAndCourseID(ctx echo.Context) error {
	idParam := ctx.Param("course-id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to convert param id to int: ", err)
	}

	userId := ctx.Get("user_id").(uint)

	response, err := courseTrackingHandler.CourseTrackingService.FindByIdMobileByUserIdAndCourseId(ctx,userId,uint(id))
	if err != nil {
		
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get course tracking, something happen", err)
	}

	return res.StatusOK(ctx, "success to get course tracking", response, nil)
}