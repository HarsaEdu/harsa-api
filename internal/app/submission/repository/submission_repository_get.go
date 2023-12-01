package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (submissionRepository *SubmissionRepositoryImpl) GetAll() ([]domain.Submissions, int64, error) {

	submission := []domain.Submissions{}
	var total int64

	query := submissionRepository.DB.Find(&submission)
	if query.Error != nil {
		return nil, 0, query.Error
	}

	query = query.Find(&submission).Count(&total)

	return submission, total, nil

}
