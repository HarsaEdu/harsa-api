package repository

import (
	"fmt"
	"sort"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (DashboardRepository *DashboardRepositoryImpl) CountStudent() (int64,error) {

	var count int64
	
	if err := DashboardRepository.DB.Model(&domain.User{}).Where("role_id = ?", 3).Count(&count).Error; err != nil {
        return 0, err
    }

	return count, nil
}


func (DashboardRepository *DashboardRepositoryImpl) CountIntructure() (int64,error) {

	var count int64
	
	if err := DashboardRepository.DB.Model(&domain.User{}).Where("role_id = ?", 2).Count(&count).Error; err != nil {
        return 0, err
    }

	return count, nil
}

func (DashboardRepository *DashboardRepositoryImpl) CountCourse() (int64,error) {

	var count int64
	
	if err := DashboardRepository.DB.Model(&domain.Course{}).Count(&count).Error; err != nil {
        return 0, err
    }

	return count, nil
}

func (DashboardRepository *DashboardRepositoryImpl) CountInterest() ([]web.CountInterest, int64, error) {
	categories := []domain.Category{}

	if err := DashboardRepository.DB.Model(&domain.Category{}).Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	if categories == nil {
		return nil, 0, fmt.Errorf("category not found")
	}

	totalStudent, err := DashboardRepository.CountStudent()

	if err != nil {
		return nil, 0, err
	}

	var total int64
	var count int64
	totalCountInterest := []web.CountInterest{}

	for _, category := range categories {
		if err := DashboardRepository.DB.Model(&domain.UserInterest{}).Where("category_id = ?", category.ID).Count(&count).Error; err != nil {
			return nil, 0, err
		}

		countInterest := web.CountInterest{
			ID:    category.ID,
			Name:  category.Name,
			Count: count,
		}

		totalCountInterest = append(totalCountInterest, countInterest)
		total = total + count
	}

	// Sort the totalCountInterest slice by Count in descending order
	sort.Slice(totalCountInterest, func(i, j int) bool {
		return totalCountInterest[i].Count > totalCountInterest[j].Count
	})

	// Take the top 3 elements, and calculate the 'other' count
	var top3CountInterest []web.CountInterest
	var other int64

	if len(totalCountInterest) > 3 {
		top3CountInterest = totalCountInterest[:3]
		other = totalStudent - (top3CountInterest[0].Count + top3CountInterest[1].Count + top3CountInterest[2].Count)
	} else {
		top3CountInterest = totalCountInterest
		other = totalStudent - total
	}

	// Add 'other' countInterest to the result
	otherCountInterest := web.CountInterest{
		ID:    0,
		Name:  "other",
		Count: other,
	}
	top3CountInterest = append(top3CountInterest, otherCountInterest)

	return top3CountInterest, totalStudent, nil
}




