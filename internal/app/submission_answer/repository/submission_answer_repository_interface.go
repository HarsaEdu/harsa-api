package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type SubmissionAnswerRepository interface {
	Create(request domain.SubmissionAnswer) error
}

type SubmissionAnswerRepositoryImpl struct {
	DB *gorm.DB
}

func NewSubmissionAnswer(db *gorm.DB) SubmissionAnswerRepository {
	return &SubmissionAnswerRepositoryImpl{DB: db}
}
