package domain
import (
	"time"

	"gorm.io/gorm"
)

type Options struct {
	ID              uint           `gorm:"type:int;primarykey" json:"id"`
	Question_id     uint           `json:"question_id"`
	Value           string         `json:"value" gorm:"type:text"`
	Is_right        bool           `json:"is_right"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeleteAt        gorm.DeletedAt `json:"delete_at"`
}