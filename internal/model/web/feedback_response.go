package web

import "time"

type FeedBackResponseForTracking struct {
	ID        uint            `json:"id"`
	Rating    int             `json:"rating"`
	Content   string          `json:"content"`
	CreatedAt time.Time       `json:"created_at"`
	User      UserForFeedBack `json:"user"`
}

type UserForFeedBack struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}