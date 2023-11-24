package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (questionHandler *QuestionsHandlereImpl) Delete(ctx echo.Context) error {

	idParam := ctx.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid question id", err)
	}

	user_id := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)
	
	err = questionHandler.QuestionsService.Delete(user_id, uint(id), roleString)

	if err != nil {
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot delete this question" ,err)
		}
		return res.StatusInternalServerError(ctx, "failed to delete question, something happen", err)

	}

	return res.StatusOK(ctx, "success to delete question", nil, nil)
}

