package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (repository *QuizzesRepositoryImpl) GetAll(moduleId uint, offset, limit int, search string) ([]domain.Quizzes, int64, error) {
    
    if offset < 0 || limit < 0 {
        return nil, 0, nil
    }
    
    quiz:= []domain.Quizzes{}
    var total int64

    query := repository.DB.Model(&quiz)

    if search != "" {
        s := "%" + search + "%"
        query = query.Where("title LIKE ? AND module_id = ?", s, moduleId)
    }

    query.Find(&quiz).Count(&total)

    query = query.Limit(limit).Offset(offset)

    result := query.Preload("Questions").Find(&quiz)

    if result.Error != nil {
        return nil, 0, result.Error
    }

    if total == 0 {
        return nil, 0, nil
    }

    if offset >= int(total) {
        return nil, 0, nil
    }

    return quiz, total, nil
}