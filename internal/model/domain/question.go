package domain

import (
	"time"

	"gorm.io/gorm"
)

type Questions struct {
	ID        uint           `gorm:"type:int;primarykey" json:"id"`
	QuizID    uint           `json:"quiz_id"`
	Question  string         `gorm:"type:text" json:"question"`
	Options   []Options      `json:"options" gorm:"foreignKey:QuestionID"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
