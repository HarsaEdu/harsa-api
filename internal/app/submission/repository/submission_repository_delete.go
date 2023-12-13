package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (submissionRepository *SubmissionRepositoryImpl) Delete(id int) error {
	result := submissionRepository.DB.Where("id = ?", id).Delete(&domain.Submissions{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
