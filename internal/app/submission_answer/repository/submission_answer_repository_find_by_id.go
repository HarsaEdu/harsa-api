package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (submissionAnswerRepository *SubmissionAnswerRepositoryImpl) FindById(id int) (*domain.SubmissionAnswer, error) {
	submissionAnswer := &domain.SubmissionAnswer{}
	result := submissionAnswerRepository.DB.Where("id=?", id).First(submissionAnswer)
	if result.Error != nil {
		return nil, result.Error
	}

	return submissionAnswer, nil
}
