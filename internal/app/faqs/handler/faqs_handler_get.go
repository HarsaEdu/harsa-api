package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (faqsHandler *FaqsHandlerImpl) GetAll(ctx echo.Context) error {
	params := ctx.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return res.StatusBadRequest(ctx, "params limit not valid", err)
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return res.StatusBadRequest(ctx, "params offset not valid", err)
	}

	response, pagiantion, err := faqsHandler.FaqsService.GetAll(offset, limit, params.Get("search"))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "faqs not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get all faqs, something happen", err)
	}

	return res.StatusOK(ctx, "success to get faqs", response, pagiantion)
}

func (faqsHandler *FaqsHandlerImpl) GetById(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	response, err := faqsHandler.FaqsService.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "faqs not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get faqs, something happen", err)
	}

	return res.StatusOK(ctx, "success to get faqs", response, err)
}
