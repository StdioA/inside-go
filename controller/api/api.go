package api

import (
	"fmt"
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
	}
	post := db.GetExistPost(id, true)
	prevID, nextID := post.PrevAndNextID()
	commentVM := make([]vm.Comment, 0)

	for _, cmt := range post.Comments {
		commentVM = append(commentVM, vm.Comment{
			Content: cmt.Content,
			Author:  cmt.Author,
		})
	}
	postVM := vm.PostAPIVM{
		Success:    true,
		PreviousID: prevID,
		NextID:     nextID,
		Post: vm.Post{
			post.ID,
			post.Content,
			post.CreatedAt,
			commentVM,
		},
	}
	c.JSON(http.StatusOK, postVM)
}

func ListComments(c *gin.Context) {

}

func Archive(c *gin.Context) {
	id, count := c.Param("id"), c.Param("count")
	// 都是 string
	fmt.Printf("%T %v\n", id, id)
	fmt.Printf("%T %v\n", count, count)

	c.Done()
}
