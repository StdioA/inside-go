package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stdioa/inside-go/controller/api"
	"github.com/stdioa/inside-go/controller/post"
)

func Register(router *gin.Engine) {
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", post.Index)
	router.GET("/archive", post.Archive)
	mblog := router.Group("/mblog")
	mblog.GET("/new", post.NewPostPage)
	mblog.POST("/new", post.NewPostHandler)
	mblog.GET("/posts/:id", post.Post)
	mblog.GET("/posts/:id/edit", post.EditPost)

	apiGroup := router.Group("/api")
	apiGroup.GET("/posts/:id", api.GetPost)
	apiGroup.PUT("/posts/:id", api.UpdatePost)
	apiGroup.DELETE("/posts/:id", api.DeletePost)
	apiGroup.GET("/comments/:id", api.ListComments)
	apiGroup.POST("/comments/:id", api.CreateComment)

	apiGroup.GET("/archive", api.Archive)
	apiGroup.GET("/archive/:id", api.Archive)
	apiGroup.GET("/archive/:id/counts/:count", api.Archive)
}
