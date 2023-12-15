package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"gorm.io/gorm"
)

type SubmissionAnswerRepository interface {
	Create(request domain.SubmissionAnswer) error
	Update(request domain.SubmissionAnswer, id, userId int) error
	FindById(id int) (*domain.SubmissionAnswer, error)
	Get(offset, limit int, search string, submissionID uint) ([]domain.SubmissionsAnswerDetail, int64, error)
	UpdateWeb(request domain.SubmissionAnswer, id int) error
	FindByIdWeb(id int) (*web.SubmissionAnswerResponseWebById, error)
}

type SubmissionAnswerRepositoryImpl struct {
	DB *gorm.DB
}

func NewSubmissionAnswer(db *gorm.DB) SubmissionAnswerRepository {
	return &SubmissionAnswerRepositoryImpl{DB: db}
}
