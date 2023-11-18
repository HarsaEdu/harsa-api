package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (moduleHandler ModuleHandlerImpl) Create(ctx echo.Context) error {
	paramCourseId := ctx.Param("courseId")
	courseId, err := strconv.Atoi(paramCourseId)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to convert course id : ", err)
	}

	moduleCreateRequest := web.ModuleCreateRequest{}
	err = ctx.Bind(&moduleCreateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request invalid", err)
	}

	err = moduleHandler.ModuleService.Create(ctx, &moduleCreateRequest, uint(courseId))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "already exist") {
			return res.StatusAlreadyExist(ctx, err.Error(), err)
		}
		return res.StatusInternalServerError(ctx, "failed to create module, something happen", err)

	}

	return res.StatusCreated(ctx, "success to create module", nil, nil)
}