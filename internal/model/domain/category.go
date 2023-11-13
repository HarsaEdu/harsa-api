package domain

type Category struct {
	ID          uint   `gorm:"type:int;primarykey" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Image_url   string `gorm:"type:varchar(255)" json:"image"`
}
