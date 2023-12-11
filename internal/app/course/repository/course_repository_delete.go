package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseRepository *CourseRepositoryImpl) Delete(course *domain.Course) error {
	tx := courseRepository.DB.Begin()

	defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

	if err := tx.Delete(course).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("course_id = ?", course.ID).Delete(&domain.CourseTracking{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
