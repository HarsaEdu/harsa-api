package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (handler *HistorySubModuleHandlerImpl) GetHistoryModule(ctx echo.Context) error {
	request := web.GetHistorySubModuleRequest{UserID: ctx.Get("user_id").(uint)}
	result, err := handler.Service.GetHistorySubModuleByUserID(&request)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "sub module history not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get sub module history, something happen", err)
	}
	return res.StatusOK(ctx, "success to get sub module history", result, nil)
}
