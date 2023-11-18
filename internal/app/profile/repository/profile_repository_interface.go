package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	CreateProfile(profile *domain.UserProfile) error
	FindByUserID(id uint) (*domain.UserProfile, error)
	UpdateProfile(profile *domain.UserProfile) error
}

type ProfileRepositoryImpl struct {
	DB *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &ProfileRepositoryImpl{
		DB: db,
	}
}
