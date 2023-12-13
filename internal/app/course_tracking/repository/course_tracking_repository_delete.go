package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseTrackingrepository *CourseTrackingRepositoryImpl) Delete(tracking *domain.CourseTracking) error {
	result := courseTrackingrepository.DB.Delete(&tracking)
	if result.Error != nil {
		return result.Error
	}

	return nil
}