package web

import "time"

type QuizResponse struct {
	Id          uint      `json:"id"`
	ModuleID    uint      `json:"module_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Durations   int       `json:"duration"`
	Created_at  time.Time `json:"created_at" `
	Updatad_at  time.Time `json:"updated_at" `
	Questions   any       `json:"questions"`
}

type GetAllQuizResponse struct {
	Id              uint      `json:"id"`
	ModuleID        uint      `json:"module_id"`
	Title           string    `json:"title"`
	CourseTitle          string    `json:"course_title"`
	Description     string    `json:"description"`
	Durations       int       `json:"duration"`
	NumberQuestions int       `json:"number_questions"`
}

type QuizResponseModule struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Durations   int       `json:"duration"`
}

type QuizResponseModuleMobile struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
}
