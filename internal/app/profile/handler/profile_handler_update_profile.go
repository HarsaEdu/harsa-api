package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (profileHandler *ProfileHandlerImpl) UpdateProfile(ctx echo.Context) error {
	profile := domain.UserProfile{}
	if err := ctx.Bind(&profile); err != nil {
		return res.StatusBadRequest(ctx, "failed to bind profile model", err)
	}

	profileID, _ := strconv.Atoi(ctx.Param("profile_id"))
	err := profileHandler.ProfileService.UpdateProfile(ctx, &profile, uint(profileID))
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
