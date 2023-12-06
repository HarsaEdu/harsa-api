package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (recommendationsHandler *RecommendationsHandlerImpl) GetRecommendations(ctx echo.Context) error {
	userId := ctx.Get("user_id").(uint)

	request := web.GetRecommendationsRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		return res.StatusBadRequest(ctx, err.Error(), err)
	}

	request.UserId = userId
	if request.MaxResults == 0 {
		request.MaxResults = 5
	}

	response, err := recommendationsHandler.RecommendationsService.GetRecommendations(&request)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "error when get user") {
			return res.StatusAlreadyExist(ctx, "user not found", err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to get recommendations, something happen", err)
	}

	return res.StatusOK(ctx, "success to get recommendations", response, nil)
}