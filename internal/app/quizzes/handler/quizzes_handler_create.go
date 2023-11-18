package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (quizzesHandler *QuizzesHandlereImpl) Create(ctx echo.Context) error {

	req := web.QuizRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind quiz request", err)
	}

	id := ctx.Get("user_id").(uint)

	req.UserId = id
	err = quizzesHandler.QuizzesService.Create(ctx, req)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		return res.StatusInternalServerError(ctx, "failed to create quiz, something happen", err)

	}

	return res.StatusCreated(ctx, "success to create quiz", nil, nil)
}
