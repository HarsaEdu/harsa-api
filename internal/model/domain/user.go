package domain

import "gorm.io/gorm"

type User struct {
	ID        uint           `gorm:"type:int;primarykey" json:"id"`
	Email     string         `gorm:"type:varchar(255)" json:"email"`
	Username  string         `gorm:"type:varchar(255)" json:"username"`
	Password  string         `gorm:"type:varchar(255)" json:"password"`
	RoleID    uint           `gorm:"type:int;primarykey" json:"role_id"`
	CreatedAt int64          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime" json:"updated_at"`
	DeleteAt  gorm.DeletedAt `gorm:"autoUpdateTime" json:"delete_at"`
}
