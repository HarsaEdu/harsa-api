package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (historyQuizRepository *HistoryQuizRepositoryImpl) Answering(userId uint,answers []web.AnswersRequest, quizId uint) error {

	tx := historyQuizRepository.DB.Begin()

	historyQuiz := domain.HistoryQuiz{}
	
	historyQuiz.UserID = userId
	historyQuiz.QuizID = quizId

	if err := tx.Create(&historyQuiz).Error; err != nil {
		tx.Rollback()
		return err
	}

	historyAnswers := []domain.HistoryQuizAnswer{}
	var optionIDs []uint 

	for _, answ := range answers {
	optionIDs = append(optionIDs, answ.Option_id)
	historyAnswers = append(historyAnswers, domain.HistoryQuizAnswer{
			HistoryQuizID: historyQuiz.ID,
			QuestionID: answ.Question_id,
			OptionID: answ.Option_id,
		})
	}

	if err := tx.Create(&historyAnswers).Error; err != nil {
		return err
	}

	var count int64

	right := true

	if err := tx.Model(&domain.Options{}).Where("id IN (?) AND is_right = ?", optionIDs, right).Count(&count).Error; err != nil {
		tx.Rollback()
		return err
	}

	score := float64(count) / float64(len(answers)) * 100

	if err := tx.Model(&domain.HistoryQuiz{}).Where("id=?", historyQuiz.ID).Updates(&domain.HistoryQuiz{Score: float32(score), IsComplete: right}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}