package auth

import (
	authHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/auth/handler"
	authRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/auth/repository"
	authRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/auth/routes"
	authServicePkg "github.com/HarsaEdu/harsa-api/internal/app/auth/service"
	"github.com/HarsaEdu/harsa-api/internal/app/notification/repository"
	userRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/go-playground/validator"
	"github.com/HarsaEdu/harsa-api/internal/pkg/firebase"
	"gorm.io/gorm"
	"github.com/HarsaEdu/harsa-api/internal/pkg/password"
)

func AuthSetup(db *gorm.DB,validate *validator.Validate ,userRepo userRepositoryPkg.UserRepository, passwordImpl password.PasswordComparer, notifRepository repository.NotificationRepository, firebaseImpl firebase.Firebase ) authRoutesPkg.AuthRoutes {

	authRepository := authRepositoryPkg.NewAuthRepository(db)
	authService := authServicePkg.NewAuthService(authRepository, userRepo, validate, passwordImpl,notifRepository, firebaseImpl)
	authHandler := authHandlerPkg.NewAuthHandler(authService)
	authRoute := authRoutesPkg.NewAuthRoutes(authHandler)

	return authRoute
}
