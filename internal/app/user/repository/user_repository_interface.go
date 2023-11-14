package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	UserAvailable(username, email string) (*domain.User, error)
	UserAvailableByID(id uint) (*domain.User, error)
	UserProfileCreate(user *domain.UserProfile) error
	UserCreate(user *domain.User) (*domain.User, error)
	UserUpdate(user *domain.User) error
	UserProfileUpdate(userProfile *domain.UserProfile) error
	UserProfileAvailableByID(id uint) (*domain.UserProfile, error)
	UserDelete(id uint) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}
