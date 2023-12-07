package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (UserHandler *UserHandlerImpl) UserUpdate(ctx echo.Context) error {
	userUpdateRequest := web.UserUpdateRequest{}

	err := ctx.Bind(&userUpdateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request not valid", err)
	}

	err = UserHandler.UserService.UserUpdate(userUpdateRequest)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "user not found", err)
		}
		if strings.Contains(err.Error(), "already exists") {
			return res.StatusAlreadyExist(ctx, "account already exists", err)
		}
		return res.StatusInternalServerError(ctx, "failed to update user, something happen", fmt.Errorf("internal server error"))
	}

	return res.StatusOK(ctx, "success to update user", nil, nil)
}
