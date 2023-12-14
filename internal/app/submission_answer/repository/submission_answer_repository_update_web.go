package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (submissionAnswerRepository *SubmissionAnswerRepositoryImpl) UpdateWeb(request domain.SubmissionAnswer, id int) error {
	result := submissionAnswerRepository.DB.Where("id=?", id).Updates(&domain.SubmissionAnswer{Status: request.Status, Feedback: request.Feedback})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
