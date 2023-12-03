package web

type QuizResponseForTracking struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	HistoryQuizID uint   `json:"history_quiz_id"`
	IsComplete    bool   `json:"is_complete"`
}

type HistoryQuizResponseMobile struct {
	ID         uint                    `json:"id"`
	IsComplete bool                    `json:"is_complete"`
	Quiz       QuizResponseForTracking `json:"quiz"`
}

type HistoryQuizTracking struct {
	HistoryQuizID uint `json:"history_quizz_id"`
	IsComplete    bool `json:"is_complete"`
}

type HistoryQuizIDTracking struct {
	ID          uint                `json:"id"`
	Title       string              `json:"title"`
	Progress    float32             `json:"progress"`
	Description string              `json:"description"`
	Order       int                 `json:"order"`
	HistoryQuiz HistoryQuizTracking `json:"history_quizz"`
	Quizz       QuizResponse        `json:"quizz"`
}

type HistoryQuizTrackingResponse struct {
	ID          uint                    `json:"id"`
	Title       string                  `json:"title"`
	Progress    float32                 `json:"progress"`
	Description string                  `json:"description"`
	Order       int                     `json:"order"`
	HistoryQuiz QuizResponseForTracking `json:"history_quizz"`
	Quizz       QuizResponse            `json:"quizz"`
}
