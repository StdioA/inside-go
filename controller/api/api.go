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
	}
	post := db.GetExistPost(id, true)
	prevID, nextID := post.PrevAndNextID()
	postVM := vm.PostAPIVM{
		Success:    true,
		PreviousID: prevID,
		NextID:     nextID,
		Post: vm.Post{
			post.ID,
			post.Content,
			post.CreatedAt,
			post.Comments,
		},
	}
	c.JSON(http.StatusOK, postVM)
}

func ListComments(c *gin.Context) {

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
	postList := make([]vm.Post, 0, len(posts))
	for _, post := range posts {
		postList = append(postList, vm.Post{
			ID:        post.ID,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
		})
	}
	postVM := vm.ArchiveAPIVM{
		Success: true,
		Posts:   postList,
	}
	c.JSON(http.StatusOK, postVM)
}
