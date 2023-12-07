package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (handler *InterestHandlerImpl) GetInterestRecommendation(ctx echo.Context) error {
	userID := ctx.Get("user_id").(uint)
	result, err := handler.Service.GetInterestRecommendation(userID)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "category not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get interest, something happen", err)
	}
	return res.StatusOK(ctx, "success to get interest", result, nil)
}
