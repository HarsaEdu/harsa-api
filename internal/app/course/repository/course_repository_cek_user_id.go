package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (courseRepository *CourseRepositoryImpl) CekIdFromCourse(userId uint, courseId uint, role string) (*domain.Course, error) {

	var course = domain.Course{}

	if err := courseRepository.DB.First(&course, courseId).Error; err != nil {
		return nil, err
	}

		if course.UserID != userId && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}
	

	return &course, nil
}
