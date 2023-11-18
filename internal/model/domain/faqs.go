package domain

type Faqs struct {
	ID       uint   `gorm:"type:int;primarykey" json:"id"`
	Question string `json:"question" gorm:"type:varchar(255)"`
	Answer   string `json:"answer" gorm:"type:varchar(255)"`
}
