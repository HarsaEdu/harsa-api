package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (handler *ProfileHandlerImpl) MyProfile(ctx echo.Context) error {
	request := web.UserGetByIDRequest{}
	request.ID = ctx.Get("user_id").(uint)

	result, err := handler.ProfileService.MyProfile(&request)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "profile not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get profile, something happen", fmt.Errorf("internal server error"))
	}

	return res.StatusOK(ctx, "success to get profile", result, nil)
}
