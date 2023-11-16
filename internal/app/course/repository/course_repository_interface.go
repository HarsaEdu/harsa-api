package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(course *domain.Course) error
	GetAll() ([]domain.Course, error)
	GetById(id uint) (*domain.Course, error)
	Update(id uint, course *domain.Course) error
	UpdateImage(course *domain.Course) error
	Delete(id uint) error
}

type CourseRepositoryImpl struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &CourseRepositoryImpl{
		DB: db,
	}
}