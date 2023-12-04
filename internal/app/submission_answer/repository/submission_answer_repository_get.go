package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (submissionAnswerRepository *SubmissionAnswerRepositoryImpl) Get(offset, limit int, search string, submissionID uint) ([]domain.SubmissionsAnswerDetail, int64, error) {
	submissionAnswer := []domain.SubmissionsAnswerDetail{}
	var count int64

	query := submissionAnswerRepository.DB.Model(&domain.SubmissionAnswer{}).Select("submission_answers.id as id,user_profiles.first_name as first_name, user_profiles.last_name as last_name  ,submission_answers.status as status").Joins("JOIN user_profiles on user_profiles.user_id=submission_answers.user_id").Where("submission_answers.submission_id = ?", submissionID)

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("user_profiles.first_name LIKE ? OR user_profiles.last_name LIKE ?", s, s)
	}

	query.Count(&count)

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
