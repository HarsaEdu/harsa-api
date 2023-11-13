package domain

type Auth struct {
	ID       uint   `gorm:"type:int;primarykey" json:"id"`
	Username string `gorm:"type:varchar(255)" json:"username"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	RoleName string `json:"role_name"`
}
