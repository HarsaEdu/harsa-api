package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"gorm.io/gorm"
)

type DashboardRepository interface {
	CountStudent() (int64,error)
	CountIntructure() (int64,error)
	CountCourse() (int64,error)
	CountInterest() ([]web.CountInterest, int64, error)
}

type DashboardRepositoryImpl struct {
	DB *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &DashboardRepositoryImpl{DB: db}
}