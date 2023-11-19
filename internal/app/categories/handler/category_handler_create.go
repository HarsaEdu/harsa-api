package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (categoryHandler *CategoryHandlerImpl) Create(ctx echo.Context) error {

	req := web.CategoryCreateRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind category request", err)
	}

	err = categoryHandler.CategoryService.Create(req)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "already exist") {
			return res.StatusAlreadyExist(ctx, "category already exist", err)
		}
		return res.StatusInternalServerError(ctx, "failed to create category, something happen", err)

	}

	return res.StatusCreated(ctx, "success to create category", nil, nil)
}
