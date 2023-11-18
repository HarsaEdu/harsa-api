package web

type GetProfileResponse struct {
	Class       string `gorm:"type:varchar(50)" json:"class"`
	ImageURL    string `gorm:"type:varchar(255)" json:"image_url"`
	FirstName   string `gorm:"type:varchar(50)" json:"first_name"`
	LastName    string `gorm:"type:varchar(50)" json:"last_name"`
	DateOfBirth int64  `gorm:"type:int" json:"date_of_birth"`
	Bio         string `gorm:"type:varchar(255)" json:"bio"`
	Gender      string `gorm:"type:enum('f','m')" json:"gender"`
	PhoneNumber string `gorm:"type:varchar(20)" json:"phone_number"`
	City        string `gorm:"type:varchar(20)" json:"city" form:"city"`
	Address     string `gorm:"type:varchar(255)" json:"address"`
	Job         string `gorm:"type:varchar(20)" json:"job"`
}
