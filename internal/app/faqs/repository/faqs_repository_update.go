package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (faqsRepository *FaqRepositoryImpl) Update(Faq *domain.Faqs, id int) error {
	result := faqsRepository.DB.Where("id=?", id).Updates(&domain.Faqs{Question: Faq.Question, Answer: Faq.Answer})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
