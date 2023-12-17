package domain

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID                 uint           `gorm:"type:int;primarykey" json:"id"`
	UserID             uint           `gorm:"type:int" json:"user_id"`
	Title              string         `gorm:"type:varchar(255)" json:"title"`
	Description        string         `gorm:"type:text" json:"description"`
	Rating             float32        `json:"rating"`
	ImageUrl           string         `gorm:"type:varchar(255);default:'https://res.cloudinary.com/dydgjkfgs/image/upload/v1702832191/harsa/courses/abd3is6os9joegg8imkj.jpg'" json:"image_url"`
	CreatedAt          time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"delete_at"`
	CategoryID         uint           `gorm:"type:int" json:"category_id"`
	User               User           `gorm:"foreignKey:UserID;references:ID"`
	Category           Category       `gorm:"foreignKey:CategoryID;references:ID"`
	Section            []Section      `gorm:"foreignKey:CourseID;references:ID"`
	Feedback           []Feedback    `gorm:"foreignKey:CourseID;references:ID"`
}

type CourseEntity struct {
	ID                 uint           `json:"id"`
	UserID             uint           `json:"user_id"`
	Title              string         `json:"title"`
	Description        string         `json:"description"`
	Enrolled           int            `json:"enrolled"`
	Rating             float32        `json:"rating"`
	ImageUrl           string         `json:"image_url"`
	CategoryID         uint           `json:"category_id"`
	CategoryName       string         `json:"category_name"`
	FirstName          string         `json:"first_name"`
	LastName           string         `json:"last_name"`
	RoleName           string         `json:"role_name"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	Job                string         `json:"job"`
}
