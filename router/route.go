package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stdioa/inside-go/controller/api"
	"github.com/stdioa/inside-go/controller/post"
)

func Register(router *gin.Engine) {
	router.Static("/static", "./static")

	router.GET("/archive", post.Archive)
	mblog := router.Group("/mblog")
	mblog.GET("/posts/:id", post.Post)

	apiGroup := router.Group("/api")
	apiGroup.GET("/posts/:id", api.GetPost)
	apiGroup.GET("/comments/:id", api.ListComments)

	apiGroup.GET("/archive", api.Archive)
	apiGroup.GET("/archive/:id", api.Archive)
	apiGroup.GET("/archive/:id/counts/:count", api.Archive)
}
