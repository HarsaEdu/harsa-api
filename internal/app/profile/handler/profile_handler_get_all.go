package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (profileHandler *ProfileHandlerImpl) GetAllProfiles(ctx echo.Context) error {
	result, err := profileHandler.ProfileService.GetAllProfiles()
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "profile not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get all profiles, something happen", err)
	}
	return res.StatusOK(ctx, "success to get user profiles data", result)
}
