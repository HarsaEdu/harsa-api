package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (courseHandler *CourseHandlerImpl) GetAll(ctx echo.Context) error {
	response, err := courseHandler.CourseService.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get all course, something happen", err)
	}

	return res.StatusOK(ctx, "success get courses", response)
}