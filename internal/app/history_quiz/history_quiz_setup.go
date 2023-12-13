package historyquiz

import (
	historyQuizHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/history_quiz/handler"
	historyQuizRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/history_quiz/repository"
	historyQuizRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/history_quiz/routes"
	historyQuizServicePkg "github.com/HarsaEdu/harsa-api/internal/app/history_quiz/service"
	subscriptionServicePkg "github.com/HarsaEdu/harsa-api/internal/app/subscription/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func HistoryQuizSetup(db *gorm.DB, validate *validator.Validate, subcriptionService subscriptionServicePkg.SubscriptionService) historyQuizRoutesPkg.HistoryQuizRoutes {
	historyQuizRepository := historyQuizRepositoryPkg.NewHistoryQuizRepository(db)
	historyQuizService := historyQuizServicePkg.NewHistoryQuizService(historyQuizRepository, validate)
	historyQuizHandler := historyQuizHandlerPkg.NewHistoryQuizHandler(historyQuizService, subcriptionService)
	historyQuizRoute := historyQuizRoutesPkg.NewHistoryQuizRoutes(historyQuizHandler)

	return historyQuizRoute
}
