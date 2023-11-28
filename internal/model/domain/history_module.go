package domain

import (
	"time"

	"gorm.io/gorm"
)

type HistoryModule struct {
	ID        uint           `gorm:"type:int;primarykey" json:"id"`
	UserID    uint           `gorm:"type:int" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}
