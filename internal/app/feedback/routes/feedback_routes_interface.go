package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/feedback/handler"
	"github.com/labstack/echo/v4"
)

type FeedbackRoutes interface {
	Feedback(apiGroup *echo.Group)
}

type FeedbackRoutesImpl struct {
	Echo            *echo.Echo
	FeedbackHandler handler.FeedbackHandler
}

func NewFeedbackRoutes(e *echo.Echo, feedbackHandler handler.FeedbackHandler) FeedbackRoutes {
	return &FeedbackRoutesImpl{
		Echo:            e,
		FeedbackHandler: feedbackHandler,
	}
}
