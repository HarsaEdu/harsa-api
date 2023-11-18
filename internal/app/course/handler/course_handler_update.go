package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (courseHandler *CourseHandlerImpl) Update(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to convert param id to int: ", err)
	}

	request := web.CourseUpdateRequest{}
	err = ctx.Bind(&request)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind request: ", err)
	}

	err = courseHandler.CourseService.Update(uint(id), &request)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to update course, something happen", err)
	}

	return res.StatusOK(ctx, "success to update course", nil, nil)
}