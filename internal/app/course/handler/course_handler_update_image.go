package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (courseHandler *CourseHandlerImpl) UpdateImage(ctx echo.Context) error {
	idParam := ctx.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid course id", err)
	}

	user_id := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	request := web.CourseUpdateImageRequest{}
	err = ctx.Bind(&request)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind course request", err)
	}

	err = courseHandler.CourseService.UpdateImage(ctx, uint(id),uint(user_id), roleString, &request)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}

		if strings.Contains(err.Error(), "file not found") {
			return res.StatusBadRequest(ctx, "file not found", err)
		}

		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot update this course" ,err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "course not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to upload image coure, something happen", err)

	}

	return res.StatusOK(ctx, "success to update image course", nil, nil)
}