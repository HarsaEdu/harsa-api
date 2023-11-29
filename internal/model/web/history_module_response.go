package web


type ModuleResponseForTracking struct {
	ID          uint           `json:"id"`
	Section     string         `json:"section"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Order       int            `json:"order"`
}

type HistoryModuleResponseMobile struct{
     ID uint `json:"id"`
	 Progress float32 `json:"progress"`
	 Module ModuleResponseForTracking `json:"module"`
	 HistorySubModule []HistorySubModuleResponseMobile `json:"history_sub_module"`
	 SubmissionAnswer []SubmissionAnswerResponseMobile `json:"submission_answer"`
	 HistoryQuiz      []HistoryQuizResponseMobile `json:"history_quiz"`
}