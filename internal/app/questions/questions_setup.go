package questions

import (
	questionsHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/questions/handler"
	questionsRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/questions/repository"
	questionsRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/questions/routes"
	questionsServicePkg "github.com/HarsaEdu/harsa-api/internal/app/questions/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func QuestionsSetup(db *gorm.DB, validate *validator.Validate)questionsRoutesPkg.QuestionsRoutes {
	questionsRepository := questionsRepositoryPkg.NewQuestionsRepository(db)
	questionsService := questionsServicePkg.NewQuestionsService(questionsRepository, validate)
	questionsHandler := questionsHandlerPkg.NewQuestionsHandler(questionsService)
	questionsRoute := questionsRoutesPkg.NewQuestionsRoutes(questionsHandler)

	return questionsRoute
}
