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

func (categoryHandler *CategoryHandlereImpl) UploadImage(ctx echo.Context) error {
	fmt.Println("handler error")
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	req := web.CategoryUploadImageRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind category request", err)
	}

	err = categoryHandler.CategoryService.UploadImage(ctx, &req, id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "categoriy not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to upload image category, something happen", fmt.Errorf("internal server error"))

	}

	return res.StatusCreated(ctx, "success to upload image", nil)
}