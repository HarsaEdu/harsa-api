package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (UserHandler *UserHandlerImpl) UserCreate(ctx echo.Context) error {
	userCreateRequest := web.UserCreateRequest{}

	err := ctx.Bind(&userCreateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request not valid", err)
	}

	err = UserHandler.UserService.UserCreate(ctx, userCreateRequest)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "already exists") {
			return res.StatusAlreadyExist(ctx, "account already exists", err)
		}
		return res.StatusInternalServerError(ctx, "failed to create user, something happen", fmt.Errorf("internal server error"))
	}

	return res.StatusCreated(ctx, "success to create user", nil, nil)
}
