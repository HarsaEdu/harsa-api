package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (SubsPlanRepository *SubsPlanRepositoryImpl) GetAllActive(offset, limit int, search string) ([]domain.SubsPlan, int64, error) {
	subsPlan := []domain.SubsPlan{}
	var total int64

	query := SubsPlanRepository.DB.Where("is_active = ?", true)

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("title LIKE ? OR price LIKE ? OR duration_days LIKE ?", s, s, s)
	}

	query = query.Find(&subsPlan).Count(&total)

	query = query.Offset(offset).Limit(limit)

	result := query.Order("duration_days ASC").Find(&subsPlan)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 {
		return nil, 0, nil
	}

	if offset >= int(total) {
		return subsPlan, 0, nil
	}

	return subsPlan, total, nil

}
