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

func (UserHandler *UserHandlerImpl) UserUpdate(ctx echo.Context) error {
	userUpdateRequest := web.UserUpdateRequest{}

	err := ctx.Bind(&userUpdateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request not valid", err)
	}

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid user id", err)
	}

	userUpdateRequest.ID = uint(id)

	err = UserHandler.UserService.UserUpdate(userUpdateRequest)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "user not found", err)
		}
		if strings.Contains(err.Error(), "already exists") {
			return res.StatusAlreadyExist(ctx, err.Error(), err)
		}
		return res.StatusInternalServerError(ctx, "failed to update user, something happen", fmt.Errorf("internal server error"))
	}

	return res.StatusOK(ctx, "success to update user", nil, nil)
}


func (UserHandler *UserHandlerImpl) UserUpdateMobile(ctx echo.Context) error {
	userUpdateRequest := web.UserUpdateRequestMobile{}

	err := ctx.Bind(&userUpdateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request not valid", err)
	}

	id := ctx.Get("user_id").(uint)

	userUpdateRequest.ID = id

	err = UserHandler.UserService.UserUpdateMobile(userUpdateRequest)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "user not found", err)
		}
		if strings.Contains(err.Error(), "already exists") {
			return res.StatusAlreadyExist(ctx, err.Error(), err)
		}
		return res.StatusInternalServerError(ctx, "failed to update user, something happen", fmt.Errorf("internal server error"))
	}

	return res.StatusOK(ctx, "success to update user", nil, nil)
}


func (UserHandler *UserHandlerImpl) UserUpdatePasswordMobile(ctx echo.Context) error {
	userUpdateRequest := web.UserUpdatePasswordRequestMobile{}

	err := ctx.Bind(&userUpdateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request not valid", err)
	}

	id := ctx.Get("user_id").(uint)

	userUpdateRequest.ID = id

	err = UserHandler.UserService.UserUpdatePasswordMobile(userUpdateRequest)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "user not found", err)
		}
		if strings.Contains(err.Error(), "invalid") {
			return res.StatusBadRequest(ctx, err.Error(), err)
		}
		return res.StatusInternalServerError(ctx, "failed to update user, something happen", fmt.Errorf("internal server error"))
	}

	return res.StatusOK(ctx, "success to update user", nil, nil)
}
