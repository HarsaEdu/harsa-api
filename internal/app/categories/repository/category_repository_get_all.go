package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (categoryRepository *CategoryRepositoryImpl) GetAll(offset, limit int, search string) ([]domain.Category, int64, error) {
	category := []domain.Category{}
	var total int64

	if limit < 0 || offset < 0 {
		return nil, 0, nil

	}

	query := categoryRepository.DB.Model(&category)

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("categories.name LIKE ? OR categories.description LIKE ?", s, s)
	}

	query.Find(&category).Count(&total)
	query = query.Limit(limit).Offset(offset)

	result := query.Find(&category)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return category, total, nil
}
