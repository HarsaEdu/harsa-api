package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/history_quiz/service"
	"github.com/labstack/echo/v4"
	subscriptionServicePkg "github.com/HarsaEdu/harsa-api/internal/app/subscription/service"
)

type HistoryQuizHandler interface {
	FindById(ctx echo.Context) error
	Create(ctx echo.Context) error
	GetAll(ctx echo.Context) error
}

type HistoryQuizHandlereImpl struct {
	SubcriptionService subscriptionServicePkg.SubscriptionService
	HistoryQuizService service.HistoryQuizService
}

func NewHistoryQuizHandler(service service.HistoryQuizService, subcriptionService subscriptionServicePkg.SubscriptionService) HistoryQuizHandler {
	return &HistoryQuizHandlereImpl{
		SubcriptionService: subcriptionService,
		HistoryQuizService: service,
	}
}