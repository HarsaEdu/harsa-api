package web


type FeedBackResponseForTracking struct {
	ID          uint           `json:"id"`
	Rating      int            `json:"rating"`
	Content     string         `json:"content"`
	User        UserForFeedBack `json:"user"`
}

type UserForFeedBack struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}