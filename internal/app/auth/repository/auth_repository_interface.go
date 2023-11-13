package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterUser(user *domain.User) (*domain.Auth, error)
	LoginUser(id uint) (*domain.Auth, error)
}

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		DB: db,
	}
}
