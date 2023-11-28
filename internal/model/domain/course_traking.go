package domain

import (
	"time"

	"gorm.io/gorm"
)

type StatusCourseTraking string

const (
	StatusCourseTrakingInProgress StatusCourseTraking = "in_progress"
	statusCourseTrakingCompleted  StatusCourseTraking = "completed"
)

type CourseTraking struct {
	ID        uint                `gorm:"type:int;primarykey" json:"id"`
	UserID    uint                `gorm:"type:int" json:"user_id"`
	CourseID  uint                `gorm:"type:int" json:"course_id"`
	Progress  float32             `gorm:"type:float" json:"progress"`
	Status    StatusCourseTraking `gorm:"type:varchar(255)" json:"status"`
	CreatedAt time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt      `gorm:"index" json:"delete_at"`
}

