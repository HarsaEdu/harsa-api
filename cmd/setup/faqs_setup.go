package setup

import (
	faqsHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/faqs/handler"
	faqsRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/faqs/repository"
	faqsRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/faqs/routes"
	faqsServicePkg "github.com/HarsaEdu/harsa-api/internal/app/faqs/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func FaqsSetup(db *gorm.DB, validate *validator.Validate) faqsRoutesPkg.FaqsRoutes{
	faqsRepository := faqsRepositoryPkg.NewFaqRepository(db)
	faqsService := faqsServicePkg.NewFaqsService(faqsRepository, validate)
	faqsHandler := faqsHandlerPkg.NewFaqsHandler(faqsService)
	faqsRoute := faqsRoutesPkg.NewFaqsRoutes(faqsHandler)

	return faqsRoute
}
