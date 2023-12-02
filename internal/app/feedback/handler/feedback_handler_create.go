package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (feedbackHandler *FeedbackHandlerImpl) Create(ctx echo.Context) error {

	req := web.FeedbackCreateRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind feedback request", err)
	}

	userId := ctx.Get("user_id").(uint)

	err = feedbackHandler.FeedbackService.Create(req, userId)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "already exist") {
			return res.StatusAlreadyExist(ctx, "feedback already exist", err)
		}
		return res.StatusInternalServerError(ctx, "failed to create feedback, something happen", err)

	}

	return res.StatusCreated(ctx, "success to create feedback", nil, nil)
}
