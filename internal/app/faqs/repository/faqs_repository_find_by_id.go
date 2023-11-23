package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (faqsRepository *FaqRepositoryImpl) FindById(id int) (*domain.Faqs, error) {
	faq := &domain.Faqs{}

	result := faqsRepository.DB.First(&faq)
	if result.Error != nil {
		return nil, result.Error
	}

	return faq, nil
}
