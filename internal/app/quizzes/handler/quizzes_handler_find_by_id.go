package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (quizzesHandler *QuizzesHandlereImpl) FindById(ctx echo.Context) error {

	idParam := ctx.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid quiz id", err)
	}

	result, err := quizzesHandler.QuizzesService.FindById(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "quiz not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get quiz by id, something happen", err)
	}

	return res.StatusOK(ctx, "success to get quiz by id", result, nil)
}