package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/HarsaEdu/harsa-api/internal/pkg/password"
	"github.com/labstack/echo/v4"
)

func (userService *UserServiceImpl) UserUpdate(ctx echo.Context, userRequest web.UserUpdateRequest) error {
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

	// check available username and email
	existingUser, _ := userService.UserRepository.UserAvailable(userRequest.Username, userRequest.Email)
	if existingUser != nil {
		return fmt.Errorf("already exists")
	}

	// convert request to user model
	userAccount := conversionRequest.UserUpdateRequestToUserModel(userRequest)

	// hash password
	if userAccount.Password != "" {
		userAccount.Password = password.HashPassword(userAccount.Password)
	}

	// update data
	err = userService.UserRepository.UserUpdate(userAccount)

	// check if error when update data
	if err != nil {
		return fmt.Errorf("error when update user %s:", err.Error())
	}

	return nil
}
