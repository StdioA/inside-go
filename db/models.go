package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Post struct {
	gorm.Model
	Content  string `json:"content"`
	Exist    bool
	Comments []Comment `json:"comments"`
}

type Comment struct {
	gorm.Model
	PostID  uint   `gorm:"index,ForeignKey:PostID"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (post *Post) PrevAndNextID() (uint, uint) {
	id := post.ID
	if id == 0 {
		return 0, 0
	}
	var pPost, nPost Post
	db.Where("id<? AND exist=?", id, true).Order("id desc").First(&pPost)
	db.Where("id>? AND exist=?", id, true).Order("id").First(&nPost)
	return pPost.ID, nPost.ID
}

func (post *Post) LoadComments() {
	db.Model(post).Related(&(post.Comments))
}

func (post *Post) CreateComment(author, content string) *Comment {
	comment := Comment{
		PostID:  post.ID,
		Author:  author,
		Content: content,
	}
	db.Create(&comment)
	return &comment
}

func GetPost(id int) *Post {
	var post Post
	db.First(&post, id)
	return &post
}

func GetExistPost(id int, exists bool) *Post {
	var post Post
	db.Where("id=? AND exist=?", id, exists).First(&post)
	fmt.Println(post)
	return &post
}

func ListPosts(id, count int) []Post {
	var result []Post
	query := map[string]interface{}{
		"exist": true,
	}

	querySet := db.Where(query)
	if id > 0 {
		querySet = querySet.Where("id <= ?", id)
	}
	querySet.Order("id desc").Limit(count).Find(&result)
	return result
}

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "inside.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
}
