package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
	"github.com/HarsaEdu/harsa-api/internal/pkg/jwt"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (authHandler *AuthHandlerImpl) RegisterUser(ctx echo.Context) error {
	registerUserRequest := web.RegisterUserRequest{}
	err := ctx.Bind(&registerUserRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request not valid", err)
	}

	response, err := authHandler.AuthService.RegisterUser(registerUserRequest)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "already exist") {
			return res.StatusAlreadyExist(ctx, "account already exist", err)
		}
		return res.StatusInternalServerError(ctx, "failed to register user, something happen", err)
	}

	tokenAccess, err := jwt.GenerateAccessToken(response)
	tokenRefresh, err := jwt.GenerateRefreshToken(response)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to register user, something happen", err)
	}

	loginResponse := conversion.AuthResponseToLoginResponse(response)
	loginResponse.AccessToken = tokenAccess
	loginResponse.RefreshToken = tokenRefresh

	return res.StatusCreated(ctx, "success to register user", loginResponse, nil)
}
