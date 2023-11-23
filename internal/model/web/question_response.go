package web

type QuestionsResForQuiz struct{
	Id              uint `json:"id"`
	Question        string `json:"question" `
	Options         []OptionsResForQuestion	`json:"options"`
}

type QuestionsResForQuizMobile struct{
	Id              uint `json:"id"`
	Question        string `json:"question" `
	Options         []OptionsResForQuestionMobile `json:"options"`
}