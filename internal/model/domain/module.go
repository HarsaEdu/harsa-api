package domain

import (
	"time"

	"gorm.io/gorm"
)

type SubModuleType string

const (
	SubModuleTypeVideo       SubModuleType = "video"
	SubModuleTypePPT         SubModuleType = "ppt"
)

type Module struct {
	ID          uint           `gorm:"type:int;primarykey" json:"id"`
	SectionID   uint           `gorm:"type:int" json:"section_id"`
	Title       string         `gorm:"type:varchar(255)" json:"title"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	OrderBy       int          `gorm:"type:int" json:"order_by"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"delete_at"`
	SubModules  []SubModule   `gorm:"foreignKey:ModuleID" json:"sub_modules"`
	Submissions []Submissions  `gorm:"foreignKey:ModuleID" json:"submissions"`
	Quizzes 	[]Quizzes 	   `gorm:"foreignKey:ModuleID" json:"quizzes"`
}

type SubModule struct {
	ID          uint           `gorm:"type:int;primarykey" json:"id"`
	ModuleID    uint           `gorm:"type:int" json:"module_id"`
	Title       string         `gorm:"type:varchar(255)" json:"title"`
	ContentUrl  string         `gorm:"type:varchar(255)" json:"content_url"`
	Type        SubModuleType  `gorm:"type:varchar(255)" json:"type"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"delete_at"`
}