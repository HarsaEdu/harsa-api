package web

import "time"


type FeedBackResponseForTracking struct {
	ID          uint           `json:"id"`
	Rating      int            `json:"rating"`
	Content     string         `json:"content"`
	User        UserForFeedBack`json:"user"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type UserForFeedBack struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
}