package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/labstack/echo/v4"
)

func (userService *UserServiceImpl) UserDelete(ctx echo.Context, userRequest web.UserDeleteRequest) error {
	// validate the request
	err := userService.Validate.Struct(userRequest)

	// check errors when validate the request
	if err != nil {
		return err
	}

	// check user
	findUser, _ := userService.UserRepository.UserAvailableByID(userRequest.ID)
	if findUser == nil {
		return fmt.Errorf("user not found")
	}

	// update data
	err = userService.UserRepository.UserDelete(userRequest.ID)

	// check if error when update data
	if err != nil {
		return fmt.Errorf("error when delete user %s:", err.Error())
	}

	return nil
}
