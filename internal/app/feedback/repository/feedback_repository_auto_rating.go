package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (feedbackRepository *FeedbackRepositoryImpl) AutoUpdateRating(courseId uint) error {
	tx := feedbackRepository.DB

	var listFeedBack []domain.Feedback

	var count int64

	if err := tx.Model(&domain.Feedback{}).Where("course_id = ?",courseId).Find(&listFeedBack).Count(&count).Error; err != nil {
		return err
	}

	var allRating float32

	if count == 0 {
		allRating = 0
	} else {
		for _, student := range listFeedBack {
			allRating = allRating + float32(student.Rating)
		}
	
		allRating = allRating / float32(count)
	}

	if err := tx.Model(&domain.Course{}).Where("id=?", courseId).Update("rating", allRating).Error; err != nil {
		return err
	}

	return nil
}