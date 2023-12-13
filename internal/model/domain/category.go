package domain

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `gorm:"type:int;primarykey" json:"id"`
	Name        string         `gorm:"type:varchar(255)" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Image_url   string         `gorm:"type:varchar(255);default:'https://api.lorem.space/image/dashboard?w=200&h=200'" json:"image"`
	Icon        string         `gorm:"type:varchar(255);default:'https://res.cloudinary.com/dydgjkfgs/image/upload/v1701967969/harsa/categories/kfzjv3btrgydvl5l9ouv.png'" json:"icon"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeleteAt    gorm.DeletedAt `json:"delete_at"`
}
