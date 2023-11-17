package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	CreateProfile(profile *domain.Profile) error
	FindByUserID(id uint) (*domain.Profile, error)
	GetAllProfiles() ([]domain.Profile, error)
	UpdateProfile(profile *domain.Profile, id uint) error
}

type ProfileRepositoryImpl struct {
	DB *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &ProfileRepositoryImpl{
		DB: db,
	}
}
