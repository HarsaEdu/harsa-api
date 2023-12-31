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

func (userHandler *UserHandlerImpl) GetAllUsers(ctx echo.Context) error {
	params := ctx.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return res.StatusBadRequest(ctx, "params limit not valid", err)
	}

	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "params offset not valid", err)
	}

	roleId, err := strconv.Atoi(params.Get("roleID"))
	if err != nil {
		roleId = 0
	}
	response, pagination, err := userHandler.UserService.UserGetAll(offset, limit, params.Get("search"), roleId)
	if err != nil {
		if strings.Contains(err.Error(), "users not found") {
			return res.StatusNotFound(ctx, "users not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get users, something happen", err)
	}

	return res.StatusOK(ctx, "success to get users", response, pagination)
}

func (userHandler *UserHandlerImpl) GetAllStudentSubscribe(ctx echo.Context) error {
	params := ctx.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return res.StatusBadRequest(ctx, "params limit not valid", err)
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return res.StatusBadRequest(ctx, "params offset not valid", err)
	}

	idParam := ctx.Param("course-id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid user id", err)
	}

	response, pagination, err := userHandler.UserService.UserGetAllStudentSubscribe(offset, limit, params.Get("search"), uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "users not found") {
			return res.StatusNotFound(ctx, "users not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get users, something happen", err)
	}

	return res.StatusOK(ctx, "success to get users", response, pagination)
}

func (UserHandler *UserHandlerImpl) GetUserDetailByID(ctx echo.Context) error {
	userGetByIDRequest := web.UserGetByIDRequest{}

	err := ctx.Bind(&userGetByIDRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request not valid", err)
	}

	response, err := UserHandler.UserService.GetUserDetail(userGetByIDRequest)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "user not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get user, something happen", fmt.Errorf("internal server error"))
	}

	return res.StatusOK(ctx, "success to get user", response, nil)
}

func (UserHandler *UserHandlerImpl) GetUserAccountByID(ctx echo.Context) error {

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid user id", err)
	}

	response, err := UserHandler.UserService.GetUserAccount(uint(id))

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "user not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get user, something happen", fmt.Errorf("internal server error"))
	}

	return res.StatusOK(ctx, "success to get user", response, nil)
}


func (UserHandler *UserHandlerImpl) GetUserMyAccount(ctx echo.Context) error {

	id := ctx.Get("user_id").(uint)

	response, err := UserHandler.UserService.GetUserAccount(id)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "user not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get user, something happen", fmt.Errorf("internal server error"))
	}

	return res.StatusOK(ctx, "success to get user", response, nil)
}
