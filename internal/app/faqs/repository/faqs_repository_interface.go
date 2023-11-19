package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type FaqRepository interface {
	Create(Faq *domain.Faqs) error
	GetAll(offset, limit int, search string) ([]domain.Faqs, int64, error)
	Update(Faq *domain.Faqs, id int) error
	Delete(id int) error
	FindById(id int) (*domain.Faqs, error)
}

type FaqRepositoryImpl struct {
	DB *gorm.DB
}

func NewFaqRepository(db *gorm.DB) FaqRepository {
	return &FaqRepositoryImpl{DB: db}
}
