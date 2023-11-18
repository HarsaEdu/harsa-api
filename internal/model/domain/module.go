package domain

import (
	"time"

	"gorm.io/gorm"
)

type SubModuleType string

const (
	SubModuleTypeVideo       SubModuleType = "video"
	SubModuleTypePPT         SubModuleType = "ppt"
	SubModuleTypePDF         SubModuleType = "pdf"
	SubModuleTypeText        SubModuleType = "text"
	SubModuleTypeHtmlElement SubModuleType = "html_element"
)

type Module struct {
	ID          uint           `gorm:"type:int;primarykey" json:"id"`
	CourseID    uint           `gorm:"type:int" json:"course_id"`
	Title       string         `gorm:"type:varchar(255)" json:"title"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	Order       int            `gorm:"type:int" json:"order"`
	Type        string         `gorm:"type:varchar(255)" json:"type"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"delete_at"`
	SubModules  []*SubModule   `gorm:"foreignKey:ModuleID" json:"sub_modules"`
}

type SubModule struct {
	ID          uint           `gorm:"type:int;primarykey" json:"id"`
	ModuleID    uint           `gorm:"type:int" json:"module_id"`
	Title       string         `gorm:"type:varchar(255)" json:"title"`
	ContentUrl  string         `gorm:"type:varchar(255)" json:"content_url"`
	ContentBody string         `gorm:"type:text" json:"content_body"`
	Type        SubModuleType  `gorm:"type:varchar(255)" json:"type"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"delete_at"`
}