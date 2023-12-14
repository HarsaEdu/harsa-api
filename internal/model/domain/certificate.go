package domain

import "time"

type Certificate struct {
	ID        string    `gorm:"varchar(255);primarykey" json:"id"`
	UserID    uint      `gorm:"int" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	CourseID  uint      `gorm:"int" json:"course_id"`
	Course    Course    `gorm:"foreignKey:CourseID" json:"course"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	IssuedAt  string    `gorm:"varchar(255)" json:"issued_at"`
}