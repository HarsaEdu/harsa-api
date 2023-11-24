package domain

import "time"

type UserInterest struct {
	ID         uint        `gorm:"type:int;primarykey" json:"id"`
	ProfileID  uint        `gorm:"type:int" json:"profile_id"`
	CategoryID uint        `gorm:"type:int" json:"category_id"`
	CreatedAt  time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	Profile    UserProfile `gorm:"foreginKey:ProfileID;references:ID"`
	Category   Category    `gorm:"foreignKey:CategoryID;references:ID"`
}
