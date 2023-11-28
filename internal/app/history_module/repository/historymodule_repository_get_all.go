package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (historyModuleRepository *HistoryModuleRepositoryImpl) GetAll(offset, limit int, search string) ([]domain.HistoryModule, int64, error) {

	if offset < 0 || limit < 0 {
		return nil, 0, nil
	}

	historyModule := []domain.HistoryModule{}
	var total int64

	query := historyModuleRepository.DB.Model(&historyModule)

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("UserID LIKE ?", s)
	}

	query.Find(&historyModule).Count(&total)

	query = query.Limit(limit).Offset(offset)

	result := query.Find(&historyModule)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 {
		return nil, 0, nil
	}

	if offset >= int(total) {
		return nil, 0, nil
	}

	return historyModule, total, nil
}
