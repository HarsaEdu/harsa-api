package domain

type HistoryModule struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ModuleID    uint   `gorm:"type:int" json:"module_id"`
	UserID      uint   `gorm:"type:int" json:"user_id"`
	IsCompleted bool   `json:"is_completed"`
	Module      Module `gorm:"foreignKey:ModuleID;references:ID"`
}
