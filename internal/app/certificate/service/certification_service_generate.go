package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	
)

func (certificateService *CertificateServiceImpl) GenerateCertificate(courseTrackingId uint, userId uint) ([]byte, *domain.Certificate, error) {
	existingCourseTracking, _ := certificateService.CourseTrackingRepository.FindById(courseTrackingId)
	if existingCourseTracking == nil {
		return nil, nil, fmt.Errorf("error when get course tracking : course tracking not found")
	}

	if existingCourseTracking.Status != "completed" {
		return nil, nil, fmt.Errorf("error when generating certificate : course not completed")
	}

	if existingCourseTracking.UserID != userId {
		return nil, nil, fmt.Errorf("unauthorized")
	}

	certificate := conversion.CourseTrackingDomainToCertificateDomain(existingCourseTracking)

	certificatePdf, err := certificateService.GenerateCertif.GenerateCertificate(certificate)
	if err != nil {
		return nil, nil, fmt.Errorf("error generating certificate: %w", err)
	}

	existingCertificate, _ := certificateService.CertificateRepository.GetById(certificate.ID)
	if existingCertificate != nil {
		return certificatePdf, existingCertificate, nil
	}
	
	err = certificateService.CertificateRepository.Create(certificate)

	certificate, _ = certificateService.CertificateRepository.GetById(certificate.ID)
	
	return certificatePdf, certificate, nil
}