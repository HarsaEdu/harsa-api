package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (categoryHandler *CategoryHandlerImpl) Update(ctx echo.Context) error {
	image := false

	fileHeader, _ := ctx.FormFile("image")
	if fileHeader != nil {
		image = true
	}

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid category id", err)
	}
	req := web.CategoryUpdateRequest{}
	err = ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind category request", err)
	}

	err = categoryHandler.CategoryService.Update(ctx, req, id, image)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "category not found", err)
		}

		if strings.Contains(err.Error(), "already exist") {
			return res.StatusAlreadyExist(ctx, "category name already exist", err)
		}

		return res.StatusInternalServerError(ctx, "failed to update category, something happen", err)

	}

	return res.StatusOK(ctx, "success to update category", nil, nil)

}
