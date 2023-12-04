package domain

type SubmissionsAnswerDetail struct {
	ID      uint   `json:"id"`
	FirstName   string `json:"first_name"`
	LastName   string `json:"last_name"`
	Status  string `json:"status"`
}
