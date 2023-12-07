package domain

type SubsPlan struct {
	ID            uint    `gorm:"type:int;primarykey" json:"id"`
	Title         string  `json:"title" gorm:"type:varchar(255)"`
	Description   string  `json:"description" gorm:"type:varchar(255)"`
	Price         float64 `json:"price"`
	Image_url     string  `json:"image" form:"image" gorm:"default:'https://res.cloudinary.com/dydgjkfgs/image/upload/v1701933791/harsa/categories/hsjfi05cx3pcvaaepglk.png'"`
	Duration_days int     `json:"duration"`
}
