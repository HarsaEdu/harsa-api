package certificate

import (
	certificateHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/certificate/handler"
	certificateRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/certificate/repository"
	certificateRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/certificate/routes"
	certificateServicePkg "github.com/HarsaEdu/harsa-api/internal/app/certificate/service"
	courseTrackingRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/course_tracking/repository"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/HarsaEdu/harsa-api/web/template"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func CertificateSetup(db *gorm.DB, validate *validator.Validate, cloudinary cloudinary.CloudinaryUploader, courseTrackingRepository courseTrackingRepositoryPkg.CourseTrackingRepository) certificateRoutesPkg.CertificateRoutes {
	certificateRepository := certificateRepositoryPkg.NewCertificateRepository(db)
	certificateService := certificateServicePkg.NewCertificateService(certificateRepository, courseTrackingRepository, cloudinary, template.CertificateBlankContent)
	certificateHandler := certificateHandlerPkg.NewCertificateHandler(certificateService)
	certificateRoutes := certificateRoutesPkg.NewCertificateRoutes(certificateHandler)

	return certificateRoutes
}