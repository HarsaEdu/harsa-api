package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (categoryHandler *CategoryHandlereImpl) Delete(ctx echo.Context) error {

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := categoryHandler.CategoryService.Delete(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "category not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to delete category, something happen", err)
	}

	return res.StatusOK(ctx, "success to delete category", nil)
}
