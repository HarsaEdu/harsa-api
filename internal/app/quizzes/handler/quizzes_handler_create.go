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

func (quizzesHandler *QuizzesHandlereImpl) Create(ctx echo.Context) error {

	req := web.QuizRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind quiz request", err)
	}

	id := ctx.Get("user_id").(uint)

	idParam := ctx.Param("module-id")

	moduleId, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid module id", err)
	}

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	req.ModuleID = uint(moduleId)
	req.UserID = id
	err = quizzesHandler.QuizzesService.Create(req, roleString)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx, "you cannot create this quiz", err)
		}
		return res.StatusInternalServerError(ctx, "failed to create quiz, something happen", err)

	}

	return res.StatusCreated(ctx, "success to create quiz", nil, nil)
}
