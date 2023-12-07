package domain

import "time"

type Subscription struct {
	ID      uint      `gorm:"type:int;primarykey" json:"id"`
	UserID  uint      `json:"user_id"`
	EndDate time.Time `json:"end_date"`
	User    User
}
