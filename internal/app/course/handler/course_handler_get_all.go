package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (courseHandler *CourseHandlerImpl) GetAllByUserId(ctx echo.Context) error {
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

	response, pagination, err := courseHandler.CourseService.GetAllByUserId(offset, limit, params.Get("search"), user_id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get all course, something happen", err)
	}

	return res.StatusOK(ctx, "success get courses", response, pagination)
}

func (courseHandler *CourseHandlerImpl) GetAll(ctx echo.Context) error {
	params := ctx.QueryParams()
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid offset", err)
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit", err)
	}

	response, pagination, err := courseHandler.CourseService.GetAll(offset, limit, params.Get("search"), params.Get("category"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get all course, something happen", err)
	}

	return res.StatusOK(ctx, "success get courses", response, pagination)
}

func (courseHandler *CourseHandlerImpl) GetAllMobile(ctx echo.Context) error {
	params := ctx.QueryParams()
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid offset", err)
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit", err)
	}

	response, pagination, err := courseHandler.CourseService.GetAllMobile(offset, limit, params.Get("search"), params.Get("category"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get all course, something happen", err)
	}

	return res.StatusOK(ctx, "success get courses", response, pagination)
}

func (courseHandler *CourseHandlerImpl) GetAllCourseByUserId(ctx echo.Context) error {
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

	response, pagination, err := courseHandler.CourseService.GetAllCourseByUserId(offset, limit, params.Get("search"), user_id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get all course, something happen", err)
	}

	return res.StatusOK(ctx, "success get courses", response, pagination)
}

func (courseHandler *CourseHandlerImpl) GetAllByCategory(ctx echo.Context) error {
	
	idParam := ctx.Param("id")
	
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

	response, pagination, err := courseHandler.CourseService.GetAllByCategory(offset, limit, params.Get("search"), uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get all course, something happen", err)
	}

	return res.StatusOK(ctx, "success get courses", response, pagination)
}

func (courseHandler *CourseHandlerImpl) GetAllMyCourse(ctx echo.Context) error {
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

	response, pagination, err := courseHandler.CourseService.GetAllMyCourse(offset, limit, params.Get("search"), params.Get("category"), user_id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get all course, something happen", err)
	}

	return res.StatusOK(ctx, "success get courses", response, pagination)
}

func (courseHandler *CourseHandlerImpl) GetAllByRating(ctx echo.Context) error {
	params := ctx.QueryParams()
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid offset", err)
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit", err)
	}

	response, pagination, err := courseHandler.CourseService.GetAllByRating(offset, limit, params.Get("search"), params.Get("category"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get all course, something happen", err)
	}

	return res.StatusOK(ctx, "success get courses", response, pagination)
}