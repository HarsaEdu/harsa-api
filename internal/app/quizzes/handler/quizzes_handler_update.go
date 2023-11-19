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

func (quizzesHandler *QuizzesHandlereImpl) Update(ctx echo.Context) error {

	req := web.QuizRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind quiz request", err)
	}

	idParam := ctx.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid quiz id", err)
	}

	user_id := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)
	
	req.UserId = user_id
	err = quizzesHandler.QuizzesService.Update(req, uint(id), roleString)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot update this quiz" ,err)
		}
		return res.StatusInternalServerError(ctx, "failed to update quiz, something happen", err)

	}

	return res.StatusOK(ctx, "success to update quiz", nil, nil)
}