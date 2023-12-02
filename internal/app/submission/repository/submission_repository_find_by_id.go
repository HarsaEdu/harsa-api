package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (submissionRepository *SubmissionRepositoryImpl) FindById(id int) (*domain.Submissions, error) {

	submission := &domain.Submissions{}

	result := submissionRepository.DB.Where("id=?", id).First(submission)
	if result.Error != nil {
		return nil, result.Error
	}

	return submission, nil
}
