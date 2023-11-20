package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (faqsRepository *FaqRepositoryImpl) Delete(id int) error {

	ifExist, _ := faqsRepository.FindById(id)
	if ifExist == nil {
		return fmt.Errorf("faqs not found")
	}

	result := faqsRepository.DB.Where("id = ?", id).Delete(&domain.Faqs{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
