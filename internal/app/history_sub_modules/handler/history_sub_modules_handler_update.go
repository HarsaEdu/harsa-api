package handler

// import (
// 	"strconv"
// 	"strings"

// 	"github.com/HarsaEdu/harsa-api/internal/model/web"
// 	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
// 	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
// 	"github.com/labstack/echo/v4"
// )

// func (handler *HistorySubModuleHandlerImpl) UpdateHistorySubModule(ctx echo.Context) error {
// 	id, _ := strconv.Atoi(ctx.Param("sub_module_id"))
// 	request := web.UpdateHistorySubModuleRequest{ID: uint(id)}
// 	err := handler.Service.UpdateHistorySubModule(&request)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "validation") {
// 			return validation.ValidationError(ctx, err)
// 		}
// 		if strings.Contains(err.Error(), "not found") {
// 			return res.StatusNotFound(ctx, "sub module history not found", err)
// 		}
// 		return res.StatusInternalServerError(ctx, "failed to update sub module history, something happen", err)
// 	}
// 	return res.StatusOK(ctx, "success to update sub module history", nil, nil)
// }
