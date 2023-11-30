package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (FaqRepository *FaqRepositoryImpl) Create(Faq *domain.Faqs) error {
	result := FaqRepository.DB.Create(&Faq)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
