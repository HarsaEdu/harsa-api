package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (historyQuizHandler *HistoryQuizHandlereImpl) Create(ctx echo.Context) error {
	
	idParam := ctx.Param("quiz-id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid quiz id", err)
	}

	var newAnswers []web.AnswersRequest
	err = ctx.Bind(&newAnswers)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind history quiz request", err)
	}

	fmt.Println(newAnswers[0].Option_id)

	userId := ctx.Get("user_id").(uint)

	err = historyQuizHandler.HistoryQuizService.Create(newAnswers, userId, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "You have completed the quiz") {
			return res.StatusBadRequest(ctx, "You have completed the quiz", err)
		}
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		return res.StatusInternalServerError(ctx, "failed to create history quiz, something happen", err)

	}

	return res.StatusCreated(ctx, "success to create history quiz", nil, nil)
}