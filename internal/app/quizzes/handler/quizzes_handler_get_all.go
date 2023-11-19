package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (quizzesHandler *QuizzesHandlereImpl) GetAll(ctx echo.Context) error {
	params := ctx.QueryParams()
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid offset", err)
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit", err)
	}

	idParam := ctx.Param("id")
	
	moduleId, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid module id", err)
	}

	response, pagination, err := quizzesHandler.QuizzesService.GetAll(uint(moduleId), offset, limit, params.Get("search"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "quiz not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get all quiz, something happen", err)
	}

	return res.StatusOK(ctx, "success get quizzes", response, pagination)
}