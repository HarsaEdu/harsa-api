package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (faqRepository *FaqRepositoryImpl) FindById(id int) (*domain.Faqs, error) {
	faq := domain.Faqs{}
	result := faqRepository.DB.First(&faq, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &faq, nil
}