package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (feedbackRepository *FeedbackRepositoryImpl) GetAllByCourseId(courseId uint, offset, limit int, search string) ([]domain.Feedback, int64, error) {

	if offset < 0 || limit < 0 {
		return nil, 0, nil
	}

	feedback := []domain.Feedback{}
	var total int64

	query := feedbackRepository.DB.Model(&feedback).Preload("User.UserProfile").Where("course_id = ?", courseId)

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("rating LIKE ?", s)
	}

	query.Find(&feedback).Count(&total)

	query = query.Limit(limit).Offset(offset)

	result := query.Find(&feedback)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 {
		return nil, 0, nil
	}

	if offset >= int(total) {
		return nil, 0, nil
	}

	return feedback, total, nil
}
