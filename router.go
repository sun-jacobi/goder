package main

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func getProbEndPoint(c *gin.Context) {
	id := c.Param("id")
	prob, err := DB.GetProblem(1)
	c.JSON()
}

func judgeEndPoint(c *gin.Context) {
	id := c.Query("id")
	src, err := c.FormFile("src")

}

func uploadEndPoint(c *gin.Context) {
	form, _ := c.MultipartForm()
	inputs := form.File["input"]
	outputs := form.File["output"]
}

// Set up the router
func SetUpRouter() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/judge", judgeEndPoint)
		api.POST("/upload", uploadEndPoint)
		api.GET("/prob/:id", getProbEndPoint)
	}
}
