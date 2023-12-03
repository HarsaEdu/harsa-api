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