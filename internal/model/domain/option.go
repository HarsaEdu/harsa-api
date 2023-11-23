package domain

import (
	"time"

	"gorm.io/gorm"
)

type Options struct {
	ID         uint           `gorm:"type:int;primarykey" json:"id"`
	QuestionId uint           `json:"question_id"`
	Value      string         `json:"value" gorm:"type:text"`
	IsRight    bool           `json:"is_right"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
