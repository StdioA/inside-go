package post

import (
	"net/http"
	"strconv"

	"github.com/stdioa/inside-go/db"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	post := db.LatestPost()
	if post.ID == 0 {
		c.Status(http.StatusNotFound)
		c.File("static/404.html")
		return
	}
	redirectURL := "/mblog/posts/" + strconv.Itoa(int(post.ID))
	c.Redirect(http.StatusFound, redirectURL)
}

func Post(c *gin.Context) {
	c.File("static/mblog/html/post.html")
}

func Archive(c *gin.Context) {
	c.File("static/mblog/html/archive.html")
}

func NewPostPage(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", gin.H{})
}

func NewPostHandler(c *gin.Context) {
	content := c.PostForm("content")
	if content == "" {
		c.Redirect(http.StatusFound, c.Request.URL.Path)
		return
	}
	post := db.CreatePost(content)
	redirectURL := "/mblog/posts/" + strconv.Itoa(int(post.ID))
	c.Redirect(http.StatusFound, redirectURL)
}

func EditPost(c *gin.Context) {
	var (
		postID int
		err    error
	)
	postIDS := c.Param("id")
	if postID, err = strconv.Atoi(postIDS); err != nil {
		c.Status(http.StatusNotFound)
		c.File("static/404.html")
		return
	}
	post := db.GetPost(postID)
	if post.ID == 0 {
		c.Status(http.StatusNotFound)
		c.File("static/404.html")
		return
	}
	post.LoadComments()
	prev, next := post.PrevAndNextID()
	c.HTML(http.StatusOK, "edit.html", gin.H{
		"post": post,
		"prev": prev,
		"next": next,
	})
}
