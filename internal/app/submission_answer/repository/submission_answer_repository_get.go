package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (submissionAnswerRepository *SubmissionAnswerRepositoryImpl) Get(offset, limit int, search string) ([]domain.SubmissionsAnswerDetail, int64, error) {
	submissionAnswer := []domain.SubmissionsAnswerDetail{}
	var count int64

	query := submissionAnswerRepository.DB.Model(&submissionAnswer).Select("users.id as peserta, submission_answers.status as status").Joins("JOIN users on users.id=submission_answers.user_id")

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("question answer submission_answer.status ?", s)
	}

	query.Find(&submissionAnswer).Count(&count)

	query = query.Limit(limit).Offset(offset)
	result := query.Find(&submissionAnswer)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if count == 0 {
		return nil, 0, nil
	}

	if offset >= int(count) {
		return nil, 0, nil
	}

	return submissionAnswer, count, nil
}
