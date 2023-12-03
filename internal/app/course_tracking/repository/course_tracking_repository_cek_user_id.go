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


