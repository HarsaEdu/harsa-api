package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (historyQuizRepository *HistoryQuizRepositoryImpl) GetAllHistoryQuizByQuizId(offset, limit int, quizID uint,search string) ([]domain.HistoryQuiz, int64, error) {
	if offset < 0 || limit <= 0 {
		return nil, 0, fmt.Errorf("invalid offset or limit")
	}

	historyQuiz := []domain.HistoryQuiz{}
	var total int64

	query := historyQuizRepository.DB.Model(&historyQuiz).Where("quiz_id = ?", quizID)

	query = query.Preload("User.UserProfile").Preload("HistoryQuizAnswer.Options")

	if search != "" {
		query = query.Joins("JOIN user_profiles ON user_profiles.user_id = history_quizzes.user_id").
		Where("user_profiles.first_name LIKE ? OR user_profiles.last_name LIKE ?  ", "%"+search+"%", "%"+search+"%")
	}

	query.Count(&total)

	result := query.Limit(limit).Offset(offset).Find(&historyQuiz)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 || offset >= int(total) {
		return nil, 0, nil
	}

	return historyQuiz, total, nil
}

func (historyQuizRepository *HistoryQuizRepositoryImpl) GetMyHistoryQuizByQuizId(userID uint, quizID uint) (*domain.HistoryQuiz, error) {

	historyQuiz := domain.HistoryQuiz{}

	if err := historyQuizRepository.DB.Model(&historyQuiz).
		Where("user_id = ? AND quiz_id = ?", userID, quizID).
		Preload("User.UserProfile").Preload("HistoryQuizAnswer.Options").
		First(&historyQuiz).Error; err != nil {
		return nil,  err
	}
	
	return &historyQuiz,  nil
}

func (historyQuizRepository *HistoryQuizRepositoryImpl) GetById(id uint) (*domain.HistoryQuiz, error) {

	historyQuiz := domain.HistoryQuiz{}

	if err := historyQuizRepository.DB.Model(&domain.HistoryQuiz{}).Where("id = ?" ,id).
		Preload("User.UserProfile").Preload("HistoryQuizAnswer.Options").First(&historyQuiz).
		Error; err != nil {
		return nil,  err
	}
	
	return &historyQuiz,  nil
}

