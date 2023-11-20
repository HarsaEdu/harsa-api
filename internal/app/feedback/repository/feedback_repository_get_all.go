package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (feedbackRepository *FeedbackRepositoryImpl) GetAll(courseid, page, pagesize int) ([]domain.Feedback, int64, error) {
	var feedbacks []domain.Feedback
	var count int64

	if page < 1 {
		page = 1
	}

	if pagesize < 1 {
		pagesize = 10
	}

	query := feedbackRepository.DB

	if courseid != 0 {
		query = query.Where("course_id LIKE ?", courseid)
	}

	query.Find(&feedbacks).Count(&count)

	offset := (page - 1) * pagesize

	if offset < 0 {
		offset = 0
	}

	query = query.Offset(offset).Limit(pagesize)

	result := query.Find(&feedbacks)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if count == 0 {
		return nil, 0, nil
	}

	return feedbacks, count, nil
}
