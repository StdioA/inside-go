package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stdioa/inside-go/db"
	"github.com/stdioa/inside-go/vm"
)

func GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "wrong id")
		return
	}
	post := db.GetExistPost(id, true)
	prevID, nextID := post.PrevAndNextID()
	postVM := vm.PostAPIVM{
		Success:    true,
		PreviousID: prevID,
		NextID:     nextID,
		Post:       vm.SerializePost(post),
	}
	c.JSON(http.StatusOK, postVM)
}

func ListComments(c *gin.Context) {
	var (
		postID int
		err    error
	)
	postIDS := c.Param("id")
	if postID, err = strconv.Atoi(postIDS); err != nil {
		c.String(http.StatusBadRequest, "ID format error")
		return
	}
	post := db.GetExistPost(postID, true)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, vm.Error("Post does not exist"))
		return
	}
	vm := vm.CommentAPIVM{
		Success:  true,
		ID:       postID,
		Comments: vm.SerializeComments(post.Comments),
	}
	c.JSON(http.StatusOK, vm)
}

func Archive(c *gin.Context) {
	idS, countS := c.Param("id"), c.Param("count")
	// 0 if id == ""
	id, _ := strconv.Atoi(idS)
	count, _ := strconv.Atoi(countS)
	if count == 0 {
		count = 6
	}
	posts := db.ListPosts(id, count)
	postVM := vm.ArchiveAPIVM{
		Success: true,
		Posts:   vm.SerializePosts(posts),
	}
	c.JSON(http.StatusOK, postVM)
}
