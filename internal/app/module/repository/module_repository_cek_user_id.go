package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (moduleRepository *ModuleRepositoryImpl) CekIdFromCourse(userId uint, courseId uint, role string) error {

	var userIDCourse uint

	if err := moduleRepository.DB.Model(&domain.Course{}).Where("id = ?", courseId).Select("user_id").Scan(&userIDCourse).Error; err != nil {
		return err
	}

	if userIDCourse != userId && role != "admin" {
		return fmt.Errorf("unauthorized")
	}

	return nil
}

func (moduleRepository *ModuleRepositoryImpl) CekIdFromSection(userId uint, sectionId uint, role string) (*domain.Section, error) {


	var section = domain.Section{} 

	if err := moduleRepository.DB.First(&section, sectionId).Error; err != nil {
		return nil, err
	}

	var userIDCourse uint

	if err := moduleRepository.DB.Model(&domain.Course{}).Where("id = ?", section.CourseID).Select("user_id").Scan(&userIDCourse).Error; err != nil {
		return nil, err
	}

	if userIDCourse != userId && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}

	return &section, nil
}

func (moduleRepository *ModuleRepositoryImpl) CekIdFromModule(userId uint, moduleId uint, role string) (*domain.Module, error) {

	var module = domain.Module{}

	if err := moduleRepository.DB.First(&module, moduleId).Error; err != nil {
		return nil, err
	}

	var courseID uint

	if err := moduleRepository.DB.Model(&domain.Section{}).Where("id = ?", module.SectionID).Select("course_id").Scan(&courseID).Error; err != nil {
		return nil, err
	}
	
	var userIDCourse uint

	if err := moduleRepository.DB.Model(&domain.Course{}).Where("id = ?", courseID ).Select("user_id").Scan(&userIDCourse).Error; err != nil {
		return nil, err
	}

	if userIDCourse != userId && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}
	return &module, nil
}

func (moduleRepository *ModuleRepositoryImpl) CekIdFromSubModule(userId uint, subModuleId uint, role string) (*domain.SubModule, error) {

	var subModule = domain.SubModule{}

	if err := moduleRepository.DB.First(&subModule, subModuleId).Error; err != nil {
		return nil, err
	}

	var sectionId uint

	if err := moduleRepository.DB.Model(&domain.Module{}).Where("id = ?", subModule.ModuleID).Select("section_id").Scan(&sectionId).Error; err != nil {
		return nil, err
	}

	var courseID uint

	if err := moduleRepository.DB.Model(&domain.Section{}).Where("id = ?", sectionId).Select("course_id").Scan(&courseID).Error; err != nil {
		return nil, err
	}
	
	var userIDCourse uint

	if err := moduleRepository.DB.Model(&domain.Course{}).Where("id = ?", courseID ).Select("user_id").Scan(&userIDCourse).Error; err != nil {
		return nil, err
	}

	if userIDCourse != userId && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}
	return &subModule, nil
}
