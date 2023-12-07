package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (faqsRepository *FaqRepositoryImpl) Delete(id int) error {

	result := faqsRepository.DB.Where("id = ?", id).Delete(&domain.Faqs{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
