package web

type QuizResponseForTracking struct {
	ID          uint           `json:"id"`
	Title       string         `json:"title"`
	HistoryQuizID uint           `json:"history_quiz_id"`
	IsComplete  bool           `json:"is_complete"` 
}

type HistoryQuizResponseMobile struct{
     ID uint `json:"id"`
	 IsComplete bool `json:"is_complete"`
	 Quiz QuizResponseForTracking `json:"quiz"`
}

