package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (categoryRepository *CategoryRepositoryImpl) GetAll(offset, limit int, search string) ([]domain.Category, int64, error) {
	category := []domain.Category{}
	var total int64

	query := categoryRepository.DB.Model(&category)

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("categories.name LIKE ? OR categories.description LIKE ?", s, s)
	}

	query = query.Limit(limit).Offset(offset)
	query.Find(&category).Count(&total)

	result := query.Find(&category)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return category, total, nil
}
