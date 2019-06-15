package backstage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stdioa/inside-go/db"
	"github.com/stdioa/inside-go/vm/export"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "data.html", gin.H{})
}

func Export(c *gin.Context) {
	// List all posts
	posts := db.AllPosts()
	data, err := json.Marshal(export.GenerateVM(posts))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	// Code from gin/render/json.go
	buffer := new(bytes.Buffer)
	for _, r := range string(data) {
		cvt := string(r)
		if r >= 128 {
			cvt = fmt.Sprintf("\\u%04x", int64(r))
		}
		buffer.WriteString(cvt)
	}

	c.DataFromReader(
		http.StatusOK,
		int64(buffer.Len()),
		"application/octet-stream",
		buffer,
		map[string]string{
			"Content-Disposition": `attachment; filename="inside-data.json"`,
		})
}

func Import(c *gin.Context) {
	header, err := c.FormFile("data")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, export.ImportResult{
			Message: "File upload error: " + err.Error(),
		})
		return
	}
	file, _ := header.Open()
	decoder := json.NewDecoder(file)

	dataModel := new(export.ExportVM)
	err = decoder.Decode(dataModel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, export.ImportResult{
			Message: "JSON decode error: " + err.Error(),
		})
		return
	}
	export.SaveVMtoDB(dataModel.Posts)
	c.JSON(http.StatusOK, export.ImportResult{
		Success: true,
	})
}
