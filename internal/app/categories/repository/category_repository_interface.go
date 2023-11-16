package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *domain.Category) error
	Update(category *domain.Category, id int) error
	UpdateImage(imageUrl *domain.Category, id int) error
	FindById(id int) (*domain.Category, error)
	FindByName(name string) (*domain.Category, error)
	GetAll() ([]domain.Category, error)
	Delete(id int) error
}

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}
