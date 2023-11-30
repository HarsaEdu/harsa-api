package domain

import (
	"time"

	"gorm.io/gorm"
)

type Feedback struct {
	ID        uint           `gorm:"type:int;primarykey" json:"id"`
	UserID    uint           `gorm:"type:int" json:"user_id"`
	CourseID  uint           `gorm:"type:int" json:"course_id"`
	Rating    int            `gorm:"type:int" json:"rating"`
	Content   string         `gorm:"type:text" json:"content"`
	User      User           `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}
