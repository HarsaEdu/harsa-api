package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (courseHandler *CourseHandlerImpl) Delete(ctx echo.Context) error {

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := courseHandler.CourseService.Delete(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to delete course, something happen", err)
	}

	return res.StatusOK(ctx, "success to delete course", nil)
}
