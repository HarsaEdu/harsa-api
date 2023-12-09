package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/HarsaEdu/harsa-api/internal/pkg/password"
)

func (userService *UserServiceImpl) UserUpdate(userRequest web.UserUpdateRequest) error {
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

	if findUser.Username != userRequest.Username && userRequest.Username != "" {
		fmt.Println("cek user name")
		existingUser, _ := userService.UserRepository.UserAvailableUsername(userRequest.Username)
		if existingUser != nil {
			return fmt.Errorf("username already exists")
		}
	}

	if findUser.Email != userRequest.Email && userRequest.Email != "" {
		fmt.Println("cek email")
		existingUser, _ := userService.UserRepository.UserAvailableEmail(userRequest.Email)
		if existingUser != nil {
			return fmt.Errorf("email already exists")
		}
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

func (userService *UserServiceImpl) UserUpdateMobile(userRequest web.UserUpdateRequestMobile) error {
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

	if findUser.Username != userRequest.Username && userRequest.Username != "" {
		existingUser, _ := userService.UserRepository.UserAvailableUsername(userRequest.Username)
		if existingUser != nil {
			return fmt.Errorf("username already exists")
		}
	}

	if findUser.Email != userRequest.Email && userRequest.Email != "" {
		existingUser, _ := userService.UserRepository.UserAvailableEmail(userRequest.Email)
		if existingUser != nil {
			return fmt.Errorf("email already exists")
		}
	}

	// convert request to user model
	userAccount := conversionRequest.UserUpdateRequestToUserModelMobile(userRequest)
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
