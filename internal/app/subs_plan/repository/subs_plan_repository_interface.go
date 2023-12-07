package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type SubsPlanRepository interface {
	Create(subsPlan *domain.SubsPlan) error
	GetAll(search string) ([]domain.SubsPlan, int64, error)
	Update(subsPlan *domain.SubsPlan, id int) error
	Delete(id int) error
	FindById(id int) (*domain.SubsPlan, error)
}

type SubsPlanRepositoryImpl struct {
	DB *gorm.DB
}

func NewSubsPlanRepository(db *gorm.DB) SubsPlanRepository {
	return &SubsPlanRepositoryImpl{DB: db}
}
