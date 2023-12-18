package service

import (

	"github.com/HarsaEdu/harsa-api/internal/app/certificate/repository"
	courseTrackingRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/course_tracking/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	wkh "github.com/HarsaEdu/harsa-api/internal/pkg/wkhtmltopdf"
)

type CertificateService interface {
	GenerateCertificate(courseTrackingId uint, userId uint) ([]byte, *domain.Certificate, error)
}

type CertificateServiceImpl struct {
	CertificateRepository repository.CertificateRepository
	CourseTrackingRepository courseTrackingRepositoryPkg.CourseTrackingRepository
	CloudinaryUploader cloudinary.CloudinaryUploader
	GenerateCertif wkh.PDFGenerator
}

func NewCertificateService(certificateRepository repository.CertificateRepository, courseTrackingRepository courseTrackingRepositoryPkg.CourseTrackingRepository, cloudinaryUploader cloudinary.CloudinaryUploader, generateCertificate wkh.PDFGenerator) CertificateService {
	return &CertificateServiceImpl{
		CertificateRepository: certificateRepository,
		CourseTrackingRepository: courseTrackingRepository,
		CloudinaryUploader: cloudinaryUploader,
		GenerateCertif: generateCertificate,
	}
}