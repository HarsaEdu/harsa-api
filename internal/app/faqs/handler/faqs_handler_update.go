package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (faqsHandler *FaqsHandlerImpl) Update(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	req := web.FaqsUpdateRequest{}
	err = ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind faqs update request", err)
	}

	err = faqsHandler.FaqsService.Update(req, id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "faqs not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed update faqs, something happen", err)

	}
	return res.StatusOK(ctx, "success to update faqs", nil, nil)

}
