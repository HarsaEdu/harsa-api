package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type SubmissionRepository interface {
	Create(submission *domain.Submissions) error
	Update(submission *domain.Submissions, submissionId int) error
	FindById(id int) error
	Delete(id int) error
}

type SubmissionRepositoryImpl struct {
	DB *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) SubmissionRepository {
	return &SubmissionRepositoryImpl{DB: db}
}
