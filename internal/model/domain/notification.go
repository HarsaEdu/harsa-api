package domain

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID        uint           `gorm:"type:int;primarykey" json:"id"`
	UserID    uint           `json:"user_id"`
	Title     string         `gorm:"type:varchar(255)" json:"title"`
	Content   string         `json:"content"`
	Sender    string         `gorm:"default:system" json:"sender"`
	IsRead    bool           `gorm:"default:false" json:"is_read"`
	IsArsip   bool           `gorm:"default:false" json:"is_arsip"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at"`
}
