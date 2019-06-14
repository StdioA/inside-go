package vm

import (
	"time"
)

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

type ArchiveAPIVM struct {
	Success bool   `json:"success"`
	Posts   []Post `json:"posts"`
}

type CommentAPIVM struct {
	Success  bool      `json:"success"`
	ID       int       `json:"id"`
	Comments []Comment `json:"comments"`
}

type SuccessVM struct {
	Success bool `json:"success"`
}

type ErrorVM struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason"`
}

func Error(reason string) *ErrorVM {
	return &ErrorVM{
		Reason: reason,
	}
}
