package web

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

type QuizRequest struct {
	ModuleId    uint               `json:"module_id" validate:"required"`
	UserId      uint               `json:"user_id" validate:"required"`
	Title       string             `gorm:"type:varchar(225)" json:"title" validate:"required"`
	Description string             `gorm:"type:text" json:"description" validate:"required"`
	Durations   int                `gorm:"type:int" json:"duration"`
	Questions   []domain.Questions `json:"questions" gorm:"foreignKey:QuizId"`
}
