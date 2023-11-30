package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (submissionRepository *SubmissionRepositoryImpl) FindById(id int) error {
	result := submissionRepository.DB.Where("id=?").First(&domain.Submissions{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
