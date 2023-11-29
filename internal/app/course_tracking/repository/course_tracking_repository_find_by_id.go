package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

func (courseTrackingrepository *CourseTrackingRepositoryImpl) FindById(courseTrackingId uint) (*domain.CourseTracking,int64,int64 ,error) {

	var courseTracking = domain.CourseTracking{}

	if err := courseTrackingrepository.DB.
	Preload("Course.User.UserProfile").
    Preload("Course.Feedback.User.UserProfile", func(db *gorm.DB) *gorm.DB {
        return db.Limit(5)
    }).
	Preload("User.UserProfile").
	Preload("HistoryModule.Module").
	Preload("HistoryModule.HistorySubModule.SubModule").
	Preload("HistoryModule.SubmissionAnswer.Submission").
	Preload("HistoryModule.HistoryQuiz.Quiz").
    First(&courseTracking, courseTrackingId).
    Error; err != nil {
    return nil, 0, 0, err
    }

	var courseTracking2 = []domain.CourseTracking{}

	var countEnroled int64

	if err := courseTrackingrepository.DB.Where("course_id = ?", courseTracking.CourseID).Find(&courseTracking2).Count(&countEnroled).Error; err != nil {
		return nil,0,0, err
	}

	var module = []domain.Module{}

	var countModule int64

	if err := courseTrackingrepository.DB.Where("course_id = ?", courseTracking.CourseID).Find(&module).Count(&countModule).Error; err != nil {
		return nil,0,0, err
	}

	
	return &courseTracking,countEnroled,countModule, nil
}