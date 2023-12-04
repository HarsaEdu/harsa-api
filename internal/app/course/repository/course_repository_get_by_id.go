package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

func (courseRepository *CourseRepositoryImpl) GetById(id uint) (*domain.Course, error) {
    course := domain.Course{} 

    if err := courseRepository.DB.Preload("Section", func(db *gorm.DB) *gorm.DB {
        return db.Order("order_by ASC, id ASC")
    }).Preload("Section.Modules", func(db *gorm.DB) *gorm.DB {
        return db.Order("modules.order_by ASC, modules.id ASC")
    }).First(&course, id).Error; err != nil {
        return nil, err
    }

    return &course, nil
}



func (courseRepository *CourseRepositoryImpl) GetByIdMobile(id uint) (*domain.Course, int64, int64,error) {
	course := &domain.Course{}

if err := courseRepository.DB.Where("id = ?", id).
	Preload("User.UserProfile").
	Preload("Section", func(db *gorm.DB) *gorm.DB {
        return db.Order("order_by ASC, id ASC")
    }).Preload("Section.Modules", func(db *gorm.DB) *gorm.DB {
        return db.Order("modules.order_by ASC, modules.id ASC")
    }).
	Preload("Feedback.User.UserProfile", func(db *gorm.DB) *gorm.DB {
		return db.Limit(5)
	}).
	Find(&course).Error; err != nil {
	return nil, 0,0, err
    }

	var countModule int64

	for _, section := range course.Section {
		countModule = countModule + int64(len(section.Modules))
	}

	var countEnroled int64

	if err := courseRepository.DB.Model(&domain.CourseTracking{}).Where("course_id  = ?", id).Count(&countEnroled).Error; err != nil {
		return nil, 0,0, err
	}

	return course, countModule, countEnroled, nil
}