package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (profileHandler *ProfileHandlerImpl) GetProfileByID(ctx echo.Context) error {
	idParam := ctx.Param("profile_id")
	id, _ := strconv.Atoi(idParam)

	result, err := profileHandler.ProfileService.GetProfileByID(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusBadRequest(ctx, "profile not found", err)
		}
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		return res.StatusInternalServerError(ctx, "failed to create category, something happen", err)
	}
	return res.StatusOK(ctx, "success to get user profile data", result)
}
