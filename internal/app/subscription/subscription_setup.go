package subscription

import (
	subsPlanRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/subs_plan/repository"
	subscriptionHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/subscription/handler"
	subscriptionRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/subscription/routes"
	subscriptionServicePkg "github.com/HarsaEdu/harsa-api/internal/app/subscription/service"
	userRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/pkg/midtrans"
	"github.com/go-playground/validator"
)

func SubscriptionSetup(validate *validator.Validate, midtransCoreApi midtrans.MidtransCoreApi, userRepository userRepositoryPkg.UserRepository, subsPlanRepository subsPlanRepositoryPkg.SubsPlanRepository) subscriptionRoutesPkg.SubscriptionRoutes {
	subscriptionService := subscriptionServicePkg.NewSubscriptionService(subsPlanRepository, userRepository, midtransCoreApi, validate)
	subscriptionHandler := subscriptionHandlerPkg.NewSubscriptionHandler(subscriptionService)
	subscriptionRoutes := subscriptionRoutesPkg.NewSubscriptionRoutes(subscriptionHandler)

	return subscriptionRoutes
}