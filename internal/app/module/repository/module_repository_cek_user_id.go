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

func (moduleRepository *ModuleRepositoryImpl) CekIdFromModule(userId uint, moduleId uint, role string) (*domain.Module, error) {

	var module = domain.Module{}

	if err := moduleRepository.DB.First(&module, moduleId).Error; err != nil {
		return nil, err
	}
	
	var course = domain.Course{}

	if err := moduleRepository.DB.First(&course, module.CourseID).Error; err != nil {
		return nil, err
	}

	if course.UserID != userId && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}

	return &module, nil
}
