package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (repository *QuizzesRepositoryImpl) GetAll(moduleId uint, offset, limit int, search string) ([]domain.Quizzes, int64, string, error) {
    
    if offset < 0 || limit < 0 {
        return nil, 0, "",fmt.Errorf("nvalid offset and limit")
    }
    
    quiz:= []domain.Quizzes{}
    var total int64

    query := repository.DB.Model(&quiz)

    if search != "" {
        s := "%" + search + "%"
        query = query.Where("title LIKE ?", s)
    }

    query.Where("module_id = ?", moduleId).Find(&quiz).Count(&total)

    query = query.Limit(limit).Offset(offset)

    result := query.Preload("Questions").Find(&quiz)

    if result.Error != nil {
        return nil, 0, "",result.Error
    }

    if offset >= int(total) {
        return nil, 0, "",nil
    }

    var courseTitle string

    if err := repository.DB.Model(&domain.Module{}).Where("modules.id = ?", moduleId).Select("courses.title as coursesTitle").
    Joins("JOIN sections ON sections.id = modules.section_id").
	Joins("JOIN courses ON courses.id = sections.course_id").
    Scan(&courseTitle).
    Error; err != nil {
		return  nil,0,"",err
	}

    return quiz, total, courseTitle,nil
}