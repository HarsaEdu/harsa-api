package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
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
		if strings.Contains(err.Error(), "incorrect") {
			return res.StatusNotFound(ctx, "failed to login", err)
		}
		return res.StatusInternalServerError(ctx, "failed to login, something happen", fmt.Errorf("internal server error"))
	}

	token, err := jwt.GenerateToken(response)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to login, something happen", fmt.Errorf("internal server error"))
	}

	loginResponse := &web.UserLoginResponse{
		ID:       response.ID,
		Username: response.Username,
		RoleName: response.RoleName,
		Token:    token,
	}

	return res.StatusOK(ctx, "success to login", loginResponse, nil)
}
