package domain
import (
	"time"

	"gorm.io/gorm"
)

type Quizzes struct {
	ID              uint           `gorm:"type:int;primarykey" json:"id"`
	Module_id       uint           `json:"module_id"`
	User_id         uint           `json:"user_id"`
	Title           string         `gorm:"type:varchar(225)" json:"title"`
	Description     string         `gorm:"type:text" json:"description"`
	Durations       int            `gorm:"type:int" json:"duration"`
	Questions       []Questions    `json:"questions" gorm:"foreignKey:Quiz_id"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"delete_at"`
}