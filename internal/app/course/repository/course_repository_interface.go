package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(course *domain.Course) error
	GetAll() ([]domain.Course, error)
}

type CourseRepositoryImpl struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &CourseRepositoryImpl{
		DB: db,
	}
}