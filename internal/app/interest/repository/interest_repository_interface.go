package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type InterestRepository interface {
	CreateInterest(interest *domain.UserInterest) error
	FindByProfileID(profileID uint) (*domain.UserInterest, error)
	GetInterestRecommendation(profileID uint) ([]domain.Course, int64, error)
}

type InterestRepositoryImpl struct {
	DB *gorm.DB
}

func NewInterestRepository(db *gorm.DB) InterestRepository {
	return &InterestRepositoryImpl{
		DB: db,
	}
}
