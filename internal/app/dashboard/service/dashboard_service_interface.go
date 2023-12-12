package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/dashboard/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

type DashboardService interface {
	CountAll() (*web.AllCountDashboard, error)
}

type DashboardServiceImpl struct {
	DashboardRepository repository.DashboardRepository
}

func NewDashboardService(faqRepository repository.DashboardRepository) DashboardService {
	return &DashboardServiceImpl{DashboardRepository: faqRepository}

}