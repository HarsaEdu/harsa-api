package certificate

import (
	certificateHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/certificate/handler"
	certificateRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/certificate/repository"
	certificateRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/certificate/routes"
	certificateServicePkg "github.com/HarsaEdu/harsa-api/internal/app/certificate/service"
	courseTrackingRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/course_tracking/repository"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	wkh "github.com/HarsaEdu/harsa-api/internal/pkg/wkhtmltopdf"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func CertificateSetup(db *gorm.DB, validate *validator.Validate, cloudinary cloudinary.CloudinaryUploader, courseTrackingRepository courseTrackingRepositoryPkg.CourseTrackingRepository, certifaicate wkh.PDFGenerator) certificateRoutesPkg.CertificateRoutes {
	certificateRepository := certificateRepositoryPkg.NewCertificateRepository(db)
	certificateService := certificateServicePkg.NewCertificateService(certificateRepository, courseTrackingRepository, cloudinary, certifaicate)
	certificateHandler := certificateHandlerPkg.NewCertificateHandler(certificateService)
	certificateRoutes := certificateRoutesPkg.NewCertificateRoutes(certificateHandler)

	return certificateRoutes
}