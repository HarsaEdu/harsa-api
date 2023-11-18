package res

import (
	"net/http"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/labstack/echo/v4"
)

func successResponse(ctx echo.Context, code int, message string, data any, pagination any) error {
	return ctx.JSON(code, web.SuccessResponse{
		Code:       code,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	})
}

func StatusCreated(ctx echo.Context, message string, data any, pagination any) error {
	return successResponse(ctx, http.StatusCreated, message, data, pagination)
}

func StatusOK(ctx echo.Context, message string, data any, pagination any) error {
	return successResponse(ctx, http.StatusOK, message, data, pagination)
}
