package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (historyQuizHandler *HistoryQuizHandlereImpl) GetAll(ctx echo.Context) error {
	idParam := ctx.Param("quiz-id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid quiz id", err)
	}
	params := ctx.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return res.StatusBadRequest(ctx, "params limit not valid", err)
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return res.StatusBadRequest(ctx, "params offset not valid", err)
	}

	response, pagination, err := historyQuizHandler.HistoryQuizService.GetAllByQuizWeb(uint(id),offset, limit, params.Get("search"))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "history Quiz not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get all history Quiz, something happen", err)
	}

	return res.StatusOK(ctx, "success get history Quizzes", response, pagination)

}

