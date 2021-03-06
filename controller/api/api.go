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
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	post := db.GetExistPost(id, true)
	post.LoadComments()
	prevID, nextID := post.PrevAndNextID()
	postVM := vm.PostAPIVM{
		Success:    true,
		PreviousID: prevID,
		NextID:     nextID,
		Post:       vm.SerializePost(post),
	}
	c.JSON(http.StatusOK, postVM)
}

func UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	post := db.GetPost(id)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, vm.Error("Post does not exist"))
		return
	}
	// TODO
	type updateForm struct {
		Content string `json:"content"`
		Exist   bool   `json:"exist"`
	}
	form := updateForm{}
	c.BindJSON(&form)
	post.Update(form.Content, form.Exist)
	c.JSON(http.StatusOK, vm.SuccessVM{true})
}

func DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	post := db.GetPost(id)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, vm.Error("Post does not exist"))
		return
	}
	post.Update(post.Content, false)
	c.JSON(http.StatusOK, vm.SuccessVM{true})
}

func ListComments(c *gin.Context) {
	postIDS := c.Param("id")
	postID, err := strconv.Atoi(postIDS)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	post := db.GetExistPost(postID, true)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, vm.Error("Post does not exist"))
		return
	}
	post.LoadComments()
	vm := vm.CommentAPIVM{
		Success:  true,
		ID:       postID,
		Comments: vm.SerializeComments(post.Comments),
	}
	c.JSON(http.StatusOK, vm)
}

func CreateComment(c *gin.Context) {
	postIDS := c.Param("id")
	postID, err := strconv.Atoi(postIDS)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	post := db.GetExistPost(postID, true)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, vm.Error("Post does not exist"))
		return
	}
	_ = post.CreateComment(c.PostForm("author"), c.PostForm("content"))
	c.JSON(http.StatusOK, vm.SuccessVM{true})
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
