package repository

import (

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (repository *QuizzesRepositoryImpl) FindById(quizId uint) (*domain.Quizzes, error) {

	var quiz = domain.Quizzes{}

	if err := repository.DB.Preload("Questions.Options").First(&quiz, quizId).Error; err != nil {
		return nil, err
	}
	
	return &quiz, nil
}