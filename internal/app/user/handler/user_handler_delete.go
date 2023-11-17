package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (UserHandler *UserHandlerImpl) UserDelete(ctx echo.Context) error {
	userDeleteRequest := web.UserDeleteRequest{}

	err := ctx.Bind(&userDeleteRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request not valid", err)
	}

	err = UserHandler.UserService.UserDelete(userDeleteRequest)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "user not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to delete user, something happen", fmt.Errorf("internal server error"))
	}

	return res.StatusOK(ctx, "success to delete user", nil, nil)
}
