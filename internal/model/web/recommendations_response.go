package web

type GetRecommendation struct {
	CourseId               uint    `json:"course_id"`
	CourseTitle            string  `json:"course_title"`
	CourseDescription      string  `json:"course_description"`
	CourseImage            string  `json:"course_image"`
	InstructorName         string  `json:"instructor_name"`
	IsInInterestCategories bool    `json:"is_in_interest_categories"`
	PredictedRating        float32 `json:"predicted_rating"`
}

type GetRecommendationsResponse struct {
	Recommendations []GetRecommendation `json:"recommendations"`
}