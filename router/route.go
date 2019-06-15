package router

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/stdioa/inside-go/controller/api"
	"github.com/stdioa/inside-go/controller/post"
)

var authHandler = gin.BasicAuth(gin.Accounts{
	"admin": "password",
})

func Register(router *gin.Engine) {
	router.Static("/static", "./static")
	router.HTMLRender = loadTemplates("./templates")

	router.GET("/", post.Index)
	router.GET("/archive", post.Archive)
	mblog := router.Group("/mblog")
	mblog.GET("/new", authHandler, post.NewPostPage)
	mblog.POST("/new", authHandler, post.NewPostHandler)
	mblog.GET("/posts/:id", post.Post)
	mblog.GET("/posts/:id/edit", authHandler, post.EditPost)

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

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
