package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (courseHandler *CourseHandlerImpl) UpdateImage(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	request := web.CourseUpdateImageRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind course request", err)
	}

	err = courseHandler.CourseService.UpdateImage(ctx, uint(id), &request)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}

		if strings.Contains(err.Error(), "file not found") {
			return res.StatusBadRequest(ctx, "file not found", err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to upload image coure, something happen", err)

	}

	return res.StatusOK(ctx, "success to update image course", nil, nil)
}