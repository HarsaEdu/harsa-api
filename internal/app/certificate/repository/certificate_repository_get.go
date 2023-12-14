package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (certificateRepository *CertificateRepositoryImpl) GetById(id string) (*domain.Certificate, error) {
	certificate := domain.Certificate{}

	result := certificateRepository.DB.Where("id = ?", id).Find(&certificate)
	if result.Error != nil {
		return nil, result.Error
	}

	return &certificate, nil
}