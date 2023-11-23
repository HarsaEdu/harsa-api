package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type OptionsRepository interface {
	CekIdFromOption(userId uint, optionId uint, role string) (*domain.Options, error)
	Delete(option *domain.Options) error
}

type OptionsRepositoryImpl struct {
	DB *gorm.DB
}

func NewOptionsRepository(db *gorm.DB) OptionsRepository {
	return &OptionsRepositoryImpl{
		DB: db,
	}
}