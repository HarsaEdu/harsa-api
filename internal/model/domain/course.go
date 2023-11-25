package domain

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID          uint           `gorm:"type:int;primarykey" json:"id"`
	UserID      uint           `gorm:"type:int" json:"user_id"`
	Title       string         `gorm:"type:varchar(255)" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Enrolled    int            `gorm:"type:int" json:"enrolled"`
	Rating      int            `gorm:"type:int" json:"rating"`
	ImageUrl    string         `gorm:"type:varchar(255)" json:"image_url"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"delete_at"`
	CategoryID  uint           `gorm:"type:int" json:"category_id"`
	User        UserProfile    `gorm:"foreignKey:UserID;references:UserID"`
	Category    Category       `gorm:"foreignKey:CategoryID;references:ID"`
	Modules     []*Module      `gorm:"foreignKey:CourseID;references:ID"`
}
