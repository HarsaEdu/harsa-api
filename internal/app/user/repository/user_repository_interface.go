package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/labstack/echo/v4"
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
	HandleTrx(ctx echo.Context, fn func(repo UserRepository) error) error
	UserGetAll(offset, limit int, search string) ([]domain.UserEntity, int64, error)
	GetUserByID(userID uint) (*domain.UserDetail, error)
	GetUserAccountByID(userID uint) (*domain.User, error)
	UserAvailableUsername(username string) (*domain.User, error)
	UserAvailableEmail(email string) (*domain.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}
