package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stdioa/inside-go/router"
)

func main() {
	app := gin.Default()
	router.Register(app)
	http.ListenAndServe(":8000", app)
}
