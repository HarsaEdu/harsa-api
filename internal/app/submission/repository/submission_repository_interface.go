package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"gorm.io/gorm"
)

type SubmissionRepository interface {
	Create(submission *domain.Submissions) error
	Update(submission *domain.Submissions, submissionId int) error
	GetAll(moduleId int) ([]domain.Submissions, int64, error)
	FindById(id int) (*domain.Submissions, error)
	Delete(id int) error
	CekUserIDfromModuleID(id uint, userId uint, role string) error
	CekUserIDfromSubmission(id uint, userId uint, role string)  error
	GetAllWeb(moduleId int) ([]web.SubmissionsResponseWeb, int64, error)
}

type SubmissionRepositoryImpl struct {
	DB *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) SubmissionRepository {
	return &SubmissionRepositoryImpl{DB: db}
}
