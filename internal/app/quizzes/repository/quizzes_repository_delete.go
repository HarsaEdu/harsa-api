package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (repository *QuizzesRepositoryImpl) Delete(quiz *domain.Quizzes) error {

	if err := repository.DB.Delete(&quiz).Error; err != nil {
		return  err
	}

	return nil
}