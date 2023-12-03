package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseTrackingrepository *CourseTrackingRepositoryImpl) Delete(id uint) error {
	result := courseTrackingrepository.DB.Where("id = ?", id).Delete(&domain.CourseTracking{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}