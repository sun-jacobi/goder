package main

import (
	"goder/judger"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func getProbEndPoint(c *gin.Context) {
	//TODO
}

func judgeEndPoint(c *gin.Context) {

	pid := c.PostForm("pid") // problem id
	// uid := c.PostForm("uid")

	src, err := c.FormFile("code")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	srcPath := filepath.Base(src.Filename)

	if err := c.SaveUploadedFile(src, srcPath); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	tests, err := DB.GetTest(pid)
	if err != nil {
		c.String(http.StatusBadRequest, "get test err: %s", err.Error())
		return
	}
	score := 0
	var failed []int
	for index, test := range tests {
		result, err := judger.Judge(srcPath, test)
		if err != nil {
			c.String(http.StatusBadRequest, "running test %d err: %s", index, err.Error())
		}
		if result == true {
			score += 1
		} else {
			failed = append(failed, index)
		}
	}

}

func uploadEndPoint(c *gin.Context) {
	//TODO
}

// Set up the router
func SetUpRouter() {
	Router = gin.Default()
	api := Router.Group("/api")
	{
		api.POST("/judge/:id", judgeEndPoint)
		api.POST("/upload/:id", uploadEndPoint)
		api.GET("/prob/:id", getProbEndPoint)
	}

}
