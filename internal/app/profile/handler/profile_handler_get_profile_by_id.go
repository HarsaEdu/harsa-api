package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (profileHandler *ProfileHandlerImpl) GetProfileByID(ctx echo.Context) error {
	profileID, err := strconv.Atoi(ctx.Param("profile_id"))
	if err != nil {
		return res.StatusInternalServerError(ctx, "cannot convert profile id to int", err)
	}
	result, err := profileHandler.ProfileService.GetProfileByID(uint(profileID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusBadRequest(ctx, "profile not found", err)
		}
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		return res.StatusInternalServerError(ctx, "failed to create category, something happen", err)
	}
	return res.StatusOK(ctx, "success to get user profile data", result, nil)
}
