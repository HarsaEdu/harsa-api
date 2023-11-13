package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (categoryHandler *CategoryHandlereImpl) FindById(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	result, err := categoryHandler.CategoryService.FindById(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "categoriy not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to create category, something happen", fmt.Errorf("internal server error"))
	}

	return res.StatusOK(ctx, "success to get category", result)
}