package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (categoryHandler *CategoryHandlereImpl) GetAll(ctx echo.Context) error {
	response, err := categoryHandler.CategoryService.GetAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "category not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get all category, something happen", fmt.Errorf("internal server error"))
	}

	return res.StatusOK(ctx, "success get categories", response)

}
