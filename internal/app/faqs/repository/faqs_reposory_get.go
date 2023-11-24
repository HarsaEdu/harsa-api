package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (faqsRepository *FaqRepositoryImpl) GetAll(offset, limit int, search string) ([]domain.Faqs, int64, error) {
	if offset < 0 || limit < 0 {
		return nil, 0, nil
	}

	faqs := []domain.Faqs{}
	var total int64

	query := faqsRepository.DB.Model(&faqs)

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("question LIKE ? OR answer LIKE ?", s, s)
	}

	query.Find(&faqs).Count(&total)

	query = query.Limit(limit).Offset(offset)

	result := query.Find(&faqs)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 {
		return nil, 0, nil
	}

	if offset >= int(total) {
		return nil, 0, nil
	}

	return faqs, total, nil
}
