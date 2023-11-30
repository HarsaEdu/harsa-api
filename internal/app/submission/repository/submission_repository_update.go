package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (submissionRepository *SubmissionRepositoryImpl) Update(submission *domain.Submissions, submissionId int) error {
	result := submissionRepository.DB.Where("id=?", submissionId).Updates(&domain.Submissions{Title: submission.Title, Description: submission.Description})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
