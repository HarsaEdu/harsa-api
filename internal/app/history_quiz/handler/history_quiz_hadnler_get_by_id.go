package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (historyQuizHandler *HistoryQuizHandlereImpl) FindById(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid history quiz id", err)
	}

	result, err := historyQuizHandler.HistoryQuizService.GetById(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "history quiz not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get history quiz, something happen", err)
	}

	return res.StatusOK(ctx, "success to get history quiz", result, nil)
}