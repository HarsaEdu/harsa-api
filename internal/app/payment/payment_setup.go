package payment

import (
	paymentHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/payment/handler"
	paymentRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/payment/repository"
	paymentRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/payment/routes"
	paymentServicePkg "github.com/HarsaEdu/harsa-api/internal/app/payment/service"
	subsPlanRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/subs_plan/repository"
	userRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/pkg/midtrans"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func PaymentSetup(db *gorm.DB, validate *validator.Validate, midtransCoreApi midtrans.MidtransCoreApi, userRepository userRepositoryPkg.UserRepository, subsPlanRepository subsPlanRepositoryPkg.SubsPlanRepository) paymentRoutesPkg.PaymentRoutes {
	paymentRepository := paymentRepositoryPkg.NewPaymentRepository(db)
	paymentService := paymentServicePkg.NewPaymentService(paymentRepository, subsPlanRepository, userRepository, midtransCoreApi, validate)
	paymentHandler := paymentHandlerPkg.NewPaymentHandler(paymentService)
	paymentRoutes := paymentRoutesPkg.NewPaymentRoutes(paymentHandler)

	return paymentRoutes
}