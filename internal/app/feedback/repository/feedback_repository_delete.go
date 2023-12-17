package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (feedbackRepository *FeedbackRepositoryImpl) DeleteByUserAndCourseId(userId, courseId uint) error {
	result := feedbackRepository.DB.Where("user_id = ? AND course_id = ?", userId, courseId).Delete(&domain.Feedback{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (feedbackRepository *FeedbackRepositoryImpl) DeleteById(id uint) error {
	result := feedbackRepository.DB.Where("id = ? ", id).Delete(&domain.Feedback{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
