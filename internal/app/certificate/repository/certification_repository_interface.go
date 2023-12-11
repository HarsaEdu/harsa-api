package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type CertificateRepository interface {
	Create(certificate *domain.Certificate) error
	GetById(id string) (*domain.Certificate, error)
}

type CertificateRepositoryImpl struct {
	DB *gorm.DB
}

func NewCertificateRepository(db *gorm.DB) CertificateRepository {
	return &CertificateRepositoryImpl{
		DB: db,
	}
}