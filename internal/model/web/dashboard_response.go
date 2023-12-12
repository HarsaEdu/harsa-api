package web

type CountInterest struct{
	ID uint `json:"id"`
	Name string `json:"name"`
	Count int64 `json:"count"`
}

type AllCountDashboard struct{
	CountInterest []CountInterest `json:"count_interest"`
	CountCourse   int64 `json:"count_course"`
	CountStudent  int64 `json:"count_student"`
	CountIntructure int64 `json:"count_intructure"`
}