package web

type QuizResponseForTracking struct {
	ID          uint           `json:"id"`
	Title       string         `json:"title"`
}

type HistoryQuizResponseMobile struct{
     ID uint `json:"id"`
	 IsComplete bool `json:"is_complete"`
	 Quiz QuizResponseForTracking `json:"quiz"`
}

