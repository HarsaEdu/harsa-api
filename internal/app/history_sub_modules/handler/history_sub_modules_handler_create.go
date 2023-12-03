package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (handler *HistorySubModuleHandlerImpl) CreateHistoryModule(ctx echo.Context) error {
	subModuleId, err := strconv.Atoi(ctx.Param("sub-module-id"))
	if err != nil {
		return res.StatusInternalServerError(ctx, "cannot convert sub module id", err)
	}
	request := web.CreateHistorySubModuleRequest{SubModuleID: uint(subModuleId)}
	err = handler.Service.CreateHistorySubModule(&request, ctx.Get("user_id").(uint))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "exists") {
			return res.StatusAlreadyExist(ctx, "sub module history already exist", err)
		}
		return res.StatusInternalServerError(ctx, "failed to create sub module history, something happen", err)
	}

	return res.StatusCreated(ctx, "success to create sub module history", nil, nil)
}
