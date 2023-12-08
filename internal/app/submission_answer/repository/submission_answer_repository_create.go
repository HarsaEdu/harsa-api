package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (submissionAnswerRepository *SubmissionAnswerRepositoryImpl) Create(request domain.SubmissionAnswer) error {
	result := submissionAnswerRepository.DB.Create(&request)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
