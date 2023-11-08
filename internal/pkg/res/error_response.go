package res

import (
	"net/http"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func errorResponse(ctx echo.Context, code int, message string, errors any) error {
	return ctx.JSON(code, web.ErrorResponse{
		Code:    code,
		Message: message,
		Errors:    errors,
	})
}

func StatusNotFound(ctx echo.Context, message string, err error) error {
	return errorResponse(ctx, http.StatusNotFound, message, err.Error())
}

func StatusInternalServerError(ctx echo.Context, message string, err error) error {
	logrus.Error(err.Error())
	return errorResponse(ctx, http.StatusInternalServerError, message, err.Error())
}

func StatusBadRequest(ctx echo.Context, message string, err error) error {
	return errorResponse(ctx, http.StatusBadRequest, message, err.Error())
}

func StatusAlreadyExist(ctx echo.Context, message string, err error) error {
	return errorResponse(ctx, http.StatusConflict, message, err.Error())
}