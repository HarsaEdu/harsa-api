package web

type AnswersRequest struct{
	Question_id uint `json:"question_id" validate:"required"`
	Option_id uint `json:"option_id" validate:"required"`
}