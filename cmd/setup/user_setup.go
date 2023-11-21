package setup

import (
	userHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/user/handler"
	userRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	userRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/user/routes"
	userServicePkg "github.com/HarsaEdu/harsa-api/internal/app/user/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func UserSetup(db *gorm.DB, validate *validator.Validate) (userRoutesPkg.UserRoutes , userRepositoryPkg.UserRepository){
	userRepository := userRepositoryPkg.NewUserRepository(db)
	userService := userServicePkg.NewUserService(userRepository, validate)
	userHandler := userHandlerPkg.NewUserHandler(userService)
	userRoute := userRoutesPkg.NewUserRoutes(userHandler)

	return userRoute, userRepository
}
