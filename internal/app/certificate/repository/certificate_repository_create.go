package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (certificateRepository *CertificateRepositoryImpl) Create(certificate *domain.Certificate) error {
	result := certificateRepository.DB.Create(certificate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}