package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (faqsHandler *FaqsHandlerImpl) Create(ctx echo.Context) error {
	req := web.FaqsRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind faqs request", err)
	}

	err = faqsHandler.FaqsService.Create(req)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		return res.StatusInternalServerError(ctx, "failed to create faqs, something happen", err)
	}
	return res.StatusCreated(ctx, "success to create faqs", nil, nil)

}
