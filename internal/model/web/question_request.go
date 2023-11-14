package web

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

type QuizRequest struct {
	Module_id       uint                  `json:"module_id"`
	User_id         uint                  `json:"user_id"`
	Title           string                `gorm:"type:varchar(225)" json:"title"`
	Description     string                `gorm:"type:text" json:"description"`
	Durations       int                   `gorm:"type:int" json:"duration"`
	Questions       []domain.Questions    `json:"questions"`
}