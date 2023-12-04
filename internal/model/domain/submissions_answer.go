package domain

type SubmissionsAnswerDetail struct {
	ID      uint   `gorm:"type:int;primarykey" json:"id"`
	Title   string `json:"title"`
	Content string `gorm:"type:text" json:"content"`
	Status  string `json:"status"`
	Peserta string `json:"peserta"`
}
