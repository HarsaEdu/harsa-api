package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (categoryHandler *CategoryHandlerImpl) UploadImage(ctx echo.Context) error {
	var icon, image bool
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return res.StatusBadRequest(ctx, "invalid category id", err)
	}
	req := web.CategoryUploadImageRequest{}
	err = ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind category request", err)
	}

	fileHeader, _ := ctx.FormFile("icon")
	if fileHeader != nil {
		icon = true
	}

	fileHeader, _ = ctx.FormFile("file")
	if fileHeader != nil {
		image = true
	}

	err = categoryHandler.CategoryService.UploadImage(ctx, &req, id, icon, image)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "file not found") {
			return res.StatusBadRequest(ctx, "file not found", err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "category not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to upload image category, something happen", err)

	}

	return res.StatusOK(ctx, "success to upload image", nil, nil)
}
