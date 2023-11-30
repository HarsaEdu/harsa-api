package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (repository *QuizzesRepositoryImpl) Create(newQuiz *domain.Quizzes) error {

	if err := repository.DB.Create(&newQuiz).Error; err != nil {
		return err
	}

	return nil
}