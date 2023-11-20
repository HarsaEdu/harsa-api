package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (feedbackRepository *FeedbackRepositoryImpl) GetAll(offset, limit int, search string) ([]domain.Feedback, int64, error) {
	var feedbacks []domain.Feedback
	var count int64

	query := feedbackRepository.DB

	if search != "" {
		searchQuery := "%" + search + "%"
		query = query.Where("title LIKE ?", searchQuery)
	}

	query.Find(&feedbacks).Count(&count)

	query = query.Offset(offset).Limit(limit)

	result := query.Find(&feedbacks)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if count == 0 {
		return nil, 0, nil
	}

	return feedbacks, count, nil
}
