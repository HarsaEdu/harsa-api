package auth

import (
	authHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/auth/handler"
	authRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/auth/repository"
	authRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/auth/routes"
	authServicePkg "github.com/HarsaEdu/harsa-api/internal/app/auth/service"
	userRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"github.com/HarsaEdu/harsa-api/internal/pkg/password"
)

func AuthSetup(db *gorm.DB,validate *validator.Validate ,userRepo userRepositoryPkg.UserRepository, passwordImpl password.PasswordComparer ) authRoutesPkg.AuthRoutes {
	authRepository := authRepositoryPkg.NewAuthRepository(db)
	authService := authServicePkg.NewAuthService(authRepository,userRepo ,validate, passwordImpl)
	authHandler := authHandlerPkg.NewAuthHandler(authService)
	authRoute := authRoutesPkg.NewAuthRoutes(authHandler)

	return authRoute
} 