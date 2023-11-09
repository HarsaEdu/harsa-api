package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/jwt"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (uc *AuthHanlderImpl) RegisterUser(ctx echo.Context) error {
	registerUserRequest := web.RegisterUserRequest{}
	err := ctx.Bind(&registerUserRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request not valid", err)
	}

	response, err := uc.AuthService.RegisterUser(ctx, registerUserRequest)

	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return res.StatusBadRequest(ctx, "validation failed", err)
		}
		if strings.Contains(err.Error(), "already exists") {
			return res.StatusAlreadyExist(ctx, "account already exists", err)
		}
		return res.StatusInternalServerError(ctx, "Failed to register user, something happen", fmt.Errorf("Internal Server Error"))
	}

	token, err := jwt.GenerateToken(response)
	if err != nil {
		return res.StatusInternalServerError(ctx, "Failed to register user, something happen", fmt.Errorf("Internal Server Error"))
	}

	loginResponse := &web.UserLoginResponse{
		ID:       response.ID,
		Username: response.Username,
		RoleName: response.RoleName,
		Token:    token,
	}

	return res.StatusOK(ctx, "Success to created user", loginResponse)
}
