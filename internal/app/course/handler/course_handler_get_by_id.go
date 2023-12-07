package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (courseHandler *CourseHandlerImpl) GetById(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to convert param id to int: ", err)
	}

	response, err := courseHandler.CourseService.GetById(uint(id))
	if err != nil {
		
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get course, something happen", err)
	}

	return res.StatusOK(ctx, "success to get course", response, nil)
}

func (courseHandler *CourseHandlerImpl) GetByIdMobile(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to convert param id to int: ", err)
	}

	response, err := courseHandler.CourseService.GetByIdMobile(uint(id))
	if err != nil {
		
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get course, something happen", err)
	}

	return res.StatusOK(ctx, "success to get course", response, nil)
}

func (courseHandler *CourseHandlerImpl) GetDetailCourseById(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to convert param id to int: ", err)
	}

	response, err := courseHandler.CourseService.GetDeatailCourse(uint(id))
	if err != nil {
		
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get course, something happen", err)
	}

	return res.StatusOK(ctx, "success to get course", response, nil)
}