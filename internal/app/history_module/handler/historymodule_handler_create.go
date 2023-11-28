package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (historyModuleHandler *HistoryModuleHandlerImpl) Create(ctx echo.Context) error {
	req := web.HistoryModuleCreateRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind history module request", err)
	}

	err = historyModuleHandler.HistoryModuleService.Create(req)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		return res.StatusInternalServerError(ctx, "failed to create history module, something happen", err)
	}
	return res.StatusCreated(ctx, "success to create history module", nil, nil)

}
