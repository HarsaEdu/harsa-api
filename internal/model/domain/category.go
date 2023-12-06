package domain

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `gorm:"type:int;primarykey" json:"id"`
	Name        string         `gorm:"type:varchar(255)" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Image_url   string         `gorm:"type:varchar(255)" json:"image" default:"https://placewaifu.com/image/200/200"`
	Icon        string         `gorm:"type:varchar(255)" json:"icon" default:"https://placewaifu.com/image/200/200"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeleteAt    gorm.DeletedAt `json:"delete_at"`
}
