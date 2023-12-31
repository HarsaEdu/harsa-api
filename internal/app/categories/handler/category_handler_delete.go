package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (categoryHandler *CategoryHandlerImpl) Delete(ctx echo.Context) error {

	idParam := ctx.Param("id")
	id, err:= strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid category id", err)
	}
	err = categoryHandler.CategoryService.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "category not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to delete category, something happen", err)
	}

	return res.StatusOK(ctx, "success to delete category", nil, nil)
}
