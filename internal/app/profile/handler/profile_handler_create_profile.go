package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (profileHandler *ProfileHandlerImpl) CreateProfile(ctx echo.Context) error {
	userID := ctx.Get("user_id")
	profile := web.CreateProfileRequest{}
	if err := ctx.Bind(&profile); err != nil {
		return res.StatusBadRequest(ctx, "failed to bind profile model", err)
	}

	err := profileHandler.ProfileService.CreateProfile(ctx, &profile, userID.(uint))

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return res.StatusBadRequest(ctx, "validation failed", err)
		}
		if strings.Contains(err.Error(), "Unauthorized") {
			return res.StatusBadRequest(ctx, "missing or malformed jwt", err)
		}
		if strings.Contains(err.Error(), "already exists") {
			return res.StatusAlreadyExist(ctx, "profile already exists", err)
		}
		if strings.Contains(err.Error(), "file not found") {
			return res.StatusNotFound(ctx, "file not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to create profile, something happen", err)
	}

	return res.StatusCreated(ctx, "Successfuly create profile", nil, nil)
}
