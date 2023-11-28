package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (historyModuleHandler *HistoryModuleHandlerImpl) Update(ctx echo.Context) error {

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	req := web.HistoryModuleUpdateRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind history module request", err)
	}

	err = historyModuleHandler.HistoryModuleService.Update(req, id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "history module not found", err)
		}

		if strings.Contains(err.Error(), "already exist") {
			return res.StatusAlreadyExist(ctx, "history module name already exist", err)
		}

		return res.StatusInternalServerError(ctx, "failed to update history module, something happen", err)

	}

	return res.StatusOK(ctx, "success to update history module", nil, nil)

}
