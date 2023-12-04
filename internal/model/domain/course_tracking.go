package domain

import (
	"time"

	"gorm.io/gorm"
)

type StatusCourseTraking string

const (
	StatusCourseTrakingInProgress StatusCourseTraking = "in progress"
	statusCourseTrakingCompleted  StatusCourseTraking = "completed"
)

type CourseTracking struct {
	ID        uint                `gorm:"type:int;primarykey" json:"id"`
	UserID    uint                `gorm:"type:int" json:"user_id"`
	User      User                `gorm:"foreignKey:UserID;references:ID" json:"user"`
	CourseID  uint                `gorm:"type:int" json:"course_id"`
	Status    StatusCourseTraking `gorm:"type:varchar(255)" json:"status"`
	Course    Course              `gorm:"foreignKey:CourseID;references:ID" json:"course"`
	CreatedAt time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt      `gorm:"index" json:"delete_at"`
}


type HistorySubModule struct{
	ID uint `gorm:"type:int;primarykey" json:"id"`
	UserID uint `gorm:"type:int" json:"user_id"`
	SubModuleID uint `gorm:"type:int" json:"sub_module_id"`
	SubModule SubModule `gorm:"foreignKey:SubModuleID;references:ID" json:"sub_module"`
	IsComplete bool `gorm:"type:boolean" json:"is_complete"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

type StatusSubmissionAnswer string

const (
	StatusSubmissionAnswerAccepted StatusSubmissionAnswer = "accepted"
	StatusSubmissionAnswerPending  StatusSubmissionAnswer = "pending"
	StatusSubmissionAnswerRejected StatusSubmissionAnswer = "rejected"
	StatusSubmissionAnswerSubmitted StatusSubmissionAnswer = "submitted"
	StatusSubmissionAnswerNotSubmitted StatusSubmissionAnswer = "not submitted"
)

type SubmissionAnswer struct{
	ID uint `gorm:"type:int;primarykey" json:"id"`
	UserID uint `gorm:"type:int" json:"user_id"`
	SubmissionID uint `gorm:"type:int" json:"submission_id"`
	Submission Submissions `gorm:"foreignKey:SubmissionID;references:ID" json:"submission"`
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






