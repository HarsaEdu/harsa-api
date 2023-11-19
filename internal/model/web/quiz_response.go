package web

import "time"

type QuizResponse struct {
	Id          uint      `json:"id"`
	ModuleId    uint      `json:"module_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Durations   int       `json:"duration"`
	Created_at  time.Time `json:"created_at" `
	Updatad_at  time.Time `json:"updated_at" `
	Questions   any       `json:"questions"`
}

type GetAllQuizResponse struct {
	Id              uint      `json:"id"`
	ModuleId        uint      `json:"module_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Durations       int       `json:"duration"`
	NumberQuestions int       `json:"number_questions"`
}
