package vm

import "time"

type Post struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"pub_date"`
	Comments  []Comment `json:"comments"`
}

type Comment struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

type PostAPIVM struct {
	Success    bool `json:"success"`
	PreviousID uint `json:"previous_id"`
	NextID     uint `json:"next_id"`
	Post       Post `json:"post"`
}
