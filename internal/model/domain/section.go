package domain

import (
	"time"

	"gorm.io/gorm"
)

type Section struct {
	ID                 uint           `gorm:"type:int;primarykey" json:"id"`
	CourseID           uint           `gorm:"type:int" json:"user_id"`
	Title            string           `gorm:"type:varchar(255)" json:"title"`
	OrderBy              int          `gorm:"type:int" json:"order_by"`
	Modules            []Module      `gorm:"foreignKey:CourseID;references:ID"`
	CreatedAt          time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"delete_at"`
}