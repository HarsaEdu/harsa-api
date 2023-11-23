package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/feedback/handler"
	"github.com/labstack/echo/v4"
)

type FeedbackRoutes interface {
	FeedbackWeb(apiGroup *echo.Group)
	FeedbackMobile(apiGroup *echo.Group)
}

type FeedbackRoutesImpl struct {
	Echo            *echo.Echo
	FeedbackHandler handler.FeedbackHandler
}

func NewFeedbackRoutes(feedbackHandler handler.FeedbackHandler) FeedbackRoutes {
	return &FeedbackRoutesImpl{
		FeedbackHandler: feedbackHandler,
	}
}
