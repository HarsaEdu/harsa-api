package domain

import (
	"time"

	"gorm.io/gorm"
)

type Submissions struct {
	ID          uint           `gorm:"type:int;primarykey" json:"id"`
	ModuleID    uint           `json:"module_id"`
	Title       string         `gorm:"type:varchar(225)" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}