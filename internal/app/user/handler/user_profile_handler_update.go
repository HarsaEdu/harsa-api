package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (UserHandler *UserHandlerImpl) UserProfileUpdate(ctx echo.Context) error {
	userProfileUpdateRequest := web.UserProfileUpdateRequest{}

	err := ctx.Bind(&userProfileUpdateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request not valid", err)
	}

	err = UserHandler.UserService.UserProfileUpdate(ctx, userProfileUpdateRequest)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusAlreadyExist(ctx, "user not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to update user profile, something happen", fmt.Errorf("internal server error"))
	}

	return res.StatusOK(ctx, "success to update user profile", nil)
}
