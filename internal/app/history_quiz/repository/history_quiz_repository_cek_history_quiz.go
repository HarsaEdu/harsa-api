package repository

import (

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (historyQuizRepository *HistoryQuizRepositoryImpl) Cek(userId uint, quizID uint) (*domain.HistoryQuiz, error) {

	historyQuiz := domain.HistoryQuiz{}

	if err := historyQuizRepository.DB.Model(&domain.HistoryQuiz{}).Where("user_id  = ? AND quiz_id = ?" ,userId,quizID).
		Preload("User.UserProfile").Preload("HistoryQuizAnswer.Options").First(&historyQuiz).
		Error; err != nil {
		return nil,  err
	}
	
	return &historyQuiz,  nil
}