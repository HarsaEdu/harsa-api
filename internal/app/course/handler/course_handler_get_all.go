package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (courseHandler *CourseHandlerImpl) GetAll(ctx echo.Context) error {
	param := ctx.QueryParams()
	offset, err := strconv.Atoi(param.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid offset", err)
	}

	limit, err := strconv.Atoi(param.Get("limit"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit", err)
	}

	response, pagination, err := courseHandler.CourseService.GetAll(offset, limit, param.Get("s"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get all course, something happen", err)
	}

	return res.StatusOK(ctx, "success get courses", response, pagination)
}