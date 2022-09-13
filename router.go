package main

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func getProbEndPoint(c *gin.Context) {
	//TODO
}

func judgeEndPoint(c *gin.Context) {
	// TODO
}

func uploadEndPoint(c *gin.Context) {
	//TODO
}

// Set up the router
func SetUpRouter() {
	Router = gin.Default()
	api := Router.Group("/api")
	{
		api.POST("/judge", judgeEndPoint)
		api.POST("/upload", uploadEndPoint)
		api.GET("/prob/:id", getProbEndPoint)
	}

}
