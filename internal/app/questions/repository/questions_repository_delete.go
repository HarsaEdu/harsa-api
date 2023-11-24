package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (repositoryQuestions *QuestionsRepositoryImpl) Delete(question *domain.Questions) error {

	if err := repositoryQuestions.DB.Delete(&question).Error; err != nil {
		return  err
	}

	return nil
}