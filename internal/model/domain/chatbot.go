package domain

import (
	"time"

	"gorm.io/gorm"
)

type UserChatTopic struct {
	ID        string         `gorm:"type:varchar(255);primarykey" json:"id"`
	UserID    uint           `gorm:"type:varchar(255)" json:"user_id"`
	Topic     string         `gorm:"type:varchar(255)" json:"topic"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
	User      User           `gorm:"foreignKey:UserID"`
}