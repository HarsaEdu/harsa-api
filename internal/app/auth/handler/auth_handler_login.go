package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
	"github.com/HarsaEdu/harsa-api/internal/pkg/jwt"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (authHandler *AuthHandlerImpl) LoginUser(ctx echo.Context) error {
	loginUserRequest := web.LoginUserRequest{}
	err := ctx.Bind(&loginUserRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request not valid", err)
	}

	response, err := authHandler.AuthService.LoginUser(ctx, loginUserRequest)

	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			return res.StatusBadRequest(ctx, err.Error(), err)
		}
		return res.StatusInternalServerError(ctx, "failed to login, something happen", err)
	}

	token, err := jwt.GenerateToken(response)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to login, something happen", err)
	}

	loginResponse := conversion.AuthResponseToLoginResponse(response)
	loginResponse.Token = token

	return res.StatusOK(ctx, "success to login", loginResponse, nil)
}
