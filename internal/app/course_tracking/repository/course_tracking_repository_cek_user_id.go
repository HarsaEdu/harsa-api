package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (courseTrackingrepository *CourseTrackingRepositoryImpl) CekIdFromCourse(userId uint, courseId uint, role string) error {

	var userIDCourse uint

	if err := courseTrackingrepository.DB.Model(&domain.Course{}).Where("id = ?", courseId).Select("user_id").Scan(&userIDCourse).Error; err != nil {
		return err
	}

	if userIDCourse != userId && role != "admin" {
		return fmt.Errorf("unauthorized")
	}

	return  nil
}

func (courseTrackingrepository *CourseTrackingRepositoryImpl) Cek(userId uint, courseId uint) (*domain.CourseTracking,error) {

	courseTraking := domain.CourseTracking{}

	if err := courseTrackingrepository.DB.Model(&domain.CourseTracking{}).Where("user_id  = ? AND course_id = ?" ,userId,courseId).First(&courseTraking).
		Error; err != nil {
		return nil,  err
	}

	return &courseTraking,  nil
}


