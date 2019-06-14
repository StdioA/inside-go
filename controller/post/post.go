package post

import (
	"github.com/gin-gonic/gin"
)

// func Post(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "wrong id")
// 	}
// 	post := db.GetPost(id)
// 	if post.id > 0 {
// 		c.String(http.StatusOK, post.Content)
// 	} else {
// 		c.String(http.StatusNotFound, "Post not found.")
// 	}
// }

func Post(c *gin.Context) {
	c.File("static/mblog/html/post.html")
}

func Archive(c *gin.Context) {
	c.File("static/mblog/html/archive.html")
}
