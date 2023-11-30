package domain

import (
	"time"

	"gorm.io/gorm"
)

type Submissions struct {
	ID                uint                `gorm:"type:int;primarykey" json:"id"`
	ModuleID          uint                `gorm:"type:int" json:"module_id"`
	UserId            uint                `json:"user_id"`
	Title             string              `gorm:"type:varchar(225)" json:"title"`
	Description       string              `gorm:"type:text" json:"description"`
	SubmissionsAnswer []SubmissionsAnswer `gorm:"foreignKey:SubsmissionsID" json:"submissions_answer"`
	CreatedAt         time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt      `json:"deleted_at"`
}
