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

func (categoryHandler *CategoryHandlereImpl) Update(ctx echo.Context) error {

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	req := web.CategoryUpdateRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind category request", err)
	}

	err = categoryHandler.CategoryService.Update(ctx, req, id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "category not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to delete category, something happen", fmt.Errorf("internal server error"))

	}

	return res.StatusCreated(ctx, "success to update category", nil)

}