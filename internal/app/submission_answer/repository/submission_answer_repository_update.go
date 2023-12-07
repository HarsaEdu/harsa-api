package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (submissionAnswerRepository *SubmissionAnswerRepositoryImpl) Update(request domain.SubmissionAnswer, id, userId int) error {
	result := submissionAnswerRepository.DB.Where("id=? and user_id=?", id, userId).Updates(&domain.SubmissionAnswer{SubmittedUrl: request.SubmittedUrl})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
