package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (courseTrackingrepository *CourseTrackingRepositoryImpl) CekIdFromCourse(userId uint, trackingID uint, role string) (*domain.CourseTracking, error) {

	courseTracking := domain.CourseTracking{}

	if err := courseTrackingrepository.DB.First(&courseTracking, trackingID).Error; err != nil {
		return nil, err
	}
	
	var userIDCourse uint

	if err := courseTrackingrepository.DB.Model(&domain.Course{}).Where("id = ?", courseTracking.CourseID).Select("user_id").Scan(&userIDCourse).Error; err != nil {
		return nil, err
	}

	if userIDCourse != userId && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}

	return  &courseTracking,nil
}

func (courseTrackingrepository *CourseTrackingRepositoryImpl) Cek(userId uint, courseId uint) (*domain.CourseTracking,error) {

	courseTraking := domain.CourseTracking{}

	if err := courseTrackingrepository.DB.Model(&domain.CourseTracking{}).Where("user_id  = ? AND course_id = ?" ,userId,courseId).First(&courseTraking).
		Error; err != nil {
		return nil,  err
	}

	return &courseTraking,  nil
}


