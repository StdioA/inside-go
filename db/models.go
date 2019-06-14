package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Post struct {
	gorm.Model
	Content  string
	Exist    bool
	Comments []Comment
}

type Comment struct {
	gorm.Model
	PostID  uint `gorm:"index,FOREIGNKEY"`
	Content string
	Author  string
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

func GetPost(id int) *Post {
	var post Post
	db.First(&post, id)
	return &post
}

func GetExistPost(id int, exists bool) *Post {
	var post Post
	db.Where("id=? AND exist=?", id, exists).First(&post)
	db.Model(&post).Related(&(post.Comments))
	return &post
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
