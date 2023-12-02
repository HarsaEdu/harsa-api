package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (feedbackRepository *FeedbackRepositoryImpl) UpdateByUserAndCourseId(userId, courseId uint, feedback *domain.Feedback) error {
	result  := feedbackRepository.DB.Where("user_id = ? AND course_id = ?", userId, courseId).Updates(&domain.Feedback{Rating: feedback.Rating, Content: feedback.Content})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
