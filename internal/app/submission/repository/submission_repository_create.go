package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (SubmissionRepository *SubmissionRepositoryImpl) Create(submission *domain.Submissions) error {
	result := SubmissionRepository.DB.Create(&submission)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
