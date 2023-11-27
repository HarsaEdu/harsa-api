package domain

type HistorySubModule struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	SubModuleID uint      `gorm:"type:int" json:"module_id"`
	UserID      uint      `gorm:"type:int" json:"user_id"`
	IsCompleted bool      `json:"is_completed"`
	SubModule   SubModule `gorm:"foreignKey:SubModuleID;references:ID"`
}
