package auth

import (
	"github.com/HarsaEdu/harsa-api/configs"
	authHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/auth/handler"
	authRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/auth/repository"
	authRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/auth/routes"
	authServicePkg "github.com/HarsaEdu/harsa-api/internal/app/auth/service"
	"github.com/HarsaEdu/harsa-api/internal/app/notification/repository"
	userRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func AuthSetup(db *gorm.DB, validate *validator.Validate, userRepo userRepositoryPkg.UserRepository, notifRepository repository.NotificationRepository, config configs.AppConfig) authRoutesPkg.AuthRoutes {
	authRepository := authRepositoryPkg.NewAuthRepository(db)
	authService := authServicePkg.NewAuthService(authRepository, userRepo, validate, notifRepository, config)
	authHandler := authHandlerPkg.NewAuthHandler(authService)
	authRoute := authRoutesPkg.NewAuthRoutes(authHandler)

	return authRoute
}
