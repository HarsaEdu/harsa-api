package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type CourseTrackingRepository interface {
	Create(enrolled *domain.CourseTracking) error
	FindById(courseTrackingId uint) (*domain.CourseTracking,int64,int64 ,error)
}

type CourseTrackingRepositoryImpl struct {
	DB *gorm.DB
}

func NewCourseTrackingRepository(db *gorm.DB) CourseTrackingRepository {
	return &CourseTrackingRepositoryImpl{DB: db}
}
