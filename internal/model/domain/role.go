package domain

type Role struct {
	ID   uint   `gorm:"type:int;primarykey" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
}
