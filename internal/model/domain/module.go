package domain

type Modules struct {
	ID          uint     `gorm:"type:int;primarykey" json:"id"`
	CourseId      uint     `gorm:"type:int" json:"course_id"`
}