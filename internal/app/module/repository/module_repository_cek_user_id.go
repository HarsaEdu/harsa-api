package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (moduleRepository *ModuleRepositoryImpl) CekIdFromCourse(userId uint, courseId uint, role string) error {

	var course = domain.Course{}

	if err := moduleRepository.DB.First(&course, courseId).Error; err != nil {
		return err
	}

	if course.UserID != userId && role != "admin" {
		return fmt.Errorf("unauthorized")
	}

	return nil
}
