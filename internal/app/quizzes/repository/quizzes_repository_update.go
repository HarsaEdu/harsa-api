package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (repository *QuizzesRepositoryImpl) Update(UpdateQuiz *domain.Quizzes, quizExist *domain.Quizzes) error {

	tx := repository.DB.Begin()

	quizExist.ModuleId = UpdateQuiz.ModuleId
	quizExist.Title = UpdateQuiz.Title
	quizExist.Description = UpdateQuiz.Description
	quizExist.Durations = UpdateQuiz.Durations
    quizExist.Questions = UpdateQuiz.Questions

	if err := tx.Save(&quizExist).Error; err != nil {
		repository.DB.Rollback()
		return err
	}

	for _, updatedQuestion := range UpdateQuiz.Questions {
        if err := tx.Save(&updatedQuestion).Error; err != nil {
            tx.Rollback()
            return err
        }

        for _, updatedOption := range updatedQuestion.Options {
            if err := tx.Save(&updatedOption).Error; err != nil {
                tx.Rollback()
                return err
            }
        }
    }

    tx.Commit()

	return nil
}