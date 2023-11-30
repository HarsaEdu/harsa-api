package repository


import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseTrackingrepository *CourseTrackingRepositoryImpl) Create(enrolled *domain.CourseTracking) error {

	if err := courseTrackingrepository.DB.Create(&enrolled).Error; err != nil {
		return err
	}

	return nil
}