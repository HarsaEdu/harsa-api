package profile

import (
	profileHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/profile/handler"
	profileRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/profile/repository"
	profileRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/profile/routes"
	profileServicePkg "github.com/HarsaEdu/harsa-api/internal/app/profile/service"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProfileSetup(db *gorm.DB, validate *validator.Validate, e *echo.Echo, cloudinary cloudinary.CloudinaryUploader) profileRoutesPkg.ProfileRoutes {
	profileRepository := profileRepositoryPkg.NewProfileRepository(db)
	profileService := profileServicePkg.NewProfileService(profileRepository, validate, cloudinary)
	profileHandler := profileHandlerPkg.NewProfileHandler(profileService)
	profileRoute := profileRoutesPkg.NewProfileRoutes(profileHandler)

	return profileRoute
}
