package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	
)

func (courseTrackingrepository *CourseTrackingRepositoryImpl) FindById(courseTrackingId uint) (*domain.CourseTracking ,error) {

	var courseTracking = domain.CourseTracking{}

	if err := courseTrackingrepository.DB.
	Preload("User.UserProfile").
    First(&courseTracking, courseTrackingId).
    Error; err != nil {
    return nil, err
    }
	
	return &courseTracking, nil
}

func (courseTrackingrepository *CourseTrackingRepositoryImpl) FindByUserIdAndCourseID(courseID uint, UserID uint) (*domain.CourseTracking ,error) {

	var courseTracking = domain.CourseTracking{}

	if err := courseTrackingrepository.DB.Model(&domain.CourseTracking{}).Where("user_id = ? AND course_id = ? ", UserID, courseID).
    Find(&courseTracking).
    Error; err != nil {
    return nil, err
    }

	var UserProfile = domain.UserProfile{}
	if err := courseTrackingrepository.DB.Model(&domain.UserProfile{}).Where("user_id = ? ", UserID).
    Find(&UserProfile).
    Error; err != nil {
    return nil, err
    }

	courseTracking.User.UserProfile = UserProfile
	courseTracking.User.UserProfile.ID = UserID
	
	return &courseTracking, nil
}