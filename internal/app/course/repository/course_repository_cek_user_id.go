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

func (courseRepository *CourseRepositoryImpl) CekIdFromUser(userId uint) (error) {

	var user = domain.User{}

	if err := courseRepository.DB.First(&user, userId).Error; err != nil {
		return err
	}
		if user.ID != userId && user.RoleID == 3 {
		return fmt.Errorf("unauthorized")
	}
	

	return nil
}
