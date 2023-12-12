package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (SubsPlanRepository *SubsPlanRepositoryImpl) GetAllActive(search string) ([]domain.SubsPlan, int64, error) {
	subsPlan := []domain.SubsPlan{}
	var total int64

	query := SubsPlanRepository.DB.Where("isActive = ?", true)

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("title LIKE ? OR price LIKE ? OR duration_days LIKE ?", s, s, s)
	}

	query = query.Find(&subsPlan).Count(&total)

	result := query.Find(&subsPlan)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 {
		return nil, 0, nil
	}

	return subsPlan, total, nil

}
