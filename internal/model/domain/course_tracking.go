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

type CourseTracking struct {
	ID        uint                `gorm:"type:int;primarykey" json:"id"`
	UserID    uint                `gorm:"type:int" json:"user_id"`
	User      User                `gorm:"foreignKey:UserID;references:ID" json:"user"`
	CourseID  uint                `gorm:"type:int" json:"course_id"`
	Progress  float32             `gorm:"type:float" json:"progress"`
	Status    StatusCourseTraking `gorm:"type:varchar(255)" json:"status"`
	Course    Course              `gorm:"foreignKey:CourseID;references:ID" json:"course"`
	HistoryModule []HistoryModule    `gorm:"foreignKey:CourseTrakingID" json:"history_module"`
	CreatedAt time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt      `gorm:"index" json:"delete_at"`
}

type HistoryModule struct{
	ID uint `gorm:"type:int;primarykey" json:"id"`
	ModuleID uint `gorm:"type:int" json:"module_id"`
	Module Module `gorm:"foreignKey:ModuleID;references:ID" json:"module"`
	CourseTrakingID uint `gorm:"type:int" json:"course_traking_id"`
	HistorySubModule []HistorySubModule `gorm:"foreignKey:HistoryModuleID" json:"history_sub_module"`
	SubmissionAnswer []SubmissionAnswer `gorm:"foreignKey:HistoryModuleID" json:"submission_answer"`
	HistoryQuiz []HistoryQuiz `gorm:"foreignKey:HistoryModuleID" json:"history_quiz"`
	Progress float32 `gorm:"type:float" json:"progress"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

type HistorySubModule struct{
	ID uint `gorm:"type:int;primarykey" json:"id"`
	SubModuleID uint `gorm:"type:int" json:"sub_module_id"`
	SubModule SubModule `gorm:"foreignKey:SubModuleID;references:ID" json:"sub_module"`
	HistoryModuleID uint `gorm:"type:int" json:"history_module_id"`
	IsComplete bool `gorm:"type:boolean" json:"is_complete"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

type StatusSubmissionAnswer string

const (
	StatusSubmissionAnswerAccepted StatusSubmissionAnswer = "accepted"
	StatusSubmissionAnswerPending  StatusSubmissionAnswer = "pending"
	StatusSubmissionAnswerReview   StatusSubmissionAnswer = "review"
	StatusSubmissionAnswerRejected StatusSubmissionAnswer = "rejected"
	StatusSubmissionAnswerNotSubmitted StatusSubmissionAnswer = "not_submitted"
)

type SubmissionAnswer struct{
	ID uint `gorm:"type:int;primarykey" json:"id"`
	UserID uint `gorm:"type:int" json:"user_id"`
	SubmissionID uint `gorm:"type:int" json:"submission_id"`
	Submission Submissions `gorm:"foreignKey:SubmissionID;references:ID" json:"submission"`
	HistoryModuleID uint `gorm:"type:int" json:"history_module_id"`
	Content string `gorm:"type:text" json:"content"`
	SubmittedUrl string `gorm:"type:text" json:"submitted_url"`
	Status  StatusSubmissionAnswer `gorm:"type:varchar(255)" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

type HistoryQuiz struct{
	ID uint `gorm:"type:int;primarykey" json:"id"`
	UserID uint `gorm:"type:int" json:"user_id"`
	User User `gorm:"foreignKey:UserID" json:"user"`
	QuizID uint `gorm:"type:int" json:"quiz_id"`
	Quiz Quizzes `gorm:"foreignKey:QuizID" json:"quiz"`
	HistoryModuleID uint `gorm:"type:int" json:"history_module_id"`
	HistoryQuizAnswer []HistoryQuizAnswer `gorm:"foreignKey:HistoryQuizID" json:"history_quiz_answer"`
	Score float32 `gorm:"type:float" json:"score"`
	IsComplete bool `gorm:"type:boolean" json:"is_complete"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

type HistoryQuizAnswer struct{
	ID uint `gorm:"type:int;primarykey" json:"id"`
	HistoryQuizID uint `gorm:"type:int" json:"history_quiz_id"`
	QuestionID uint `gorm:"type:int" json:"question_id"`
	OptionID uint `gorm:"type:int" json:"option_id"`
	Options Options `gorm:"foreignKey:OptionID" json:"options"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}





