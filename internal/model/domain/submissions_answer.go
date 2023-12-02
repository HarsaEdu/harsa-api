package domain

type Status string

const (
	Accepted  Status = "Accepted"
	Pending   Status = "Pending"
	Review    Status = "Review"
	Rejected  Status = "Rejected"
	NotSubmit Status = "NotSubmit"
)

type SubmissionsAnswer struct {
	ID             uint   `gorm:"type:int;primarykey" json:"id"`
	SubsmissionsID uint   `json:"subsmission_id" `
	Content        string `gorm:"type:text" json:"content"`
	Submited_url   string `gorm:"type:text" json:"submited_url"`
	Status         Status `json:"status" gorm:"type:varchar(255)"`
	Feedback       string `json:"feedback"`
}
