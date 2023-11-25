package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
	"github.com/HarsaEdu/harsa-api/internal/pkg/jwt"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (authHandler *AuthHandlerImpl) GetAccessToken(ctx echo.Context) error {
	accessTokenRequest := web.AccessTokenRequest{}
	err := ctx.Bind(&accessTokenRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request not valid", err)
	}

	userData, err := authHandler.AuthService.GetAccessToken(accessTokenRequest)

	if err != nil {
		return res.StatusBadRequest(ctx, "failed to get access token!", err)
	}

	loginResponse := conversion.AuthResponseToLoginResponse(userData)

	tokenAccess, err := jwt.GenerateAccessToken(userData)
	tokenRefresh, err := jwt.GenerateRefreshToken(userData)

	loginResponse.AccessToken = tokenAccess
	loginResponse.RefreshToken = tokenRefresh

	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to get access token, something happen", err)
	}

	return res.StatusOK(ctx, "success to get access token", loginResponse, nil)
}
