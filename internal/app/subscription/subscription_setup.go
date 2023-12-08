package subscription

import (
	subscriptionRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/subscription/repository"
	subscriptionServicePkg "github.com/HarsaEdu/harsa-api/internal/app/subscription/service"
	"gorm.io/gorm"
)

func SubscriptionSetup(db *gorm.DB) (subscriptionServicePkg.SubscriptionService) {
	subscriptionRepository := subscriptionRepositoryPkg.NewSubscriptionRepository(db)
	subscriptionService := subscriptionServicePkg.NewSubscriptionService(subscriptionRepository)

	return subscriptionService
}