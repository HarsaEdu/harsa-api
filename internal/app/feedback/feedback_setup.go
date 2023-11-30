package feedback

import (
	feedbackHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/feedback/handler"
	feedbackRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/feedback/repository"
	feedbackRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/feedback/routes"
	feedbackServicePkg "github.com/HarsaEdu/harsa-api/internal/app/feedback/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func FeedbackSetup(db *gorm.DB, validate *validator.Validate) feedbackRoutesPkg.FeedbackRoutes {
	feedbackRepository := feedbackRepositoryPkg.NewFeedbackRepository(db)
	feedbackService := feedbackServicePkg.NewFeedbackService(feedbackRepository, validate)
	feedbackHandler := feedbackHandlerPkg.NewFeedbackHandler(feedbackService)
	feedbackRoute := feedbackRoutesPkg.NewFeedbackRoutes(feedbackHandler)

	return feedbackRoute
}
