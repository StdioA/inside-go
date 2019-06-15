package export

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stdioa/inside-go/db"
	"github.com/stdioa/inside-go/vm"
)

type ImportResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ExportVM struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	ID       uint         `json:"id"`
	Exist    bool         `json:"exist"`
	Content  string       `json:"content"`
	PubDate  int64        `json:"pub_date"`
	Comments []vm.Comment `json:"comments"`
}

func GenerateVM(posts []db.Post) *ExportVM {
	postVMs := make([]Post, 0, len(posts))
	for _, post := range posts {
		post.LoadComments()
		postVMs = append(postVMs, Post{
			post.ID,
			post.Exist,
			post.Content,
			post.CreatedAt.Unix(),
			vm.SerializeComments(post.Comments),
		})
	}
	return &ExportVM{
		postVMs,
	}
}

func SaveVMtoDB(posts []Post) {
	for _, postVM := range posts {
		post := db.Post{
			Model: gorm.Model{
				CreatedAt: time.Unix(postVM.PubDate, 0),
			},
			Content: postVM.Content,
			Exist:   postVM.Exist,
		}
		post.Save()

		for _, commentVM := range postVM.Comments {
			comment := db.Comment{
				PostID:  post.ID,
				Content: commentVM.Content,
				Author:  commentVM.Author,
			}
			comment.Save()
		}
	}
}
