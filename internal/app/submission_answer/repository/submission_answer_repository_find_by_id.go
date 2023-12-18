package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (submissionAnswerRepository *SubmissionAnswerRepositoryImpl) FindById(id int) (*domain.SubmissionAnswer, error) {
	submissionAnswer := &domain.SubmissionAnswer{}
	result := submissionAnswerRepository.DB.Where("id=?", id).First(submissionAnswer)
	if result.Error != nil {
		return nil, result.Error
	}

	return submissionAnswer, nil
}

func (submissionAnswerRepository *SubmissionAnswerRepositoryImpl) FindByuserIDAndSubmissionID(submissionId int, userid int) (*domain.SubmissionAnswer, error) {
	submissionAnswer := &domain.SubmissionAnswer{}
	result := submissionAnswerRepository.DB.Where("submission_id = ? AND user_id = ?", submissionId, userid).First(submissionAnswer)
	if result.Error != nil {
		return nil, result.Error
	}

	return submissionAnswer, nil
}

func (submissionAnswerRepository *SubmissionAnswerRepositoryImpl) FindByIdWeb(id int) (*web.SubmissionAnswerResponseWebById, error) {
	submissionAnswer := &web.SubmissionAnswerResponseWebById{}
	result := submissionAnswerRepository.DB.Model(&domain.SubmissionAnswer{}).Where("submission_answers.id=?", id).Select("submission_answers.feedback as feedback, submission_answers.id as id, submission_answers.user_id as student_id, submission_answers.status as status,submission_answers.submitted_url as submitted_url, submissions.content as description, concat(user_profiles.first_name, ' ', user_profiles.last_name) as student_name").
	Joins("LEFT JOIN submissions ON submissions.id = submission_answers.submission_id").
	Joins("LEFT JOIN user_profiles ON user_profiles.id = submission_answers.user_id").
	Find(&submissionAnswer)
	if result.Error != nil {
		return nil, result.Error
	}

	return submissionAnswer, nil
}
