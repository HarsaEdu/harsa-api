package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	CreateProfile(profile *domain.UserProfile) error
	FindByUserID(userID uint) (*domain.UserDetail, error)
	UpdateProfile(profile *domain.UserProfile) error
	GetProfileByID(profileID uint) (*domain.UserProfile, error)
}

type ProfileRepositoryImpl struct {
	DB *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &ProfileRepositoryImpl{
		DB: db,
	}
}
