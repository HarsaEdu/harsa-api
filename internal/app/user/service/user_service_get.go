package service

import (
	"fmt"
	"strconv"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/labstack/echo/v4"
)

func (userService *UserServiceImpl) UserGetAll(ctx echo.Context) ([]domain.UserEntity, int64, error) {
	params := ctx.QueryParams()
	search := params.Get("s")
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return nil, 0, fmt.Errorf("params limit not valid")
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return nil, 0, fmt.Errorf("params offset not valid")
	}

	users, total, err := userService.UserRepository.UserGetAll(offset, limit, search)
	if len(search) > 0 && total <= 0 {
		return nil, total, fmt.Errorf("users not found")
	}
	if err != nil {
		return nil, total, fmt.Errorf("internal Server Error")
	}
	return users, total, nil
}

func (userService *UserServiceImpl) GetUserDetail(ctx echo.Context, userRequest web.UserGetByIDRequest) (*domain.UserDetail, error) {
	err := userService.Validate.Struct(userRequest)

	if err != nil {
		return nil, err
	}

	users, err := userService.UserRepository.GetUserByID(userRequest.ID)

	if err != nil {
		return nil, err
	}

	return users, nil
}
