package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (courseHandler *CourseHandlerImpl) Create(ctx echo.Context) error {
	courseCreateRequest := web.CourseCreateRequest{}
	err := ctx.Bind(&courseCreateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind course request", err)
	}

	var instructorId uint

	fmt.Println(courseCreateRequest.UserId)

	if courseCreateRequest.UserId == 0 {
		instructorId = ctx.Get("user_id").(uint)
	} else {
		instructorId = courseCreateRequest.UserId
	}

	err = courseHandler.CourseService.Create(ctx, &courseCreateRequest, instructorId)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "already exist") {
			return res.StatusAlreadyExist(ctx, "course already exist", err)
		}
		return res.StatusInternalServerError(ctx, "failed to create course, something happen", err)
	}

	return res.StatusCreated(ctx, "success to create course", nil, nil)
}
