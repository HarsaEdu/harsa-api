package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (profileHandler *ProfileHandlerImpl) UpdateProfile(ctx echo.Context) error {
	profileID, err := strconv.Atoi(ctx.Param("profile_id"))
	if err != nil {
		return res.StatusInternalServerError(ctx, "cannot convert profile id to int", err)
	}

	profile := web.UpdateProfileRequest{}
	if err = ctx.Bind(&profile); err != nil {
		return res.StatusBadRequest(ctx, "failed to bind profile model", err)
	}

	err = profileHandler.ProfileService.UpdateProfile(ctx, &profile, uint(profileID))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "profile not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to update profile, something happen", err)
	}
	return res.StatusOK(ctx, "success to update profile", nil, nil)
}

func (profileHandler *ProfileHandlerImpl) UpdateMyProfile(ctx echo.Context) error {
	userID := ctx.Get("user_id").(uint)

	profile := web.UpdateProfileRequest{}
	if err := ctx.Bind(&profile); err != nil {
		return res.StatusBadRequest(ctx, "failed to bind profile model", err)
	}

	err := profileHandler.ProfileService.UpdateProfile(ctx, &profile, userID)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "profile not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get all profiles, something happen", err)
	}
	return res.StatusOK(ctx, "success", nil, nil)
}
