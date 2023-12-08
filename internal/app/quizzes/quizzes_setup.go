package quizzes

import (
	quizzesHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/quizzes/handler"
	quizzesRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/quizzes/repository"
	quizzesRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/quizzes/routes"
	quizzesServicePkg "github.com/HarsaEdu/harsa-api/internal/app/quizzes/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func QuizzesSetup(db *gorm.DB, validate *validator.Validate) (quizzesRoutesPkg.QuizzesRoutes, quizzesServicePkg.QuizzesService) {
	quizzesRepository := quizzesRepositoryPkg.NewQuizzesRepository(db)
	quizzesService := quizzesServicePkg.NewQuizzesService(quizzesRepository, validate)
	quizzesHandler := quizzesHandlerPkg.NewQuizzesHandler(quizzesService)
	quizzesRoute := quizzesRoutesPkg.NewQuizzesRoutes(quizzesHandler)

	return quizzesRoute, quizzesService
}
