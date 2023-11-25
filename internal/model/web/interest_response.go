package web

type InterestResponse struct {
	CourseID   uint   `json:"course_id"`
	Title      string `json:"title"`
	Rating     int    `json:"rating"`
	ImageUrl   string `json:"image_url"`
	Instructor string `json:"intructor"`
}
